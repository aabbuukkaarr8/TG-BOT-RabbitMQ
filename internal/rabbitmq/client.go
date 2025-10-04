package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aabbuukkaarr8/TG-BOT/internal/models"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Client обертка для работы с RabbitMQ
type Client struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func New(url string) (*Client, error) {

	conn, err := amqp.Dial(url) // "amqp://guest:guest@localhost:5672/"
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open channel: %w", err)
	}

	err = channel.ExchangeDeclare(
		"notifications_exchange",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare exchange: %w", err)
	}

	// 4. СОЗДАЕМ ОЧЕРЕДЬ - "почтовый ящик" где лежат сообщения
	_, err = channel.QueueDeclare(
		"notifications_queue", // название очереди
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare queue: %w", err)
	}

	err = channel.QueueBind(
		"notifications_queue",    // очередь
		"notification",           // routing key - "адрес" сообщения
		"notifications_exchange", // exchange
		false,                    // no-wait
		nil,                      // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to bind queue: %w", err)
	}

	log.Println("RabbitMQ connected and configured")
	return &Client{conn: conn, channel: channel}, nil
}

// PublishNotification КИДАЕТ сообщение в RabbitMQ
func (c *Client) PublishNotification(ctx context.Context, notification *models.Notification) error {
	// ПРЕОБРАЗУЕМ структуру в JSON
	body, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("failed to marshal notification: %w", err)
	}

	// ПУБЛИКУЕМ сообщение в RabbitMQ
	err = c.channel.PublishWithContext(
		ctx,
		"notifications_exchange", // exchange - куда отправляем
		"notification",           // routing key - "адрес" сообщения
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json", // тип содержимого
			Body:        body,               // само сообщение
			// Сообщение будет persistent - сохранится на диск
			DeliveryMode: amqp.Persistent,
		},
	)

	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	log.Printf("📨 Notification %s published to RabbitMQ", notification.ID)
	return nil
}

// ConsumeNotifications ПОЛУЧАЕТ сообщения из очереди
// Возвращает канал из которого можно читать сообщения
func (c *Client) ConsumeNotifications() (<-chan amqp.Delivery, error) {

	messages, err := c.channel.Consume(
		"notifications_queue",
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to consume messages: %w", err)
	}

	log.Println("👂 Started consuming messages from RabbitMQ")
	return messages, nil
}

// Close закрывает соединение с RabbitMQ
func (c *Client) Close() {
	c.channel.Close()
	c.conn.Close()
}
