package main

import (
	"coffeh/api"
	server "coffeh/bot"
	"coffeh/config"
	"coffeh/db"
	"context"
	"log"
	"os"
	"os/signal"
)

func main() {
	config := config.LoadEnv()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	dbClient, disconnect := db.NewStore()
	defer disconnect()

	bot, err := server.NewBot(dbClient, config)
	if err != nil {
		log.Fatalf("Failed to start bot: %v", err)
	}

	apiServer, err := api.NewApiServer(dbClient, config)
	if err != nil {
		log.Fatalf("Failed to start API server: %v", err)
	}

	go func() {
		err = apiServer.Start()
		if err != nil {
			log.Fatalf("Failed to start API server: %v", err)
		}
	}()

	log.Println("🔥 Hello from Coffe(EH) bot")
	bot.Start(ctx)
}
