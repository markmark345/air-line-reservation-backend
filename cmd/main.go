package main

import (
	v1 "air-line-reservation-backend/api/example/v1"
	"air-line-reservation-backend/dbs"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	laddr = flag.String("addr", ":8880", "Local address for the HTTP API")
)

func main() {

	// Load .env from current working directory
	curDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	loadErr := godotenv.Load(curDir + "/.env")
	if loadErr != nil {
		log.Fatalln("can't load env file from current directory: " + curDir)
	}

	// Auto load .env -> Caused an error when debugging
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// Initializing Database connections/Instances
	dbErr := dbs.IntializeDatabase()
	if dbErr != nil {
		log.Fatal(dbErr)
	}

    fmt.Println("Hello, world.")
	
	// Instantiate & Initialize Rest API Router
    api := v1.NewExampleRESTApi()

	// Start the configured application
	api.Serve(*laddr)
}