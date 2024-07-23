package ticket

import (
	"time"
	//"strconv"
	//"sort"
	//"github.com/sfaizh/ticket-management-system/internal/structs/defaults"
	"github.com/sfaizh/ticket-management-system/internal/structs"
)

// Create a new ticket - requester is an email address
func CreateTicket(requester, subject, text string) structs.Ticket {
	// create new entry
	entry := structs.Entry{
		Time: time.Now(),
		User: requester,
		Text: text,
	}

	var entries []structs.Entry
	entries = append(entries, entry)

	// write to file

	// return the ticket
	return structs.Ticket{
		ID:        "0",
		Subject:   subject,
		Status:    structs.New,
		User:      structs.User{},
		Requester: requester,
		Entries:   entries,
	}
}
