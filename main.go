package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	token := "5991212255:AAEElsQFxXumVZ5rmDHSf8hKKijJc5AQcUg"
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			if strings.Contains(update.Message.Text, "Ready") {
				message, err := os.ReadFile("template/GetReady.txt")
				if err != nil {
					log.Fatal(err)
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(message))
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
			} else if strings.Contains(update.Message.Text, "Exit") {
				message, err := os.ReadFile("template/Exit.txt")
				if err != nil {
					log.Fatal(err)
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(message))
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
			} else {
				welmessage := fmt.Sprintf("Halo %s, a", update.Message.From.UserName)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, welmessage)
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
			}
		}
	}
}
