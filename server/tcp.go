package server

import (
	"bufio"
	"log"
	"net"
	"strings"

	"github.com/PI2-Irri/goberry/common"
	"github.com/PI2-Irri/goberry/measurement"
)

// Socket type holds the informations necessary
// for configuring a TCP socket
type Socket struct {
	Host     string
	Port     string
	listener net.Listener
}

const network = "tcp"

// Create instantiates a socket properly configures
// according to a cfg.json
func Create() *Socket {
	config := common.LoadConfiguration()

	s := &Socket{
		Host: config.SocketServer["host"],
		Port: config.SocketServer["port"],
	}

	list := []string{s.Host, s.Port}
	hostport := strings.Join(list, ":")
	listener, err := net.Listen(network, hostport)

	if err != nil {
		log.Fatal(err)
	}
	s.listener = listener

	return s
}

// AcceptConnections starts the socket activa phase where it
// continuously accepts connections
func (s *Socket) AcceptConnections() {
	log.Println("Socket accepting connection")
	defer close(measurement.Queue)
	for {
		conn, err := s.listener.Accept()
		remote := conn.RemoteAddr().String()
		log.Println("Starting connection with:", remote)
		if err != nil {
			log.Fatal(err)
		}
		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		msg := scanner.Text()
		log.Println("Received:", msg)
		sender := measurement.ParseMessage(msg)
		measurement.Queue <- sender
		conn.Write([]byte("OK\n"))
	}

	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
}
