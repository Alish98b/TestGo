package repositories

import (
	"fmt"
	_ "fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"hotel/internal/models"
	_ "hotel/internal/models"

	_ "github.com/jmoiron/sqlx"
	_ "github.com/sirupsen/logrus"
)

type SessionPostgres struct {
	db *sqlx.DB
}

func (r *SessionPostgres) GetAllSessions() (models.Session, error) {
	var session []models.Session
	query := `SELECT id, movie_id, hall_id, start_time FROM session`

	err := r.db.Select(&session, query)
	if err != nil {
		logrus.Errorf("Failed to fetch all halls: %v", err)
		return models.Session{}, nil
	}
	logrus.Infof("Fetched %d halls", len(session))
	return models.Session{}, nil
}

func NewSessionPostgres(db *sqlx.DB) *SessionPostgres {
	return &SessionPostgres{db: db}
}

func (r *SessionPostgres) GetSessionById(id int) (models.Session, error) {
	var session models.Session

	query := fmt.Sprintf("SELECT r.movie_id,  r.hall_id , r.start_time FROM %s r WHERE r.id=$1", "Session")
	err := r.db.Get(&session, query, id)

	if err != nil {
		return session, err
	}

	return session, nil
}

func (r *SessionPostgres) CreateSession(session models.SessionCreate) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (MovieID, HallID, StartTime) VALUES ($1, $2, $3) RETURNING id", "Session")
	row := r.db.QueryRow(query, session.MovieID, session.HallID, session.StartTime)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *SessionPostgres) GetAllSession() interface{} {
	var session []models.SessionCreate
	query := `SELECT movie_id, hall_id, start_time FROM session`

	err := r.db.Select(&session, query)
	if err != nil {
		logrus.Errorf("Failed to fetch all halls: %v", err)
		return err
	}
	logrus.Infof("Fetched %d halls", len(session))
	return session
}

func (r *SessionPostgres) DeleteSession(id int) (int error) {
	query := `DELETE FROM hall WHERE id=$1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		logrus.Errorf("Failed to delete hall with ID %d: %v", id, err)
		return err
	}
	logrus.Infof("Deleted hall with ID %d", id)
	return nil
}

func (r *SessionPostgres) UpdateSession(id int, session models.Session) error {
	query := `UPDATE hall SET start_time=$1  WHERE id=$2`

	_, err := r.db.Exec(query, session.StartTime, id)
	if err != nil {
		logrus.Errorf("Failed to update hall with ID %d: %v", session, err)
		return err
	}
	logrus.Infof("Updated hall with ID %d", session)
	return err
}
