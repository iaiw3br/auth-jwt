package main

import (
	"log"
	authError "main/internal/error"
	"main/internal/server"
)

func main() {
	err := server.Run()
	if err != nil {
		log.Fatal(authError.ErrFailedStartServer)
	}
}
