package repositories

import (
	"fmt"
	"hotel/internal/models"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type HallPostgres struct {
	db *sqlx.DB
}

func NewHallPostgres(db *sqlx.DB) *HallPostgres {
	return &HallPostgres{db: db}
}

func (r *HallPostgres) GetHallById(id int) (models.Hall, error) {
	var hall models.Hall

	query := fmt.Sprintf("SELECT r.id, r.name, r.capacity FROM %s r WHERE r.id=$1", "Hall")
	err := r.db.Get(&hall, query, id)

	if err != nil {
		return hall, err
	}

	return hall, nil
}

func (r *HallPostgres) CreateHall(hall models.HallCreate) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, capacity) VALUES ($1, $2) RETURNING id", "hall")
	row := r.db.QueryRow(query, hall.Name, hall.Capacity)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
func (r *HallPostgres) GetAllHalls() interface{} {
	var hall []models.Hall
	query := `SELECT id, name, capacity FROM hall`

	err := r.db.Select(&hall, query)
	if err != nil {
		logrus.Errorf("Failed to fetch all halls: %v", err)
		return err
	}
	logrus.Infof("Fetched %d halls", len(hall))
	return hall
}

func (r *HallPostgres) DeleteHall(id int) (int error) {
	query := `DELETE FROM hall WHERE id=$1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		logrus.Errorf("Failed to delete hall with ID %d: %v", id, err)
		return err
	}
	logrus.Infof("Deleted hall with ID %d", id)
	return nil
}

func (r *HallPostgres) UpdateHall(id int, hall models.Hall) error {
	query := `UPDATE hall SET name=$1, capacity=$2 WHERE id=$3`

	_, err := r.db.Exec(query, hall.Name, hall.Capacity, id)
	if err != nil {
		logrus.Errorf("Failed to update hall with ID %d: %v", hall, err)
		return err
	}
	logrus.Infof("Updated hall with ID %d", hall)
	return err
}

//func (r *HallPostgres) CheckHallExists(id int) (bool, error) {
//	var exists bool
//	query := "SELECT EXISTS(SELECT 1 FROM halls WHERE id = $1)"
//	err := r.db.QueryRow(query, id).Scan(&exists)
//	if err != nil {
//		return false, err
//	}
//	return exists, nil
//}
