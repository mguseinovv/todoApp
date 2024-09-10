package main

import (
	todo "github.com/mguseinovv/todoApp"
	"log"
)

func main() {
	server := new(todo.Server)
	if err := server.Run(8000); err != nil {
		log.Fatalf("Failed to start server %s", err.Error())
	}
}
