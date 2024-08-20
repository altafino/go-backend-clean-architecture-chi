package controller

import (
	"encoding/json"
	"net/http"

	"github.com/altafino/ivisual/domain"
)

type ProfileController struct {
	ProfileUsecase domain.ProfileUsecase
}

func (pc *ProfileController) Fetch(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("x-user-id").(string)

	profile, err := pc.ProfileUsecase.GetProfileByID(r.Context(), userID)
	if err != nil {
		http.Error(w, jsonError(err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profile)
}
