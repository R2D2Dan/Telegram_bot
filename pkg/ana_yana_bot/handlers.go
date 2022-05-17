package bot

import t_bot_api "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const (
	commandStart   = "start"
	commandHelp    = "help"
	commandWeather = "weather"
)

func (b *Bot) handleCommand(message *t_bot_api.Message) error {
	switch message.Command() {
	case commandStart:

	}
	return nil
}

func (b *Bot) handleStartCommand(message *t_bot_api.Message) error {
	return nil
}
