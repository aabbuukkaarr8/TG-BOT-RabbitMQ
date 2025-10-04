package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (h *Handler) Status(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		logrus.WithError(err).Errorf("[strconv.Atoi] Your ID is not a number")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	status, err := h.srv.Status(id)
	if err != nil {
		logrus.WithError(err).Errorf("[h.Status] Cannot get status")
	}

	if status == "scheduled" {
		status = "готовится к отправке"
	} else if status == "deleted" {
		status = "напоминание удалено"
	} else if status == "sent" {
		status = "напоминание отправлено"
	}

	c.JSON(http.StatusOK, status)

}
