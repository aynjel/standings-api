package tabulationDB

import (
	"database/sql"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	// use pure-Go sqlite driver to avoid cgo requirement when building with CGO_ENABLED=0
	_ "modernc.org/sqlite"
)

var Client *sql.DB
// Driver exports the currently used SQL driver name (eg "sqlite" or "postgres").
var Driver string

func init() {
    env := os.Getenv("APP_ENV")
    if env == "" {
        env = "development"
    }

    var err error
    if env == "development" {
        inMemory := strings.ToLower(os.Getenv("SQLITE_IN_MEMORY"))
        if inMemory == "true" || inMemory == "1" {
            Client, err = sql.Open("sqlite", ":memory:")
            Driver = "sqlite"
            if err != nil {
                log.Println("failed to open sqlite in-memory DB:", err)
                panic(err)
            }
        } else {
            dbPath := os.Getenv("SQLITE_DB_PATH")
            if dbPath == "" {
                // keep DB file inside the project database folder by default
                dbPath = "./database/standings_dev.db"
            }
            Client, err = sql.Open("sqlite", dbPath)
            Driver = "sqlite"
            if err != nil {
                log.Println("failed to open sqlite DB (path=", dbPath, "):", err)
                panic(err)
            }
        }
    } else {
        dbUrl := os.Getenv("DATABASE_URL")
        if dbUrl == "" {
            // default postgres connection used previously in the project
            dbUrl = "postgresql://postgres:*d5FFFcbAE516-3E-C5B4afB6-fDD3BD@viaduct.proxy.rlwy.net:49558/railway"
        }
        Client, err = sql.Open("postgres", dbUrl)
        Driver = "postgres"
        if err != nil {
            log.Println("failed to open postgres DB:", err)
            panic(err)
        }
    }

    if err = Client.Ping(); err != nil {
        log.Println("failed to ping DB:", err)
        panic(err)
    }

    log.Printf("Connected to database (env=%s) driver=%s\n", env, Driver)
}
