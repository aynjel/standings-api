package tabulationdb

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client        *sql.DB
	MYSQLDATABASE = "railway"
	MYSQLHOST     = "monorail.proxy.rlwy.net"
	MYSQLPASSWORD = "123GeH5dafcDGe5A5FaF6DcCAbdDHBg6"
	MYSQLUSER     = "root"
)

func init() {
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parsetime=true", MYSQLUSER, MYSQLPASSWORD, MYSQLHOST, MYSQLDATABASE)
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Database successfully configured")
}
