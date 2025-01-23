package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	token := getToken()
	fmt.Printf("Bot token = %s\n", token)
	// tgClient = telegram.New(token)

	// fetcher = fetcher.New()

	// processor = processor.New()

	// consumer.Start(fetcher, processor)
}

func getToken() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("token(): Error loading .env file.")
	}

	token := os.Getenv("TOKEN")
	return token
}
