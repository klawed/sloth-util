#!/bin/bash

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
cd ../..