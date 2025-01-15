package subscriptions

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var subscriptionColl *mongo.Collection

func RepoInit(client *mongo.Client) {
	// initializes the collection
	subscriptionColl = client.Database("trackrplus").Collection("subscriptions")
}

func GetSubscriptions(userID primitive.ObjectID) (*Subscriptions, error) {
	var subscriptions Subscriptions

	// gets the subscriptions from db
	cursor, err := subscriptionColl.Find(context.TODO(), bson.M{"userID": userID})
	if err != nil {
		return nil, err
	}

	// decodes all subscriptions
	if err := cursor.All(context.TODO(), &subscriptions); err != nil {
		return nil, err
	}

	return &subscriptions, nil
}

func GetSubscription(userID primitive.ObjectID, subscriptionID primitive.ObjectID) (*Subscription, error) {
	var subscription Subscription

	// gets the subscription from db
	sub := subscriptionColl.FindOne(context.TODO(), bson.M{"userID": userID, "_id": subscriptionID})
	if sub.Err() != nil {
		return nil, sub.Err()
	}

	// decodes the subscription
	if err := sub.Decode(&subscription); err != nil {
		return nil, err
	}

	return &subscription, nil
}

func CreateSubscription(subscription *Subscription) error {
	// add subscription to the db
	_, err := subscriptionColl.InsertOne(context.TODO(), subscription)
	if err != nil {
		return err
	}

	return nil
}

func DeleteSubscription(subscriptionID primitive.ObjectID) error {
	_, err := subscriptionColl.DeleteOne(context.TODO(), bson.M{"_id": subscriptionID})
	if err != nil {
		return err
	}

	return nil
}

func UpdateSubscription(subscriptionID primitive.ObjectID, subscription *Subscription) error {
	_, err := subscriptionColl.ReplaceOne(context.TODO(), bson.M{"_id": subscriptionID}, subscription)
	if err != nil {
		return err
	}

	return nil
}
