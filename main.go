package main

import (
	"log"

	"github.com/PI2-Irri/goberry/api"
	"github.com/PI2-Irri/goberry/common"
	"github.com/PI2-Irri/goberry/controller"
)

func init() {
	log.SetFlags(log.Ltime)
}

func main() {
	log.Println("Starting GoBerry")

	common.SetFlags()

	// 1. do controller sync stuff
	// Creates the API object and logs in
	api := api.Create()
	api.Login()
	// Creates the Controller
	ctr := controller.Create(api)
	log.Println(ctr)
	// 2. start polling the api for commands
	// 3. start websocket for listening elet
}
