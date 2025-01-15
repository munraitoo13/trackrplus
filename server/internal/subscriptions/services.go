package subscriptions

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetSubscriptionsService(ctx context.Context) (*Subscriptions, error) {
	// gets the user id from the token
	userID := ctx.Value("userID").(primitive.ObjectID)

	// gets the subscriptions
	subscriptions, err := GetSubscriptions(userID)
	if err != nil {
		return nil, errors.New("error getting subscriptions")
	}

	// checks if the user has any subscriptions
	if len(*subscriptions) == 0 {
		return &Subscriptions{}, nil
	}

	return subscriptions, nil
}

func GetSubscriptionService(ctx context.Context, subscriptionID primitive.ObjectID) (*Subscription, error) {
	// gets the user id from the token
	userID := ctx.Value("userID").(primitive.ObjectID)

	// gets the subscription
	subscription, err := GetSubscription(userID, subscriptionID)
	if err != nil {
		return nil, errors.New("error getting subscriptions")
	}

	// checks if the subscription exists
	if subscription == nil {
		return nil, errors.New("subscription not found")
	}

	if subscription.UserID != userID {
		return nil, errors.New("unauthorized")
	}

	return subscription, nil
}

func CreateSubscriptionService(ctx context.Context, payload *Subscription) error {
	userID := ctx.Value("userID").(primitive.ObjectID)

	subscription := &Subscription{
		UserID:        userID,
		Name:          payload.Name,
		Price:         payload.Price,
		PaymentMethod: payload.PaymentMethod,
		CardLast4:     payload.CardLast4,
		Paid:          payload.Paid,
		RenewalDate:   payload.RenewalDate,
		Notes:         payload.Notes,
		CreatedAt:     payload.CreatedAt,
	}

	if err := CreateSubscription(subscription); err != nil {
		return errors.New("error creating subscription")
	}

	return nil
}

func UpdateSubscriptionService(ctx context.Context, payload *Subscription, subscriptionID primitive.ObjectID) error {
	userID := ctx.Value("userID").(primitive.ObjectID)

	existingSubscription, err := GetSubscription(userID, subscriptionID)
	if err != nil {
		return errors.New("error getting subscriptions")
	}

	if existingSubscription == nil || existingSubscription.UserID != userID {
		return errors.New("unauthorized")
	}

	subscription := &Subscription{
		UserID:        userID,
		Name:          payload.Name,
		Price:         payload.Price,
		PaymentMethod: payload.PaymentMethod,
		CardLast4:     payload.CardLast4,
		Paid:          payload.Paid,
		RenewalDate:   payload.RenewalDate,
		Notes:         payload.Notes,
		CreatedAt:     existingSubscription.CreatedAt,
	}

	if err := UpdateSubscription(subscriptionID, subscription); err != nil {
		return errors.New("error creating subscription")
	}

	return nil
}

func DeleteSubscriptionService(ctx context.Context, subscriptionID primitive.ObjectID) error {
	userID := ctx.Value("userID").(primitive.ObjectID)

	existingSubscription, err := GetSubscription(userID, subscriptionID)
	if err != nil {
		return errors.New("error getting subscriptions")
	}

	if existingSubscription == nil || existingSubscription.UserID != userID {
		return errors.New("unauthorized")
	}

	if err := DeleteSubscription(subscriptionID); err != nil {
		return errors.New("error deleting subscription")
	}

	return nil
}
