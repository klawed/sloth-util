package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	projectName = "sloth-util"
	groupId     = "com.slothutil"
	javaVersion = "17"
	springBootVersion = "3.2.0"
	springCloudVersion = "2023.0.0"
)

type ProjectStructure struct {
	Path     string
	IsDir    bool
	Content  string
	Template bool
}

func main() {
	fmt.Println("🚀 Setting up Sloth Util SpringBoot Project Structure...")
	
	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get current directory:", err)
	}
	
	fmt.Printf("📁 Working directory: %s\n", cwd)
	
	// Define the complete project structure
	structure := getProjectStructure()
	
	// Create directories and files
	for _, item := range structure {
		fullPath := filepath.Join(cwd, item.Path)
		
		if item.IsDir {
			if err := createDirectory(fullPath); err != nil {
				log.Printf("❌ Failed to create directory %s: %v", item.Path, err)
				continue
			}
			fmt.Printf("📁 Created directory: %s\n", item.Path)
		} else {
			if err := createFile(fullPath, item.Content, item.Template); err != nil {
				log.Printf("❌ Failed to create file %s: %v", item.Path, err)
				continue
			}
			fmt.Printf("📄 Created file: %s\n", item.Path)
		}
	}
	
	fmt.Println("\n🎉 Project structure created successfully!")
	fmt.Println("\n📋 Next Steps:")
	fmt.Println("1. Run: cd packages/functions/quote-generator && mvn clean install")
	fmt.Println("2. Run: docker-compose up -d")
	fmt.Println("3. Run: npx sst dev")
	fmt.Println("4. Check README.md for detailed setup instructions")
}

func createDirectory(path string) error {
	return os.MkdirAll(path, 0755)
}

func createFile(path, content string, isTemplate bool) error {
	// Create parent directory if it doesn't exist
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	
	// Don't overwrite existing files
	if _, err := os.Stat(path); err == nil {
		fmt.Printf("⚠️  File exists, skipping: %s\n", path)
		return nil
	}
	
	// Process template if needed
	if isTemplate {
		content = processTemplate(content)
	}
	
	return os.WriteFile(path, []byte(content), 0644)
}

func processTemplate(content string) string {
	content = strings.ReplaceAll(content, "{{PROJECT_NAME}}", projectName)
	content = strings.ReplaceAll(content, "{{GROUP_ID}}", groupId)
	content = strings.ReplaceAll(content, "{{JAVA_VERSION}}", javaVersion)
	content = strings.ReplaceAll(content, "{{SPRING_BOOT_VERSION}}", springBootVersion)
	content = strings.ReplaceAll(content, "{{SPRING_CLOUD_VERSION}}", springCloudVersion)
	return content
}

func getProjectStructure() []ProjectStructure {
	return []ProjectStructure{
		// Root directories
		{Path: "packages", IsDir: true},
		{Path: "packages/functions", IsDir: true},
		{Path: "packages/core", IsDir: true},
		{Path: "packages/infrastructure", IsDir: true},
		{Path: "scripts", IsDir: true},
		{Path: "docs", IsDir: true},
		{Path: "tools", IsDir: true},
		{Path: ".github", IsDir: true},
		{Path: ".github/workflows", IsDir: true},
		
		// Quote Generator Service
		{Path: "packages/functions/quote-generator", IsDir: true},
		{Path: "packages/functions/quote-generator/src", IsDir: true},
		{Path: "packages/functions/quote-generator/src/main", IsDir: true},
		{Path: "packages/functions/quote-generator/src/main/java", IsDir: true},
		{Path: "packages/functions/quote-generator/src/main/java/com", IsDir: true},
		{Path: "packages/functions/quote-generator/src/main/java/com/slothutil", IsDir: true},
		{Path: "packages/functions/quote-generator/src/main/java/com/slothutil/quotes", IsDir: true},
		{Path: "packages/functions/quote-generator/src/main/java/com/slothutil/quotes/config", IsDir: true},
		{Path: "packages/functions/quote-generator/src/main/java/com/slothutil/quotes/service", IsDir: true},
		{Path: "packages/functions/quote-generator/src/main/java/com/slothutil/quotes/model", IsDir: true},
		{Path: "packages/functions/quote-generator/src/main/resources", IsDir: true},
		{Path: "packages/functions/quote-generator/src/test", IsDir: true},
		{Path: "packages/functions/quote-generator/src/test/java", IsDir: true},
		{Path: "packages/functions/quote-generator/src/test/java/com", IsDir: true},
		{Path: "packages/functions/quote-generator/src/test/java/com/slothutil", IsDir: true},
		{Path: "packages/functions/quote-generator/src/test/java/com/slothutil/quotes", IsDir: true},
		
		// Auth Service
		{Path: "packages/functions/auth-service", IsDir: true},
		{Path: "packages/functions/auth-service/src", IsDir: true},
		{Path: "packages/functions/auth-service/src/main", IsDir: true},
		{Path: "packages/functions/auth-service/src/main/java", IsDir: true},
		{Path: "packages/functions/auth-service/src/main/java/com", IsDir: true},
		{Path: "packages/functions/auth-service/src/main/java/com/slothutil", IsDir: true},
		{Path: "packages/functions/auth-service/src/main/java/com/slothutil/auth", IsDir: true},
		{Path: "packages/functions/auth-service/src/main/java/com/slothutil/auth/config", IsDir: true},
		{Path: "packages/functions/auth-service/src/main/java/com/slothutil/auth/service", IsDir: true},
		{Path: "packages/functions/auth-service/src/main/java/com/slothutil/auth/model", IsDir: true},
		{Path: "packages/functions/auth-service/src/main/resources", IsDir: true},
		{Path: "packages/functions/auth-service/src/test", IsDir: true},
		{Path: "packages/functions/auth-service/src/test/java", IsDir: true},
		{Path: "packages/functions/auth-service/src/test/java/com", IsDir: true},
		{Path: "packages/functions/auth-service/src/test/java/com/slothutil", IsDir: true},
		{Path: "packages/functions/auth-service/src/test/java/com/slothutil/auth", IsDir: true},
		
		// JWKS Service
		{Path: "packages/functions/jwks-service", IsDir: true},
		{Path: "packages/functions/jwks-service/src", IsDir: true},
		{Path: "packages/functions/jwks-service/src/main", IsDir: true},
		{Path: "packages/functions/jwks-service/src/main/java", IsDir: true},
		{Path: "packages/functions/jwks-service/src/main/java/com", IsDir: true},
		{Path: "packages/functions/jwks-service/src/main/java/com/slothutil", IsDir: true},
		{Path: "packages/functions/jwks-service/src/main/java/com/slothutil/jwks", IsDir: true},
		{Path: "packages/functions/jwks-service/src/main/java/com/slothutil/jwks/config", IsDir: true},
		{Path: "packages/functions/jwks-service/src/main/java/com/slothutil/jwks/service", IsDir: true},
		{Path: "packages/functions/jwks-service/src/main/java/com/slothutil/jwks/model", IsDir: true},
		{Path: "packages/functions/jwks-service/src/main/resources", IsDir: true},
		{Path: "packages/functions/jwks-service/src/test", IsDir: true},
		{Path: "packages/functions/jwks-service/src/test/java", IsDir: true},
		{Path: "packages/functions/jwks-service/src/test/java/com", IsDir: true},
		{Path: "packages/functions/jwks-service/src/test/java/com/slothutil", IsDir: true},
		{Path: "packages/functions/jwks-service/src/test/java/com/slothutil/jwks", IsDir: true},
		
		// Common/Shared utilities
		{Path: "packages/functions/common", IsDir: true},
		{Path: "packages/functions/common/src", IsDir: true},
		{Path: "packages/functions/common/src/main", IsDir: true},
		{Path: "packages/functions/common/src/main/java", IsDir: true},
		{Path: "packages/functions/common/src/main/java/com", IsDir: true},
		{Path: "packages/functions/common/src/main/java/com/slothutil", IsDir: true},
		{Path: "packages/functions/common/src/main/java/com/slothutil/common", IsDir: true},
		{Path: "packages/functions/common/src/main/java/com/slothutil/common/exception", IsDir: true},
		{Path: "packages/functions/common/src/main/java/com/slothutil/common/util", IsDir: true},
		{Path: "packages/functions/common/src/main/java/com/slothutil/common/model", IsDir: true},
		{Path: "packages/functions/common/src/main/resources", IsDir: true},
		{Path: "packages/functions/common/src/test", IsDir: true},
		{Path: "packages/functions/common/src/test/java", IsDir: true},
		
		// Maven POM files  
		{Path: "packages/functions/pom.xml", Content: getParentPom(), Template: true},
		
		// Configuration files
		{Path: ".env.example", Content: getEnvExample(), Template: true},
		{Path: "docker-compose.yml", Content: getDockerCompose(), Template: true},
		
		// Scripts
		{Path: "scripts/deploy.sh", Content: getDeployScript(), Template: true},
		{Path: "scripts/test.sh", Content: getTestScript(), Template: true},
		{Path: "scripts/build.sh", Content: getBuildScript(), Template: true},
	}
}

// Template functions (basic versions for now)
func getParentPom() string {
	return `<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 
         http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>

    <groupId>{{GROUP_ID}}</groupId>
    <artifactId>{{PROJECT_NAME}}-functions</artifactId>
    <version>1.0.0-SNAPSHOT</version>
    <packaging>pom</packaging>

    <name>Sloth Util Functions</name>
    <description>AWS Lambda utilities for microservices</description>

    <properties>
        <maven.compiler.source>{{JAVA_VERSION}}</maven.compiler.source>
        <maven.compiler.target>{{JAVA_VERSION}}</maven.compiler.target>
        <java.version>{{JAVA_VERSION}}</java.version>
        <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
        <spring-boot.version>{{SPRING_BOOT_VERSION}}</spring-boot.version>
        <spring-cloud.version>{{SPRING_CLOUD_VERSION}}</spring-cloud.version>
        <aws-lambda-java.version>1.2.3</aws-lambda-java.version>
        <aws-sdk.version>2.21.29</aws-sdk.version>
    </properties>

    <modules>
        <module>common</module>
        <module>quote-generator</module>
        <module>auth-service</module>
        <module>jwks-service</module>
    </modules>

    <dependencyManagement>
        <dependencies>
            <!-- Spring Boot BOM -->
            <dependency>
                <groupId>org.springframework.boot</groupId>
                <artifactId>spring-boot-dependencies</artifactId>
                <version>${spring-boot.version}</version>
                <type>pom</type>
                <scope>import</scope>
            </dependency>
            
            <!-- Spring Cloud BOM -->
            <dependency>
                <groupId>org.springframework.cloud</groupId>
                <artifactId>spring-cloud-dependencies</artifactId>
                <version>${spring-cloud.version}</version>
                <type>pom</type>
                <scope>import</scope>
            </dependency>
            
            <!-- AWS SDK BOM -->
            <dependency>
                <groupId>software.amazon.awssdk</groupId>
                <artifactId>bom</artifactId>
                <version>${aws-sdk.version}</version>
                <type>pom</type>
                <scope>import</scope>
            </dependency>
            
            <!-- AWS Lambda Java Core -->
            <dependency>
                <groupId>com.amazonaws</groupId>
                <artifactId>aws-lambda-java-core</artifactId>
                <version>${aws-lambda-java.version}</version>
            </dependency>
        </dependencies>
    </dependencyManagement>

    <build>
        <pluginManagement>
            <plugins>
                <plugin>
                    <groupId>org.springframework.boot</groupId>
                    <artifactId>spring-boot-maven-plugin</artifactId>
                    <version>${spring-boot.version}</version>
                </plugin>
                <plugin>
                    <groupId>org.apache.maven.plugins</groupId>
                    <artifactId>maven-compiler-plugin</artifactId>
                    <version>3.11.0</version>
                    <configuration>
                        <source>${java.version}</source>
                        <target>${java.version}</target>
                    </configuration>
                </plugin>
                <plugin>
                    <groupId>org.apache.maven.plugins</groupId>
                    <artifactId>maven-surefire-plugin</artifactId>
                    <version>3.0.0</version>
                </plugin>
            </plugins>
        </pluginManagement>
    </build>
</project>`
}

func getEnvExample() string {
	return `# Environment Configuration for Sloth Util
# Copy this file to .env and update with your values

# Stage/Environment
STAGE=local
ARCHITECTURE_TYPE=cloud-agnostic

# AWS Configuration  
AWS_REGION=us-east-1
BEDROCK_MODEL_ID=anthropic.claude-instant-v1

# Database Configuration (Cloud-Agnostic)
DATABASE_URL=postgresql://username:password@host:5432/slothutil
REDIS_URL=redis://host:6379

# Authentication Configuration
JWT_SECRET=your-secret-key-change-in-production-must-be-at-least-32-characters
JWT_ISSUER=sloth-util-local
JWT_EXPIRY=3600

# Logging
LOG_LEVEL=DEBUG

# Local Development Services
DYNAMODB_ENDPOINT=http://localhost:8000
POSTGRES_URL=postgresql://localhost:5432/slothutil
REDIS_URL=redis://localhost:6379

# Cloudflare Configuration (for production)
CLOUDFLARE_API_TOKEN=your-cloudflare-api-token
CLOUDFLARE_ZONE_ID=your-zone-id
CLOUDFLARE_DOMAIN=your-domain.com`
}

func getDockerCompose() string {
	return `version: '3.8'

services:
  # PostgreSQL for cloud-agnostic architecture
  postgres:
    image: postgres:15-alpine
    container_name: sloth-util-postgres
    environment:
      POSTGRES_DB: slothutil
      POSTGRES_USER: slothutil
      POSTGRES_PASSWORD: password
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data
      - ./tools/db/init.sql:/docker-entrypoint-initdb.d/init.sql
    networks:
      - sloth-util-network

  # Redis for caching
  redis:
    image: redis:7-alpine
    container_name: sloth-util-redis
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data
    networks:
      - sloth-util-network

  # DynamoDB Local for AWS services simulation
  dynamodb-local:
    image: amazon/dynamodb-local:latest
    container_name: sloth-util-dynamodb
    ports:
      - "8000:8000"
    command: ["-jar", "DynamoDBLocal.jar", "-sharedDb", "-inMemory"]
    networks:
      - sloth-util-network

  # LocalStack for AWS services simulation (optional)
  localstack:
    image: localstack/localstack:latest
    container_name: sloth-util-localstack
    ports:
      - "4566:4566"
    environment:
      - SERVICES=lambda,s3,iam,logs
      - DEBUG=1
      - DATA_DIR=/tmp/localstack/data
    volumes:
      - localstack_data:/tmp/localstack
    networks:
      - sloth-util-network

volumes:
  postgres_data:
  redis_data:
  localstack_data:

networks:
  sloth-util-network:
    driver: bridge`
}

func getDeployScript() string {
	return `#!/bin/bash

# Deployment script for Sloth Util
set -e

STAGE=${1:-dev}
REGION=${AWS_REGION:-us-east-1}

echo "🚀 Deploying Sloth Util to stage: $STAGE"

# Build all Maven projects
echo "📦 Building Maven projects..."
cd packages/functions
mvn clean install -q
cd ../..

# Deploy with SST
echo "🌐 Deploying with SST..."
npx sst deploy --stage $STAGE

echo "✅ Deployment completed successfully!"
echo "📋 Next steps:"
echo "  1. Check AWS Console for deployed resources"
echo "  2. Test the deployed endpoints"
echo "  3. Monitor CloudWatch logs"`
}

func getTestScript() string {
	return `#!/bin/bash

# Test script for Sloth Util
set -e

echo "🧪 Running Sloth Util Tests..."

# Unit tests
echo "📋 Running unit tests..."
cd packages/functions
mvn test

# Integration tests
echo "🔗 Running integration tests..."
mvn verify -P integration-tests

echo "✅ All tests passed!"
cd ../..`
}

func getBuildScript() string {
	return `#!/bin/bash

# Build script for Sloth Util
set -e

echo "🔨 Building Sloth Util..."

# Install Node.js dependencies
echo "📦 Installing Node.js dependencies..."
npm install

# Build all Maven projects
echo "☕ Building Maven projects..."
cd packages/functions
mvn clean compile
cd ../..

echo "✅ Build completed successfully!"`
}
