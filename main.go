package main

import (
	"log"
	"os"

	bot_api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := bot_api.NewBotAPI(get_api())
	if err != nil {
		log.Println("Error take bot:", err)
	}

	bot.Debug = true
}

func get_api() string {
	f, err := os.ReadFile("./Data/api_key.txt")
	if err != nil {
		log.Println("Error read file:", err)
	}

	return string(f)

}
