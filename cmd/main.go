package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

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

	route.Setup(env, timeout, db, r)

	http.ListenAndServe(env.ServerAddress, r)
}
