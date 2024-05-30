package handler

import (
	"BackDOM/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		api.Use(h.userIdentity)
		users := api.Group("/users")
		{
			users.GET("/:id", h.getUserById)
			users.PUT("/:id", h.updateUser)
		}

		userStatistics := api.Group("userStatistics")
		{
			userStatistics.GET("/:id", h.getUserStatisticsById)
		}

		drinkRecords := api.Group("/drinkRecords")
		{
			drinkRecords.POST("/:id", h.createDrinkRecord)
		}

		analytics := api.Group("/analytics")
		{
			analytics.POST("/norm/:id", h.getNormResult)
			analytics.GET("/tracker/:id", h.getTrackerScore)
		}
	}

	return router
}
