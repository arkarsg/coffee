package main

import (
	server "coffeh/bot"
	"coffeh/db"
	"context"
	"log"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	dbClient, disconnect := db.NewStore()
	defer disconnect()

	server, err := server.NewBot(dbClient)
	if err != nil {
		log.Fatalf("Failed to start bot: %v", err)
	}

	log.Println("🔥 Hello from Coffe(EH) bot")
	server.Start(ctx)
}
