package main

import (
	"log"

	"github.com/ssenthilnathan3/rpxy/internal/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatalf("could not start the server: %v", err)
	}
}
