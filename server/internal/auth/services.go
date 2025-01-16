package auth

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"server/internal/common"
)

type AuthService struct {
	repo *AuthRepository
}

func NewAuthService(repo *AuthRepository) *AuthService {
	return &AuthService{repo}
}

func (s *AuthService) LoginService(payload LoginPayload) (string, error) {
	// gets user by its email
	user, err := s.repo.GetUserByEmail(payload.Email)
	if err != nil {
		return "", fmt.Errorf("user not found: %v", err)
	}

	// compare pass input with hashed one in db
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return "", fmt.Errorf("invalid password: %v", err)
	}

	// generate the jwt token
	token, err := common.GenerateToken(user.ID)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	// return the token
	return token, nil
}

func (s *AuthService) RegisterService(payload RegisterPayload) error {
	// check if the email already exists
	_, err := s.repo.GetUserByEmail(payload.Email)
	if err == nil {
		return errors.New("email already exists")
	}

	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	// create a new user object
	user := User{
		Name:      payload.Name,
		Username:  payload.Username,
		Email:     payload.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),

		Admin:   false,
		Premium: false,
	}

	// register the user
	return s.repo.RegisterUser(user)
}
