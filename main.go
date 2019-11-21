package main

import (
	"log"
	"sync"

	"github.com/PI2-Irri/goberry/api"
	"github.com/PI2-Irri/goberry/common"
	"github.com/PI2-Irri/goberry/controller"
	"github.com/PI2-Irri/goberry/measurement"
	"github.com/PI2-Irri/goberry/socket"
)

func init() {
	log.SetFlags(log.Ltime)
}

func main() {
	log.Println("Starting GoBerry")

	common.SetFlags()

	// Creates the API object and logs in
	api := api.Instance()
	api.Login()
	// Creates the Controller
	ctr := controller.Create()
	log.Println(ctr)
	// TODO: Starts controller http polling

	// Creates TCP server socket
	serverSocket := socket.CreateServer()
	// Runs it in another thread
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		serverSocket.AcceptConnections()
		wg.Done()
	}()
	defer close(measurement.Queue)

	// Creates a TCP client socket
	clientSocket := socket.CreateClient()
	// Runs client thread
	wg.Add(1)
	go func() {
		clientSocket.Start()
		wg.Done()
	}()
	defer close(socket.ClientQueue)

	// Starts the HTTP Polling
	wg.Add(1)
	go func() {
		ctr.Poll()
		wg.Done()
	}()

	// Wait for all threads to be finished
	wg.Wait()
}
