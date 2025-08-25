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
	r.HandleFunc("/users/{user-ID}/Parse-Kindle-File", s.handleParseKindleFile).Methods("POST")
	r.HandleFunc("/cloud/Send-Daily-Insights", s.handleSendDailyInsights).Methods("GET")
}

func (s *Service) handleParseKindleFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)                //this is a map of the variables in the request(extracted using gorilla mux)
	userID := vars["user-ID"]          //this is the userID from the request
	file, _, err := r.FormFile("file") //this is the file from the request(extracted using gorilla mux)
	println(userID)                    //remove this at the end
	println(file)

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
	println(rawExtractBook.Title)      //remove this at the end
	println(rawExtractBook.Authors)    //remove this at the end
	println(rawExtractBook.Highlights) //remove this at the end
	//save the highlights to the database

}

func (s *Service) handleSendDailyInsights(w http.ResponseWriter, r *http.Request) {

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
	return nil
}
