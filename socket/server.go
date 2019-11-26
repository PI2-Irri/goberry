package socket

import (
	"bufio"
	"log"
	"net"
	"strings"

	"github.com/PI2-Irri/goberry/common"
	"github.com/PI2-Irri/goberry/measurement"
)

// Server type holds the informations necessary
// for configuring a TCP socket
type Server struct {
	Host     string
	Port     string
	listener net.Listener
}

const network = "tcp"

// CreateServer instantiates a socket properly configures
// according to a cfg.json
func CreateServer() *Server {
	config := common.LoadConfiguration()

	s := &Server{
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
func (s *Server) AcceptConnections() {
	log.Println("Server accepting connections")
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		remote := conn.RemoteAddr().String()
		log.Println("Server connected with:", remote)

		handleServerConnection(conn)

		log.Println("Server disconnected with:", remote)
	}
}

func handleServerConnection(conn net.Conn) {
	defer conn.Close()

	// conn.Write([]byte("HELO\n"))
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		msg := scanner.Text()
		log.Println("Server received:", msg)
		sender := measurement.ParseMessage(msg)
		measurement.Queue <- sender
		// conn.Write([]byte("OK\n"))
	}

	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
}
