package handler

import (
	todo "BackDOM"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getUserById(c *gin.Context) {
	idStr := c.Param("id")

	// Пытаемся преобразовать строку в целое число
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Получаем данные пользователя из базы данных по ID
	user, err := h.services.User.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Возвращаем данные пользователя в ответе
	c.JSON(http.StatusOK, user)
}

func (h *Handler) updateUser(c *gin.Context) {
	var input todo.UserResponse

	idStr := c.Param("id")

	// Пытаемся преобразовать строку в целое число
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	input.ID = id

	// Парсим JSON из тела запроса
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	// Обновляем данные пользователя в базе данных по ID
	updatedUser, err := h.services.User.Update(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userResponse := todo.UserResponse{
		ID:          updatedUser.ID,
		Username:    updatedUser.Username,
		Email:       updatedUser.Email,
		Nationality: updatedUser.Nationality,
		Height:      updatedUser.Height,
		Weight:      updatedUser.Weight,
		BodyType:    updatedUser.BodyType,
		Allergies:   updatedUser.Allergies,
		Goal:        updatedUser.Goal,
	}
	c.JSON(http.StatusOK, userResponse)
}
