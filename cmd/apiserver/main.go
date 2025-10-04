package main

import (
	"encoding/json"
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/aabbuukkaarr8/TG-BOT/cmd/apiserver/rabbit"
	"github.com/aabbuukkaarr8/TG-BOT/cmd/apiserver/tg_bot"
	"github.com/aabbuukkaarr8/TG-BOT/internal/apiserver"
	"github.com/aabbuukkaarr8/TG-BOT/internal/handler"
	"github.com/aabbuukkaarr8/TG-BOT/internal/models"
	"github.com/aabbuukkaarr8/TG-BOT/internal/rabbitmq"
	"github.com/aabbuukkaarr8/TG-BOT/internal/repository"
	"github.com/aabbuukkaarr8/TG-BOT/internal/service"
	"github.com/aabbuukkaarr8/TG-BOT/internal/store"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"time"
)

var srv *service.Service

var (
	configPath string
)

func main() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
	flag.Parse()
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	db := store.New()
	err = db.Open(config.Store.DatabaseURL)
	if err != nil {
		log.Fatal(err)
		return
	}
	rabbit.ConnectRabbit()
	rab, _ := rabbitmq.New("amqp://guest:guest@localhost:5672/")
	bot, err := tg_bot.Init()

	//repo
	repo := repository.NewRepository(db)
	//service
	srv = service.NewService(repo, rab)
	//handler
	handler := handler.NewHandler(srv)

	s := apiserver.New(config)
	s.ConfigureRouter(handler)
	go startWorker(rab, bot)

	if err = s.Run(); err != nil {
		panic(err)
	}

}

func startWorker(rabbit *rabbitmq.Client, bot *tgbotapi.BotAPI) {
	log.Println("Starting worker")
	messages, _ := rabbit.ConsumeNotifications()

	for msg := range messages {
		// Парсим сообщение из RabbitMQ
		var notification models.Notification
		json.Unmarshal(msg.Body, &notification)
		status, err := srv.Status(notification.ID)
		if err != nil {
			log.Println(err)
		}

		// Проверяем время
		if time.Now().Before(notification.ScheduledAt) {
			msg.Nack(false, true) // возвращаем в очередь
			continue
		}
		//удаляем из очереди удаленные
		if status == "deleted" {
			msg.Ack(false)
		}
		//проверяем статус
		if status == "scheduled" {
			// Отправляем в Telegram
			err := sendToTelegram(bot, notification.TelegramChatID, notification.Message)
			if err != nil {
				// Ошибка - возвращаем в очередь
				msg.Nack(false, true)
			} else {
				msg.Ack(false)
				srv.Sent(notification.ID)
			}

		}

	}
}

func sendToTelegram(bot *tgbotapi.BotAPI, chatID int64, message string) error {
	msg := tgbotapi.NewMessage(chatID, message)
	_, err := bot.Send(msg)
	return err
}
