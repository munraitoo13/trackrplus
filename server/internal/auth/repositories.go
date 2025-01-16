package auth

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// repository for users
type AuthRepository struct {
	coll *mongo.Collection
}

// creates a new user repository
func NewAuthRepository(client *mongo.Client) *AuthRepository {
	return &AuthRepository{
		coll: client.Database("trackrplus").Collection("users"),
	}
}

func (r *AuthRepository) GetUserByEmail(email string) (User, error) {
	// create a user object
	user := User{}

	// find the user by email
	err := r.coll.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return User{}, fmt.Errorf("user with email %s not found: %w", email, err)
	}

	return user, nil
}

func (r *AuthRepository) RegisterUser(user User) error {
	// insert the user to the db
	_, err := r.coll.InsertOne(context.TODO(), user)
	if err != nil {
		return fmt.Errorf("failed to register the user: %w", err)
	}

	return nil
}
