package database

import (
    "database/sql"
    "log"
    "os"

    _ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
    dsn := os.Getenv("DATABASE_URL")
    var err error
    DB, err = sql.Open("postgres", dsn)
    if err != nil {
        log.Fatalf("Failed to connect to DB: %v", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatalf("Failed to ping DB: %v", err)
    }

    log.Println("Database connection established")
}
