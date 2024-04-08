package app

import (
	"anggi.tabulation/controllers/posts"
	"anggi.tabulation/controllers/users"
	"github.com/gin-gonic/gin"
)

func MapUrls() {
	router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(200, gin.H{
			"message": "Welcome to the API",
		})
	})

	router.PUT("/v2/bot/channel/webhook/endpoint", func(c *gin.Context) {
		c.IndentedJSON(200, gin.H{
			"message": "Hello, playground",
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
