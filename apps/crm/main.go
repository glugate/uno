package main

//go:generate sqlboiler --config ./gen.toml --wipe sqlite3

import (
	"fmt"

	"github.com/glugate/uno/pkg/uno"
	"github.com/glugate/uno/pkg/uno/log"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//var menu models.Menu
	//var menuItems []models.MenuItem

	fmt.Println("CRM application!")

	// Load env vars from file
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	// Head up!
	log.Default().Success("You are a superstar!")

	app := uno.NewUno()
	app.Metrics()

	if err := app.Server.Run(); err != nil {
		panic(err)
	}

}
