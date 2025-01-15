package subscriptions

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetSubscriptionsHandler(w http.ResponseWriter, r *http.Request) {
	// gets the user's subscriptions
	subscriptions, err := GetSubscriptionsService(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// returns the subscriptions
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(subscriptions); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetSubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	// get the url param
	subscriptionID, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// gets the user's specified subscription
	subscription, err := GetSubscriptionService(r.Context(), subscriptionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// returns the subscription
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(subscription); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CreateSubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	var subscriptionPayload *Subscription

	// tries to decode the body data to payload variable
	if err := json.NewDecoder(r.Body).Decode(&subscriptionPayload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// tries to create the subscription
	if err := CreateSubscriptionService(r.Context(), subscriptionPayload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func UpdateSubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	var subscriptionPayload *Subscription

	// get the url param
	subscriptionID, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// tries to decode the body data to payload variable
	if err := json.NewDecoder(r.Body).Decode(&subscriptionPayload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// tries to update the subscription
	if err := UpdateSubscriptionService(r.Context(), subscriptionPayload, subscriptionID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func DeleteSubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	// get the url param
	subscriptionID, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// tries to delete the subscription
	if err := DeleteSubscriptionService(r.Context(), subscriptionID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
