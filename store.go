package main

type Storage interface {
	CreateBook(book *Book) error
	SaveHighlights(rawExtractBook *RawExtractBook, userID string) error
	GetBook(userID string) (*Book, error)
}

// MySQLStore in db.go implements this Storage interface
