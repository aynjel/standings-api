package app

import (
	"bytes"
	"encoding/json"
	"io"

	"anggi.tabulation/controllers/posts"
	"anggi.tabulation/controllers/users"
	"github.com/gin-gonic/gin"
)

func MapUrls() {
	router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(200, gin.H{
			"status": "ok",
		})
	})

	router.POST("/webhook", func(c *gin.Context) {
		body, _ := io.ReadAll(c.Request.Body)
		println(string(body))

		c.Request.Body = io.NopCloser(bytes.NewReader(body))
		c.IndentedJSON(200, gin.H{
			"status":   "ok",
			"response": json.RawMessage(body),
		})
	})

	apiRouter := router.Group("/api")
	{
		apiRouter.POST("/register", users.Register)
		apiRouter.POST("/login", users.Login)
		apiRouter.GET("/me", users.Me)
		apiRouter.GET("/user/:id", users.GetUser)
		apiRouter.GET("/users", users.GetAllUsers)
		apiRouter.GET("/logout", users.Logout)

		blogRouter := apiRouter.Group("/blog")
		{
			blogRouter.GET("/", posts.GetAll)
			blogRouter.POST("/", posts.Create)
			blogRouter.GET("/:id", posts.Get)
			blogRouter.PUT("/:id", posts.Update)
			blogRouter.DELETE("/:id", posts.Delete)
		}
	}
}
