package handler

import (
	"github.com/gin-gonic/gin"
	"hotel/internal/models"
	"net/http"
)

func (h *Handler) CreateFilm(c *gin.Context) {
	var input models.Film

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateFilm(input)
	if err != "" {
		newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
