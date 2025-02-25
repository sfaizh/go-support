package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sfaizh/ticket-management-system/internal/util/database"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Storage database.Storage

type APIServer struct {
	listenAddr string
	store      Storage
}

func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()

	router.HandleFunc("/", handleHTTP(s.handleLanding))
	router.HandleFunc("/tickets", handleHTTP(s.handleGetTickets))
	router.HandleFunc("/tickets/{id}", handleHTTP(s.handleGetTicketByID))
	router.HandleFunc("/entries/{id}", handleHTTP(s.handleGetTicketEntries))
	router.HandleFunc("/create", handleHTTP(s.handleCreateTicket))
	log.Println("API rebuilt server running on port:", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)

	return nil
}

func (s *APIServer) handleGetTickets(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		tickets, err := s.store.GetTickets()
		if err != nil {
			return err
		}

		return ToJSON(w, http.StatusOK, tickets)
	}

	return fmt.Errorf("Method no allowed %s", r.Method)
}

func (s *APIServer) handleGetTicketEntries(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		id, err := getID(r)
		if err != nil {
			return err
		}
		entries, err := s.store.GetTicketEntries(id)
		if err != nil {
			return err
		}

		return ToJSON(w, http.StatusOK, entries)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetTicketByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		id, err := getID(r)
		if err != nil {
			return err
		}

		ticket, err := s.store.GetTicketByID(id)
		if err != nil {
			return err
		}

		return ToJSON(w, http.StatusOK, ticket)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleCreateTicket(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		// build ticket here
		// if i use database.Ticket, it will have extra fields User, Entries
		var t database.Ticket

		// get fields from body
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			return fmt.Errorf("Failed to decode JSON: %w", err)
		}

		t.CreatedAt = time.Now().UTC()

		// store the ticket
		if err := s.store.StoreTicket(&t); err != nil {
			log.Fatal(err)
		}

		return ToJSON(w, http.StatusOK, t)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
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

func getID(r *http.Request) (int, error) {
	idString := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		return id, fmt.Errorf("Invalid id given %s", idString)
	}
	return id, nil
}
