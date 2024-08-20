package route

import (
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/altafino/ivisual/api/middleware"
	"github.com/altafino/ivisual/bootstrap"
	"github.com/altafino/ivisual/mongo"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, r *chi.Mux) {
	// All Public APIs
	publicRouter := chi.NewRouter()
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)
	r.Mount("/public", publicRouter)

	// Middleware to verify AccessToken
	protectedRouter := chi.NewRouter()
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	NewProfileRouter(env, timeout, db, protectedRouter)
	NewTaskRouter(env, timeout, db, protectedRouter)
	r.Mount("/protected", protectedRouter)
}
