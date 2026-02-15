#!/bin/bash

# Script to generate Swagger documentation
echo "Generating Swagger documentation..."

# Install swag if not present
if ! command -v swag &> /dev/null; then
    echo "Installing swag..."
    go install github.com/swaggo/swag/cmd/swag@latest
fi

# Generate docs
swag init

if [ $? -eq 0 ]; then
    echo "✓ Swagger documentation generated successfully!"
    echo "  - docs/docs.go"
    echo "  - docs/swagger.json"
    echo "  - docs/swagger.yaml"
else
    echo "✗ Failed to generate Swagger documentation"
    exit 1
fi
