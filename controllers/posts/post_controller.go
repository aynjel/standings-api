package posts

import (
	"anggi.tabulation/domain/posts"
	"anggi.tabulation/services"
	"anggi.tabulation/utils/errors"
	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary      Create a new post
// @Description  Create a new blog post
// @Tags         blog
// @Accept       json
// @Produce      json
// @Param        post  body      posts.Post  true  "Post data"
// @Success      201   {object}  posts.Post
// @Failure      400   {object}  errors.RestErr
// @Router       /blog/ [post]
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

// Get godoc
// @Summary      Get post by ID
// @Description  Retrieve a specific blog post by ID
// @Tags         blog
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Post ID"
// @Success      200  {object}  posts.Post
// @Failure      400  {object}  errors.RestErr
// @Failure      404  {object}  errors.RestErr
// @Router       /blog/{id} [get]
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

// Update godoc
// @Summary      Update a post
// @Description  Update an existing blog post
// @Tags         blog
// @Accept       json
// @Produce      json
// @Param        id    path      int         true  "Post ID"
// @Param        post  body      posts.Post  true  "Updated post data"
// @Success      200   {object}  posts.Post
// @Failure      400   {object}  errors.RestErr
// @Failure      404   {object}  errors.RestErr
// @Router       /blog/{id} [put]
func Update(c *gin.Context) {
	var post posts.Post

	if err := c.ShouldBindJSON(&post); err != nil {
		getErr := errors.NewBadRequestError(err.Error() + " Invalid JSON")
		c.JSON(getErr.Status, getErr)
// Delete godoc
// @Summary      Delete a post
// @Description  Delete a blog post by ID
// @Tags         blog
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Post ID"
// @Success      204  "No Content"
// @Failure      400  {object}  errors.RestErr
// @Failure      404  {object}  errors.RestErr
// @Router       /blog/{id} [delete]
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
// GetAll godoc
// @Summary      Get all posts
// @Description  Retrieve a list of all blog posts
// @Tags         blog
// @Accept       json
// @Produce      json
// @Success      200  {array}   posts.Post
// @Failure      500  {object}  errors.RestErr
// @Router       /blog/ [get]
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
