package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MapUrls() {

	// allow all origins
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(200, gin.H{
			"status": "ok",
		})
	})

	// Telegram webhook
	router.POST("/webhook", func(c *gin.Context) {
		var buf bytes.Buffer
		_, err := io.Copy(&buf, c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var update map[string]interface{}
		if err := json.Unmarshal(buf.Bytes(), &update); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Println(update)

		c.JSON(http.StatusOK, gin.H{"status": "ok", "update": update})
	})

	// apiRouter := router.Group("/api")
	// {
	// 	apiRouter.POST("/register", users.Register)
	// 	apiRouter.POST("/login", users.Login)
	// 	apiRouter.GET("/me", users.Me)
	// 	apiRouter.GET("/user/:id", users.GetUser)
	// 	apiRouter.GET("/users", users.GetAllUsers)
	// 	apiRouter.GET("/logout", users.Logout)

	// 	blogRouter := apiRouter.Group("/blog")
	// 	{
	// 		blogRouter.GET("/", posts.GetAll)
	// 		blogRouter.POST("/", posts.Create)
	// 		blogRouter.GET("/:id", posts.Get)
	// 		blogRouter.PUT("/:id", posts.Update)
	// 		blogRouter.DELETE("/:id", posts.Delete)
	// 	}
	// }
}
