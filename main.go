package main

import (
	"log"

	"github.com/PI2-Irri/goberry/api"
	"github.com/PI2-Irri/goberry/common"
)

func init() {
	log.SetFlags(log.Ltime)
}

func main() {
	log.Println("Starting GoBerry")

	common.SetFlags()

	// 1. do controller sync stuff
	api := api.Create()
	api.Login()
	api.GetControllers()
	// 2. start polling the api for commands
	// 3. start websocket for listening elet
}
