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
	movieQuery := `SELECT id, title, description, release_date, duration FROM movies WHERE id=$1`
	err := c.db.Get(&movie, movieQuery, id)
	if err != nil {
		return nil, err
	}

	genresQuery := `SELECT g.id, g.name FROM genres g
                    INNER JOIN movie_genres mg ON mg.genre_id = g.id
                    WHERE mg.movie_id = $1`
	err = c.db.Select(&movie.Genres, genresQuery, id)
	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (c *CinemaDB) CreateMovie(movie models.MovieCreate) (int, error) {
	tx, err := c.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	movieQuery := `INSERT INTO movies (title, description, release_date, duration) VALUES ($1, $2, $3, $4) RETURNING id`
	row := tx.QueryRow(movieQuery, movie.Title, movie.Description, movie.ReleaseDate, movie.Duration)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		logrus.Errorf("Failed to create movie: %v", err)
		return 0, err
	}

	genreQuery := `INSERT INTO movie_genres (movie_id, genre_id) VALUES ($1, $2)`
	for _, genreID := range movie.GenreIDs {
		_, err := tx.Exec(genreQuery, id, genreID)
		if err != nil {
			tx.Rollback()
			logrus.Errorf("Failed to associate movie with genre: %v", err)
			return 0, err
		}
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	logrus.Infof("Created movie with ID %d", id)
	return id, nil
}

func (c *CinemaDB) UpdateMovie(id int, movie models.MovieCreate) error {
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}

	movieQuery := `UPDATE movies SET title=$1, description=$2, release_date=$3, duration=$4 WHERE id=$5`
	_, err = tx.Exec(movieQuery, movie.Title, movie.Description, movie.ReleaseDate, movie.Duration, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	deleteGenresQuery := `DELETE FROM movie_genres WHERE movie_id=$1`
	_, err = tx.Exec(deleteGenresQuery, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	insertGenreQuery := `INSERT INTO movie_genres (movie_id, genre_id) VALUES ($1, $2)`
	for _, genreID := range movie.GenreIDs {
		_, err := tx.Exec(insertGenreQuery, id, genreID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
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
	moviesQuery := `SELECT id, title, description, release_date, duration FROM movies`
	err := c.db.Select(&movies, moviesQuery)
	if err != nil {
		return nil, err
	}

	for i := range movies {
		genresQuery := `SELECT g.id, g.name FROM genres g
                        INNER JOIN movie_genres mg ON mg.genre_id = g.id
                        WHERE mg.movie_id = $1`
		err := c.db.Select(&movies[i].Genres, genresQuery, movies[i].ID)
		if err != nil {
			return nil, err
		}
	}

	return movies, nil
}
