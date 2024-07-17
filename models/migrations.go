package models

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func Migrate(db *sqlx.DB) {
	schema := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        username TEXT UNIQUE NOT NULL,
        password TEXT NOT NULL
    );

    CREATE TABLE IF NOT EXISTS tasks (
        id SERIAL PRIMARY KEY,
        title TEXT NOT NULL,
        status TEXT NOT NULL,
        user_id INTEGER REFERENCES users(id)
    );
    `
	_, err := db.Exec(schema)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
