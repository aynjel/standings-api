# Docker Helper Script for Tabulation API (Windows PowerShell)
# Usage: .\docker-helper.ps1 <command>

param(
    [string]$Command = "help"
)

# Colors
$Colors = @{
    Red = "Red"
    Green = "Green"
    Yellow = "Yellow"
    Blue = "Cyan"
}

function Write-Success {
    param([string]$Message)
    Write-Host "✓ $Message" -ForegroundColor $Colors.Green
}

function Write-Error-Custom {
    param([string]$Message)
    Write-Host "✗ $Message" -ForegroundColor $Colors.Red
}

function Write-Info {
    param([string]$Message)
    Write-Host "→ $Message" -ForegroundColor $Colors.Blue
}

function Write-Warning-Custom {
    param([string]$Message)
    Write-Host "⚠ $Message" -ForegroundColor $Colors.Yellow
}

function Print-Help {
    Write-Host "=== Tabulation API Docker Helper ===" -ForegroundColor $Colors.Blue
    Write-Host ""
    Write-Host "Usage: .\docker-helper.ps1 <command>" -ForegroundColor $Colors.Yellow
    Write-Host ""
    Write-Host "Commands:" -ForegroundColor $Colors.Yellow
    Write-Host "  build           Build Docker image"
    Write-Host "  up              Start containers (development with SQLite)"
    Write-Host "  down            Stop containers"
    Write-Host "  logs            View application logs"
    Write-Host "  logs-postgres   View PostgreSQL logs"
    Write-Host "  shell           Open shell in running container"
    Write-Host "  test            Run quick health check"
    Write-Host "  restart         Restart containers"
    Write-Host "  clean           Remove containers and volumes"
    Write-Host "  postgres        Start with PostgreSQL database"
    Write-Host "  sqlite          Start with SQLite (default)"
    Write-Host "  help            Show this help message"
    Write-Host ""
    Write-Host "Examples:" -ForegroundColor $Colors.Yellow
    Write-Host "  .\docker-helper.ps1 build"
    Write-Host "  .\docker-helper.ps1 up"
    Write-Host "  .\docker-helper.ps1 logs"
    Write-Host "  .\docker-helper.ps1 postgres"
}

function Build-Image {
    Write-Warning-Custom "Building Docker image..."
    docker-compose build
    Write-Success "Build complete!"
}

function Start-Dev {
    Write-Warning-Custom "Starting development containers (SQLite)..."
    docker-compose up -d
    Start-Sleep -Seconds 2
    Write-Success "Containers started!"
    Write-Info "Application URL: http://localhost:4300"
    Write-Host "View logs with: .\docker-helper.ps1 logs"
}

function Start-Postgres {
    Write-Warning-Custom "Starting containers with PostgreSQL..."
    docker-compose --profile postgres up -d
    Start-Sleep -Seconds 3
    Write-Success "Containers started!"
    Write-Info "Application URL: http://localhost:4300"
    Write-Info "PostgreSQL URL: localhost:5432"
    Write-Host "View logs with: .\docker-helper.ps1 logs"
}

function Stop-Containers {
    Write-Warning-Custom "Stopping containers..."
    docker-compose down
    Write-Success "Containers stopped!"
}

function View-Logs {
    Write-Info "=== Application Logs ==="
    docker-compose logs -f app
}

function View-Postgres-Logs {
    Write-Info "=== PostgreSQL Logs ==="
    docker-compose logs -f postgres
}

function Open-Shell {
    Write-Info "Opening shell in app container..."
    docker-compose exec app sh
}

function Health-Check {
    Write-Warning-Custom "Running health check..."
    Start-Sleep -Seconds 1
    try {
        $response = Invoke-WebRequest -Uri "http://localhost:4300/" -ErrorAction Stop
        Write-Success "Application is running!"
        Write-Host "Response: $($response.Content)"
    }
    catch {
        Write-Error-Custom "Application is not responding"
        exit 1
    }
}

function Restart-Containers {
    Write-Warning-Custom "Restarting containers..."
    docker-compose restart
    Start-Sleep -Seconds 2
    Write-Success "Containers restarted!"
}

function Clean-All {
    Write-Host "WARNING: This will remove containers and volumes!" -ForegroundColor $Colors.Red
    $response = Read-Host "Continue? (y/N)"
    if ($response -eq "y" -or $response -eq "Y") {
        Write-Warning-Custom "Cleaning up..."
        docker-compose down -v
        docker-compose --profile postgres down -v
        Write-Success "Cleanup complete!"
    }
    else {
        Write-Host "Cancelled."
    }
}

# Execute command
switch -CaseSensitive -Exact ($Command.ToLower()) {
    "build" {
        Build-Image
    }
    "up" {
        Start-Dev
    }
    "dev" {
        Start-Dev
    }
    "sqlite" {
        Start-Dev
    }
    "postgres" {
        Start-Postgres
    }
    "pg" {
        Start-Postgres
    }
    "down" {
        Stop-Containers
    }
    "stop" {
        Stop-Containers
    }
    "logs" {
        View-Logs
    }
    "logs-postgres" {
        View-Postgres-Logs
    }
    "logs-pg" {
        View-Postgres-Logs
    }
    "shell" {
        Open-Shell
    }
    "bash" {
        Open-Shell
    }
    "sh" {
        Open-Shell
    }
    "test" {
        Health-Check
    }
    "check" {
        Health-Check
    }
    "health" {
        Health-Check
    }
    "restart" {
        Restart-Containers
    }
    "clean" {
        Clean-All
    }
    "help" {
        Print-Help
    }
    "--help" {
        Print-Help
    }
    "-h" {
        Print-Help
    }
    default {
        Write-Error-Custom "Unknown command: $Command"
        Write-Host "Run '.\docker-helper.ps1 help' for usage information"
        exit 1
    }
}
