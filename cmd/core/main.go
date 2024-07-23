package main

import (
	// "flag"
	// "fmt"
	// "github.com/sfaizh/ticket-management-system/internal/structs/defaults"
	"log"
	// "github.com/sfaizh/ticket-management-system/internal/ticket"
	"github.com/sfaizh/ticket-management-system/internal/util/api"
)

// var (
// 	port     = flag.Uint("port", uint(defaults.ServerPort), "server `port`")
// 	tickets  = flag.String("tickets", defaults.ServerTickets, "tickets `directory`")
// 	users    = flag.String("users", defaults.ServerUsers, "users `file`")
// 	emails   = flag.String("emails", defaults.ServerEmails, "emails `directory`")
// 	verbose  = flag.Bool("verbose", defaults.LogVerbose, "enable `verbose` logs")
// 	logLevel = flag.String("logLevel", defaults.LogLevel, "logs verbose level either 'info' | 'debug'")
// )

func main() {
	// var testTicket = ticket.CreateTicket("example-customer@test.com", "New test support request", "Hi, i need support!")
	// fmt.Println(testTicket)
	// server := tcpserver.NewServer(":3000")
	apiServer := api.NewAPIServer(":3000")
	// go func() {
	// 	for msg := range server.Req {
	// 		fmt.Printf("Client request from: %s\n", msg.Requester)
	// 		ticket.CreateTicket(string(msg.Requester), string(msg.Subject), string(msg.Text))
	// 	}
	// }()

	log.Fatal(apiServer.Run())
}
