package main

import (
	"go_tg_bot/clients/telegram"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	host := getHost()
	token := getToken()

	tgClient := telegram.New(host, token)

	// fetcher = fetcher.New()

	// processor = processor.New()

	// consumer.Start(fetcher, processor)
}

func getHost() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("token(): Error loading .env file.")
	}

	host := os.Getenv("HOST")
	return host
}

func getToken() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("token(): Error loading .env file.")
	}

	token := os.Getenv("TOKEN")
	return token
}
