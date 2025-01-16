package subscriptions

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SubscriptionHandler struct {
	service *SubscriptionService
}

func NewSubscriptionHandler(service *SubscriptionService) *SubscriptionHandler {
	return &SubscriptionHandler{service}
}

func (h *SubscriptionHandler) GetSubscriptionsHandler(w http.ResponseWriter, r *http.Request) {
	// gets the user's subscriptions
	subscriptions, err := h.service.GetSubscriptionsService(r.Context())
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

func (h *SubscriptionHandler) GetSubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	// get the url param
	subscriptionID, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// gets the user's specified subscription
	subscription, err := h.service.GetSubscriptionService(r.Context(), subscriptionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// returns the subscription
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(subscription); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

func (h *SubscriptionHandler) CreateSubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	subscriptionPayload := Subscription{}

	// tries to decode the body data to payload variable
	if err := json.NewDecoder(r.Body).Decode(&subscriptionPayload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// tries to create the subscription
	if err := h.service.CreateSubscriptionService(r.Context(), subscriptionPayload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *SubscriptionHandler) UpdateSubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	subscriptionPayload := Subscription{}

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
	defer r.Body.Close()

	// tries to update the subscription
	if err := h.service.UpdateSubscriptionService(r.Context(), subscriptionPayload, subscriptionID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *SubscriptionHandler) DeleteSubscriptionHandler(w http.ResponseWriter, r *http.Request) {
	// get the url param
	subscriptionID, err := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// tries to delete the subscription
	if err := h.service.DeleteSubscriptionService(r.Context(), subscriptionID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
