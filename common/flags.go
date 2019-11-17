package common

import (
	"flag"
	"log"
	"strings"
)

var (
	// JSONPath is the variable holding the filesystem path to the cfg.json
	JSONPath string
	// Pin is the variable holding the controller's pin code
	Pin string
)

// SetFlags sets the command line interface flags
func SetFlags() {
	flag.StringVar(&JSONPath, "cfg", "./cfg.json", "Path to the cfg.json file")
	flag.StringVar(&Pin, "pin", "", "Pin code for the controller")
	flag.Parse()
	fixFlags()
}

func fixFlags() {
	JSONPath = strings.Trim(JSONPath, " ")
	Pin = strings.Trim(Pin, " ")
	if Pin == "" {
		log.Fatal("Pin not provided through -pin flag")
	}
}
