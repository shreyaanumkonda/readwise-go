package main

import (
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

func main() {
	// MySQL configuration from environment variables
	cfg := mysql.Config{
		User:                 getEnv("DB_USER", "root"),
		Passwd:               getEnv("DB_PASSWORD", ""),
		Net:                  "tcp",
		Addr:                 getEnv("DB_HOST", "localhost:3306"),
		DBName:               getEnv("DB_NAME", "highlights"),
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	storage := NewMySQLStore(cfg)

	// Initialize database tables
	if err := storage.Init(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Create a new API server with database
	server := NewAPIServer(":8080", storage)

	// Start the server
	log.Println("Starting Readwise-like API server...")
	if err := server.Run(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

// getEnv gets environment variable or returns default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
