package handler

import (
	todo "BackDOM"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getNormResult(c *gin.Context) {
	idStr := c.Param("id")

	// Пытаемся преобразовать строку в целое число
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var input todo.Norm
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.services.CalculateUserNorm(id, input.AlcoholPercentage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}

func (h *Handler) getTrackerScore(c *gin.Context) {
	idStr := c.Param("id")

	// Пытаемся преобразовать строку в целое число
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	trackerScore, err := h.services.User.GetUserTrackerScore(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tracker_score": trackerScore})
}
