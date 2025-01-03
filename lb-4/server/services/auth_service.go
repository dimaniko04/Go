package services

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/dimaniko04/Go/lb-4/server/models"
	"github.com/dimaniko04/Go/lb-4/server/responses"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(email string, password string) (responses.AuthResponse, error)
	Login(email string, password string) (responses.AuthResponse, error)
}

type authService struct {
	db     *sql.DB
	secret string
}

func (s *authService) Register(email string, password string) (responses.AuthResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return responses.AuthResponse{}, errors.New("failed to hash password")
	}

	query := "INSERT INTO users (email, password_hash, role) VALUES (?, ?, 'user')"
	res, err := s.db.Exec(query, email, hashedPassword)
	if err != nil {
		log.Println(err)
		return responses.AuthResponse{}, errors.New("failed to register user")
	}
	userId, err := res.LastInsertId()
	if err != nil {
		return responses.AuthResponse{}, errors.New("failed to return registered user")
	}
	user := models.User{Id: int(userId), Email: email, Role: "user"}
	token, err := s.generateJWT(user)
	if err != nil {
		return responses.AuthResponse{}, err
	}

	return responses.AuthResponse{
		User:  user,
		Token: token,
	}, nil
}

func (s *authService) Login(email string, password string) (responses.AuthResponse, error) {
	query := "SELECT id, email, password_hash, role FROM users WHERE email = ?"
	row := s.db.QueryRow(query, email)

	user := models.User{}
	if err := row.Scan(&user.Id, &user.Email, &user.PasswordHash, &user.Role); err != nil {
		if err == sql.ErrNoRows {
			return responses.AuthResponse{}, errors.New("invalid email or password")
		}

		return responses.AuthResponse{}, errors.New("failed to fetch user")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return responses.AuthResponse{}, errors.New("invalid username or password")
	}
	token, err := s.generateJWT(user)
	if err != nil {
		return responses.AuthResponse{}, err
	}

	return responses.AuthResponse{
		User:  user,
		Token: token,
	}, nil
}

func (s *authService) generateJWT(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.Id,
		"role":  user.Role,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.secret))

	if err != nil {
		log.Print(err)
		return "", errors.New("failed to sign jwt token")
	}

	return tokenString, nil
}
