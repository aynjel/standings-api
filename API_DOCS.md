# API Documentation with Scalar

This application now includes automatic API documentation using **Scalar** - a modern, interactive API documentation interface.

## 🚀 Features

- **Scalar UI**: Beautiful, modern API documentation interface
- **Auto-Open Browser**: Automatically opens the API docs in your default browser when the app starts
- **OpenAPI/Swagger Spec**: Full OpenAPI 3.0 specification
- **Interactive Testing**: Test API endpoints directly from the documentation

## 📖 Accessing the Documentation

When you start the application, it will automatically:
1. Start the API server on port 4300 (or your configured PORT)
2. Open your default browser with the Scalar API documentation
3. Display the documentation at `http://localhost:4300/docs`

### Manual Access

You can also manually access the documentation:
- **Scalar UI**: http://localhost:4300/docs
- **Swagger JSON**: http://localhost:4300/swagger/doc.json
- **Swagger UI (Raw)**: http://localhost:4300/swagger/index.html

## 🔧 Generating Documentation

The API documentation is automatically generated from code annotations.

### Automatic Generation (Docker)

When using Docker, the documentation is automatically generated during the build process.

### Manual Generation

If you need to regenerate the documentation locally:

**Windows (PowerShell):**
```powershell
.\generate-docs.ps1
```

**Linux/Mac:**
```bash
./generate-docs.sh
```

**Using Go directly:**
```bash
swag init
```

## 📝 Adding Annotations

To document your API endpoints, add Swagger annotations to your controllers:

```go
// GetUser godoc
// @Summary      Get user by ID
// @Description  Retrieve a specific user by their ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  users.User
// @Failure      404  {object}  errors.RestErr
// @Router       /user/{id} [get]
func GetUser(c *gin.Context) {
    // ... handler implementation
}
```

## 🐳 Docker Usage

The application is Docker-ready with documentation generation:

```bash
# Development (with live reload)
docker-compose -f docker-compose.dev.yml up --build

# Production
docker-compose up --build
```

## 🎨 Customizing Scalar Theme

You can customize the Scalar UI theme by modifying the `app/scalar.go` file. Available themes:
- `default`
- `purple` (current)
- `bluePlanet`
- `deepSpace`
- `alternate`
- `moon`
- `mars`
- `kepler`
- `saturn`

## 📚 Documentation Structure

- **Tags**: Endpoints are grouped by tags (users, blog, etc.)
- **Security**: JWT Bearer authentication is documented
- **Models**: Request/response models are automatically generated from Go structs

## 🔗 References

- [Scalar Documentation](https://github.com/scalar/scalar)
- [Swaggo Documentation](https://github.com/swaggo/swag)
- [OpenAPI Specification](https://swagger.io/specification/)
