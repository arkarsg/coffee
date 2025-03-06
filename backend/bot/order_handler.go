package bot

import (
	"context"
	"encoding/json"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (s *CoffeeBot) registerOrder() {
	s.bot.RegisterHandler(
		bot.HandlerTypeCallbackQueryData,
		"view_orders",
		bot.MatchTypeExact,
		s.ViewOrderHandler,
	)
}

func (s *CoffeeBot) ViewOrderHandler(ctx context.Context, b *bot.Bot, update *models.Update) {

	telegramId := update.CallbackQuery.From.ID
	orders, err := s.store.FindAllOrdersByTelegramId(telegramId)
	if err != nil {
		log.Printf("Could not retrieve orders %s", err.Error())
		b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
			CallbackQueryID: update.CallbackQuery.ID,
			ShowAlert:       false,
		})
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.CallbackQuery.Message.Message.Chat.ID,
			Text:   "No orders found",
		})
	}

	// Format the orders instead
	jsonBytes, err := json.Marshal(orders)
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	})
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.CallbackQuery.Message.Message.Chat.ID,
		Text:   string(jsonBytes),
	})

}
