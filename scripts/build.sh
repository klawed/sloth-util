#!/bin/bash

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

echo "✅ Build completed successfully!"