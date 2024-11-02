package repositories

import (
	"hotel/internal/models"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type CinemaDB struct {
	db *sqlx.DB
}

func NewCinemaDB(db *sqlx.DB) *CinemaDB {
	return &CinemaDB{db: db}
}

func (c *CinemaDB) GetMovieById(id int) (interface{}, error) {
	var movie models.Movie
	query := `SELECT id, title, description, release_date, duration FROM movies WHERE id=$1`

	err := c.db.Get(&movie, query, id)
	if err != nil {
		logrus.Errorf("Failed to get movie with ID %d: %v", id, err)
		return movie, err
	}
	logrus.Infof("Fetched movie with ID %d", id)
	return movie, nil
}

func (c *CinemaDB) CreateMovie(movie models.MovieCreate) (int, error) {
	var id int
	query := `INSERT INTO movies (title, description, release_date, duration) VALUES ($1, $2, $3, $4) RETURNING id`

	row := c.db.QueryRow(query, movie.Title, movie.Description, movie.ReleaseDate, movie.Duration)
	if err := row.Scan(&id); err != nil {
		logrus.Errorf("Failed to create movie: %v", err)
		return 0, err
	}
	logrus.Infof("Created movie with ID %d", id)
	return id, nil
}

func (c *CinemaDB) UpdateMovie(id int, movie models.MovieCreate) error {
	query := `UPDATE movies SET title=$1, description=$2, release_date=$3, duration=$4 WHERE id=$5`

	_, err := c.db.Exec(query, movie.Title, movie.Description, movie.ReleaseDate, movie.Duration, id)
	if err != nil {
		logrus.Errorf("Failed to update movie with ID %d: %v", id, err)
		return err
	}
	logrus.Infof("Updated movie with ID %d", id)
	return nil
}

func (c *CinemaDB) DeleteMovie(id int) error {
	query := `DELETE FROM movies WHERE id=$1`

	_, err := c.db.Exec(query, id)
	if err != nil {
		logrus.Errorf("Failed to delete movie with ID %d: %v", id, err)
		return err
	}
	logrus.Infof("Deleted movie with ID %d", id)
	return nil
}

func (c *CinemaDB) GetAllMovies() (interface{}, error) {
	var movies []models.Movie
	query := `SELECT id, title, description, release_date, duration FROM movies`

	err := c.db.Select(&movies, query)
	if err != nil {
		logrus.Errorf("Failed to fetch all movies: %v", err)
		return nil, err
	}
	logrus.Infof("Fetched %d movies", len(movies))
	return movies, nil
}
