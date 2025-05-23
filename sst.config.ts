/// <reference path="./.sst/platform/config.d.ts" />

export default $config({
  app(input) {
    return {
      name: "sloth-util",
      removal: input?.stage === "production" ? "retain" : "remove",
      home: "aws",
      providers: {
        aws: {
          region: "us-east-1",
        },
      },
    };
  },
  async run() {
    // Environment variables and configurations
    const stage = $app.stage;
    const isProduction = stage === "production";
    const isDevelopment = stage === "development" || stage === "dev";
    
    // Architecture selection based on stage or environment variable
    const useAwsNative = process.env.ARCHITECTURE_TYPE !== "cloud-agnostic";
    
    // Common Lambda configuration
    const lambdaConfig = {
      runtime: "java17",
      timeout: "30 seconds",
      memory: "1024 MB",
      environment: {
        STAGE: stage,
        AWS_REGION: aws.getRegionOutput().name,
        BEDROCK_MODEL_ID: process.env.BEDROCK_MODEL_ID || "anthropic.claude-instant-v1",
        LOG_LEVEL: isDevelopment ? "DEBUG" : "INFO",
      },
    };

    // Storage configuration based on architecture
    let storage: any = {};
    
    if (useAwsNative) {
      // AWS-Native: Use DynamoDB and S3
      storage.dynamodb = new aws.dynamodb.Table("sloth-util-quotes", {
        name: `sloth-util-quotes-${stage}`,
        billingMode: "PAY_PER_REQUEST",
        hashKey: "id",
        attributes: [{
          name: "id",
          type: "S",
        }, {
          name: "category",
          type: "S",
        }],
        globalSecondaryIndexes: [{
          name: "CategoryIndex",
          hashKey: "category",
          projectionType: "ALL",
        }],
        tags: {
          Environment: stage,
        },
        ttl: {
          attributeName: "ttl",
          enabled: true,
        },
      });

      storage.s3 = new aws.s3.Bucket("sloth-util-config", {
        bucket: `sloth-util-config-${stage}-${Math.random().toString(36).substr(2, 9)}`,
        versioning: {
          enabled: true,
        },
        serverSideEncryptionConfiguration: {
          rule: {
            applyServerSideEncryptionByDefault: {
              sseAlgorithm: "AES256",
            },
          },
        },
        tags: {
          Environment: stage,
        },
      });
    } else {
      // Cloud-agnostic: External database references (not created here)
      storage.databaseUrl = process.env.DATABASE_URL || "postgresql://localhost:5432/slothutil";
      storage.redisUrl = process.env.REDIS_URL || "redis://localhost:6379";
    }

    // IAM role for Lambda functions
    const lambdaRole = new aws.iam.Role("sloth-util-lambda-role", {
      assumeRolePolicy: JSON.stringify({
        Version: "2012-10-17",
        Statement: [{
          Action: "sts:AssumeRole",
          Effect: "Allow",
          Principal: {
            Service: "lambda.amazonaws.com",
          },
        }],
      }),
      managedPolicyArns: [
        "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole",
      ],
    });

    // IAM policy for Bedrock access
    const bedrockPolicy = new aws.iam.RolePolicy("sloth-util-bedrock-policy", {
      role: lambdaRole.id,
      policy: JSON.stringify({
        Version: "2012-10-17",
        Statement: [{
          Effect: "Allow",
          Action: [
            "bedrock:InvokeModel",
            "bedrock:InvokeModelWithResponseStream",
          ],
          Resource: [
            `arn:aws:bedrock:us-east-1::foundation-model/*`,
          ],
        }],
      }),
    });

    // Additional IAM policies for AWS-native architecture
    if (useAwsNative && storage.dynamodb) {
      new aws.iam.RolePolicy("sloth-util-dynamodb-policy", {
        role: lambdaRole.id,
        policy: storage.dynamodb.arn.apply(arn => JSON.stringify({
          Version: "2012-10-17",
          Statement: [{
            Effect: "Allow",
            Action: [
              "dynamodb:GetItem",
              "dynamodb:PutItem",
              "dynamodb:Query",
              "dynamodb:Scan",
              "dynamodb:UpdateItem",
              "dynamodb:DeleteItem",
            ],
            Resource: [arn, `${arn}/index/*`],
          }],
        })),
      });

      new aws.iam.RolePolicy("sloth-util-s3-policy", {
        role: lambdaRole.id,
        policy: storage.s3.arn.apply(arn => JSON.stringify({
          Version: "2012-10-17",
          Statement: [{
            Effect: "Allow",
            Action: [
              "s3:GetObject",
              "s3:PutObject",
            ],
            Resource: `${arn}/*`,
          }],
        })),
      });
    }

    // Quote Generator Lambda Function
    const quoteGenerator = new sst.aws.Function("QuoteGenerator", {
      handler: "com.slothutil.quotes.QuoteHandler::handleRequest",
      ...lambdaConfig,
      environment: {
        ...lambdaConfig.environment,
        ...(useAwsNative ? {
          DYNAMODB_TABLE_NAME: storage.dynamodb?.name || "",
          S3_CONFIG_BUCKET: storage.s3?.bucket || "",
        } : {
          DATABASE_URL: storage.databaseUrl,
          REDIS_URL: storage.redisUrl,
        }),
        ARCHITECTURE_TYPE: useAwsNative ? "aws-native" : "cloud-agnostic",
      },
      role: lambdaRole.arn,
    });

    // Authentication Lambda Function (for cloud-agnostic architecture)
    const authService = !useAwsNative ? new sst.aws.Function("AuthService", {
      handler: "com.slothutil.auth.AuthHandler::handleRequest",
      ...lambdaConfig,
      environment: {
        ...lambdaConfig.environment,
        JWT_SECRET: process.env.JWT_SECRET || "your-secret-key-change-in-production",
        JWT_ISSUER: `sloth-util-${stage}`,
        JWT_EXPIRY: "3600", // 1 hour
        DATABASE_URL: storage.databaseUrl,
      },
      role: lambdaRole.arn,
    }) : undefined;

    // JWKS Lambda Function (for cloud-agnostic architecture)
    const jwksService = !useAwsNative ? new sst.aws.Function("JWKSService", {
      handler: "com.slothutil.auth.JWKSHandler::handleRequest",
      ...lambdaConfig,
      environment: {
        ...lambdaConfig.environment,
        JWT_SECRET: process.env.JWT_SECRET || "your-secret-key-change-in-production",
        JWT_ISSUER: `sloth-util-${stage}`,
      },
      role: lambdaRole.arn,
    }) : undefined;

    // API Gateway setup (AWS-Native architecture only)
    let api: any = undefined;
    if (useAwsNative) {
      api = new sst.aws.ApiGatewayV2("SlothUtilApi", {
        cors: {
          allowCredentials: true,
          allowOrigins: ["*"],
          allowMethods: ["GET", "POST", "PUT", "DELETE", "OPTIONS"],
          allowHeaders: ["Content-Type", "Authorization"],
        },
      });

      api.route("GET /quotes/random", quoteGenerator.arn);
      api.route("OPTIONS /quotes/random", quoteGenerator.arn);

      // Cognito User Pool for authentication
      const userPool = new aws.cognito.UserPool("sloth-util-users", {
        name: `sloth-util-users-${stage}`,
        passwordPolicy: {
          minimumLength: 8,
          requireNumbers: true,
          requireSymbols: true,
          requireUppercase: true,
          requireLowercase: true,
        },
        autoVerifiedAttributes: ["email"],
        usernameAttributes: ["email"],
        tags: {
          Environment: stage,
        },
      });

      const userPoolClient = new aws.cognito.UserPoolClient("sloth-util-client", {
        name: `sloth-util-client-${stage}`,
        userPoolId: userPool.id,
        explicitAuthFlows: [
          "ADMIN_NO_SRP_AUTH",
          "USER_PASSWORD_AUTH",
        ],
        generateSecret: false,
      });
    } else {
      // Function URLs for cloud-agnostic architecture (simpler than ALB)
      const quoteGeneratorUrl = new aws.lambda.FunctionUrl("quote-generator-url", {
        functionName: quoteGenerator.name,
        authorizationType: "NONE",
        cors: {
          allowCredentials: true,
          allowOrigins: ["*"],
          allowMethods: ["GET", "POST", "OPTIONS"],
          allowHeaders: ["Content-Type", "Authorization"],
          maxAge: 300,
        },
      });

      const authServiceUrl = new aws.lambda.FunctionUrl("auth-service-url", {
        functionName: authService!.name,
        authorizationType: "NONE",
        cors: {
          allowCredentials: true,
          allowOrigins: ["*"],
          allowMethods: ["POST", "OPTIONS"],
          allowHeaders: ["Content-Type", "Authorization"],
          maxAge: 300,
        },
      });

      const jwksServiceUrl = new aws.lambda.FunctionUrl("jwks-service-url", {
        functionName: jwksService!.name,
        authorizationType: "NONE",
        cors: {
          allowCredentials: true,
          allowOrigins: ["*"],
          allowMethods: ["GET", "OPTIONS"],
          allowHeaders: ["Content-Type"],
          maxAge: 3600,
        },
      });
    }

    // Outputs
    return {
      architecture: useAwsNative ? "aws-native" : "cloud-agnostic",
      stage: stage,
      ...(useAwsNative ? {
        apiUrl: api?.url,
        userPoolId: userPool?.id,
        userPoolClientId: userPoolClient?.id,
        dynamoTableName: storage.dynamodb?.name,
        s3BucketName: storage.s3?.bucket,
      } : {
        quoteGeneratorUrl: quoteGeneratorUrl?.functionUrl,
        authServiceUrl: authServiceUrl?.functionUrl,
        jwksServiceUrl: jwksServiceUrl?.functionUrl,
      }),
      region: aws.getRegionOutput().name,
      lambdaRoleArn: lambdaRole.arn,
    };
  },
});