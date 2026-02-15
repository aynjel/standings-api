---
title: Tabulation API
description: A RESTful API for event management and blog functionality built with Gin and Go
tags:
  - gin
  - golang
  - rest-api
  - postgresql
  - sqlite
---

# Tabulation API

A comprehensive RESTful API built with [Gin](https://gin-gonic.com/) framework in Go for managing users, authentication, blog posts, and event-related data. This application supports both SQLite (development) and PostgreSQL (production) databases.

## ✨ Features

- **Interactive API Documentation**
  - 🎨 Modern Scalar UI for API documentation
  - 📖 Auto-generates OpenAPI/Swagger specification
  - 🚀 Automatically opens in browser when app starts
  - 🧪 Interactive endpoint testing
  - Access at: `http://localhost:4300/docs`

- **User Authentication & Management**
  - User registration with password hashing (bcrypt)
  - Login/Logout functionality
  - User profile retrieval
  - Get user by ID
  - List all users

- **Blog Management**
  - Create blog posts
  - Read/Retrieve posts
  - Update post content
  - Delete posts
  - Retrieve all posts

- **Database Flexibility**
  - SQLite support for development (in-memory or file-based)
  - PostgreSQL support for production environments
  - Automatic database selection based on environment

- **API Features**
  - CORS enabled for cross-origin requests
  - Structured error handling with custom error responses
  - RESTful endpoint design
  - Environment-based configuration
  - JSON request/response format

- **Architecture**
  - Service layer pattern for business logic
  - Domain-driven design with DTOs and DAOs
  - Centralized controller layer for HTTP handling
  - Logging utilities for debugging

## 📋 Prerequisites

- Go 1.21.6 or higher
- PostgreSQL 12+ (for production)
- Git
- A terminal/command prompt

## 🚀 Quick Start with Docker (Recommended)

The fastest way to get started without installing Go locally is using Docker.

### Prerequisites
- Docker Desktop installed and running

### Quick Start

```bash
# 1. Clone the repository
git clone <repository-url>
cd standings-api

# 2. Create .env file
cp .env.example .env

# 3. Build and start the application
docker-compose up -d

# 4. Verify it's running
curl http://localhost:4300/
```

That's it! The application is now running at `http://localhost:4300`.

### Using the Helper Script (Windows)

```powershell
# View all available commands
.\docker-helper.ps1 help

# Start development environment
.\docker-helper.ps1 up

# View logs
.\docker-helper.ps1 logs

# Stop containers
.\docker-helper.ps1 down
```

### Using the Helper Script (Linux/Mac)

```bash
# Make the script executable
chmod +x docker-helper.sh

# View all available commands
./docker-helper.sh help

# Start development environment
./docker-helper.sh up

# View logs
./docker-helper.sh logs
```

---

## 🚀 Installation & Setup

### 1. Clone the Repository

```bash
git clone <repository-url>
cd standings-api
```

### 2. Install Dependencies (Local Go Installation)

```bash
go mod download
go mod tidy
```

### 3. Configure Environment Variables

Create a `.env` file in the project root directory. You can use the provided `.env.example` as a template:

```bash
cp .env.example .env
```

Edit the `.env` file with your configuration:

```dotenv
# Server Configuration
PORT=4300
SECRET_KEY=your-secret-key-here

# Database Configuration (Development)
APP_ENV=development
SQLITE_IN_MEMORY=false
SQLITE_DB_PATH=./database/standings_dev.db

# Or for Production (PostgreSQL)
# APP_ENV=production
# DATABASE_URL=postgresql://username:password@host:port/database_name
```

### 4. Environment Variables Explained

| Variable | Description | Default | Example |
|----------|-------------|---------|---------|
| `PORT` | Server port number | 8080 | 4300 |
| `SECRET_KEY` | Secret key for application (used for auth/security) | secret | your-secure-key |
| `APP_ENV` | Application environment | development | production |
| `SQLITE_IN_MEMORY` | Use in-memory SQLite (dev only) | false | true/false |
| `SQLITE_DB_PATH` | Path to SQLite database file | ./database/standings_dev.db | custom/path.db |
| `DATABASE_URL` | PostgreSQL connection string (prod) | - | postgresql://user:pass@localhost:5432/dbname |

## 🗄️ Database Setup

### Development Setup (SQLite)

SQLite is the default for development mode. You have two options:

**Option 1: File-based SQLite (recommended for local development)**

```shell
# Update .env
APP_ENV=development
SQLITE_IN_MEMORY=false
SQLITE_DB_PATH=./database/standings_dev.db
```

The database file will be created automatically at the specified path on first run.

**Option 2: In-memory SQLite (for testing)**

```shell
APP_ENV=development
SQLITE_IN_MEMORY=true
```

### Production Setup (PostgreSQL)

For production environments, set up PostgreSQL:

```shell
APP_ENV=production
DATABASE_URL=postgresql://username:password@localhost:5432/dbname
```

**Create the database schema:**

```sql
-- Create users table
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(255) UNIQUE NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create posts table
CREATE TABLE posts (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  content TEXT NOT NULL,
  user_id INTEGER REFERENCES users(id),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create other domain tables as needed (events, judges, participants, etc.)
```

## ▶️ Running the Application

### Using Railway (Recommended for deployment)

```bash
railway link
railway run go run main.go
```

### Local Development

**Option 1: Direct execution**

```bash
go run main.go
```

**Option 2: Using go build**

```bash
go build -o tabulation-api
./tabulation-api
```

The server will start on the port specified in your `.env` file (default: 4300).

You should see output like:
```
[GIN-debug] Loaded HTML Templates (0): 
[GIN-debug] Listening and serving HTTP on :4300
```

### Verify the Server is Running

```bash
curl http://localhost:4300/
```

Expected response:
```json
{
  "status": "ok"
}
```

## 📡 API Endpoints

### Base URL
```
http://localhost:4300/api
```

### Authentication Endpoints

| Method | Endpoint | Description | Body |
|--------|----------|-------------|------|
| `POST` | `/register` | Register a new user | `{ "username": "user", "email": "email@example.com", "password": "pass" }` |
| `POST` | `/login` | Login user | `{ "username": "user", "password": "pass" }` |
| `GET` | `/me` | Get current user profile | - |
| `GET` | `/logout` | Logout current user | - |
| `GET` | `/user/:id` | Get user by ID | - |
| `GET` | `/users` | Get all users | - |

### Blog Endpoints

| Method | Endpoint | Description | Body |
|--------|----------|-------------|------|
| `GET` | `/blog/` | Get all posts | - |
| `POST` | `/blog/` | Create new post | `{ "title": "Post Title", "content": "Post content" }` |
| `GET` | `/blog/:id` | Get post by ID | - |
| `PUT` | `/blog/:id` | Update post | `{ "title": "New Title", "content": "New content" }` |
| `DELETE` | `/blog/:id` | Delete post | - |

### Testing Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/logs/test` | Test logging functionality |

## 🔐 Authentication

The application uses bcrypt for password hashing. When users register or modify passwords:

1. Password is hashed using bcrypt with default cost
2. Hashed password is stored in the database
3. On login, the provided password is compared against the stored hash
4. The `SECRET_KEY` environment variable can be used for additional security layers (tokens, etc.)

## 📦 Project Structure

```
standings-api/
├── app/
│   ├── application.go      # Application initialization and startup
│   ├── environment.go      # Environment variable loading
│   └── router.go           # Route definitions and middleware setup
├── controllers/
│   ├── auth/               # Authentication controllers
│   ├── logs/               # Logging controllers
│   ├── posts/              # Blog post controllers
│   └── users/              # User controllers
├── database/
│   ├── tabulationDB/       # Database connection and initialization
│   └── standings_dev.db    # SQLite development database (auto-created)
├── domain/
│   ├── events/             # Event models (DAO/DTO)
│   ├── judges/             # Judge models (DAO/DTO)
│   ├── organizersAuth/     # Organizer auth models (DAO/DTO)
│   ├── participants/       # Participant models (DAO/DTO)
│   ├── posts/              # Post models (DAO/DTO)
│   └── users/              # User models (DAO/DTO)
├── services/
│   ├── authService.go      # Authentication business logic
│   ├── postService.go      # Post business logic
│   └── userService.go      # User business logic
├── utils/
│   ├── errors/             # Custom error handling
│   └── logs/               # Logging utilities
├── main.go                 # Application entry point
├── go.mod                  # Go module definition
└── .env.example            # Environment variables template
```

### Architecture Pattern

The application follows a **layered architecture**:

1. **Controllers Layer**: Handle HTTP requests/responses
2. **Services Layer**: Contain business logic
3. **Domain Layer**: Contains domain models with Data Access Objects (DAO) and Data Transfer Objects (DTO)
4. **Database Layer**: Manages database connections and queries
5. **Utilities Layer**: Shared functions like error handling and logging

## 🛠️ Development

### Adding New Features

1. **Create Domain Models**: Add DAO/DTO files in `domain/`
2. **Create Service Layer**: Add business logic in `services/`
3. **Create Controller**: Add HTTP handlers in `controllers/`
4. **Register Routes**: Add routes in `app/router.go`

### Example: Adding a new endpoint

```go
// 1. Add route in app/router.go
apiRouter.GET("/example", controllers.ExampleHandler)

// 2. Create controller in controllers/example/
func ExampleHandler(c *gin.Context) {
    result := services.GetExample()
    c.JSON(200, result)
}

// 3. Add service logic in services/
func GetExample() (*ExampleData, *RestErr) {
    // Business logic here
}
```

### Dependencies

- `github.com/gin-gonic/gin` - Web framework
- `github.com/lib/pq` - PostgreSQL driver
- `modernc.org/sqlite` - Pure Go SQLite driver
- `golang.org/x/crypto` - Password hashing (bcrypt)
- `github.com/joho/godotenv` - Environment variable management

## 🐛 Troubleshooting

### Issue: "failed to open sqlite DB"
**Solution**: Ensure the `./database/` directory exists or update `SQLITE_DB_PATH` to a valid path.

### Issue: "DATABASE_URL environment variable not set"
**Solution**: Set the `DATABASE_URL` in your `.env` file for production environment.

### Issue: "failed to ping DB"
**Solution**: 
- For PostgreSQL: Verify database connection string and that PostgreSQL is running
- For SQLite: Check file permissions and disk space

### Issue: Port already in use
**Solution**: Change the `PORT` in `.env` to an available port number.

### Issue: CORS errors in browser
**Solution**: The API has CORS enabled for all origins. Ensure requests include proper headers.

## 📝 Logging

The application includes logging utilities located in `utils/logs/`. Use these for debugging:

Test logging:
```bash
curl http://localhost:4300/api/logs/test
```

## � Running with Docker

Docker allows you to run the application without installing Go locally. This section covers various Docker setup options.

### Docker Setup Files

- `Dockerfile` - Multi-stage build configuration
- `docker-compose.yml` - Orchestration for app and optional PostgreSQL
- `.dockerignore` - Excludes unnecessary files from the Docker build
- `.env.docker` - Example environment configuration for Docker
- `docker-helper.ps1` - PowerShell helper script for Windows
- `docker-helper.sh` - Bash helper script for Linux/Mac

### Basic Docker Commands

#### Build the image

```bash
docker build -t tabulation-api .
```

#### Using docker-compose (Recommended)

```bash
# Start the application (SQLite by default)
docker-compose up -d

# View logs
docker-compose logs -f app

# Stop the application
docker-compose down

# Remove containers and volumes
docker-compose down -v
```

### Development with SQLite (Docker)

This is the default configuration:

```bash
# For production-optimized image
docker-compose up -d

# For development with Go tools (can run `go get` inside container)
docker-compose -f docker-compose.dev.yml up -d
```

**Development image advantages:**
- Includes Go compiler for running `go get` inside the container
- Perfect for adding packages without local Go installation
- Source code mounting for easier debugging

Application will be available at `http://localhost:4300`
Database: `./database/standings_dev.db` (persisted in a Docker volume)

Run a quick test:
```bash
curl http://localhost:4300/
```

View logs in real-time:
```bash
docker-compose logs -f app
# or with dev image
docker-compose -f docker-compose.dev.yml logs -f app
```

Access the container shell:
```bash
docker-compose exec app sh
# or with dev image
docker-compose -f docker-compose.dev.yml exec app sh
```

### Production with PostgreSQL (Docker)

For a more robust production setup with PostgreSQL:

```bash
# Create or update .env file with PostgreSQL configuration
# APP_ENV=production
# DATABASE_URL=postgresql://postgres:postgres@postgres:5432/tabulation_db

# Start with PostgreSQL profile
docker-compose --profile postgres up -d

# Verify both services are running
docker-compose ps

# View application logs
docker-compose logs -f app

# View PostgreSQL logs
docker-compose logs -f postgres

# Access PostgreSQL
docker-compose exec postgres psql -U postgres -d tabulation_db
```

### Helper Script Usage (Windows PowerShell)

For easier Docker management on Windows, use the provided helper script:

```powershell
# Show available commands
.\docker-helper.ps1 help

# Build the image
.\docker-helper.ps1 build

# Start with SQLite (default)
.\docker-helper.ps1 up

# Start with PostgreSQL
.\docker-helper.ps1 postgres

# View application logs
.\docker-helper.ps1 logs

# View PostgreSQL logs (if running)
.\docker-helper.ps1 logs-postgres

# Open shell in container
.\docker-helper.ps1 shell

# Health check
.\docker-helper.ps1 test

# Restart containers
.\docker-helper.ps1 restart

# Stop containers
.\docker-helper.ps1 down

# Clean up everything
.\docker-helper.ps1 clean
```

### Helper Script Usage (Linux/Mac Bash)

```bash
# Make executable
chmod +x docker-helper.sh

# Show available commands
./docker-helper.sh help

# Build the image
./docker-helper.sh build

# Start with SQLite (default)
./docker-helper.sh up

# Start with PostgreSQL
./docker-helper.sh postgres

# View logs
./docker-helper.sh logs

# And other commands are the same as PowerShell version...
```

### Docker Environment Configuration

Docker uses environment variables from `.env` file. Key variables for Docker:

```dotenv
# Port mapping
PORT=4300

# Security
SECRET_KEY=your-secret-key-change-in-production

# Database Mode - Development (SQLite)
APP_ENV=development
SQLITE_IN_MEMORY=false
SQLITE_DB_PATH=./database/standings_dev.db

# Database Mode - Production (PostgreSQL)
# APP_ENV=production
# DATABASE_URL=postgresql://postgres:postgres@postgres:5432/tabulation_db

# PostgreSQL specific (if using postgres service)
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=tabulation_db
DB_PORT=5432
```

### Docker Volumes & Persistence

The `docker-compose.yml` defines two volumes:

1. **db_data** - Persists SQLite database across container restarts
2. **postgres_data** - Persists PostgreSQL data across container restarts

To view volumes:
```bash
docker volume ls
docker volume inspect standings-api_db_data
```

### Docker Networking

The compose file creates a custom bridge network (`tabulation-network`) that allows:
- The `app` service to communicate with `postgres` service using hostname `postgres`
- Isolation from other containers

### Debugging Docker Containers

```bash
# View running containers
docker-compose ps

# View detailed logs
docker-compose logs --tail=100 app

# Follow logs in real-time
docker-compose logs -f app

# Execute commands in running container
docker-compose exec app ls -la

# Get container statistics
docker stats

# Inspect container details
docker inspect <container_id>
```

### Building for Different Architectures

The current Dockerfile uses `CGO_ENABLED=0` for a static binary compatible across Linux architectures. To build for specific platforms:

```bash
# Build for Apple Silicon (arm64)
docker buildx build --platform linux/arm64 -t tabulation-api:arm64 .

# Build for x86_64
docker buildx build --platform linux/amd64 -t tabulation-api:amd64 .

# Build for both (requires buildx setup)
docker buildx build --platform linux/amd64,linux/arm64 -t tabulation-api:latest .
```

### Docker Troubleshooting

**Issue: Port already in use**
```bash
# Change port in .env
PORT=4301

# Restart containers
docker-compose restart
```

**Issue: Database connection refused**
```bash
# Ensure PostgreSQL profile is running when needed
docker-compose --profile postgres up -d

# Check if postgres service is healthy
docker-compose ps
```

**Issue: Volume permission denied**
```bash
# On Linux, the container runs as root, so volume permissions are usually fine
# On Mac/Windows with Docker Desktop, volumes are managed by Docker, no action needed
```

**Issue: Out of disk space**
```bash
# Clean up unused Docker resources
docker system prune -a --volumes
```

**Issue: Container won't start**
```bash
# Check logs for errors
docker-compose logs app

# Rebuild the image
docker-compose build --no-cache
docker-compose up -d
```

## �🚀 Deployment

### Deploying to Railway

```bash
# Link to Railway project
railway link

# Run the application
railway run go run main.go

# View logs
railway logs
```

### Docker Deployment (Other Platforms)

The Docker image can be deployed to:
- **Kubernetes** - Use the image in a deployment manifest
- **AWS ECS** - Push to ECR and create ECS task definition
- **Google Cloud Run** - Push to GCR and deploy
- **Docker Hub/Registries** - Push the image for distribution

Example pushing to Docker Hub:
```bash
docker build -t your-username/tabulation-api .
docker push your-username/tabulation-api
```

## 📄 License

This project is part of the Anggi Tabulation system.

## 🤝 Contributing

When contributing to this project:

1. Follow the existing code structure and patterns
2. Use the service layer for business logic
3. Add proper error handling
4. Update the README if adding new features
5. Test endpoints before submitting changes

## 📞 Support

For issues or questions:
- Check the Troubleshooting section above
- Review the `.env.example` for configuration options
- Check application logs for detailed error messages
