# Sloth Util

**Cloud-Agnostic AWS Lambda utilities for microservices** - A collection of reusable Lambda functions designed to provide common functionality across multiple applications using external services and Cloudflare's free tier.

## 🚀 Quick Start

**1. Run the setup script:**
```bash
go run scripts/setup.go
```

**2. Start local services:**
```bash
docker-compose up -d
```

**3. Build and test:**
```bash
cd packages/functions
mvn clean install
```

**4. Deploy to development:**
```bash
npx sst dev
```

## Overview

Sloth Util provides a suite of serverless utility functions deployed as AWS Lambda functions using Spring Boot Java. The functions are designed to be lightweight, reusable, and easily integrated into your existing microservices architecture. 

**✅ Now focusing exclusively on cloud-agnostic architecture** - uses external PostgreSQL, Redis, and Cloudflare's free tier for optimal cost efficiency and vendor independence.

### Current Features

- **Random Quote Generator**: Leverages AWS Bedrock LLM to generate inspirational, motivational, or contextual quotes on demand
- **JWT Authentication Service**: Custom JWT-based authentication with PostgreSQL user management
- **JWKS Service**: JSON Web Key Set endpoint for public key distribution and JWT verification
- **Common Utilities**: Shared models, exception handling, and response utilities

### Planned Features

- Data transformation functions
- Notification services  
- File processing utilities
- Multi-region deployment support

## Architecture

This project uses a **cloud-agnostic architecture** to avoid vendor lock-in while leveraging AWS Lambda for compute and external services for data storage.

```mermaid
graph TB
    subgraph "Client Applications"
        A[Mobile App]
        B[Web App]
        C[Other Services]
    end
    
    subgraph "Cloudflare (Free Tier)"
        CF[Cloudflare<br/>CDN + Load Balancer<br/>SSL/TLS + DDoS Protection]
    end
    
    subgraph "Authentication Layer"
        AUTH[JWT Auth Service<br/>Lambda Function]
        JWKS[JWKS Endpoint<br/>Lambda Function]
    end
    
    subgraph "Core Services"
        L1[Quote Generator<br/>Lambda Function]
        L2[Future Lambda<br/>Functions]
    end
    
    subgraph "AWS AI Service"
        BR[AWS Bedrock<br/>LLM Service]
    end
    
    subgraph "External Storage"
        DB[(External PostgreSQL<br/>Managed Database)]
        CACHE[(External Redis<br/>Managed Cache)]
    end
    
    subgraph "Monitoring"
        MON[Custom Monitoring<br/>Prometheus/Grafana]
    end
    
    A --> CF
    B --> CF
    C --> CF
    CF --> AUTH
    AUTH --> JWKS
    AUTH --> L1
    AUTH --> L2
    L1 --> BR
    L1 --> DB
    L1 --> CACHE
    L1 --> MON
    
    style AUTH fill:#4CAF50
    style L1 fill:#4CAF50
    style BR fill:#ff9900
    style DB fill:#9C27B0
    style CF fill:#f39c12
```

### Architecture Benefits

✅ **Cost Effective**: Uses Cloudflare free tier + external databases  
✅ **Scalable**: Lambda functions scale automatically  
✅ **Maintainable**: Clean separation of concerns  
✅ **Testable**: Comprehensive testing setup with TestContainers  
✅ **Portable**: Not locked into AWS-specific services  
✅ **Secure**: Custom JWT implementation with full control

## Technology Stack

- **Runtime**: Java 17+ with Spring Boot 3.2
- **Build Tool**: Maven 3.8+
- **Deployment**: SST (Serverless Stack) v3
- **Infrastructure**: AWS Lambda with Function URLs
- **AI/ML**: AWS Bedrock (Claude/Titan models)
- **Database**: External PostgreSQL (cloud-agnostic)
- **Cache**: External Redis (cloud-agnostic)
- **Authentication**: Custom JWT implementation
- **CDN/Load Balancer**: Cloudflare (free tier)
- **Testing**: JUnit 5 + TestContainers

## Cost Analysis

### Traffic Scenarios

| Scenario | Requests/Month | Concurrent Users | Peak RPS |
|----------|----------------|------------------|----------|
| **Low** | 100K | 10-50 | 5 |
| **Medium** | 1M | 100-500 | 50 |
| **High** | 10M | 1K-5K | 500 |

### AWS-Native Architecture Costs

| Component | Low Traffic | Medium Traffic | High Traffic |
|-----------|-------------|----------------|--------------|
| **Lambda Invocations** | $0.20 | $2.00 | $20.00 |
| **Lambda Duration** (1GB, 500ms avg) | $0.83 | $8.33 | $83.33 |
| **API Gateway** | $0.35 | $3.50 | $35.00 |
| **AWS Cognito** | $0.00 | $5.50 | $55.00 |
| **AWS Bedrock** (Claude Instant) | $15.00 | $150.00 | $1,500.00 |
| **DynamoDB** (On-Demand) | $2.50 | $12.50 | $62.50 |
| **S3** (Config storage) | $0.25 | $0.50 | $2.50 |
| **CloudWatch Logs** | $0.50 | $2.50 | $12.50 |
| **X-Ray Tracing** | $0.50 | $5.00 | $50.00 |
| **Data Transfer** | $0.90 | $4.50 | $22.50 |
| **Total Monthly Cost** | **$21.03** | **$194.33** | **$1,843.33** |

### Cloud-Agnostic Architecture Costs

| Component | Low Traffic | Medium Traffic | High Traffic |
|-----------|-------------|----------------|--------------|
| **Lambda Invocations** | $0.20 | $2.00 | $20.00 |
| **Lambda Duration** (1GB, 500ms avg) | $0.83 | $8.33 | $83.33 |
| **Function URLs** (Free) | $0.00 | $0.00 | $0.00 |
| **Cloudflare Free** | $0.00 | $0.00 | $0.00 |
| **AWS Bedrock** (Claude Instant) | $15.00 | $150.00 | $1,500.00 |
| **External PostgreSQL** (managed) | $25.00 | $100.00 | $400.00 |
| **External Redis** (managed) | $15.00 | $60.00 | $240.00 |
| **CloudWatch Logs** | $0.50 | $2.50 | $12.50 |
| **Monitoring Stack** | $0.00 | $25.00 | $100.00 |
| **Data Transfer** | $0.90 | $5.40 | $27.00 |
| **Total Monthly Cost** | **$57.43** | **$353.23** | **$2,382.83** |

### Cost Comparison Summary

| Traffic Level | AWS-Native | Cloud-Agnostic | Difference |
|---------------|------------|----------------|------------|
| **Low** | $21.03 | $57.43 | +$36.40 (+173%) |
| **Medium** | $194.33 | $353.23 | +$158.90 (+82%) |
| **High** | $1,843.33 | $2,382.83 | +$539.50 (+29%) |

**Key Insights:**
- AWS-Native is significantly cheaper at low traffic volumes
- Cost difference decreases as traffic increases  
- Cloud-agnostic provides more control but at higher base cost
- Both architectures scale cost-effectively with traffic
- Cloud-agnostic eliminates vendor lock-in and provides operational flexibility

**Cloud-Agnostic Benefits Despite Higher Cost:**
- No vendor lock-in - can switch providers easily
- Full control over database and caching infrastructure
- Better performance with dedicated external services
- Simplified architecture without VPC complexity
- Cloudflare free tier includes enterprise-grade CDN and security
- More predictable pricing with external managed services

## Local Development Environment

### Prerequisites

- **Java 17+** (Amazon Corretto recommended)
- **Maven 3.8+**
- **Node.js 18+** (for SST)
- **Docker & Docker Compose** (for local services)
- **AWS CLI** configured with appropriate permissions
- **Go 1.19+** (for setup scripts)

### Local Setup

```bash
# 1. Clone repository
git clone https://github.com/klawed/sloth-util.git
cd sloth-util

# 2. Setup project structure
go run scripts/setup.go

# 3. Install dependencies
npm install
cd packages/functions && mvn clean install && cd ../..

# 4. Configure environment
cp .env.example .env
# Edit .env with your configuration

# 5. Start local development services
docker-compose up -d

# 6. Start SST development mode
npx sst dev
```

### Local Development Stack

```yaml
# docker-compose.yml services
services:
  - PostgreSQL (port 5432)
  - Redis (port 6379)
  - LocalStack (AWS services simulation)
```

### Development Workflow

1. **Code Changes**: Edit Java functions in `packages/functions/`
2. **Hot Reload**: SST automatically rebuilds and redeploys on changes
3. **Local Testing**: Use local endpoints provided by `sst dev`
4. **Integration Tests**: Run against local Docker services
5. **Debug**: Attach debugger to running Lambda functions

### Environment Configuration

```bash
# .env.local (for local development)
STAGE=local
ARCHITECTURE_TYPE=cloud-agnostic
AWS_REGION=us-east-1
DATABASE_URL=postgresql://localhost:5432/slothutil
REDIS_URL=redis://localhost:6379
BEDROCK_MODEL_ID=anthropic.claude-instant-v1
JWT_SECRET=your-jwt-secret-key-min-32-chars
LOG_LEVEL=DEBUG
```

## CI/CD Pipeline

### GitHub Actions Workflow

The project uses GitHub Actions for continuous integration and deployment:

- **Automated Testing**: Unit tests, integration tests, security scans
- **Quality Checks**: SonarCloud code quality analysis
- **Development Deployment**: Auto-deploy on `develop` branch
- **Production Deployment**: Auto-deploy on `main` branch with approvals
- **Security**: Snyk vulnerability scanning

### Deployment Stages

1. **Development** (`develop` branch)
   - Automatic deployment on push
   - Full test suite execution
   - Integration with external services
   - Performance baseline testing

2. **Production** (`main` branch)
   - Automatic deployment on push
   - Comprehensive monitoring
   - Rollback capabilities

### Required GitHub Secrets

```bash
AWS_ACCESS_KEY_ID           # AWS deployment credentials
AWS_SECRET_ACCESS_KEY       # AWS deployment credentials
DEV_DATABASE_URL           # Development PostgreSQL connection
DEV_REDIS_URL              # Development Redis connection
DEV_JWT_SECRET             # Development JWT secret
PROD_DATABASE_URL          # Production PostgreSQL connection
PROD_REDIS_URL             # Production Redis connection
PROD_JWT_SECRET            # Production JWT secret
BEDROCK_MODEL_ID           # AWS Bedrock model identifier
SONAR_TOKEN                # SonarCloud integration
SNYK_TOKEN                 # Security scanning
```

### Testing Strategy

```bash
# Unit Tests
mvn test

# Integration Tests (with TestContainers)
mvn verify -P integration-tests

# End-to-End Tests
npm run test:e2e

# Performance Tests
mvn test -P performance-tests

# Security Tests
npm run test:security
```

## Project Structure

```
sloth-util/
├── sst.config.ts                    # SST configuration (cloud-agnostic)
├── docker-compose.yml               # Local development services
├── packages/
│   ├── functions/                   # Lambda function code
│   │   ├── quote-generator/         # Quote generator service
│   │   │   └── src/main/java/com/slothutil/quotes/
│   │   │       ├── QuoteHandler.java
│   │   │       ├── config/
│   │   │       ├── service/
│   │   │       └── model/
│   │   ├── auth-service/            # JWT authentication service
│   │   ├── jwks-service/            # JWKS endpoint service
│   │   └── common/                  # Shared utilities
│   ├── core/                        # Core business logic
│   └── infrastructure/              # SST stack definitions
├── scripts/
│   ├── setup.go                     # Project setup script
│   ├── deploy.sh                    # Deployment scripts
│   └── test.sh                      # Testing scripts
├── .github/
│   └── workflows/
│       └── ci-cd.yml                # GitHub Actions pipeline
├── docs/                            # Additional documentation
└── tools/                           # Development tools and configs
```

## API Endpoints

### Quote Generator

**GET** `/quotes/random`

Generate a random inspirational quote using AI.

**Query Parameters:**
- `category` (optional): Category of quote (motivational, tech, business, life)
- `length` (optional): Preferred length (short, medium, long)
- `cache` (optional): Use cached quotes if available (default: true)

**Headers:**
- `Authorization: Bearer <jwt-token>` (for custom JWT auth)
- `Content-Type: application/json`

**Response:**
```json
{
  "quote": "The best way to predict the future is to create it.",
  "author": "AI Generated",
  "category": "motivational",
  "length": "short",
  "timestamp": "2025-05-23T19:18:03Z",
  "cached": false,
  "requestId": "req-123-456-789"
}
```

### Authentication

**POST** `/auth/login` - User authentication  
**POST** `/auth/refresh` - Token refresh  
**GET** `/.well-known/jwks.json` - JWKS public keys

## Monitoring and Observability

### Cloud-Agnostic Monitoring
- Custom Prometheus metrics
- Grafana dashboards
- Structured logging with JSON format
- Custom health check endpoints
- Application performance monitoring
- CloudWatch logs (for Lambda functions)

## Security

- JWT-based authentication with custom implementation
- Rate limiting per client/IP (via Cloudflare)
- Input validation and sanitization
- Secrets managed securely via environment variables
- Regular security dependency updates
- OWASP compliance testing
- SSL/TLS termination via Cloudflare

## Getting Started

### Quick Setup

```bash
# 1. Clone repository
git clone https://github.com/klawed/sloth-util.git
cd sloth-util

# 2. Setup project structure
go run scripts/setup.go

# 3. Install dependencies
npm install
cd packages/functions && mvn clean install && cd ../..

# 4. Configure environment
cp .env.example .env
# Edit .env with your configuration

# 5. Start local development
docker-compose up -d
npx sst dev

# 6. Deploy to development
npx sst deploy --stage dev
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Follow the coding standards (Checkstyle + SpotBugs)
4. Write tests for new functionality
5. Commit your changes (`git commit -m 'Add amazing feature'`)
6. Push to the branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Roadmap

- [x] Project setup and cloud-agnostic architecture design
- [x] SpringBoot project structure with Go setup script
- [x] Complete Maven configuration with parent/child modules
- [x] Docker Compose for local development
- [x] CI/CD pipeline implementation
- [ ] Quote Generator Lambda implementation
- [ ] Custom JWT authentication service
- [ ] JWKS endpoint implementation
- [ ] Monitoring and alerting setup
- [ ] Performance optimization
- [ ] Additional utility functions
- [ ] Multi-region deployment support
- [ ] Cloudflare integration and configuration

## Support

For support and questions:
- Create an issue in this repository
- Check the [documentation](docs/)
- Review existing issues and discussions

---

**Note**: This project uses a **cloud-agnostic architecture** designed to be modular and extensible. Each Lambda function is independently deployable and can be used across different applications in your ecosystem while avoiding vendor lock-in.
