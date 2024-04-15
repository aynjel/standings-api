package app

import (
	"bytes"
	"encoding/json"
	"io"

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

	// Line event webhook
	router.POST("/webhook", func(c *gin.Context) {
		body, _ := io.ReadAll(c.Request.Body)
		println(string(body))

		c.Request.Body = io.NopCloser(bytes.NewReader(body))
		c.IndentedJSON(200, gin.H{
			"status":   "ok",
			"response": json.RawMessage(body),
		})
	})

	// Line login callback url
	router.GET("/callback", func(c *gin.Context) {
		println(c.Request.URL.Query().Encode())
		c.IndentedJSON(200, gin.H{
			"status": "ok",
			"params": c.Request.URL.Query(),
		})
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
