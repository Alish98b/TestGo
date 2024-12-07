package handler

import (
	"github.com/gin-gonic/gin"
	"hotel/internal/models"
	"net/http"
	"strconv"
)

func (h *Handler) GetSessionById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		newErrorResponse(c, http.StatusBadRequest, "id parameter is required and must be a positive integer")
		return
	}

	session, err := h.services.GetSessionById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, session)
}
func (h *Handler) CreateSession(c *gin.Context) {
	var input models.SessionCreate

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	// Add field-specific validations
	if input.MovieID == 0 {
		newErrorResponse(c, http.StatusBadRequest, "Field 'MovieID' is required and must be greater than 0")
		return
	}
	if input.HallID == 0 {
		newErrorResponse(c, http.StatusBadRequest, "Field 'HallID' is required and must be greater than 0")
		return
	}
	if input.StartTime.IsZero() {
		newErrorResponse(c, http.StatusBadRequest, "Field 'StartTime' is required and must be a valid date")
		return
	}

	id, err := h.services.CreateSession(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

//func (h *Handler) CreateSession(c *gin.Context) {
//	var input models.SessionCreate
//
//	if err := c.BindJSON(&input); err != nil {
//		newErrorResponse(c, http.StatusBadRequest, err.Error())
//		return
//	}
//
//	id, err := h.services.CreateSession(input)
//	if err != nil {
//		newErrorResponse(c, http.StatusInternalServerError, err.Error())
//		return
//	}
//
//	c.JSON(http.StatusOK, map[string]interface{}{
//		"id": id,
//	})
//}

func (h *Handler) GetAllSessions(c *gin.Context) {
	// No specific validations needed for fetching all sessions
	sessions, err := h.services.GetAllSessions()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, sessions)
}

func (h *Handler) DeleteSession(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "id param is not correct")
		return
	}
	if err := h.services.DeleteSession(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Session deleted successfully",
	})
}
func (h *Handler) UpdateSession(c *gin.Context) {

	var input models.Session
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "id param is not correct")
		return
	}
	_, err = h.services.GetSessionById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	errUpHall := h.services.UpdateSession(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, errUpHall.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Session update successfully",
	})
}
