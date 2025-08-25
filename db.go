package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

type MySQLStore struct {
	db *sql.DB
}

func NewMySQLStore(cfg mysql.Config) *MySQLStore {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected to MySQL!")
	return &MySQLStore{db: db}
}

// NewMySQLStoreWithDB creates MySQLStore with an existing database connection
func NewMySQLStoreWithDB(db *sql.DB) *MySQLStore {
	return &MySQLStore{db: db}
}

func (s *MySQLStore) CreateBook(book *Book) error {
	// TODO: Implement MySQL book creation
	return nil
}

func (s *MySQLStore) SaveHighlights(rawExtractBook *RawExtractBook, userID string) error {
	// TODO: Implement MySQL highlights saving
	return nil
}

func (s *MySQLStore) GetBook(userID string) (*Book, error) {
	// TODO: Implement MySQL book retrieval
	return nil, nil
}
