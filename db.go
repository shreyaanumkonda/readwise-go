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
	query := `INSERT INTO books (id, asin, title, authors, user_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := s.db.Exec(query, book.ID, book.ASIN, book.Title, book.Authors, book.UserID, book.CreatedAt, book.UpdatedAt)
	return err
}

func (s *MySQLStore) SaveHighlights(rawExtractBook *RawExtractBook, userID string) error {
	// Save each highlight individually
	for _, highlight := range rawExtractBook.Highlights {
		query := `INSERT INTO highlights (id, book_id, highlight, user_id, created_at, updated_at) VALUES (?, ?, ?, ?, NOW(), NOW())`
		_, err := s.db.Exec(query, highlight.Text, rawExtractBook.ASIN, highlight.Text, userID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *MySQLStore) SaveHighlightsWithBookID(rawExtractBook *RawExtractBook, userID string, bookID string) error {
	// Save each highlight individually with the specified book ID
	for _, highlight := range rawExtractBook.Highlights {
		query := `INSERT INTO highlights (id, book_id, highlight, user_id, created_at, updated_at) VALUES (?, ?, ?, ?, NOW(), NOW())`
		_, err := s.db.Exec(query, highlight.Text, bookID, highlight.Text, userID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *MySQLStore) GetBook(userID string) (*Book, error) {
	query := `SELECT * FROM books WHERE user_id = ? LIMIT 1`
	row := s.db.QueryRow(query, userID)

	var book Book
	if err := row.Scan(&book.ID, &book.ASIN, &book.Title, &book.Authors, &book.UserID, &book.CreatedAt, &book.UpdatedAt); err != nil {
		return nil, err
	}
	return &book, nil
}

func (s *MySQLStore) GetBooks(userID string) ([]*Book, error) {
	query := `SELECT * FROM books WHERE user_id = ?`
	rows, err := s.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*Book
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.ASIN, &book.Title, &book.Authors, &book.UserID, &book.CreatedAt, &book.UpdatedAt); err != nil {
			return nil, err
		}
		books = append(books, &book)
	}
	return books, nil
}
func (s *MySQLStore) CreateUserTable() error {
	query := `CREATE TABLE IF NOT EXISTS users (
		id VARCHAR(255) PRIMARY KEY,
		email VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	)`
	_, err := s.db.Exec(query)
	return err
}

func (s *MySQLStore) CreateBookTable() error {
	query := `CREATE TABLE IF NOT EXISTS books (
        id VARCHAR(255) PRIMARY KEY,
        asin VARCHAR(255) NOT NULL,
        title TEXT NOT NULL,
        authors TEXT NOT NULL,
        user_id VARCHAR(255) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        UNIQUE KEY (asin)
    )`

	_, err := s.db.Exec(query)
	return err
}
func (s *MySQLStore) CreateHighlightTable() error {
	query := `CREATE TABLE IF NOT EXISTS highlights (
		id VARCHAR(255) PRIMARY KEY,
		book_id VARCHAR(255) NOT NULL,
		highlight TEXT NOT NULL,
		user_id VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		FOREIGN KEY (book_id) REFERENCES books(id),
		FOREIGN KEY (user_id) REFERENCES users(id)
	)`
	_, err := s.db.Exec(query)
	return err
}
func (s *MySQLStore) Init() error {
	if err := s.CreateUserTable(); err != nil {
		return err
	}
	if err := s.CreateBookTable(); err != nil {
		return err
	}
	if err := s.CreateHighlightTable(); err != nil {
		return err
	}
	return nil
}
