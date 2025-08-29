package service

import (
	"context"
	"errors"
	"time"

	"jp-song-trainer/internal/model"
	"jp-song-trainer/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthService interface {
	Register(email, password string) (*model.User, string, error)
	Login(email, password string) (*model.User, string, error)
	GenerateToken(user *model.User) (string, error)
	ParseToken(tokenString string) (*jwt.RegisteredClaims, error)
}

type authService struct {
	userRepo  repository.UserRepository
	jwtSecret string
}

func NewAuthService(userRepo repository.UserRepository, jwtSecret string) AuthService {
	return &authService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (s *authService) Register(email, password string) (*model.User, string, error) {
	// Check if user already exists
	_, err := s.userRepo.GetUserByEmail(context.Background(), email)
	if err == nil {
		return nil, "", errors.New("user already exists")
	}

	// Create new user
	user := &model.User{
		ID:        uuid.New(),
		Email:     email,
		Role:      "user",
		CreatedAt: time.Now(),
	}

	if err := user.SetPassword(password); err != nil {
		return nil, "", err
	}

	if err := s.userRepo.CreateUser(context.Background(), user); err != nil {
		return nil, "", err
	}

	token, err := s.GenerateToken(user)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (s *authService) Login(email, password string) (*model.User, string, error) {
	user, err := s.userRepo.GetUserByEmail(context.Background(), email)
	if err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	if !user.CheckPassword(password) {
		return nil, "", errors.New("invalid credentials")
	}

	token, err := s.GenerateToken(user)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (s *authService) GenerateToken(user *model.User) (string, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		ID:        user.ID.String(),
		Subject:   user.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}

func (s *authService) ParseToken(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
