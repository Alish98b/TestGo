package repositories

import (
	"hotel/internal/models"

	"github.com/jmoiron/sqlx"
)

type GenreDB struct {
	db *sqlx.DB
}

func NewGenreDB(db *sqlx.DB) *GenreDB {
	return &GenreDB{db: db}
}

func (g *GenreDB) CreateGenre(genre models.GenreCreate) (int, error) {
	var id int
	query := `INSERT INTO genres (name) VALUES ($1) RETURNING id`	
	row := g.db.QueryRow(query, genre.Name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (g *GenreDB) GetGenreById(id int) (models.Genre, error) {
	var genre models.Genre
	query := `SELECT id, name FROM genres WHERE id=$1`
	err := g.db.Get(&genre, query, id)
	return genre, err
}

func (g *GenreDB) GetAllGenres() ([]models.Genre, error) {
	var genres []models.Genre
	query := `SELECT id, name FROM genres`
	err := g.db.Select(&genres, query)
	return genres, err
}

func (g *GenreDB) UpdateGenre(id int, genre models.GenreCreate) error {
	query := `UPDATE genres SET name=$1 WHERE id=$2`
	_, err := g.db.Exec(query, genre.Name, id)
	return err
}

func (g *GenreDB) DeleteGenre(id int) error {
	query := `DELETE FROM genres WHERE id=$1`
	_, err := g.db.Exec(query, id)
	return err
}
