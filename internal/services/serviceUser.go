package services

import (
	"errors"
	"hotel/internal/models"
	"hotel/internal/repositories"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	signingKey = "213easdxz1c856eq"
)

type UserService struct {
	repo *repositories.Repo
}

func NewUserService(repo *repositories.Repo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserById(id int) (interface{}, error) {
	return s.repo.GetUserById(id)
}

func (s *UserService) CreateUser(user models.UserCreate) (int, error) {
	return s.repo.CreateUser(user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.repo.DeleteUser(id)
}

func (s *UserService) GetAllUsers() (interface{}, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) UpdateUser(id int, user models.UserCreate) error {
	return s.repo.UpdateUser(id, user)
}

func (s *UserService) GenerateToken(id int) (string, error) {
	user, err := s.repo.GetUserById(id)

	if err != nil {
		return "", err
	}

	//Добавить в конгиф файл время жизни токенов!!!
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
		Subject:   strconv.Itoa(int(user.ID)),
	})

	return token.SignedString([]byte(signingKey))
}

func (s *UserService) GenerateRefreshToken(id int) (string, error) {
	user, err := s.repo.GetUserById(id)
	if err != nil {
		return "", err
	}

	// Создаем refresh токен с длительным сроком действия (например, 7 дней)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(),
		Subject:   strconv.Itoa(int(user.ID)),
	})

	return refreshToken.SignedString([]byte(signingKey))
}

func (s *UserService) ParseToken(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims["sub"].(string), nil
}
