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

	movie := router.Group("/movie", h.userIdentity)
	{
		movie.GET("/:id", h.GetMovieById)
		movie.POST("/", h.CreateMovie)
		movie.PUT("/:id", h.UpdateMovie)
		movie.GET("/", h.GetAllMovies)
		movie.DELETE("/:id", h.DeleteMovie)
	}

	user := router.Group("/user")
	{
		user.GET("/:id", h.GetUserById)
		user.POST("/", h.CreateUser)
		user.PUT("/:id", h.UpdateUser)
		user.GET("/", h.GetAllUsers)
		user.DELETE("/:id", h.DeleteUser)
		user.POST("/login", h.Login)
		//TODO: сделать метод который будет запрашивать пользователя по email
	}

	//
	return router
}
