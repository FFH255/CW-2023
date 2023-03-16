package main

import (
	"fmt"
	"log"

	"github.com/FFH255/CW-2023/internal/store"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}
}

func main() {
	config := store.NewConfig()
	store := store.New(config)

	if err := store.Open(); err != nil {
		log.Fatal("Can not ping postres")
	}

	fmt.Printf("Connection to postgres established")

}
