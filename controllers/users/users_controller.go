package users

import (
	"net/http"
	"strconv"
	"time"

	"anggi.tabulation/domain/users"
	"anggi.tabulation/services"
	"anggi.tabulation/utils/errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const (
	SecretKey = "thisisasecret"
)

func Register(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		err := errors.NewBadRequestError(err.Error())
		c.IndentedJSON(err.Status, err)
		return
	}

	result, err := services.CreateUser(user)
	if err != nil {
		c.IndentedJSON(err.Status, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, result)
}

func Login(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		err := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(err.Status, err)
		return
	}

	result, getErr := services.GetUserByEmailOrUsername(user)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(result.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		err := errors.NewInternalServerError("Error when trying to generate token")
		c.JSON(err.Status, err)
		return
	}

	c.SetCookie("jwt", token, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, result)
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
	result, err := services.GetUsers()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result)
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
