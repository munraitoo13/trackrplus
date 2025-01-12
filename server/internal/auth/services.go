package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"server/configs"
)

func Login(payload *UserLogin) (string, error) {
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
	token, err := GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func Register(payload *UserRegister) error {
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

func GenerateToken(userID primitive.ObjectID) (string, error) {
	secretKey := configs.GetEnv("SECRET_KEY")

	// jwt claims
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	// sign the claims with secret key
	token, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}
