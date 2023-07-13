package main

import (
	v1 "air-line-reservation-backend/api/example/v1"
	"air-line-reservation-backend/dbs"
	"flag"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

var (
	laddr = flag.String("addr", ":8880", "Local address for the HTTP API")
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	err := dbs.IntializeDatabase()
	if err != nil {
		log.Fatal(err)
	}

    fmt.Println("Hello, world.")
	
    api := v1.NewExampleRESTApi()
	api.Serve(*laddr)
}