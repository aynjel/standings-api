package posts

import (
	"anggi.tabulation/database/postgreSQL/tabulationDB"
	"anggi.tabulation/utils/errors"
)

var (
	createTablePost = `
		CREATE TABLE IF NOT EXISTS posts (
			id SERIAL PRIMARY KEY,
			title VARCHAR(100) NOT NULL,
			content TEXT NOT NULL,
			author_id INT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (author_id) REFERENCES users(id)
		);
	`
	// dropTablePost    = `DROP TABLE IF EXISTS posts;`
	queryInsertPost  = `INSERT INTO posts (title, content, author_id) VALUES ($1, $2, $3) RETURNING id`
	queryGetPostByID = `SELECT id, title, content, author_id, created_at FROM posts WHERE id = $1`
	queryGetAllPosts = `SELECT id, title, content, author_id, created_at FROM posts`
	queryUpdatePost  = `UPDATE posts SET title = $1, content = $2 WHERE id = $3`
	queryDeletePost  = `DELETE FROM posts WHERE id = $1`
)

func init() {
	// tabulationDB.Client.Exec(dropTablePost)
	if _, err := tabulationDB.Client.Exec(createTablePost); err != nil {
		panic(err)
	}
}

func (post *Post) Save() *errors.RestErr {
	stmt, err := tabulationDB.Client.Prepare(queryInsertPost)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	var id int
	if err := stmt.QueryRow(post.Title, post.Content, post.AuthorID).Scan(&id); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func (post *Post) Get() *errors.RestErr {
	stmt, err := tabulationDB.Client.Prepare(queryGetPostByID)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(post.ID)
	if getErr := result.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.CreatedAt); getErr != nil {
		return errors.NewInternalServerError(getErr.Error())
	}

	return nil
}

func (post *Post) GetAll() ([]Post, *errors.RestErr) {
	stmt, err := tabulationDB.Client.Prepare(queryGetAllPosts)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.CreatedAt); err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (post *Post) Update() *errors.RestErr {
	stmt, err := tabulationDB.Client.Prepare(queryUpdatePost)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err := stmt.Exec(post.Title, post.Content, post.ID); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func (post *Post) Delete() *errors.RestErr {
	stmt, err := tabulationDB.Client.Prepare(queryDeletePost)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err := stmt.Exec(post.ID); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}
