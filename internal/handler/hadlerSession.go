package handler

import (
	"github.com/gin-gonic/gin"
	"hotel/internal/models"
	"net/http"
	"strconv"
)

func (h *Handler) GetSessionById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "id param is not correct")
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
		newErrorResponse(c, http.StatusBadRequest, err.Error())
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

func (h *Handler) GetAllSessions(c *gin.Context) {
	session, _ := h.services.GetAllSessions()
	c.JSON(http.StatusOK, session)
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
