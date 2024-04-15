package app

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"

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
		const LINE_API = "https://api.line.me/oauth2/v2.1/"
		const redirect_uri = "https://chat.line.biz/"

		code, _ := c.GetQuery("code")
		state, _ := c.GetQuery("state")
		println("code: ", code)
		println("state: ", state)

		// Get access token
		url := LINE_API + "token"
		payload := strings.NewReader("grant_type=authorization_code&code=" + code + "&redirect_uri=" + redirect_uri + "&client_id=" + "1656327446" + "&client_secret=" + "d782240c1a7ecb6ab9950be288e5068a")
		req, _ := http.NewRequest("POST", url, payload)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			println(err.Error())
		}
		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)
		println(string(body))

		// Get user profile
		url = LINE_API + "profile"
		req, _ = http.NewRequest("GET", url, nil)
		req.Header.Add("Authorization", "Bearer "+string(body))

		res, err2 := http.DefaultClient.Do(req)
		if err2 != nil {
			println(err2.Error())
		}
		defer res.Body.Close()
		body, _ = io.ReadAll(res.Body)
		println(string(body))

		c.IndentedJSON(200, gin.H{
			"status":   "ok",
			"response": json.RawMessage(body),
		})

		// Issue access token
		// data := map[string]string{
		// 	"client_id":     "1656327446",
		// 	"client_secret": "d782240c1a7ecb6ab9950be288e5068a",
		// 	"code":          code,
		// 	"grant_type":    "client_credentials",
		// 	"redirect_uri":  "https://chat.line.biz/",
		// }
		// payload, _ := json.Marshal(data)
		// resp, err := http.Post(LINE_API+"token", "application/x-www-form-urlencoded", bytes.NewBuffer(payload))
		// if err != nil {
		// 	println(err.Error())
		// }
		// defer resp.Body.Close()

		// body, _ := io.ReadAll(resp.Body)
		// println(string(body))

		// c.IndentedJSON(200, gin.H{
		// 	"status":   "ok",
		// 	"response": json.RawMessage(body),
		// })
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
