package main

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/gorilla/mux"
)

type Service struct {
	database Storage
}

func NewService(database Storage) *Service {
	return &Service{database: database}
}

func (s *Service) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/users/{user_id}/parse-kindle-file", s.handleParseKindleFile).Methods("POST")
	r.HandleFunc("/cloud/send-daily-insights", s.handleSendDailyInsights).Methods("GET")
}

func (s *Service) handleParseKindleFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)                //this is a map of the variables in the request(extracted using gorilla mux)
	userID := vars["user_id"]          //this is the userID from the request
	file, _, err := r.FormFile("file") //this is the file from the request(extracted using gorilla mux)

	if err != nil {
		WriteJSON(w, http.StatusBadRequest, fmt.Sprintf("Error parsing file: %v", err))
		return
	}
	defer file.Close()
	//parse the multipart file
	rawExtractBook, err := ParseKindleExtractFile(file)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, fmt.Sprintf("Error parsing file: %v", err))
		return
	}

	//save the highlights to the database
	if err := s.saveHighlightsToDatabase(rawExtractBook, userID); err != nil {
		WriteJSON(w, http.StatusInternalServerError, fmt.Sprintf("Error saving highlights: %v", err))
		return
	}

	WriteJSON(w, http.StatusOK, "Highlights saved successfully")
}

func (s *Service) handleSendDailyInsights(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement daily insights logic
	WriteJSON(w, http.StatusOK, map[string]string{
		"message": "Daily insights feature coming soon",
		"status":  "not implemented yet",
	})
}

func ParseKindleExtractFile(file multipart.File) (*RawExtractBook, error) {
	decoder := json.NewDecoder(file)
	var rawExtractBook RawExtractBook
	err := decoder.Decode(&rawExtractBook)
	if err != nil {
		return nil, err
	}
	return &rawExtractBook, nil
}

func (s *Service) saveHighlightsToDatabase(rawExtractBook *RawExtractBook, userID string) error {
	// First save the book
	book := &Book{
		ID:        rawExtractBook.ASIN, // Use ASIN as ID since it's unique
		ASIN:      rawExtractBook.ASIN,
		Title:     rawExtractBook.Title,
		Authors:   rawExtractBook.Authors,
		UserID:    userID,
		CreatedAt: "", // Will use database default
		UpdatedAt: "", // Will use database default
	}

	if err := s.database.CreateBook(book); err != nil {
		return fmt.Errorf("failed to create book: %w", err)
	}

	// Then save the highlights
	if err := s.database.SaveHighlights(rawExtractBook, userID); err != nil {
		return fmt.Errorf("failed to save highlights: %w", err)
	}

	return nil
}
