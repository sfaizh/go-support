// Ticket Management System
// By Faizan Hussain
// Package implements system-wide structs used for the server and cli tool

package structs

import (
	"net"
	"strconv"
	"time"
)

type ServerConfig struct {
	// port the server is listening on
	Port uint16

	// directory tickets are stored to
	Tickets string

	// path to users.json
	Users string

	// path to mails
	EmailPath string
}

type Request struct {
	Requester []byte
	Subject   []byte
	Text      []byte
}

type Message struct {
	From    string
	Payload []byte
}

type Server struct {
	ListenAddr string
	Ln         net.Listener
	QuitCh     chan struct{}
	Msg        chan Message
	Req        chan Request
}

type APIServer struct {
	ListenAddr string
}

// provides the CLI config parameters set on startup
type CLIConfig struct {
	Host string
	Port uint16
}

// provides logging configuration flags set on startup
type LogConfig struct {
	LogLevel LogLevel
	Verbose  bool
}

const TicketIdLength int = 10

// custom type via enumerated constant defining log levels
type LogLevel int

const (
	LevelInfo LogLevel = iota
	LevelDebug
)

// converts a log level to corresponding output string which will be used in the log
func (level LogLevel) String() string {
	switch level {
	case LevelInfo:
		return "[INFO]"

	case LevelDebug:
		return "[DEBUG]"
	}

	return "undefined"
}

// converts given logging output string to log level
// func AsLogLevel(LogLevel string) LogLevel {
//   switch LogLevel {
//   case "info":
//     return LevelInfo
//
//   case "debug":
//     return LevelDebug
//   }
//
//   return LogLevel(-1)
// }

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// user session
type Session struct {
	ID         string
	User       User
	Created    time.Time
	IsLoggedIn bool
}

// holds sessions
type SessionManager struct {
	Name    string
	Session Session
	TTL     int64
}

// user assigned to ticket
type Assignee struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Hash     string `json:"hash"`
}

type Status int

const (
	New Status = iota
	Open
	Pending
	OnHold
	Solved
)

// converts ticket status to description (string output)
func (status Status) String() string {
	switch status {
	case New:
		return "New"

	case Open:
		return "Open"

	case Pending:
		return "Pending"

	case OnHold:
		return "On Hold"

	case Solved:
		return "Solved"
	}

	return "Status not Defined"
}

type Ticket struct {
	ID        string  `json:"id"`
	Subject   string  `json:"subject"`
	Status    Status  `json:"status"`
	User      User    `json:"user"`
	Requester string  `json:"requester"`
	Entries   []Entry `json:"entries"`
}

type Entry struct {
	Text     string    `json:"text"`
	User     string    `json:"user"`
	Time     time.Time `json:"time"`
	Internal bool      `json:"internal"`
}

// an email for a ticket entry
type Email struct {
	ID      string `json:"id"`
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

// key:value mappings for the mail API
type EmailJSONMap map[string]interface{}

// converts command to string output
func (c Command) String() string {
	return strconv.Itoa(int(c))
}

// CLI command to interact with mail
type Command int

const (
	Fetch Command = iota
	Submit
	Exit
)
