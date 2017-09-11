package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

//Settings struct stores global server settings
type Settings struct {
	Host string
	Port string
}

func loadSettings() Settings {
	f, err := os.Open("settings.json")
	if err != nil {
		generateSettings()
		fmt.Println("Could not load settings.json, default config generated.")
		os.Exit(1)
	}
	dec := json.NewDecoder(f)

	var settings Settings
	err = dec.Decode(&settings)
	if err != nil {
		log.Fatal("Could not decode settings.json: ", err)
	}

	return settings
}

func generateSettings() {
	f, err := os.Create("settings.json")
	if err != nil {
		log.Fatal("Could not create settings.json: ", err)
	}

	enc := json.NewEncoder(f)
	err = enc.Encode(Settings{Host: "0.0.0.0", Port: "80"})
	if err != nil {
		log.Fatal("Could not encode json to settings.json: ", err)
	}
}
