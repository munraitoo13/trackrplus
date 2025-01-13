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

func GetSubscriptionService(ctx context.Context, subID primitive.ObjectID) (*Subscription, error) {
	// gets the user id from the token
	userID := ctx.Value("userID").(primitive.ObjectID)

	// gets the subscription
	subscription, err := GetSubscription(userID, subID)
	if err != nil {
		return nil, errors.New("error getting subscriptions")
	}

	// checks if the subscription exists
	if subscription == nil {
		return nil, errors.New("subscription not found")
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
	}

	if err := CreateSubscription(subscription); err != nil {
		return errors.New("error creating subscription")
	}

	return nil
}

func UpdateSubscriptionService(ctx context.Context, payload *Subscription, subID primitive.ObjectID) error {
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

	if err := UpdateSubscription(subID, subscription); err != nil {
		return errors.New("error creating subscription")
	}

	return nil
}

func DeleteSubscriptionService(subID primitive.ObjectID) error {
	if err := DeleteSubscription(subID); err != nil {
		return errors.New("error deleting subscription")
	}

	return nil
}
