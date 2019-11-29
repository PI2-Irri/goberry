package command

import (
	"log"
	"strconv"
	"strings"

	"github.com/PI2-Irri/goberry/socket"
)

// Command is a struct representation of the command that will be sent to
// the socket client
type Command struct {
	Status bool
	Timer  int
}

// ToString converts a command to a string that
// will be sent over the tcp socket
func (c *Command) ToString() string {
	var cmd string
	fields := make([]string, 2)

	if c.Status {
		cmd = "on"
	} else {
		cmd = "off"
	}

	fields[0] = "status:" + cmd
	fields[1] = "timer:" + strconv.Itoa(c.Timer)

	result := strings.Join(fields, ",")

	return result
}

// Send parses and send the command to the tcp socket
func (c *Command) Send() {
	log.Println("CMD:", c.Status, c.Timer)
	cmd := c.ToString()
	socket.ClientQueue <- cmd
}
