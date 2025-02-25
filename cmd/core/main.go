package main

import (
	"flag"
	"fmt"
	"github.com/sfaizh/ticket-management-system/internal/structs"
	"github.com/sfaizh/ticket-management-system/internal/util/api"
	"github.com/sfaizh/ticket-management-system/internal/util/database"
	"log"
	"time"
)

type Storage database.Storage

// var (
// 	port     = flag.Uint("port", uint(defaults.ServerPort), "server `port`")
// 	tickets  = flag.String("tickets", defaults.ServerTickets, "tickets `directory`")
// 	users    = flag.String("users", defaults.ServerUsers, "users `file`")
// 	emails   = flag.String("emails", defaults.ServerEmails, "emails `directory`")
// 	verbose  = flag.Bool("verbose", defaults.LogVerbose, "enable `verbose` logs")
// 	logLevel = flag.String("logLevel", defaults.LogLevel, "logs verbose level either 'info' | 'debug'")
// )

// Seed tickets
func SeedTicket(store Storage, requester, subject, description string, entries []structs.Entry) *database.Ticket {
	t := &database.Ticket{
		Subject:   subject,
		Requester: requester,
		Entries:   entries,
		CreatedAt: time.Now().UTC(),
	}

	if err := store.StoreTicket(t); err != nil {
		log.Fatal(err)
	}

	fmt.Println("New ticket created: ", t.Requester)

	return t
}

func SeedTickets(s Storage) {
	entries := []structs.Entry{
		{
			TicketID: "1",
			Text:     "I need some help with this go code",
			User:     "TestUser",
			Time:     time.Now().UTC(),
			Internal: false,
		},
		{
			TicketID: "1",
			Text:     "Some other support ticket request",
			User:     "TestUser2",
			Time:     time.Now().UTC(),
			Internal: false,
		},
	}
	SeedTicket(s, "faizan@gmail.com", "Support request ticket", "I need help setting up APM for Node.js.", entries)
}

func main() {
	// Seed setup for database
	seed := flag.Bool("seed", true, "seed the db")
	flag.Parse()

	// Create new db
	store, err := database.NewStore()

	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	if *seed {
		fmt.Println("seeding initial value")
		SeedTickets(store)
	}

	apiServer := api.NewAPIServer(":3000", store)
	log.Fatal(apiServer.Run())
}
