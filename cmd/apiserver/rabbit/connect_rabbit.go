package rabbit

import (
	"github.com/aabbuukkaarr8/TG-BOT/internal/rabbitmq"
	"log"
)

func ConnectRabbit() {
	rabbit, err := rabbitmq.New("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}
	defer rabbit.Close()
}
