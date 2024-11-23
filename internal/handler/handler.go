package handler

import (
	"hotel/internal/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *services.ServicesCinema
}

func NewHandler(srv *services.ServicesCinema) *Handler {
	return &Handler{
		services: srv,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// base middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	movie := router.Group("/movie")
	{
		movie.GET("/:id", h.GetMovieById)
		movie.POST("/", h.CreateMovie)
		movie.PUT("/:id", h.UpdateMovie)
		movie.GET("/", h.GetAllMovies)
		movie.DELETE("/:id", h.DeleteMovie)
	}

	genre := router.Group("/genre")
    {
        genre.POST("/", h.CreateGenre)
        genre.GET("/:id", h.GetGenreById)
        genre.GET("/", h.GetAllGenres)
        genre.PUT("/:id", h.UpdateGenre)
        genre.DELETE("/:id", h.DeleteGenre)
    }

	//
	return router
}
