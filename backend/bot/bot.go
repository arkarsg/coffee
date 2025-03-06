package bot

import (
	"coffeh/config"
	"coffeh/db"
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type CoffeeBot struct {
	config *config.Env
	bot    *bot.Bot
	store  *db.Store
}

func NewBot(store *db.Store, config *config.Env) (*CoffeeBot, error) {
	coffeeBot := &CoffeeBot{store: store, config: config}

	opts := []bot.Option{
		bot.WithDefaultHandler(coffeeBot.defaultHandler),
	}

	b, err := bot.New(coffeeBot.config.TelegramToken, opts...)
	if err != nil {
		return nil, err
	}
	coffeeBot.bot = b
	return coffeeBot, nil
}

func (s *CoffeeBot) Start(ctx context.Context) {
	s.registerStart()
	s.registerOrder()
	s.bot.Start(ctx)
}

func (s *CoffeeBot) defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Hello from Coffe(EH) bot! Type /start to get started",
		})
	}
}
