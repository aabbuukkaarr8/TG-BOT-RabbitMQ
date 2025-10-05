package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wb-go/wbf/zlog"
)

func (h *Handler) Status(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zlog.Logger.Error().Err(err).Msg("[strconv.Atoi] Your ID is not a number")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	status, err := h.srv.Status(c.Request.Context(), id)
	if err != nil {
		zlog.Logger.Error().Err(err).Msg("[h.Status] Cannot get status")
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
