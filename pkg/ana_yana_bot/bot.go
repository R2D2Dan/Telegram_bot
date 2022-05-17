package bot

import (
	config "Telegram_bot/pkg/configs"

	t_bot_api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot      *t_bot_api.BotAPI
	messages config.Messages
	api      string
}
