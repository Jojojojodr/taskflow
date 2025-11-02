package main

import (
	"flag"
	"log"
	"os"

	"github.com/Jojojojodr/taskflow/api"
	"github.com/Jojojojodr/taskflow/api/database"
	"github.com/Jojojojodr/taskflow/api/internal"
)

func main() {
	var port = flag.String("p", "", "Port to run the server on (e.g., 8080)")
	var dbType = flag.String("d", "", "Database type: sqlite or postgres")
	flag.Parse()

	selectedDBType := checkFlag(*dbType, "DB_TYPE", "Database type not specified via -d flag or DB_TYPE environment variable")

	log.Println("Initializing database...")
	db := database.InitDB(selectedDBType)
	database.SetDB(db)

	selectedPort := checkFlag(*port, "PORT", "Port not specified via -p flag or PORT environment variable")

	log.Println("Initializing server...")
	api.StartServer(selectedPort)
}

func checkFlag(str, env, msg string) string {
	selected := str
	if selected == "" {
		selected = internal.Env(env)
		if selected == "" {
			log.Println(msg)
			log.Println("Usage: ./app -p 8080 -d sqlite")
			os.Exit(1)
		}
	}
	return selected
}
