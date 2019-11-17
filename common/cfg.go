package common

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

// Cfg is the type which holds all configurations
// variables for the project
type Cfg struct {
	API map[string]interface{}
}

var configPath string // path to json file cfg.json

func init() {
	log.SetFlags(log.Ltime)

	if JSONPath == "" {
		log.Println("JSON path not set")
		relativePath := filepath.Dir(os.Args[0])
		absolutePath, err := filepath.Abs(relativePath)
		if err != nil {
			log.Fatal(err)
		}
		JSONPath = path.Join(absolutePath, "cfg.json")
	}
}

// LoadConfiguration loads the cfg.json file
// and returns its data as a Cfg type
func LoadConfiguration() *Cfg {
	cfg, err := ioutil.ReadFile(JSONPath)
	if err != nil {
		log.Fatal(err)
	}

	var config Cfg
	err = json.Unmarshal(cfg, &config)
	if err != nil {
		log.Fatal(err)
	}

	return &config
}
