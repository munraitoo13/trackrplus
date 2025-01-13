package common

import (
	"errors"
	"server/configs"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var secretKey = configs.GetEnv("SECRET_KEY")

func GenerateToken(userID primitive.ObjectID) (string, error) {
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

func GetUserIdFromToken(tokenString string) (primitive.ObjectID, error) {
	// parse the token
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil || !token.Valid {
		return primitive.NilObjectID, errors.New("invalid jwt token")
	}

	// get the claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return primitive.NilObjectID, errors.New("invalid claims")
	}

	// get the userID field from claims
	userID, ok := claims["userID"].(string)
	if !ok {
		return primitive.NilObjectID, errors.New("userID not in the token")
	}

	userIDObj, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return primitive.NilObjectID, errors.New("invalid userID format")
	}

	return userIDObj, nil
}
