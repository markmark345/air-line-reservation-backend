package main

import (
	v1 "air-line-reservation-backend/api/example/v1"
	"flag"
	"fmt"
)

var (
	laddr           = flag.String("addr", ":8880", "Local address for the HTTP API")
)

func main() {
    fmt.Println("Hello, world.")
    api := v1.NewExampleRESTApi()
	api.Serve(*laddr)
}