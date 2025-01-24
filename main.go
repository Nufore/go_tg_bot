package main

import (
	tg_client "go_tg_bot/clients/telegram"
	event_consumer "go_tg_bot/consumer/event-consumer"
	tg_proc "go_tg_bot/events/telegram"
	"go_tg_bot/storage/files"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	storagePath = "storage"
	batchSize   = 100
)

func main() {
	host := getHost()
	token := getToken()

	tgClient := tg_client.New(host, token)

	eventsProcessor := tg_proc.New(tgClient, files.New(storagePath))

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
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
