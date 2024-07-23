package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sfaizh/ticket-management-system/internal/structs"
)

type APIServer structs.APIServer

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		ListenAddr: listenAddr,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()

	router.HandleFunc("/", handleHTTP(s.handleLanding))
	log.Println("API server running on port:", s.ListenAddr)

	http.ListenAndServe(s.ListenAddr, router)

	return nil
}

func (s *APIServer) handleLanding(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return ToJSON(w, http.StatusOK, "Landing page for Go API Server")
	}

	return fmt.Errorf("Method no allowed %s", r.Method)
}

type httpFunc func(http.ResponseWriter, *http.Request) error

type httpError struct {
	Error string `json:"error"`
}

func handleHTTP(f httpFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			ToJSON(w, http.StatusBadRequest, httpError{Error: err.Error()})
		}
	}
}

func ToJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
