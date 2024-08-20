package route

import (
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/altafino/ivisual/api/controller"
	"github.com/altafino/ivisual/bootstrap"
	"github.com/altafino/ivisual/domain"
	"github.com/altafino/ivisual/mongo"
	"github.com/altafino/ivisual/repository"
	"github.com/altafino/ivisual/usecase"
)

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, router chi.Router) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pc := &controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	}
	router.Get("/profile", pc.Fetch)
}
