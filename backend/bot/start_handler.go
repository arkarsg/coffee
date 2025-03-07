package bot

import (
	"context"
	"encoding/json"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (s *CoffeeBot) registerStart() {
	s.bot.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, s.StartHandler)
	s.bot.RegisterHandler(bot.HandlerTypeMessageText, "/whoami", bot.MatchTypeExact, s.WhoamiHandler)
	s.bot.RegisterHandler(bot.HandlerTypeCallbackQueryData, "button", bot.MatchTypePrefix, s.CallbackHandler)

}

func (s *CoffeeBot) StartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	s.store.CreateUser(ctx, *update.Message.From)

	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "View my orders", CallbackData: "view_orders"},
				{Text: "Opening hours", CallbackData: "button_opening_hours"},
			},
			{
				{Text: "Order", CallbackData: "button_order"},
			},
		},
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "☕️ Welcome to Coffe(EH)!",
		ReplyMarkup: kb,
	})
}

func (s *CoffeeBot) WhoamiHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	log.Printf("%d %s", update.Message.From.ID, update.Message.From.Username)
	user, err := s.store.FindUserByTelegramID(ctx, update.Message.From.ID)

	if err != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID:      update.Message.Chat.ID,
			Text:        "Could not get you",
			ReplyMarkup: models.ParseModeMarkdown,
		})
	}

	jsonBytes, err := json.Marshal(user)

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   string(jsonBytes),
	})
}

func (s *CoffeeBot) CallbackHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	})

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.CallbackQuery.Message.Message.Chat.ID,
		Text:   "Selected " + update.CallbackQuery.Data,
	})
}
