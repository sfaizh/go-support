// Ticket Management System
// By Faizan Hussain
// Package implements tcp module used for sending and recieving messages from the client server

package tcpserver

import (
	"fmt"
	"github.com/sfaizh/ticket-management-system/internal/structs"
	"io"
	"net"
)

// define local type backed by structs.Server
type Server structs.Server

// defines the server
func NewServer(listenAddr string) *Server {
	return &Server{
		ListenAddr: listenAddr,
		QuitCh:     make(chan struct{}),
		Msg:        make(chan structs.Message, 10),
		Req:        make(chan structs.Request, 10),
	}
}

// starts the server
func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.Ln = ln

	//start accepting connections
	go s.acceptLoop()
	//what does <- syntax mean? | this line seems to be closing all channels
	<-s.QuitCh
	//close message channel
	close(s.Msg)

	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.Ln.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}
		fmt.Println("new connection:", conn.RemoteAddr())
		go s.readLoop(conn)
	}
}

func (s *Server) CLIPrompt(buffer *[]byte, conn net.Conn, prompt string) []byte {
	conn.Write([]byte(prompt))
	n, err := conn.Read(*buffer)

	if err == io.EOF {
		return nil
	}
	if err != nil {
		fmt.Println("read error:", err)
		return nil
	}

	return (*buffer)[:n]
}

func (s *Server) readLoop(conn net.Conn) {
	defer conn.Close()

	conn.Write([]byte("***** Ticket management system by Faizan *****\n"))

	for {
		conn.Write([]byte("***** Submit a new Ticket *****\n"))
		// buf := make([]byte, 2048)
		requester := make([]byte, 2048)
		subject := make([]byte, 2048)
		text := make([]byte, 2048)

		s.CLIPrompt(&requester, conn, "Email address: ")
		s.CLIPrompt(&subject, conn, "Subject: ")
		s.CLIPrompt(&text, conn, "Description: ")

		s.Req <- structs.Request{
			Requester: requester,
			Subject:   subject,
			Text:      text,
		}
		conn.Write([]byte("Thank you for your request!\n"))
	}
}
