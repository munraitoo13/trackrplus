package routes

import (
	"server/internal/auth"
	"server/internal/subscriptions"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r chi.Router, auth *auth.AuthHandler, subscriptions *subscriptions.SubscriptionHandler) {
	r.Route("/api", func(r chi.Router) {
		// auth
		r.Route("/auth", func(r chi.Router) {
			r.Post("/login", auth.LoginHandler)
			r.Post("/register", auth.RegisterHandler)
		})

		// subscriptions
		r.Route("/subscriptions", func(r chi.Router) {
			r.Get("/", subscriptions.GetSubscriptionsHandler)
			r.Post("/", subscriptions.CreateSubscriptionHandler)

			r.Get("/{id}", subscriptions.GetSubscriptionHandler)
			r.Put("/{id}", subscriptions.UpdateSubscriptionHandler)
			r.Delete("/{id}", subscriptions.DeleteSubscriptionHandler)
		})
	})
}
