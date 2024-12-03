package handler

import (
	"github.com/gin-gonic/gin"
	"hotel/internal/models"
	"net/http"
	"strconv"
)

func (h *Handler) GetHallById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "id param is not correct")
		return
	}

	hall, err := h.services.GetHallById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, hall)
}

func (h *Handler) CreateHall(c *gin.Context) {
	var input models.HallCreate

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateHall(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) GetAllHalls(c *gin.Context) {
	hall := h.services.GetAllHalls()
	c.JSON(http.StatusOK, hall)
}

func (h *Handler) DeleteHall(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "id param is not correct")
		return
	}
	if err := h.services.DeleteHall(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Hall deleted successfully",
	})
}
func (h *Handler) UpdateHall(c *gin.Context) {

	var input models.Hall
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "id param is not correct")
		return
	}
	_, err = h.services.GetHallById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	errUpHall := h.services.UpdateHall(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, errUpHall.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Hall update successfully",
	})
}
