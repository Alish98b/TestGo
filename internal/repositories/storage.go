package repositories

import (
	"hotel/internal/models"
	"time"
)

type Storage struct {
	Films map[int]models.Film
}

func NewStorage() *Storage {
	films := make(map[int]models.Film)
	films[1] = models.Film{1,
		"Теперь Эдди и Веном в бегах. Преследуемые обоими мирами и приближающейся сетью, Эдди и Веном вынуждены принять роковое решение, которое опустит занавес на их последнем танце.",
		"Venom: The Last Dance",
		"триллер, экшн",
		time.Date(2024, time.October, 24, 0, 0, 0, 0, time.UTC)}
	return &Storage{
		Films: films,
	}
}

func (s *Storage) GetFilmById(id int) (models.Film, error) {
	//Если нет ключа в мап, то будет возвращен нулевое значение. Это у нас пустая структура Room
	return s.Films[id], nil
}

func (s *Storage) CreateFilm(film models.Film) (int, string) { //пока возвращаем стринг так как мы сами создаем ошибку и отправляем ввиду строки
	if _, exist := s.Films[film.Id]; exist {
		err := "Can not create room, room already exist"
		return 0, err
	}

	s.Films[film.Id] = film
	return film.Id, ""
}

func (s *Storage) UpdateFilm(film models.Film) (interface{}, string) { //пока возвращаем стринг так как мы сами создаем ошибку и отправляем ввиду строки
	if _, exist := s.Films[film.Id]; exist {
		s.Films[film.Id] = film
		return s.Films[film.Id], ""
	} else {
		err := "Room id not found"
		return nil, err
		//sefsefes
	}
}

func (s *Storage) GetAllFilms() interface{} {
	var films []models.Film
	for _, room := range s.Films {
		films = append(films, room)
	}

	return films
}

func (s *Storage) DeleteFilm(id int) (int, string) {
	if _, exist := s.Films[id]; exist {
		err := "Can not delete room, room does not exist"
		return 0, err
	}

	delete(s.Films, id)
	return id, ""
}
