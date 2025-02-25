package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sfaizh/ticket-management-system/internal/structs"
)

// type Ticket structs.Ticket
type Ticket structs.Ticket

// type Entry []structs.Entry

type Storage interface {
	GetTickets() ([]*Ticket, error)
	GetTicketByID(int) (*Ticket, error)
	StoreTicket(*Ticket) error
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
	q1 := `CREATE TABLE IF NOT EXISTS ticket (
    id SERIAL PRIMARY KEY,
    subject VARCHAR(100),
    requester VARCHAR(100),
    created_at TIMESTAMP
);`

	q2 := `CREATE TABLE IF NOT EXISTS ticket_entry (
    id SERIAL PRIMARY KEY,
    ticket_id INTEGER REFERENCES ticket(id),
    text TEXT,
    "user" VARCHAR(100),
    "time" TIMESTAMP,
    internal BOOLEAN
);`

	// Then execute q1 and q2 separately:
	if _, err := s.db.Exec(q1); err != nil {
		// handle error
	}
	if _, err := s.db.Exec(q2); err != nil {
		// handle error
	}
	// q := `create table if not exists ticket (
	//    id serial primary key,
	//    subject varchar(100),
	//    statusid integer,
	//    userid integer,
	//    requester varchar(100),
	//    entryid integer,
	//    created_at timestamp,
	//     foreign key (statusid) references status(id),
	//     foreign key (userid) references "users"(id),
	//     foreign key (entryid) references entries(id)
	//  )`

	return nil
}

// Ticket functionality
type CreateTicketRequest struct {
	Requester string          `json:"requester"`
	Subject   string          `json:"subject"`
	Text      string          `json:"text"`
	Entries   []structs.Entry `json:"entries"`
}

// Validate
// func (t *Ticket) ValidatePassword(p string) bool {
// 	return bcrypt.CompareHashAndPassword([]byte(t.EncryptedPassword), []byte(p)) == nil
// }

func (s *dbStore) StoreTicket(t *Ticket) error {
	// there should be a link to entries here in the form of a list of varchar(100)
	q1 := `insert into ticket
  (subject, requester, created_at)
  values ($1, $2, $3)
  returning id`

	q2 := `insert into ticket_entry
  (ticket_id, text, "user", "time", internal)
  values ($1, $2, $3, $4, $5)`

	// Then execute q1 and q2 separately:
	// Use Scan to store id from ticket into t.ID
	if err := s.db.QueryRow(q1, t.Subject, t.Requester, t.CreatedAt).Scan(&t.ID); err != nil {
		// handle error
	}

	for _, entry := range t.Entries {
		if _, err := s.db.Exec(q2, t.ID, entry.Text, entry.User, entry.Time, entry.Internal); err != nil {
			// handle error
		}
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

// entries are stored in a normalised fashion
func buildTicketsList(rows *sql.Rows) (*Ticket, error) {
	ticket := new(Ticket)
	err := rows.Scan(
		&ticket.ID,
		&ticket.Subject,
		// &ticket.Status,
		// &ticket.User,
		&ticket.Requester,
		&ticket.CreatedAt,
	)

	return ticket, err
}
