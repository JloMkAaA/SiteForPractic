package handler

import (
	"net/http"

	"github.com/JloMkAaA/SiteForPractic"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createUser(c *gin.Context) {
	var input SiteForPractic.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Phone_number, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	bearer := token
	// Cookie
	c.SetCookie("token", bearer, 60*60, "/", "localhost", false, true)

	// InternalStorage
	// locstor.SetItem("token", bearer)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Phone_number int    `json:"phone_number" binding:"required"`
	Password     string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := c.Cookie("token")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
