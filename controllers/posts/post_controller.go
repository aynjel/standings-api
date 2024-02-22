package posts

import (
	"anggi.tabulation/domain/posts"
	"anggi.tabulation/services"
	"anggi.tabulation/utils/errors"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var post posts.Post

	if err := c.ShouldBindJSON(&post); err != nil {
		getErr := errors.NewBadRequestError(err.Error() + " Invalid JSON")
		c.JSON(getErr.Status, getErr)
		return
	}

	result, err := services.CreatePost(post)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(201, result)
}

func Get(c *gin.Context) {
	var post posts.Post

	if err := c.ShouldBindUri(&post); err != nil {
		getErr := errors.NewBadRequestError(err.Error() + " Invalid URI")
		c.JSON(getErr.Status, getErr)
		return
	}

	result, err := services.GetPost(post.ID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(200, result)
}

func Update(c *gin.Context) {
	var post posts.Post

	if err := c.ShouldBindJSON(&post); err != nil {
		getErr := errors.NewBadRequestError(err.Error() + " Invalid JSON")
		c.JSON(getErr.Status, getErr)
		return
	}

	result, err := services.UpdatePost(post)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(200, result)
}

func Delete(c *gin.Context) {
	var post posts.Post

	if err := c.ShouldBindUri(&post); err != nil {
		getErr := errors.NewBadRequestError(err.Error() + " Invalid URI")
		c.JSON(getErr.Status, getErr)
		return
	}

	err := services.DeletePost(post.ID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(204, nil)
}

func GetAll(c *gin.Context) {
	result, err := services.GetAllPosts()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(200, result)
}
