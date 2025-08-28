package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   Storage
}

func NewAPIServer(addr string, db Storage) *APIServer {
	return &APIServer{addr: addr, db: db}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// Create service using the storage interface
	service := NewService(s.db)
	service.RegisterRoutes(subrouter)

	// Add some basic routes
	router.HandleFunc("/", s.handleHome).Methods("GET")
	router.HandleFunc("/health", s.handleHealth).Methods("GET")

	log.Printf("Server starting on %s", s.addr)
	return http.ListenAndServe(s.addr, router)
}

func (s *APIServer) handleHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Welcome to Readwise-like API", "endpoints": ["/users/{user_id}/parse-kindle-file", "/cloud/send-daily-insights"]}`))
}

func (s *APIServer) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "healthy", "service": "readwise-api"}`))
}
