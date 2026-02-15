package app

import (
	"anggi.tabulation/utils/browser"
	"anggi.tabulation/utils/logs"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	LoadEnv()
}

func StartApplication() {
	logs.Info.Println("Starting the application...")

	router = gin.Default()
	MapUrls()

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "4300"
	}

	// Get host - check if running in Docker
	host := "localhost"
	if os.Getenv("APP_ENV") == "development" {
		// In Docker, we still want to open localhost in the host browser
		host = "localhost"
	}

	docsURL := fmt.Sprintf("http://%s:%s/docs", host, port)
	
	// Open browser after a short delay to ensure server is ready
	logs.Info.Printf("API Documentation will be available at: %s\n", docsURL)
	browser.OpenBrowserAfterDelay(docsURL, 2*time.Second)

	// Start the server
	logs.Info.Printf("Server starting on port %s\n", port)
	router.Run(":" + port)
}
