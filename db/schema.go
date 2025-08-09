package db

import "database/sql"

func InitializeSchema(db *sql.DB) error {
    createTableQuery := `
    CREATE TABLE IF NOT EXISTS users (
        id VARCHAR(36) PRIMARY KEY,
        email VARCHAR(255) UNIQUE NOT NULL,
        password VARCHAR(255) NOT NULL,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    );`

    _, err := db.Exec(createTableQuery)
    return err
}
