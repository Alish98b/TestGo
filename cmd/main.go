package main

import (
	"hotel/internal/handler"
	"hotel/internal/repositories"
	"hotel/internal/services"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	url := "postgresql://cinema_6i7q_user:URS5LXBh4NDZNJUAbSgwJQkZpKKuANxv@dpg-csebeddsvqrc73evunbg-a.frankfurt-postgres.render.com/cinema_6i7q"
	db, err := repositories.NewPostgresDB(url)
	if err != nil {
		logrus.Fatal("Error from db: ", err)
		return
	}

	repository := repositories.CinemaRepo(db)
	service := services.NewServicesCinema(repository)
	handler := handler.NewHandler(service)

	server := &http.Server{
		Addr:    ":5050",
		Handler: handler.InitRoutes(),
	}

	logrus.Info("Server starting on port 5050...")

	if err := server.ListenAndServe(); err != nil {
		logrus.Fatal(err)
	}
}
