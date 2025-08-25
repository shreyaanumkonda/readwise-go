package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{addr: addr, db: db}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// Create MySQL store using the existing database connection
	store := NewMySQLStoreWithDB(s.db)
	service := NewService(store)
	service.RegisterRoutes(subrouter)

	// Add some basic routes
	router.HandleFunc("/", s.handleHome).Methods("GET")
	router.HandleFunc("/health", s.handleHealth).Methods("GET")

	log.Printf("Server starting on %s", s.addr)
	return http.ListenAndServe(s.addr, subrouter)
}

func (s *APIServer) handleHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Welcome to Readwise-like API", "endpoints": ["/users/{user-ID}/Parse-Kindle-File", "/cloud/Send-Daily-Insights"]}`))
}

func (s *APIServer) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "healthy", "service": "readwise-api"}`))
}
