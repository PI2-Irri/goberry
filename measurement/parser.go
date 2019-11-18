package measurement

import (
	"log"
	"strings"
)

// ParseMessage parses a message as string and returns
// the proper Sender
func ParseMessage(msg string) Sender {
	list := strings.Split(msg, ",")
	data := make(map[string]string, 5)

	for _, item := range list {
		keyValue := strings.Split(item, ":")
		data[keyValue[0]] = keyValue[1]
	}

	// for k, v := range data {
	// 	log.Println("\t", k, ":", v)
	// }

	if data["type"] == "0" {
		return CreateActuator(data)
	} else if data["type"] == "1" {
		return CreateModule(data)
	}

	log.Fatal("Type", data["type"], "not recognized")
	return nil
}
