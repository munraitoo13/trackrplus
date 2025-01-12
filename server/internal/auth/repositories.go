package auth

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userColl *mongo.Collection

func RepoInit(client *mongo.Client) {
	userColl = client.Database("trackrplus").Collection("users")
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	err := userColl.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func RegisterUser(user *User) error {
	_, err := userColl.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	return nil
}
