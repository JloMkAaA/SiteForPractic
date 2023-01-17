package handler

import (
	"net/http"
	"strconv"

	"github.com/JloMkAaA/SiteForPractic"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getUserById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if userId == id {
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "invalid id param")
			return
		}

		user, err := h.services.GetUserById(id)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, user)
	}

}

func (h *Handler) updateUser(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if userId == id {

		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "invalid id param")
			return
		}

		var input SiteForPractic.UpdateUser
		if err := c.BindJSON(&input); err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		if err := h.services.Update(userId, input); err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, statusResponse{
			Status: "ok",
		})
	}

}

func (h *Handler) deleteUser(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if userId == id {
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "invalid id param")
			return
		}

		err = h.services.Delete(userId)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, statusResponse{
			Status: "Ok",
		})
	}

}
