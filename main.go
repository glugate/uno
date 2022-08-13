package main

import (
	"log"

	cmd "github.com/glugate/uno/cmd/uno"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load .env file")
	}
	cmd.Execute()
}
