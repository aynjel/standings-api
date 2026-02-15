# Script to generate Swagger documentation
Write-Host "Generating Swagger documentation..." -ForegroundColor Green

# Check if running in Docker container
if (Test-Path "/.dockerenv") {
    Write-Host "Running in Docker container..." -ForegroundColor Yellow
    # Install swag if not present
    if (-not (Get-Command swag -ErrorAction SilentlyContinue)) {
        Write-Host "Installing swag..." -ForegroundColor Yellow
        go install github.com/swaggo/swag/cmd/swag@latest
    }
}

# Generate docs
swag init

if ($?) {
    Write-Host "✓ Swagger documentation generated successfully!" -ForegroundColor Green
    Write-Host "  - docs/docs.go" -ForegroundColor Cyan
    Write-Host "  - docs/swagger.json" -ForegroundColor Cyan
    Write-Host "  - docs/swagger.yaml" -ForegroundColor Cyan
} else {
    Write-Host "✗ Failed to generate Swagger documentation" -ForegroundColor Red
    exit 1
}
