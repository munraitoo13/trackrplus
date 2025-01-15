package auth

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"

	"server/internal/common"
)

func LoginService(payload *LoginPayload) (string, error) {
	var err error

	// gets user by its email
	user, err := GetUserByEmail(payload.Email)
	if err != nil {
		return "", err
	}

	// compare pass input with hashed one in db
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// generate the jwt token
	token, err := common.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func RegisterService(payload *RegisterPayload) error {
	// check if the email already exists
	_, err := GetUserByEmail(payload.Email)
	if err == nil {
		return errors.New("email already exists")
	}

	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// create a new user object
	user := &User{
		Name:      payload.Name,
		Username:  payload.Username,
		Email:     payload.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),

		Admin:   false,
		Premium: false,
	}

	// register the user
	if err := RegisterUser(user); err != nil {
		return err
	}

	return nil
}
