package app

import (
	"anggi.tabulation/utils/logs"
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
	router.Run()
}
