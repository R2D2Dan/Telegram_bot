package bot

import (
	"log"
	"os"

	t_bot_api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	Tbot *t_bot_api.BotAPI
}

func Start() {
	ana_yanabot := get_bot()

	updateconfig := t_bot_api.NewUpdate(0)
	updateconfig.Limit = 30

	updates := ana_yanabot.GetUpdatesChan(updateconfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := t_bot_api.NewMessage(update.Message.Chat.ID, update.Message.Text)

		msg.ReplyToMessageID = update.Message.MessageID

		if _, err := ana_yanabot.Send(msg); err != nil {
			log.Println("Error send message:")
			log.Fatal(err)
		}
	}
}

func get_bot() *t_bot_api.BotAPI {
	bot, err := t_bot_api.NewBotAPI(get_api())
	if err != nil {
		log.Println("Error take bot:", err)
	}

	bot.Debug = true
	return bot
}

func get_api() string {
	f, err := os.ReadFile("./Data/api_key.txt")
	if err != nil {
		log.Println("Error read file:", err)
	}

	return string(f)

}
