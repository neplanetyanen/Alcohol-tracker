package handler

import (
	todo "BackDOM"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createDrinkRecord(c *gin.Context) {
	idStr := c.Param("id")

	// Пытаемся преобразовать строку в целое число
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var input todo.DrinkRecord

	input.UserId = id

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.DrinkRecords.CreateDrinkRecord(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.services.UpdateLastConsumptionDate(input.UserId, input.Date); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.services.UpdateUserTrackerScore(input.UserId, input.Quantity, input.AlcoholPercentage); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.services.UpdateNumberOfDrinks(input.UserId, input.AlcoholPercentage); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
