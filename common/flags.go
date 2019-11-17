package common

import (
	"flag"
	"strings"
)

// JSONPath is the variable holding the filesystem
// path to the cfg.json
var JSONPath string

// SetFlags sets the command line interface flags
func SetFlags() {
	flag.StringVar(&JSONPath, "cfg", "./cfg.json", "Path to the cfg.json file")
	flag.Parse()
	fixFlags()
}

func fixFlags() {
	JSONPath = strings.Trim(JSONPath, " ")
}
