package main

import (
	"log"
	"main/kudoserver"
)

func main() {
	if err := kudoserver.SartServer(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
