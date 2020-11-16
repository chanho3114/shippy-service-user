package main

import (
    "github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

// NewConnection returns a new database connection instance
func NewConnection() (*sqlx.DB, error) {
//	host := os.Getenv("DB_HOST")
//	user := os.Getenv("DB_USER")
//	dbName := os.Getenv("DB_NAME")
//	password := os.Getenv("DB_PASSWORD")
//	conn := fmt.Sprintf("host=%s port=5432 user=%s dbname=%s password=%s sslmode=disable", host, user, dbName, password)
//	db, err := sqlx.Connect("postgres", conn)
//  db, err := sql.Open("postgres", "postgresql://admin:password@database:5432/postgres?sslmode=disable")
    db, err := sqlx.Connect("postgres", "postgresql://admin:password@database:5432/postgres?sslmode=disable")
	if err != nil {
		return nil, err
	}

	return db, nil
}
