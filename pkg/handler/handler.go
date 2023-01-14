package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-Up")
		auth.POST("/sign-Ip")
	}

	api := router.Group("/api")
	{
		user := api.Group("/user")
		{
			user.POST("/")
			user.GET("/")
			user.GET("/:id")
			user.PUT("/:id")
			user.DELETE("/:id")
		}
	}
	return router
}
