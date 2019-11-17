package main

import (
	"log"

	"github.com/PI2-Irri/goberry/api"
	"github.com/PI2-Irri/goberry/common"
	"github.com/PI2-Irri/goberry/controller"
	"github.com/PI2-Irri/goberry/socket"
)

func init() {
	log.SetFlags(log.Ltime)
}

func main() {
	log.Println("Starting GoBerry")

	common.SetFlags()

	// Creates the API object and logs in
	api := api.Create()
	api.Login()
	// Creates the Controller
	ctr := controller.Create(api)
	log.Println(ctr)
	// TODO: Starts controller http polling

	// Creates TCP socket
	tcpSocket := socket.Create()
	tcpSocket.AcceptConnections()
}
