package main

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

//Settings struct stores global server settings
type Settings struct {
	Host string
	Port string
}

func (s *Settings) serverAddr() string {
	return s.Host + ":" + s.Port
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
	err = enc.Encode(Settings{Host: "0.0.0.0", Port: "3000"})
	if err != nil {
		log.Fatal("Could not encode json to settings.json: ", err)
	}
}

func loadDB() *sqlx.DB {
	_, err := os.Stat("./lxgg.db")
	genTables := os.IsNotExist(err)

	db, err := sqlx.Open("sqlite3", "./lxgg.db")
	if err != nil {
		log.Fatal("Could not load DB: ", err)
	}

	//Database file wasn't found, generate tables
	if genTables {
		fmt.Println("DB not found, generating tables")

		f, err := os.Open("./schema.sql")
		if err != nil {
			log.Fatal("Couldn't open schema.sql, which is needed to generate tables: ", err)
		}

		buf := new(bytes.Buffer)
		buf.ReadFrom(f)

		schema := buf.String()
		db.MustExec(schema)

		//Generate default user
		username := "lxgg"
		password := generateRandomPW()

		fmt.Println("Generating default admin user:")
		fmt.Printf("Username: %s\nPassword: %s\n", username, password)

		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			deleteDBFile()
			log.Fatal("Could not generate from hash: ", err)
		}

		_, err = db.Exec(`INSERT INTO users (username, password, admin) VALUES (?, ?, ?)`, username, string(hash), true)

		if err != nil {
			deleteDBFile()
			log.Fatal("Error inserting default user: ", err)
		}
	}

	return db
}

func generateRandomPW() string {
	bytes := make([]byte, 10)
	_, err := rand.Read(bytes)

	if err != nil {
		deleteDBFile()
		log.Fatal("Error generating random password, delete lxgg.db and try again: ", err)
	}

	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"

	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}

	return string(bytes)
}

func deleteDBFile() {
	err := os.Remove("./lxgg.db")
	if err != nil {
		fmt.Println("Error deleting lxgg.db, delete this file before running LXGG again. message: ", err)
	}
}
