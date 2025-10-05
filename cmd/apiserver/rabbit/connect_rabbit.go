package rabbit

import (
	"github.com/aabbuukkaarr8/TG-BOT/internal/rabbitmq"
	"github.com/wb-go/wbf/zlog"
)

func ConnectRabbit() {
	rabbit, err := rabbitmq.New("amqp://guest:guest@localhost:5672/")
	if err != nil {
		zlog.Logger.Fatal().Err(err).Msg("Failed to connect to RabbitMQ")
	}
	defer rabbit.Close()
}
