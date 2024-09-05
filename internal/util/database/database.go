package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sfaizh/ticket-management-system/internal/structs"
	"time"
)

// type Ticket structs.Ticket
type Ticket structs.Ticket

type Storage interface {
	GetTickets() ([]*Ticket, error)
	GetTicketByID(int) (*Ticket, error)
	CreateTicket(*Ticket) error
}

type dbStore struct {
	db *sql.DB
}

func NewStore() (*dbStore, error) {
	conn := "user=postgres dbname=supportdb password=gosupport sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &dbStore{
		db: db,
	}, nil
}

func (s *dbStore) Init() error {
	return s.createTicketTable()
}

func (s *dbStore) createTicketTable() error {
	// q := `create table if not exists ticket (
	//    id serial primary key,
	//    subject varchar(100),
	//    requester varchar(100),
	//    created_at timestamp
	//  )`

	q := `create table if not exists ticket (
	   id serial primary key,
	   subject varchar(100),
	   statusid integer,
	   userid integer,
	   requester varchar(100),
	   entryid integer,
	   created_at timestamp,
     foreign key (statusid) references status(id),
     foreign key (userid) references "users"(id),
     foreign key (entryid) references entries(id)
	 )`

	_, err := s.db.Exec(q)
	return err
}

// Ticket functionality
type CreateTicketRequest struct {
	Requester string `json:"requester"`
	Subject   string `json:"subject"`
	Text      string `json:"text"`
}

// Validate
// func (t *Ticket) ValidatePassword(p string) bool {
// 	return bcrypt.CompareHashAndPassword([]byte(t.EncryptedPassword), []byte(p)) == nil
// }

// Create a new ticket - requester is an email address
func NewTicket(requester, subject, text string) (*Ticket, error) {
	// encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// if err != nil {
	// 	return nil, err
	// }

	// create new entry
	entry := structs.Entry{
		Time: time.Now().UTC(),
		User: requester,
		Text: text,
	}

	var entries []structs.Entry
	entries = append(entries, entry)

	// write to file

	// return the ticket
	return &Ticket{
		Subject:   subject,
		Status:    structs.New,
		User:      structs.User{},
		Requester: requester,
		Entries:   entries,
		CreatedAt: time.Now().UTC(),
	}, nil
}

func (s *dbStore) CreateTicket(t *Ticket) error {
	q := `insert into ticket
  (subject, statusid, requester, userid, entryid, created_at)
  values ($1, $2, $3, $4, $5, $6)`

	_, err := s.db.Query(
		q,
		t.Subject,
		t.Status,
		t.Requester,
		t.User,
		t.Entries,
		t.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

// error here
// If err!= nil then res==nil and res.Body panics.
func (s *dbStore) GetTickets() ([]*Ticket, error) {
	rows, err := s.db.Query("select * from ticket")
	if err != nil {
		return nil, err
	}

	tickets := []*Ticket{}
	for rows.Next() {
		ticket, err := buildTicketsList(rows)
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, ticket)
	}

	return tickets, nil
}

func (s *dbStore) GetTicketByID(id int) (*Ticket, error) {
	rows, err := s.db.Query("select * from ticket where id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		// returns single ticket
		return buildTicketsList(rows)
	}

	return nil, fmt.Errorf("Ticket ID %d not found", id)
}

// needs adjustment - database does not support entries or user object, these should be changed
func buildTicketsList(rows *sql.Rows) (*Ticket, error) {
	ticket := new(Ticket)
	err := rows.Scan(
		&ticket.ID,
		&ticket.Subject,
		&ticket.Status,
		&ticket.User,
		&ticket.Requester,
		&ticket.Entries,
		&ticket.CreatedAt,
	)

	return ticket, err
}
