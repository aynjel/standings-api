SQLite usage (development)

This project initializes the database connection from `database/tabulationDB/tabulationDB.go`.

Environment variables supported:

- APP_ENV

  - If set to `development` (or unset), the application uses SQLite by default.
  - Any other value will cause the initializer to attempt to use PostgreSQL by reading `DATABASE_URL`.

- SQLITE_DB_PATH

  - Path to the SQLite database file used in development mode.
  - Default: `./database/standings_dev.db` (inside the repo).

- SQLITE_IN_MEMORY
  - If set to `true` or `1` and running in development, the DB will be opened in-memory (`:memory:`) which is
    useful for tests or ephemeral runs.

Examples (PowerShell):

# Run using default dev DB file

$env:APP_ENV = 'development'; go run main.go

# Run using a custom DB file

$env:APP_ENV = 'development'; $env:SQLITE_DB_PATH = 'D:\data\my_dev.db'; go run main.go

# Run with an in-memory DB

$env:APP_ENV = 'development'; $env:SQLITE_IN_MEMORY = 'true'; go run main.go

Notes:

- For production, set APP_ENV to something other than `development` and provide `DATABASE_URL` with your
  Postgres connection string.
- The initializer panics on DB open/ping errors (consistent with the project's previous behavior).

Additional notes:

- This project uses the pure-Go SQLite driver `modernc.org/sqlite` so it works when building with
  `CGO_ENABLED=0`. If you prefer the cgo-backed driver `github.com/mattn/go-sqlite3`, set `CGO_ENABLED=1` in
  your environment and replace the sqlite import in `database/tabulationDB/tabulationDB.go`.
