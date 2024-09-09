package route

import (
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/altafino/go-backend-clean-architecture-chi/api/middleware"
	"github.com/altafino/go-backend-clean-architecture-chi/bootstrap"
	"github.com/altafino/go-backend-clean-architecture-chi/mongo"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, r *chi.Mux) {
	// Public APIs
	r.Group(func(r chi.Router) {
		NewSignupRouter(env, timeout, db, r)
		NewLoginRouter(env, timeout, db, r)
		NewRefreshTokenRouter(env, timeout, db, r)
	})

	// Protected routes
	r.Group(func(r chi.Router) {
		r.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
		NewTaskRouter(env, timeout, db, r)
		NewProfileRouter(env, timeout, db, r)
	})
}
