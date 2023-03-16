package main

import (
	"log"

	"github.com/FFH255/CW-2023/internal/apiserver"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}
}

func main() {

	s := apiserver.New()

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
