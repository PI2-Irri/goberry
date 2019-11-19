package socket

import (
	"bufio"
	"log"
	"net"
	"strings"

	"github.com/PI2-Irri/goberry/common"
)

// ClientQueue receives a string to be sent to the server
// with a tcp socket connection
var ClientQueue chan string

// Client provides data loaded from the configuration
// to create a tcp socket client
type Client struct {
	Host string
	Port string
}

// CreateClient creates a Client and set all fields
// acording to the configuration json
func CreateClient() *Client {
	config := common.LoadConfiguration()

	client := &Client{
		Host: config.SocketClient["host"],
		Port: config.SocketClient["port"],
	}

	return client
}

// Start opens a thread which listens for the Queue channel
// receiving strings to send to the server
func (c *Client) Start() {
	for msg := range ClientQueue {
		log.Println("Client queue received:", msg)
	}
}

func (c *Client) connect() {
	list := []string{c.Host, c.Port}
	hostport := strings.Join(list, ":")

	log.Println("Client connecting to:", hostport)

	conn, err := net.Dial(network, hostport)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Client connected with:", conn.RemoteAddr())

	handleClientConnection(conn)

	log.Println("Client disconnected with:", conn.RemoteAddr())
}

func handleClientConnection(conn net.Conn) {
	connBuffer := bufio.NewReader(conn)
	for {
		str, err := connBuffer.ReadString('\n')
		if err != nil {
			log.Println("Connection closed with ", conn.RemoteAddr())
			break
		}
		if len(str) > 0 {
			// TODO: parse message
			log.Println("Client received:", str)
		}
	}

}
