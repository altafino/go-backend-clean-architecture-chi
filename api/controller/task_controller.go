package controller

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/altafino/ivisual/domain"
)

type TaskController struct {
	TaskUsecase domain.TaskUsecase
}

func (tc *TaskController) Create(w http.ResponseWriter, r *http.Request) {
	var task domain.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, jsonError(err.Error()), http.StatusBadRequest)
		return
	}

	userID := r.Context().Value("x-user-id").(string)
	task.ID = primitive.NewObjectID()

	task.UserID, err = primitive.ObjectIDFromHex(userID)
	if err != nil {
		http.Error(w, jsonError(err.Error()), http.StatusBadRequest)
		return
	}

	err = tc.TaskUsecase.Create(r.Context(), &task)
	if err != nil {
		http.Error(w, jsonError(err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(domain.SuccessResponse{
		Message: "Task created successfully",
	})
}

func (u *TaskController) Fetch(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("x-user-id").(string)

	tasks, err := u.TaskUsecase.FetchByUserID(r.Context(), userID)
	if err != nil {
		http.Error(w, jsonError(err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}
