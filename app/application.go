package app

import (
	"time"

	"anggi.tabulation/utils/logs"
	"github.com/gin-contrib/cors"
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
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://tabulation-hkzf5t7h3-aynjel.vercel.app/home", "http://localhost:4201", "http://localhost:4200"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Authorization", "Content-Type", "Content-Length", "X-CSRF-Token", "Token"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	router.Use(cors.New(config))
	MapUrls()
	// router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
	router.Run()
}
