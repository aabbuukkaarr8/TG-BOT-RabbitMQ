package tg_bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

func Init() (*tgbotapi.BotAPI, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		return nil, err
	}

	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)

	go func() {
		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60
		updates := bot.GetUpdatesChan(u)

		for update := range updates {
			if update.Message == nil {
				continue
			} else if strings.HasPrefix(update.Message.Text, "/remind") { // If we got a message
				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
				parts := strings.SplitN(update.Message.Text, " ", 5)
				if len(parts) != 5 {
					bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "wrong format"))
					continue
				}
				timePart := strings.Join(parts[1:4], "")
				text := parts[4]
				scheduledTime := parseTime(timePart)
				if scheduledTime.IsZero() {
					bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "не понимаю время. Используй : через 2 часа, через 30 минут"))
					continue
				}
				go sendToAPI(update.Message.Chat.ID, text, scheduledTime)
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Напоминание установлено!"))
			} else if strings.HasPrefix(update.Message.Text, "/delete") {
				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
				parts := strings.SplitN(update.Message.Text, " ", 2)
				if len(parts) != 2 {
					bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Формат: /delete [id]"))
					continue
				}

				id := parts[1] // "1"

				// Отправляем DELETE запрос API
				go deleteNotification(update.Message.Chat.ID, id)
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Удаляем напоминание..."))
			} else if strings.HasPrefix(update.Message.Text, "/status") {
				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
				parts := strings.SplitN(update.Message.Text, " ", 2)
				if len(parts) != 2 {
					bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Формат: /status [id]"))
					continue
				}

				id := parts[1] // "1"

				// Отправляем GET запрос  API
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Проверяем статус..."))
				go func(chatID int64, id string) {
					status := getStatus(chatID, id) // синхронно в горутине
					bot.Send(tgbotapi.NewMessage(chatID, status))
				}(update.Message.Chat.ID, id)

			}
		}
	}()
	return bot, nil
}
