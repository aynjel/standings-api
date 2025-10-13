package logs

import (
	"database/sql"
	"time"

	"anggi.tabulation/database/tabulationDB"
)

// ensureLogsTable creates the logs table if it does not exist.
func ensureLogsTable(db *sql.DB) error {
	create := `
	CREATE TABLE IF NOT EXISTS logs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		level TEXT NOT NULL,
		message TEXT NOT NULL,
		file TEXT,
		created_at DATETIME NOT NULL
	);
	`
	_, err := db.Exec(create)
	return err
}

// InsertLog inserts a log entry into the DB. If the DB driver does not support AUTOINCREMENT the
// SQL will still work for sqlite and postgres (serial behavior for postgres is handled by the DB).
func InsertLog(level, message, file string) error {
	if tabulationDB.Client == nil {
		return nil
	}

	// try to ensure table exists; ignore error if DB is read-only
	if err := ensureLogsTable(tabulationDB.Client); err != nil {
		return err
	}

	insert := `INSERT INTO logs (level, message, file, created_at) VALUES ($1, $2, $3, $4)`
	// postgres uses $1 placeholders, sqlite uses ?; adapt based on driver
	var _, err = func() (sql.Result, error) {
		if tabulationDB.Driver == "sqlite" {
			return tabulationDB.Client.Exec("INSERT INTO logs (level, message, file, created_at) VALUES (?, ?, ?, ?)", level, message, file, time.Now())
		}
		// default to postgres-style placeholders
		return tabulationDB.Client.Exec(insert, level, message, file, time.Now())
	}()
	return err
}
