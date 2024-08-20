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

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, router chi.Router) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	router.Get("/task", tc.Fetch)
	router.Post("/task", tc.Create)
}
