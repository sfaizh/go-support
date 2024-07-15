// Ticket Management System
// By Faizan Hussain
// Package implements tcp module used for sending and recieving messages from the client server

package tcpserver

import (
	"fmt"
	"github.com/sfaizh/ticket-management-system/internal/structs"
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

func (s *Server) readLoop(conn net.Conn) {
	defer conn.Close()

	conn.Write([]byte("***** Ticket management system by Faizan *****"))

	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read error:", err)
			continue
		}

		//read input into message channel
		s.Msg <- structs.Message{
			From:    conn.RemoteAddr().String(),
			Payload: buf[:n],
		}

		conn.Write([]byte("Message sent"))
	}
}
