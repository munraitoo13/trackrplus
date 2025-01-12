package routes

import (
	"server/internal/auth"
	"server/internal/subscriptions"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r chi.Router) {
	r.Route("/api", func(api chi.Router) {
		// auth
		api.Post("/login", auth.LoginHandler)
		api.Post("/register", auth.RegisterHandler)

		// subscriptions
		api.Route("/subscriptions", func(subs chi.Router) {
			subs.Get("/", subscriptions.GetSubscriptionsHandler)
			subs.Post("/", subscriptions.CreateSubscriptionHandler)

			subs.Get("/{id}", subscriptions.GetSubscriptionHandler)
			subs.Put("/{id}", subscriptions.UpdateSubscriptionHandler)
			subs.Delete("/{id}", subscriptions.DeleteSubscriptionHandler)
		})
	})
}
