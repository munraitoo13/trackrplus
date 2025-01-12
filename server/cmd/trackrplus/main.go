package main

import (
	"context"
	"log"
	"net/http"

	"server/configs"
	"server/internal/auth"
	"server/internal/routes"

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

	// setting up the routes
	routes.SetupRoutes(r)

	// setting up the mongo client
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	// initializing repos
	auth.RepoInit(client)

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
