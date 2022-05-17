package bot

import (
	"log"
	"os"

	t_bot_api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) Start() {
	u := t_bot_api.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)

	for update := range updates {
		//Ignore if not message
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {

		}

	}

}

func newBot() *Bot {
	b, err := t_bot_api.NewBotAPI(getApi())
	if err != nil {
		log.Println("Error take bot:", err)
	}

	return &Bot{
		bot: b,
	}
}

func getApi() string {
	f, err := os.ReadFile("./Data/api_key.txt")
	if err != nil {
		log.Println("Error read file:", err)
	}

	return string(f)
}
