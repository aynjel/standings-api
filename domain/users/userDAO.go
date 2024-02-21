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
			username VARCHAR(50) UNIQUE NOT NULL,
			email VARCHAR(100) UNIQUE NOT NULL,
			password TEXT NOT NULL
		);
	`
	// dropTableUser            = `DROP TABLE IF EXISTS users;`
	queryInsertUser          = `INSERT INTO users (firstname, lastname, username, email, password) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	queryGetUserByID         = `SELECT id, firstname, lastname, email, username, password FROM users WHERE id = $1`
	queryGetUsernameAndEmail = `SELECT id, firstname, lastname, email, username FROM users WHERE username = $1 AND email = $2`
	queryGetAllUsers         = `SELECT id, firstname, lastname, email FROM users`
	queryUpdate              = `UPDATE users SET firstname = $1, lastname = $2, username = $3 WHERE id = $4`
	queryDelete              = `DELETE FROM users WHERE id = $1`
)

func init() {
	// tabulationDB.Client.Exec(dropTableUser)
	if _, err := tabulationDB.Client.Exec(createTableUser); err != nil {
		panic(err)
	}
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
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.UserName, &user.Password); getErr != nil {
		return errors.NewInternalServerError(getErr.Error())
	}

	return nil
}

func (user *User) GetByUsernameAndEmail() *errors.RestErr {
	stmt, err := tabulationDB.Client.Prepare(queryGetUsernameAndEmail)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.UserName, user.Email)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.UserName); getErr != nil {
		return errors.NewInternalServerError(getErr.Error())
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
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email); err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		users = append(users, user)
	}

	if len(users) == 0 {
		return nil, errors.NewNotFoundError("no users found")
	}

	return users, nil
}

func (user *User) Update(currentUser *User) *errors.RestErr {
	stmt, err := tabulationDB.Client.Prepare(queryUpdate)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.FirstName, user.LastName, user.UserName, user.ID); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := tabulationDB.Client.Prepare(queryDelete)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.ID); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}
