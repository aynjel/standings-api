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
			email VARCHAR(100) NOT NULL,
			password TEXT NOT NULL
		);
	`
	dropTableUser            = `DROP TABLE IF EXISTS users;`
	queryInsertUser          = `INSERT INTO users (firstname, lastname, username, email, password) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	queryGetUserByID         = `SELECT id, username, email, password FROM users WHERE id = $1`
	queryGetUsernameAndEmail = `SELECT id, username, email, password FROM users WHERE username = $1 OR email = $2`
	queryGetAllUsers         = `SELECT id, username, email, password FROM users`
)

func init() {
	// tabulationDB.Client.Exec(dropTableUser)
	tabulationDB.Client.Exec(createTableUser)
}

func (user *User) Save() *errors.RestErr {
	stmt, err := tabulationDB.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	var id int
	if err := stmt.QueryRow(user.FirstName, user.LastName, user.UserName, user.Email, user.Password).Scan(&id); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func (user *User) Get() *errors.RestErr {
	stmt, err := tabulationDB.Client.Prepare(queryGetUserByID)
	if err != nil {
		return errors.NewInternalServerError("error when trying to get user")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.UserName, &user.Email, &user.Password); getErr != nil {
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

	result := stmt.QueryRow(user.UserName, user.Email)
	if getErr := result.Scan(&user.ID, &user.UserName, &user.Email, &user.Password); getErr != nil {
		return errors.NewInternalServerError("error when trying to find user")
	}
	return nil
}

func (user *User) GetAll() ([]User, *errors.RestErr) {
	stmt, err := tabulationDB.Client.Prepare(queryGetAllUsers)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.UserName, &user.Email, &user.Password); err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		users = append(users, user)
	}
	if len(users) == 0 {
		return nil, errors.NewInternalServerError(err.Error())
	}
	return users, nil
}
