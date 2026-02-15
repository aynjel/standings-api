# Docker Quick Start Guide

This guide will help you run the Tabulation API using Docker without needing Go installed on your machine.

## Prerequisites

- Docker Desktop installed and running
- Port 4300 available (or change in `.env`)
- Terminal/PowerShell access

## � Ultra-Quick Start (2 minutes)

### Windows (PowerShell)

```powershell
# 1. Navigate to project directory
cd standings-api

# 2. Create .env file
Copy-Item .env.example .env

# 3. Start the application
docker-compose up -d

# 4. Test it
curl http://localhost:4300/
```

### Linux/Mac (Bash)

```bash
# 1. Navigate to project directory
cd standings-api

# 2. Create .env file
cp .env.example .env

# 3. Start the application
docker-compose up -d

# 4. Test it
curl http://localhost:4300/
```

## 📌 Choosing the Right Docker Setup

This project includes TWO Docker configurations:

### 1. **Production Setup** (Default)
**File:** `docker-compose.yml` + `Dockerfile`

Use when:
- Just running the application
- You don't need to add packages in the running container
- Deploying to production

Features:
- ✅ Optimized small image size
- ✅ Fast startup
- ✅ Best for deployment
- ❌ No `go get` inside container

Start with:
```powershell
docker-compose up -d
```

### 2. **Development Setup** (For Active Development)
**File:** `docker-compose.dev.yml` + `Dockerfile.dev`

Use when:
- Adding packages with `go get`
- Debugging inside container
- Need Go compiler tools
- Active development

Features:
- ✅ Includes Go compiler
- ✅ Can run `go get` inside container
- ✅ Hot reload with source mounting
- ❌ Larger image size
- ❌ Not suitable for production

Start with:
```powershell
docker-compose -f docker-compose.dev.yml up -d
```

## 🚀 Ultra-Quick Start (2 minutes)

## Common Docker Commands

### Start and Stop

```bash
# Start containers in background
docker-compose up -d

# Stop containers
docker-compose down

# View running containers
docker-compose ps

# Restart containers
docker-compose restart
```

### Logs and Debugging

```bash
# View application logs
docker-compose logs -f app

# View last 50 lines
docker-compose logs --tail=50 app

# View PostgreSQL logs (if running)
docker-compose logs -f postgres

# Stop following logs (press Ctrl+C)
```

### Database Management

```bash
# Access SQLite database shell
# (Database file location: ./database/standings_dev.db)

# Access PostgreSQL shell
docker-compose exec postgres psql -U postgres -d tabulation_db

# List all tables
\dt

# Exit PostgreSQL
\q
```

### Container Shell Access

```bash
# Open shell in app container (production image - no Go)
docker-compose exec app sh

# Open shell in dev container (development image - has Go)
docker-compose -f docker-compose.dev.yml exec app sh

# Run commands in container
docker-compose exec app ls -la

# Exit shell
exit
```

## Development Workflow

### Using SQLite (Default)

Perfect for local development:

```bash
# For production-optimized image (standard setup)
docker-compose up -d

# For development with Go tools (recommended for development)
docker-compose -f docker-compose.dev.yml up -d
```

### Adding New Go Packages - FIXED SOLUTION

I've created `Dockerfile.dev` and `docker-compose.dev.yml` specifically for development. This includes Go compiler tools!

#### ✅ Recommended: Use Development Docker Compose

```powershell
# 1. Start with the development compose file (has Go installed)
docker-compose -f docker-compose.dev.yml up -d

# 2. Verify it's running
docker-compose -f docker-compose.dev.yml ps

# 3. Open shell in the container (Go IS available now)
docker-compose -f docker-compose.dev.yml exec app sh

# 4. Inside container, add the package
go get github.com/example/example-package

# 5. Exit shell
exit

# 6. Restart to apply changes
docker-compose -f docker-compose.dev.yml restart

# 7. View logs
docker-compose -f docker-compose.dev.yml logs -f app
```

#### What's Different?

**Standard docker-compose.yml** (production):
- Optimized final image (small, no Go compiler)
- Great for deployment
- Can't run `go get` inside container
- Run code as compiled binary

**docker-compose.dev.yml** (development):
- Includes Go compiler and tools
- Perfect for development
- Can run `go get` inside container
- Runs code with `go run main.go` for live changes

### Method 1: Quick Local Update (Still Recommended if Go Installed)

```powershell
# 1. Add package locally
go get github.com/example/example-package

# 2. Rebuild Docker image
docker-compose build

# 3. Restart containers
docker-compose up -d
```

### Method 2: Using Development Compose (Recommended if No Local Go)

```powershell
# 1. Start development environment
docker-compose -f docker-compose.dev.yml up -d

# 2. Shell into container
docker-compose -f docker-compose.dev.yml exec app sh

# 3. Add package inside container
go get github.com/example/example-package

# 4. Exit and restart
exit
docker-compose -f docker-compose.dev.yml restart

# 5. Verify
docker-compose -f docker-compose.dev.yml logs -f app
```

### Method 3: Update go.mod Manually

If the above doesn't work:

```powershell
# 1. Open go.mod file in your editor
notepad go.mod
# or use VSCode
code go.mod

# 2. Add the import line manually in the require section:
# require (
#     github.com/sirupsen/logrus v1.9.3
# )

# 3. Rebuild Docker image
docker-compose -f docker-compose.dev.yml build

# 4. Restart
docker-compose -f docker-compose.dev.yml up -d
```

### Example: Adding a Popular Package

Let's say you want to add the `gorilla/mux` router:

#### Step 1: Add the Package

```powershell
# Option A: If you have Go installed
go get github.com/gorilla/mux

# Option B: Inside Docker container
docker-compose exec app sh
go get github.com/gorilla/mux
exit

# Option C: Edit go.mod manually
# Add "github.com/gorilla/mux v1.8.0" to require section
```

#### Step 2: Rebuild Docker Image

```powershell
docker-compose build
```

Output will show:
```
Building app
Step 1/12 : FROM golang:1.21.6-alpine AS builder
...
Step 3/12 : RUN go mod download
...
Successfully built [hash]
```

#### Step 3: Restart Container

```powershell
docker-compose restart
# or
docker-compose up -d
```

#### Step 4: Use the Package in Your Code

```go
package main

import (
    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()
    // Use router...
}
```

### Viewing Current Dependencies

```powershell
# Inside container, view all dependencies
docker-compose exec app go list -m all

# View specific package
docker-compose exec app go list -m github.com/example/package

# View go.mod
docker-compose exec app cat go.mod

# View go.sum
docker-compose exec app cat go.sum
```

### Cleaning Dependencies

```powershell
# Remove unused dependencies
# Option 1: Locally (if Go installed)
go mod tidy

# Option 2: Inside container
docker-compose exec app go mod tidy

# Rebuild and restart
docker-compose build
docker-compose restart
```

### Common Packages to Add

```bash
# Database
go get github.com/jmoiron/sqlx          # SQL utilities
go get gorm.io/gorm                      # ORM
go get github.com/jmoiron/sqlc           # Type-safe SQL

# HTTP/Routing
go get github.com/gorilla/mux            # Router

# Validation
go get github.com/go-playground/validator # Struct validation

# JSON/Serialization
go get google.golang.org/protobuf        # Protocol buffers

# Logging
go get github.com/example/example-package        # Structured logging
go get go.uber.org/zap                   # Performant logging

# Testing
go get github.com/stretchr/testify       # Testing toolkit

# Configuration
go get github.com/spf13/viper            # Configuration management

# JWT/Auth
go get github.com/golang-jwt/jwt         # JWT tokens
```

### Troubleshooting Package Issues

**Issue: Package not found after rebuild**
```powershell
# Clear Docker build cache and rebuild
docker-compose build --no-cache

# Start fresh
docker-compose up -d
```

**Issue: Version conflicts**
```powershell
# Update go.mod to specific version
go get github.com/example/package@v1.2.3

# Or inside container
docker-compose exec app go get github.com/example/package@v1.2.3

# Rebuild
docker-compose build
docker-compose restart
```

**Issue: go.mod not updating**
```powershell
# Check current go.mod
cat go.mod

# Verify package was added
go list -m all | findstr package-name

# If issues persist:
# 1. Delete go.sum
# 2. Run go mod tidy
# 3. Rebuild Docker image

docker-compose exec app sh
rm go.sum
go mod tidy
exit
docker-compose build
docker-compose up -d
```

### Testing API Endpoints

```bash
# Register a user
curl -X POST http://localhost:4300/api/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "test@example.com",
    "password": "securepass123"
  }'

# Login
curl -X POST http://localhost:4300/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "securepass123"
  }'

# Get all users
curl http://localhost:4300/api/users

# Create a blog post
curl -X POST http://localhost:4300/api/blog \
  -H "Content-Type: application/json" \
  -d '{
    "title": "My First Post",
    "content": "This is my first blog post"
  }'

# Get all posts
curl http://localhost:4300/api/blog

# Get specific post
curl http://localhost:4300/api/blog/1

# Update post
curl -X PUT http://localhost:4300/api/blog/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Updated Title",
    "content": "Updated content"
  }'

# Delete post
curl -X DELETE http://localhost:4300/api/blog/1
```

## Using Helper Scripts

### Windows PowerShell

The `docker-helper.ps1` script provides convenient shortcuts:

```powershell
# View available commands
.\docker-helper.ps1 help

# Build Docker image
.\docker-helper.ps1 build

# Start with SQLite (production-optimized)
.\docker-helper.ps1 up
# or
.\docker-helper.ps1 sqlite

# Start with development image (includes Go tools)
.\docker-helper.ps1 dev

# Start with PostgreSQL
.\docker-helper.ps1 postgres

# View logs
.\docker-helper.ps1 logs

# Open container shell
.\docker-helper.ps1 shell

# Health check
.\docker-helper.ps1 test

# Restart
.\docker-helper.ps1 restart

# Stop
.\docker-helper.ps1 down

# Clean up everything
.\docker-helper.ps1 clean
```

### Linux/Mac Bash

The `docker-helper.sh` script provides similar functionality:

```bash
# Make executable first
chmod +x docker-helper.sh

# View commands
./docker-helper.sh help

# Start with SQLite
./docker-helper.sh up

# Start with PostgreSQL
./docker-helper.sh postgres

# View logs
./docker-helper.sh logs

# Other commands...
./docker-helper.sh shell
./docker-helper.sh test
./docker-helper.sh restart
./docker-helper.sh down
./docker-helper.sh clean
```

## Environment Configuration

### Development (SQLite)

The `.env` file controls configuration:

```dotenv
PORT=4300
SECRET_KEY=your-secret-key
APP_ENV=development
SQLITE_IN_MEMORY=false
SQLITE_DB_PATH=./database/standings_dev.db
```

### Production (PostgreSQL)

To use PostgreSQL instead:

```dotenv
PORT=4300
SECRET_KEY=production-secret-key-change-this
APP_ENV=production
DATABASE_URL=postgresql://postgres:postgres@postgres:5432/tabulation_db
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=tabulation_db
DB_PORT=5432
```

Then start with:
```bash
docker-compose --profile postgres up -d
```

## Docker Compose Services

The `docker-compose.yml` defines:

1. **app** - The Go application
   - Runs on port 4300
   - Auto-restarts on failure
   - Uses SQLite by default

2. **postgres** (optional)
   - PostgreSQL database
   - Port 5432
   - Only starts with `--profile postgres`
   - Useful for production testing

## Data Persistence

### SQLite Data

Stored in Docker volume `standings-api_db_data`:
- File-based: `./database/standings_dev.db`
- Persists between container restarts
- Shared with host filesystem

### PostgreSQL Data

Stored in Docker volume `standings-api_postgres_data`:
- Database files persisted
- Survives container restarts

View volumes:
```bash
docker volume ls
docker volume inspect standings-api_db_data
```

## Troubleshooting

### **"go: not found" Error in Container Shell**

This happens when you try `go get` in the production container, which doesn't include the Go compiler.

**Solution: Use the development Docker setup**

```powershell
# 1. Stop current containers
docker-compose down

# 2. Start with development Compose file (includes Go tools)
docker-compose -f docker-compose.dev.yml up -d

# 3. Now open shell in container
docker-compose -f docker-compose.dev.yml exec app sh

# 4. Inside container, Go is available
go get github.com/example/example-package
go list -m all

# 5. Exit and restart
exit
docker-compose -f docker-compose.dev.yml restart

# 6. View logs
docker-compose -f docker-compose.dev.yml logs -f app
```

**Why this happens:**
- `docker-compose.yml` uses `Dockerfile` (production-optimized, no Go compiler)
- `docker-compose.dev.yml` uses `Dockerfile.dev` (includes Go for development)

**Quick comparison:**
| Image | Size | Go Tool | Use Case |
|-------|------|---------|----------|
| Dockerfile (production) | ~50MB | ❌ No | Deployment |
| Dockerfile.dev | ~500MB | ✅ Yes | Development |

### Port Already in Use

```bash
# Change port in .env
PORT=4301

# Restart
docker-compose restart
```

### Database Connection Error

```bash
# Check services are running
docker-compose ps

# Verify PostgreSQL is set to start
docker-compose --profile postgres up -d

# Check logs
docker-compose logs app
```

### Container Won't Start

```bash
# Check logs for error details
docker-compose logs app

# Rebuild the image
docker-compose build --no-cache

# Restart
docker-compose up -d
```

### Need to Reset Everything

```bash
# Stop and remove everything
docker-compose --profile postgres down -v

# Remove all volumes
docker volume rm standings-api_db_data standings-api_postgres_data

# Rebuild and start fresh
docker-compose build
docker-compose up -d
```

### View Resource Usage

```bash
# See CPU, memory, network stats
docker stats

# For specific container
docker stats tabulation-api
```

## Performance Tips

1. **Use SQLite for development** - Faster startup, no database setup
2. **Use PostgreSQL for production testing** - More realistic
3. **Mount source code** - Auto-hot reload during development (see docker-compose.yml)
4. **Build once, run many times** - `docker-compose build` is fast after changes

## Next Steps

1. Read the main [README.md](README.md) for full API documentation
2. Check [API Endpoints](README.md#-api-endpoints) section
3. Explore the [Project Structure](README.md#-project-structure)
4. Review [Deployment](README.md#-deployment) options

## Need Help?

- View logs: `docker-compose logs -f app`
- Test endpoint: `curl http://localhost:4300/`
- Check Docker: `docker-compose ps`
- Read main README: See README.md

## File Reference

- `Dockerfile` - Production container image (optimized, minimal)
- `Dockerfile.dev` - Development container image (includes Go compiler)
- `docker-compose.yml` - Production orchestration (use by default)
- `docker-compose.dev.yml` - Development orchestration (includes Go tools)
- `.env` - Environment configuration
- `.env.example` - Configuration template
- `.env.docker` - Docker-specific example
- `.dockerignore` - Files excluded from build
- `docker-helper.ps1` - Windows PowerShell helper
- `docker-helper.sh` - Linux/Mac Bash helper
