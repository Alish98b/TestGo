package handler

import (
	"github.com/gin-gonic/gin"
	"hotel/internal/models"
	"net/http"
	"strconv"
)

func (h *Handler) GetHallById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		newErrorResponse(c, http.StatusBadRequest, "id parameter is required and must be a positive integer")
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
		newErrorResponse(c, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	// Add field-specific validations
	if input.Name == "" {
		newErrorResponse(c, http.StatusBadRequest, "Field 'Name' is required")
		return
	}
	if input.Capacity <= 0 {
		newErrorResponse(c, http.StatusBadRequest, "Field 'Capacity' is required and must be greater than 0")
		return
	}
	//existingHall, err := h.services.GetHallByNmae(input.Name)
	//if err == nil && existingHall != nil {
	//	newErrorResponse(c, http.StatusConflict, "A hall with the same name already exists")
	//	return
	//}

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
	c.JSON(http.StatusOK, h.services.GetAllHalls())
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
