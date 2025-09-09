package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {

	connStr := "postgres://postgres:Lokesh@1129@localhost:5432/postgres?sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil || connStr == "" {
		panic(fmt.Sprintf("DB connection failed: %v", err))
	}

	// checking connection
	err = db.Ping()
	if err != nil {
		panic(fmt.Sprintf("DB ping failed: %v", err))
	}

	// Assigning the db instance to global variable
	DB = db

	fmt.Println("Connected to PostgreSQL successfully!")

}
