package subscriptions

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// repository for subscriptions
type SubscriptionRepository struct {
	coll *mongo.Collection
}

// creates a new user repository
func NewSubscriptionRepository(client *mongo.Client) *SubscriptionRepository {
	return &SubscriptionRepository{
		coll: client.Database("trackrplus").Collection("subscriptions"),
	}
}

func (r *SubscriptionRepository) GetSubscriptions(userID primitive.ObjectID) (Subscriptions, error) {
	subscriptions := Subscriptions{}

	// gets the subscriptions from db
	cursor, err := r.coll.Find(context.TODO(), bson.M{"userID": userID})
	if err != nil {
		return nil, fmt.Errorf("failed to get subscriptions: %w", err)
	}
	defer cursor.Close(context.Background())

	// decodes all subscriptions
	if err := cursor.All(context.TODO(), &subscriptions); err != nil {
		return nil, fmt.Errorf("failed to decode subscriptions: %w", err)
	}

	return subscriptions, nil
}

func (r *SubscriptionRepository) GetSubscription(userID primitive.ObjectID, subscriptionID primitive.ObjectID) (Subscription, error) {
	subscription := Subscription{}

	// gets the subscription from db
	if err := r.coll.FindOne(context.TODO(), bson.M{"userID": userID, "_id": subscriptionID}).Decode(&subscription); err != nil {
		return Subscription{}, fmt.Errorf("failed to get subscription: %w", err)
	}

	return subscription, nil
}

func (r *SubscriptionRepository) CreateSubscription(subscription Subscription) error {
	// add subscription to the db
	_, err := r.coll.InsertOne(context.TODO(), subscription)
	if err != nil {
		return fmt.Errorf("failed to create subscription: %w", err)
	}

	return nil
}

func (r *SubscriptionRepository) DeleteSubscription(subscriptionID primitive.ObjectID, userID primitive.ObjectID) error {
	_, err := r.coll.DeleteOne(context.TODO(), bson.M{"userID": userID, "_id": subscriptionID})
	if err != nil {
		return fmt.Errorf("failed to delete subscription: %w", err)
	}

	return nil
}

func (r *SubscriptionRepository) UpdateSubscription(subscriptionID primitive.ObjectID, subscription Subscription, userID primitive.ObjectID) error {
	_, err := r.coll.ReplaceOne(context.TODO(), bson.M{"userID": userID, "_id": subscriptionID}, subscription)
	if err != nil {
		return fmt.Errorf("failed to update subscription: %w", err)
	}

	return nil
}
