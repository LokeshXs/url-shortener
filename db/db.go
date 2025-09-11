package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {

	connStr := os.Getenv("POSTGRES_URL")

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

	// Creating the required tables
	err = createTables();
	if(err !=nil){
		panic(err);
	}

	fmt.Println("Connected to PostgreSQL successfully!")

}


func createTables() error{

	// Creating the required tables

	// users table
	usersTableQuery := 	`
	CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	name TEXT,
	email TEXT UNIQUE NOT NULL,
	created_at TIMESTAMPTZ DEFAULT NOW()
	);
	`

	_,err := DB.Exec(usersTableQuery);

	if(err != nil){
		return errors.New(fmt.Sprintf("Failed to create USERS TABLE. Error:%v",err))
	}

	// urls table
	urlsTableQuery := `
	CREATE TABLE IF NOT EXISTS urls (
	id SERIAL PRIMARY KEY,
	original_url TEXT NOT NULL,
	shortcode VARCHAR(10) UNIQUE NOT NULL, 
	created_at TIMESTAMPTZ DEFAULT NOW(),
	user_id INT REFERENCES users(id) ON DELETE CASCADE
	);
	`

	_,err = DB.Exec(urlsTableQuery);

	if(err != nil){
		return errors.New(fmt.Sprintf("Failed to create URLS TABLE. Error:%v",err))
	}

	return nil
	

}
