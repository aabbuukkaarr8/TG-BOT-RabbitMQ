package tg_bot

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func sendToAPI(chatID int64, message string, scheduledTime time.Time) {
	// 1. Формируем JSON
	requestBody := fmt.Sprintf(`{
        "telegram_chat_id": %d,
        "message": "%s",
        "scheduled_at": "%s"
    }`, chatID, message, scheduledTime.Format(time.RFC3339))

	// 2. Отправляем POST на API
	resp, err := http.Post(
		"http://localhost:8080/notify",
		"application/json",
		strings.NewReader(requestBody),
	)
	if err != nil {
		log.Printf("API error: %v", err)
		return
	}
	defer resp.Body.Close()

	// 3. Проверяем ответ
	if resp.StatusCode != http.StatusCreated {
		log.Printf("API bad status: %d", resp.StatusCode)
	}
}

func deleteNotification(chatID int64, id string) {
	// HTTP DELETE запрос к  API
	req, _ := http.NewRequest("DELETE", "http://localhost:8080/notify/"+id, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Delete error: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		// Успешно удалено
	}
}

func getStatus(chatID int64, id string) string {

	resp, err := http.Get("http://localhost:8080/notify/" + id)
	if err != nil {
		log.Printf("Status error: %v", err)
		return ""
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	str := string(body)

	return str

}
