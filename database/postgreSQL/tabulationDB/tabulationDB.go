package tabulationDB

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	Client *sql.DB
	dbUrl  = "postgresql://postgres:*d5FFFcbAE516-3E-C5B4afB6-fDD3BD@viaduct.proxy.rlwy.net:49558/railway"
)

func init() {
	var err error
	Client, err = sql.Open("postgres", dbUrl)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Connected to the database")
}
