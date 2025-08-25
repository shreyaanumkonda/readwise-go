package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	// TODO: Create MySQL database connection
	// For now, using nil to test server without database
	var db *sql.DB = nil

	// MySQL configuration (commented out for testing)
	/*
		cfg := mysql.Config{
			User:                 "root",     // TODO: Use environment variables
			Passwd:               "password", // TODO: Use environment variables
			Net:                  "tcp",
			Addr:                 "localhost:3306", // TODO: Use environment variables
			DBName:               "highlights",     // TODO: Use environment variables
			AllowNativePasswords: true,
			ParseTime:            true,
		}

		// Create MySQL database connection
		db, err := sql.Open("mysql", cfg.FormatDSN())
		if err != nil {
			log.Fatal("Failed to connect to database:", err)
		}
		defer db.Close()

		// Test the connection
		if err := db.Ping(); err != nil {
			log.Fatal("Failed to ping database:", err)
		}
		log.Println("Connected to MySQL database!")
	*/

	// Create a new API server with database
	server := NewAPIServer(":8080", db)

	// Start the server
	log.Println("Starting Readwise-like API server...")
	if err := server.Run(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
