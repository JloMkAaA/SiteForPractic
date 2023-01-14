package handler

import (
	"github.com/JloMkAaA/SiteForPractic/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		user := api.Group("/user")
		{
			user.POST("/", h.createUser)
			user.GET("/", h.getAllUsers)
			user.GET("/:id", h.getUserById)
			user.PUT("/:id", h.updateUser)
			user.DELETE("/:id", h.deleteUser)
		}
	}
	return router
}
