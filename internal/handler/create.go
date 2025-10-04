package handler

import (
	"github.com/aabbuukkaarr8/TG-BOT/internal/service"
	"github.com/aabbuukkaarr8/TG-BOT/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) Create(c *gin.Context) {
	var req CreateNotificationRequest
	err := validator.BindJSON(&req, c.Request)
	if err != nil {
		logrus.WithError(err).Warn("Create: invalid request JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.srv.Create(c.Request.Context(), service.CreateNotification{
		TelegramChatID: req.TelegramChatID,
		Message:        req.Message,
		ScheduledAt:    req.ScheduledAt,
	})
	if err != nil {
		logrus.WithError(err).Error("Create: service error")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, "created!")

}
