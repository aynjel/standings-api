package users

import (
	"anggi.tabulation/database/postgreSQL/tabulationDB"
	"anggi.tabulation/utils/errors"
)

var (
	createTableUser = `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		firstname VARCHAR(50) NOT NULL,
		lastname VARCHAR(50) NOT NULL,
		username VARCHAR(50) NOT NULL,
		email VARCHAR(50) NOT NULL,
		password VARCHAR(50) NOT NULL
	);
	`
	queryInsertUser          = `INSERT INTO users (firstname, lastname, username, email, password) VALUES (?, ?, ?, ?, ?)`
	queryGetUserByID         = "SELECT id, username, email, password FROM users WHERE id = ?"
	queryGetUsernameAndEmail = "SELECT id, username, email, password FROM users WHERE username = ? OR email = ?"
)

func init() {
	tabulationDB.Client.Exec(createTableUser)
}

func (user *User) Save() *errors.RestErr {
	stmt, err := tabulationDB.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError("error when trying to save user")
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Username, user.Email, user.Password)
	if saveErr != nil {
		return errors.NewInternalServerError("error when trying to save user")
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError("error when trying to save user")
	}
	user.ID = userId
	return nil
}

func (user *User) Get() *errors.RestErr {
	stmt, err := tabulationDB.Client.Prepare(queryGetUserByID)
	if err != nil {
		return errors.NewInternalServerError("error when trying to get user")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.Username, &user.Email, &user.Password); getErr != nil {
		return errors.NewInternalServerError("error when trying to get user")
	}
	return nil
}

func (user *User) FindByEmailOrUsername() *errors.RestErr {
	stmt, err := tabulationDB.Client.Prepare(queryGetUsernameAndEmail)
	if err != nil {
		return errors.NewInternalServerError("error when trying to find user")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Username, user.Email)
	if getErr := result.Scan(&user.ID, &user.Username, &user.Email, &user.Password); getErr != nil {
		return errors.NewInternalServerError("error when trying to find user")
	}
	return nil
}
