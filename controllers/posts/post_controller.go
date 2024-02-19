package posts

import "github.com/gin-gonic/gin"

func Create(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
		"message": "Create",
	})
}

func Get(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
		"message": "Get",
	})
}

func Update(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
		"message": "Update",
	})
}

func Delete(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
		"message": "Delete",
	})
}

func GetAll(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
		"message": "GetPosts",
	})
}
