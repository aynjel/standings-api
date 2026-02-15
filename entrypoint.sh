#!/bin/sh
set -e

echo "🚀 Starting Tabulation API..."

# Check if swagger docs exist, if not generate them
if [ ! -f "/app/docs/swagger.json" ]; then
    echo "📚 Generating API documentation..."
    swag init
    echo "✓ API documentation generated"
else
    echo "✓ API documentation already exists"
fi

# Start the application
echo "🎯 Starting server..."
exec "$@"
