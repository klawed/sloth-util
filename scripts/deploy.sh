#!/bin/bash

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
echo "  3. Monitor CloudWatch logs"