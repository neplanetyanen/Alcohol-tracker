package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getUserStatisticsById(c *gin.Context) {
	idStr := c.Param("id")

	// Пытаемся преобразовать строку в целое число
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Получаем данные пользователя из базы данных по ID
	userStatistics, err := h.services.UserStatistics.GetUserStatisticsById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Возвращаем данные пользователя в ответе
	c.JSON(http.StatusOK, userStatistics)
}
