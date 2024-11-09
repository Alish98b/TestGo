package repositories

import (
	"hotel/internal/models"
	//"os/user"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type UserDB struct {
	db *sqlx.DB
}

func NewUserDB(db *sqlx.DB) *UserDB {
	return &UserDB{db: db}
}

func (c *UserDB) GetUserById(id int) (interface{}, error) {
	var user models.User
	query := `SELECT id, name, password, email FROM users WHERE id=$1`

	err := c.db.Get(&user, query, id)
	if err != nil {
		logrus.Errorf("Failed to get user with ID %d: %v", id, err)
		return user, err
	}
	logrus.Infof("Fetched user with ID %d", id)
	return user, nil
}

func (c *UserDB) CreateUser(user models.UserCreate) (int, error) {
	var id int
	query := `INSERT INTO users (name, password, email) VALUES ($1, $2, $3) RETURNING id`

	row := c.db.QueryRow(query, user.Name, user.Password, user.Email)
	if err := row.Scan(&id); err != nil {
		logrus.Errorf("Failed to create user: %v", err)
		return 0, err
	}
	logrus.Infof("Created user with ID %d", id)
	return id, nil
}

func (c *UserDB) UpdateUser(id int, user models.UserCreate) error {
	query := `UPDATE users SET name=$1, password=$2, email=$3 WHERE id=$4`

	_, err := c.db.Exec(query, user.Name, user.Password, user.Email, id)
	if err != nil {
		logrus.Errorf("Failed to update user with ID %d: %v", id, err)
		return err
	}
	logrus.Infof("Updated user with ID %d", id)
	return nil
}

func (c *UserDB) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id=$1`

	_, err := c.db.Exec(query, id)
	if err != nil {
		logrus.Errorf("Failed to delete user with ID %d: %v", id, err)
		return err
	}
	logrus.Infof("Deleted user with ID %d", id)
	return nil
}

func (c *UserDB) GetAllUsers() (interface{}, error) {
	var users []models.User
	query := `SELECT id, name, password, email FROM users`

	err := c.db.Select(&users, query)
	if err != nil {
		logrus.Errorf("Failed to fetch all users: %v", err)
		return nil, err
	}
	logrus.Infof("Fetched %d users", len(users))
	return users, nil
}
