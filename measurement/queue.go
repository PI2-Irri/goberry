package measurement

// Queue is a channel which receives all Senders which
// need to run the Send() function
var Queue chan Sender

func init() {
	Queue = make(chan Sender, 1024)
	go consume()
}

func consume() {
	for sndr := range Queue {
		sndr.Send()
	}
}
