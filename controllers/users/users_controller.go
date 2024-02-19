package users

import (
	"os"

	"github.com/gin-gonic/gin"
)

var (
	SecretKey = os.Getenv("SECRET_KEY")
)

func Register(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
		"message": "Register",
	})
}

func Login(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
		"message": "Login",
	})
}

func GetUser(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
		"message": "GetUser",
	})
}

func Logout(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
		"message": "Logout",
	})
}

func UpdateUser(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
		"message": "UpdateUser",
	})
}

func DeleteUser(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
		"message": "DeleteUser",
	})
}

func GetUsers(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
		"message": "GetUsers",
	})
}

func GetUserById(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
		"message": "GetUserById",
	})
}

func GetUserByUsername(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
		"message": "GetUserByUsername",
	})
}

func GetUserByEmail(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
		"message": "GetUserByEmail",
	})
}
