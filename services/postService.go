package services

import (
	"anggi.tabulation/domain/posts"
	"anggi.tabulation/utils/errors"
)

func CreatePost(post posts.Post) (*posts.Post, *errors.RestErr) {
	if err := post.Save(); err != nil {
		return nil, err
	}

	return &post, nil
}

func GetPost(postID int64) (*posts.Post, *errors.RestErr) {
	post := &posts.Post{ID: postID}
	if err := post.Get(); err != nil {
		return nil, err
	}

	return post, nil
}

func GetAllPosts() ([]posts.Post, *errors.RestErr) {
	post := &posts.Post{}
	posts, err := post.GetAll()
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func UpdatePost(post posts.Post) (*posts.Post, *errors.RestErr) {
	if err := post.Update(); err != nil {
		return nil, err
	}

	return &post, nil
}

func DeletePost(postID int64) *errors.RestErr {
	post := &posts.Post{ID: postID}
	if err := post.Delete(); err != nil {
		return err
	}

	return nil
}
