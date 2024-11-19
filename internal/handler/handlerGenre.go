package handler

import (
	"hotel/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateGenre(c *gin.Context) {
    var input models.GenreCreate
    if err := c.BindJSON(&input); err != nil {
        newErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }

    id, err := h.services.Genre.CreateGenre(input)
    if err != nil {
        newErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }

    c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

func (h *Handler) GetGenreById(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        newErrorResponse(c, http.StatusBadRequest, "Invalid genre ID")
        return
    }

    genre, err := h.services.Genre.GetGenreById(id)
    if err != nil {
        newErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }

    c.JSON(http.StatusOK, genre)
}

func (h *Handler) GetAllGenres(c *gin.Context) {
    genres, err := h.services.Genre.GetAllGenres()
    if err != nil {
        newErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }

    c.JSON(http.StatusOK, genres)
}

func (h *Handler) UpdateGenre(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        newErrorResponse(c, http.StatusBadRequest, "Invalid genre ID")
        return
    }

    var input models.GenreCreate
    if err := c.BindJSON(&input); err != nil {
        newErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }

    if err := h.services.Genre.UpdateGenre(id, input); err != nil {
        newErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }

    c.JSON(http.StatusOK, map[string]interface{}{"message": "Genre updated successfully"})
}

func (h *Handler) DeleteGenre(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        newErrorResponse(c, http.StatusBadRequest, "Invalid genre ID")
        return
    }

    if err := h.services.Genre.DeleteGenre(id); err != nil {
        newErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }

    c.JSON(http.StatusOK, map[string]interface{}{"message": "Genre deleted successfully"})
}