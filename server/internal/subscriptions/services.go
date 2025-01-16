package subscriptions

import (
	"context"
	"errors"
	"fmt"
	"server/internal/middlewares"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SubscriptionService struct {
	repo *SubscriptionRepository
}

func NewSubscriptionService(repo *SubscriptionRepository) *SubscriptionService {
	return &SubscriptionService{repo}
}

func (s *SubscriptionService) GetSubscriptionsService(ctx context.Context) (Subscriptions, error) {
	// gets the user id from the token
	userID := ctx.Value(middlewares.UserIDKey).(primitive.ObjectID)

	// gets the subscriptions
	subscriptions, err := s.repo.GetSubscriptions(userID)
	if err != nil {
		return nil, fmt.Errorf("error getting subscriptions: %v", err)
	}

	// checks if the user has any subscriptions
	if len(subscriptions) == 0 {
		return Subscriptions{}, errors.New("no subscriptions found")
	}

	return subscriptions, nil
}

func (s *SubscriptionService) GetSubscriptionService(ctx context.Context, subscriptionID primitive.ObjectID) (Subscription, error) {
	// gets the user id from the token
	userID := ctx.Value(middlewares.UserIDKey).(primitive.ObjectID)

	// gets the subscription
	subscription, err := s.repo.GetSubscription(userID, subscriptionID)
	if err != nil {
		return Subscription{}, errors.New("error getting subscriptions")
	}

	// checks if the user is allowed to access the subscription
	if subscription.UserID != userID {
		return Subscription{}, errors.New("unauthorized: user is not allowed to access this subscription")
	}

	return subscription, nil
}

func (s *SubscriptionService) CreateSubscriptionService(ctx context.Context, payload Subscription) error {
	userID := ctx.Value(middlewares.UserIDKey).(primitive.ObjectID)

	// creates a new subscription object
	subscription := Subscription{
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

	// tries creates the subscription
	if err := s.repo.CreateSubscription(subscription); err != nil {
		return fmt.Errorf("error creating subscription: %v", err)
	}

	return nil
}

func (s *SubscriptionService) UpdateSubscriptionService(ctx context.Context, payload Subscription, subscriptionID primitive.ObjectID) error {
	userID := ctx.Value(middlewares.UserIDKey).(primitive.ObjectID)

	// gets the existing subscription
	existingSubscription, err := s.repo.GetSubscription(userID, subscriptionID)
	if err != nil {
		return fmt.Errorf("error getting existing subscriptions: %v", err)
	}

	// checks if the user is allowed to update the subscription
	if existingSubscription.UserID != userID {
		return errors.New("unauthorized: user is not allowed to update this subscription")
	}

	// creates a new subscription object
	subscription := Subscription{
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

	// tries to update the subscription
	if err := s.repo.UpdateSubscription(subscriptionID, subscription, userID); err != nil {
		return fmt.Errorf("error updating subscription: %v", err)
	}

	return nil
}

func (s *SubscriptionService) DeleteSubscriptionService(ctx context.Context, subscriptionID primitive.ObjectID) error {
	userID := ctx.Value(middlewares.UserIDKey).(primitive.ObjectID)

	// gets the existing subscription
	existingSubscription, err := s.repo.GetSubscription(userID, subscriptionID)
	if err != nil {
		return fmt.Errorf("error getting existing subscriptions: %v", err)
	}

	// checks if the user is allowed to delete the subscription
	if existingSubscription.UserID != userID {
		return errors.New("unauthorized: user is not allowed to delete this subscription")
	}

	// tries to delete the subscription
	if err := s.repo.DeleteSubscription(subscriptionID, userID); err != nil {
		return fmt.Errorf("error deleting subscription: %v", err)
	}

	return nil
}
