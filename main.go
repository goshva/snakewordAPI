package main

import (
    "log"
    "os"

    "github.com/joho/godotenv"
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
    webAppURL := "https://snakeword.ru/"

    bot, err := tgbotapi.NewBotAPI(botToken)
    if err != nil {
        log.Fatal(err)
    }

    bot.Debug = true

    u := tgbotapi.NewUpdate(0)
    u.Timeout = 60

    updates := bot.GetUpdatesChan(u)

    for update := range updates {
        if update.Message == nil {
            continue
        }

        if update.Message.IsCommand() {
            switch update.Message.Command() {
            case "start":
                msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome! Use /play to open the web app.")
                bot.Send(msg)
            case "play":
                keyboard := tgbotapi.NewInlineKeyboardMarkup(
                    tgbotapi.NewInlineKeyboardRow(
                        tgbotapi.NewInlineKeyboardButtonURL("Open Web App", webAppURL),
                    ),
                )

                msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Click the button below to open the web app:")
                msg.ReplyMarkup = keyboard
                bot.Send(msg)
            }
        }
    }
}
