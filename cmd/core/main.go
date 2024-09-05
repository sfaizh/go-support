package main

import (
	"flag"
	"fmt"
	"github.com/sfaizh/ticket-management-system/internal/util/api"
	"github.com/sfaizh/ticket-management-system/internal/util/database"
	"log"
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
func SeedTicket(store Storage, requester, subject, description string) *database.Ticket {
	t, err := database.NewTicket(requester, subject, description)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.CreateTicket(t); err != nil {
		log.Fatal(err)
	}

	fmt.Println("New ticket created: ", t.Requester)

	return t
}

func SeedTickets(s Storage) {
	SeedTicket(s, "faizan@gmail.com", "Support request ticket", "I need help setting up APM for Node.js.")
	// need to seed the rest of the data - status, users, entries
}

func main() {
	// var testTicket = ticket.CreateTicket("example-customer@test.com", "New test support request", "Hi, i need support!")
	// fmt.Println(testTicket)
	// server := tcpserver.NewServer(":3000")

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
