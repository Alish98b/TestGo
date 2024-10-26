package repositories

import (
	"hotel/internal/models"
)

type Storage struct {
	Films map[int]models.Film
}

func NewStorage() *Storage {
	films := make(map[int]models.Film)
	//films[1] = models.Film{Id: 1, Description: "100", Details: "Single", Genre: "ASDASDASDAS", Year: time.Time{}}
	return &Storage{
		Films: films,
	}
}

func (s *Storage) CreateFilm(film models.Film) (int, string) { //пока возвращаем стринг так как мы сами создаем ошибку и отправляем ввиду строки
	if _, exist := s.Films[film.Id]; exist {
		err := "Can not create film, film already exist"
		return 0, err
	}
	s.Films[film.Id] = film
	return film.Id, ""
}
