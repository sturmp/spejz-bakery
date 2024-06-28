package configuration

import (
	"encoding/json"
	"log"
	"os"
)

const appConfigFileName = "appconfig.json"

type Configuration struct {
	Email struct {
		From string
		To   string
		Smtp struct {
			Host string
			Port int
			User string
			Pass string
		}
	}
}

var AppConfig Configuration

func LoadConfiguration() {
	file, err := os.Open(appConfigFileName)
	if err != nil {
		log.Fatalf("Failed to open %s: %s", appConfigFileName, err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&AppConfig); err != nil {
		log.Fatalf("Failed to decode %s: %s", appConfigFileName, err)
	}
}