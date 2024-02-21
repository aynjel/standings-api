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
		err := errors.NewBadRequestError(err.Error())
		c.JSON(err.Status, err)
		return
	}

	result, getErr := services.GetUserByUsernameAndEmail(user)
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
		err := errors.NewInternalServerError(err.Error())
		c.JSON(err.Status, err)
		return
	}

	c.SetCookie("jwt", token, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, result)
}

func GetUser(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		err := errors.NewBadRequestError(err.Error())
		c.JSON(err.Status, err)
		return
	}

	result, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetAllUsers(c *gin.Context) {
	result, getErr := services.GetAllUsers()
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

func UpdateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		err := errors.NewBadRequestError(err.Error())
		c.JSON(err.Status, err)
		return
	}

	result, getErr := services.UpdateUser(user)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

func DeleteUser(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		err := errors.NewBadRequestError(err.Error())
		c.JSON(err.Status, err)
		return
	}

	getErr := services.DeleteUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func Logout(c *gin.Context) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		err := errors.NewBadRequestError(err.Error())
		c.JSON(err.Status, err)
		return
	}
	c.SetCookie("jwt", cookie, -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

func Me(c *gin.Context) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		err := errors.NewBadRequestError(err.Error())
		c.JSON(err.Status, err)
		return
	}

	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		err := errors.NewBadRequestError(err.Error())
		c.JSON(err.Status, err)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err := errors.NewBadRequestError(err.Error())
		c.JSON(err.Status, err)
		return
	}

	userID, err := strconv.ParseInt(claims["iss"].(string), 10, 64)
	if err != nil {
		err := errors.NewBadRequestError(err.Error())
		c.JSON(err.Status, err)
		return
	}

	result, getErr := services.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, result)
}
