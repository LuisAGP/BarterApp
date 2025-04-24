package main

import (
	"log"
	"os"

	"github.com/LuisAGP/BarterApp/app/routes"
	"github.com/LuisAGP/BarterApp/database"
)

func main() {
	// Catch every command from the command line
	if len(os.Args) > 1 {
		// For now only migrations
		if os.Args[1] == "migrate" {
			database.AutoMigrate()
			log.Output(1, "migration successfully!")
			os.Exit(0)
		}
	}

	r := routes.API()
	r.Run("0.0.0.0:8080")
}
