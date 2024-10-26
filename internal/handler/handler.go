package handler

import (
	"hotel/internal/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *services.Service
}

func NewHandler(srv *services.Service) *Handler {
	return &Handler{
		services: srv,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// base middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	room := router.Group("/film")
	{
		room.GET("/:id", h.GetFilmById)
		room.POST("/", h.CreateFilm)
		room.PUT("/:id", h.UpdateFilm)
		room.GET("/", h.GetAllFilms)
		room.DELETE("/:id", h.DeleteFilm)
	}

	return router
}
