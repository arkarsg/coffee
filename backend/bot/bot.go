package bot

import (
	"coffeh/config"
	"coffeh/db"
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var (
	env = config.LoadEnv()
)

type CoffeeBot struct {
	bot   *bot.Bot
	store *db.Store
}

func NewBot(store *db.Store) (*CoffeeBot, error) {
	coffeeBot := &CoffeeBot{store: store}

	opts := []bot.Option{
		bot.WithDefaultHandler(coffeeBot.defaultHandler),
	}

	b, err := bot.New(env.TelegramToken, opts...)
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
