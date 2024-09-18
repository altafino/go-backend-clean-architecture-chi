package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	"github.com/altafino/go-backend-clean-architecture-chi/api/route"
	"github.com/altafino/go-backend-clean-architecture-chi/bootstrap"
)

func main() {

	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	r := chi.NewRouter()

	// CORS middleware
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"}, // Use your allowed origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Origin", "X-Requested-With"},
		ExposedHeaders:   []string{"Link", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	route.Setup(env, timeout, db, r)

	http.ListenAndServe(env.ServerAddress, r)
}
