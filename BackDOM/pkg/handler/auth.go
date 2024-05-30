package handler

import (
	todo "BackDOM"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Создаем экземпляр структуры User с заполненным ID
	userResponse := todo.UserResponse{
		ID:          id,
		Username:    input.Username,
		Email:       input.Email,
		Nationality: input.Nationality,
		Height:      input.Height,
		Weight:      input.Weight,
		BodyType:    input.BodyType,
		Allergies:   input.Allergies,
		Goal:        input.Goal,
	}

	if err := h.services.Authorization.CreateUserStatistics(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	// Используем структуру Response для отправки ответа
	c.JSON(http.StatusOK, todo.Response{
		User:  userResponse,
		Token: token,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userId, err := h.services.User.GetId(input.Username)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.services.User.GetById(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	userResponse := todo.UserResponse{
		ID:          userId,
		Username:    input.Username,
		Email:       user.Email,
		Nationality: user.Nationality,
		Height:      user.Height,
		Weight:      user.Weight,
		BodyType:    user.BodyType,
		Allergies:   user.Allergies,
		Goal:        user.Goal,
	}
	c.JSON(http.StatusOK, todo.Response{
		User:  userResponse,
		Token: token,
	})
}
