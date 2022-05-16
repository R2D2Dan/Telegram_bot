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

	ana_yanabot.Debug = true

	log.Printf("Бот - %s: вкл", ana_yanabot.Self.UserName)

	updates := ana_yanabot.GetUpdatesChan(updateconfig)
	msg := t_bot_api.NewMessage(0, "")
	for update := range updates {

		if !update.Message.IsCommand() && update.Message.Text != "" {
			msg.Text = update.Message.Text
			msg.ChatID = update.Message.Chat.ID

		} else if update.Message.IsCommand() {
			msg.ChatID = update.Message.Chat.ID
			switch update.Message.Command() {
			case "start":
				msg.Text = "Привет, меня зовут ana_yanabot 🖖.\nМой создатель @R2D2Dan, он дал мне задачу облегчать жизнь тебе🙂\nВ будущем я смогу много чего делать.\nЕсли ты хочешь узнать что я умею, введи команду /help и я покажу тебе свои возможности😏"
			case "help":
				msg.Text = "Мои возможности:\n1. Я умею предсказывать погоду, команда(/weather) 🌤\n2. Пока все😂  А как ты хотел он только сегодня меня создал☺️\nЕсли у тебя есть идее то чего не хватает мне для совершенства, то напиши мое создателю, он подумает над твоим предложением"
			case "weather":
				msg.Text = `Мой создатель исправляет пока еще не запустил😔,но в скором времени все будет работать, нужно немного подождать😇`
			default:
				msg.Text = `Я не знаю такой команды`
			}
		} else {
			continue
		}

		if _, err := ana_yanabot.Send(msg); err != nil {
			log.Println("Error send message", err)
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
