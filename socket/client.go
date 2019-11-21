package socket

import (
	"bufio"
	"log"
	"net"
	"strings"
	"time"

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
	c.connect()
}

func (c *Client) connect() {
	list := []string{c.Host, c.Port}
	hostport := strings.Join(list, ":")

	var conn net.Conn
	var err error

	log.Println("Client starting connection with:", hostport)
	for {

		conn, err = net.Dial(network, hostport)
		if err != nil {
			log.Println("Client could not connect:\n\t", err)
			time.Sleep(time.Second * 5)
		} else {
			break
		}
	}

	log.Println("Client connected successfully with:", conn.RemoteAddr())

	for msg := range ClientQueue {
		log.Println("Client queue received:", msg)
		handleClientConnection(conn, msg)
	}

	log.Println("Client disconnected with:", conn.RemoteAddr())
}

func handleClientConnection(conn net.Conn, msg string) {
	connBuffer := bufio.NewReader(conn)

	str, _ := readString(connBuffer, true)
	if str != "HELO" {
		log.Fatal("Client didnt receive HELO from the server")
	}

	for {
		conn.Write([]byte(msg)) // TODO: test this
		str, err := readString(connBuffer, false)
		if err {
			break
		}
		if str != "OK" {
			log.Fatal("Client didnt receive OK from the server")
		}
	}

}

func readString(buffer *bufio.Reader, isErrorFatal bool) (str string, hadError bool) {
	var err error
	str, err = buffer.ReadString('\n')

	if err != nil && isErrorFatal {
		log.Fatal(err)
	} else if err != nil {
		hadError = true
	}

	log.Println("Client received:", str)
	return
}
