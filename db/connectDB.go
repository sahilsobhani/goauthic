package db

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/lib/pq" // Postgres driver
)

func Connect() *sql.DB {
    url := os.Getenv("DATABASE_URL")
    if url == "" {
        log.Fatal("DATABASE_URL is not set")
    }

    db, err := sql.Open("postgres", url)
    if err != nil {
        log.Fatalf("Failed to open database: %v", err)
    }

    if err = db.Ping(); err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    fmt.Println("✅ Connected to Postgres!")
    return db
}
