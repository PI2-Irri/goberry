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

func init() {
	ClientQueue = make(chan string, 1024)
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

	log.Println("Client trying connection with:", hostport)
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

	// connBuffer := bufio.NewReader(conn)
	// str, _ := readString(connBuffer, true)
	// if !strings.Contains(str, "HELO") {
	// 	log.Fatal("Client didn't receive HELO from the server")
	// }

	for msg := range ClientQueue {
		handleClientConnection(conn, msg)
	}

	log.Println("Client disconnected with:", conn.RemoteAddr())
}

func handleClientConnection(conn net.Conn, msg string) {
	// connBuffer := bufio.NewReader(conn)

	message := []byte(msg + "\n")
	n, err := conn.Write(message)
	if err != nil || n != len(message) {
		log.Println("Error while trying to send message to server")
		log.Println("\tError:", err)
		log.Fatal("\tBytes:", n)
	}
	// str, _ := readString(connBuffer, false)
	// str = str[:len(str)-1]
	// log.Println("Client received:", str)
	// if !strings.Contains(str, "OK") {
	// 	log.Fatal("Client didn't receive OK from the server")
	// }

}

func readString(buffer *bufio.Reader, isErrorFatal bool) (str string, hadError bool) {
	var err error
	str, err = buffer.ReadString('\n')

	if err != nil && isErrorFatal {
		log.Fatal(err)
	} else if err != nil {
		hadError = true
	}

	return str, hadError
}
