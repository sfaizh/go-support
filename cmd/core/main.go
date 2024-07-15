package main

import (
	"flag"
	"fmt"
	"log"
	// "log"
	// "math"
	// "os"
	"github.com/sfaizh/ticket-management-system/internal/structs/defaults"
	// "github.com/sfaizh/ticket-management-system/internal/ticket"
	"github.com/sfaizh/ticket-management-system/internal/util/tcpserver"
)

var (
	port     = flag.Uint("port", uint(defaults.ServerPort), "server `port`")
	tickets  = flag.String("tickets", defaults.ServerTickets, "tickets `directory`")
	users    = flag.String("users", defaults.ServerUsers, "users `file`")
	emails   = flag.String("emails", defaults.ServerEmails, "emails `directory`")
	verbose  = flag.Bool("verbose", defaults.LogVerbose, "enable `verbose` logs")
	logLevel = flag.String("logLevel", defaults.LogLevel, "logs verbose level either 'info' | 'debug'")
)

func main() {
	// var testTicket = ticket.CreateTicket("example-customer@test.com", "New test support request", "Hi, i need support!")
	// fmt.Println(testTicket)
	fmt.Println(port)
	server := tcpserver.NewServer(":3000")
	go func() {
		for msg := range server.Msg {
			fmt.Printf("sender:%s\nmsg:%s\n", msg.From, string(msg.Payload))
		}
	}()

	log.Fatal(server.Start())
}
