package posts

import (
	"strings"

	"anggi.tabulation/utils/errors"
)

type Post struct {
	ID        int64  `json:"ID"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	AuthorID  int64  `json:"author_id"`
	CreatedAt string `json:"created_at"`
}

func (post *Post) Validate() *errors.RestErr {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)

	if post.Title == "" {
		return errors.NewBadRequestError("invalid title")
	}
	if post.Content == "" {
		return errors.NewBadRequestError("invalid content")
	}
	return nil
}
