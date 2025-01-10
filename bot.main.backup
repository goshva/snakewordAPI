
package main

import (
	"os"
	"log"
"github.com/joho/godotenv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	// Replace with the chat ID to send the message
	chatID := int64(190404167)

	// Create the InlineKeyboardMarkup with the web app button
	webAppButton := tgbotapi.NewInlineKeyboardButtonURL("Запустить SnakeWord", "https://snakeword.ru/")
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(webAppButton),
	)

	// Create the message
	message := tgbotapi.NewMessage(chatID, "Играй в SnakeWord прямо сейчас!")
	message.ReplyMarkup = keyboard

	// Send the message
	if _, err := bot.Send(message); err != nil {
		log.Panic(err)
	}
}


