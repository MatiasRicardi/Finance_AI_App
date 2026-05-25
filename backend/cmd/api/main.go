package main

import (
	"log"
	"os"
)

func main() {
	//run function is defined in cmd/api/server.go and is responsible for starting the API server
	if err := run(); err != nil {
		log.Printf("App stopped with error: %v", err)
		os.Exit(1)
	}
}
