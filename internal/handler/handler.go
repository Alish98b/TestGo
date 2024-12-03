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
	hall := router.Group("/hall")
	{
		hall.POST("/", h.CreateHall)
		hall.GET("/:id", h.GetHallById)
		hall.GET("/", h.GetAllHalls)
		hall.PUT("/:id", h.UpdateHall)
		hall.DELETE("/:id", h.DeleteHall)

	}
	session := router.Group("/session")
	{
		session.POST("/", h.CreateSession)
		session.GET("/:id", h.GetSessionById)
		session.GET("/", h.GetAllSessions)
		session.PUT("/:id", h.UpdateSession)
		session.DELETE("/:id", h.DeleteSession)
	}
	//
	return router
}
