package main

import (
	"context"
	"log"
	"net/http"

	"server/configs"
	"server/internal/auth"
	"server/internal/middlewares"
	"server/internal/routes"
	"server/internal/subscriptions"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// port of the app and mongo uri
	port := configs.GetEnv("PORT")
	uri := configs.GetEnv("MONGODB_URI")

	// initializing the router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middlewares.AuthMiddleware)

	// setting up the mongo client
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	// initializing repos
	authRepo := auth.NewAuthRepository(client)
	subscriptionRepo := subscriptions.NewSubscriptionRepository(client)

	// initializing services
	authService := auth.NewAuthService(authRepo)
	subscriptionService := subscriptions.NewSubscriptionService(subscriptionRepo)

	// initializing handlers
	authHandler := auth.NewAuthHandler(authService)
	subscriptionHandler := subscriptions.NewSubscriptionHandler(subscriptionService)

	// setting up the routes
	routes.SetupRoutes(r, authHandler, subscriptionHandler)

	// closing the mongo client
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	// starting the server
	log.Println("Server running on port: " + port + "\n" + "http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
