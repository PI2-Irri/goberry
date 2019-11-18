package measurement

// Sender is the type which contains a method Send,
// the method responsible for sending data to the API
type Sender interface {
	Send()
}
