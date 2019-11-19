package measurement

import "log"

// Queue is a channel which receives all Senders which
// need to run the Send() function
var Queue chan Sender

const threads = 4

func init() {
	Queue = make(chan Sender, 1024)
	log.Println("Starting", threads, "consumers")
	for i := 0; i < threads; i++ {
		go consume()
	}
}

func consume() {
	log.Println("Consumer up")
	for sndr := range Queue {
		sndr.Send()
	}
}
