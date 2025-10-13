package app

import (
	"anggi.tabulation/controllers/logs"
	"anggi.tabulation/controllers/posts"
	"anggi.tabulation/controllers/users"
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

	apiRouter := router.Group("/api")
	{
		apiRouter.POST("/register", users.Register)
		apiRouter.POST("/login", users.Login)
		apiRouter.GET("/me", users.Me)
		apiRouter.GET("/user/:id", users.GetUser)
		apiRouter.GET("/users", users.GetAllUsers)
		apiRouter.GET("/logout", users.Logout)

	// log testing route
	apiRouter.GET("/logs/test", logs.TestLogs)

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
