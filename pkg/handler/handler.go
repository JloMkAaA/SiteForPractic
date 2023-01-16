package handler

import (
	"github.com/JloMkAaA/SiteForPractic/pkg/service"
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
		auth.POST("/createUser", h.createUser)
		auth.GET("/sign-in", h.signIn)

	}

	user := router.Group("/user", h.userIdentity)
	{
		user.GET("/:id", h.getUserById)
		user.PUT("/:id", h.updateUser)
		user.DELETE("/:id", h.deleteUser)
	}

	return router
}
