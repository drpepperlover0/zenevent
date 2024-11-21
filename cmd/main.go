package main

import (
	"log"

	"github.com/drpepperlover0/internal/server"
)

func main() {

	s := server.NewServer()

	if err := s.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
