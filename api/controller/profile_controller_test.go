package controller_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/altafino/go-backend-clean-architecture-chi/api/controller"
	"github.com/altafino/go-backend-clean-architecture-chi/domain"
	"github.com/altafino/go-backend-clean-architecture-chi/domain/mocks"
)

func setUserID(userID string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "x-user-id", userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func TestFetch(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		mockProfile := &domain.Profile{
			Name:  "Test Name",
			Email: "test@gmail.com",
		}

		userObjectID := primitive.NewObjectID()
		userID := userObjectID.Hex()

		mockProfileUsecase := new(mocks.ProfileUsecase)

		mockProfileUsecase.On("GetProfileByID", mock.Anything, userID).Return(mockProfile, nil)

		router := chi.NewRouter()

		rec := httptest.NewRecorder()

		pc := &controller.ProfileController{
			ProfileUsecase: mockProfileUsecase,
		}

		router.Use(setUserID(userID))
		router.Get("/protected/profile", pc.Fetch)

		body, err := json.Marshal(mockProfile)
		assert.NoError(t, err)

		bodyString := string(body)

		req := httptest.NewRequest(http.MethodGet, "/protected/profile", nil)
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		assert.JSONEq(t, bodyString, rec.Body.String())

		mockProfileUsecase.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		userObjectID := primitive.NewObjectID()
		userID := userObjectID.Hex()

		mockProfileUsecase := new(mocks.ProfileUsecase)

		customErr := errors.New("Unexpected")

		mockProfileUsecase.On("GetProfileByID", mock.Anything, userID).Return(nil, customErr)

		router := chi.NewRouter()

		rec := httptest.NewRecorder()

		pc := &controller.ProfileController{
			ProfileUsecase: mockProfileUsecase,
		}

		router.Use(setUserID(userID))
		router.Get("/protected/profile", pc.Fetch)

		body, err := json.Marshal(domain.ErrorResponse{Message: customErr.Error()})
		assert.NoError(t, err)

		bodyString := string(body)

		req := httptest.NewRequest(http.MethodGet, "/protected/profile", nil)
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		assert.JSONEq(t, bodyString, rec.Body.String())

		mockProfileUsecase.AssertExpectations(t)
	})

}
