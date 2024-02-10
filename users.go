package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/PailosNicolas/BlogAggregatorGo/helpers"
	"github.com/PailosNicolas/BlogAggregatorGo/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) HandlerCreateNewUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	ctx := r.Context()

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error decoding parameters")
		return
	}

	if params.Name == "" {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Name can not be empty")
		return
	}

	newUser, err := cfg.DB.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
	})

	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error creating user")
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, newUser)
}

func (cfg *apiConfig) HandlerGetUserByApiKey(w http.ResponseWriter, r *http.Request, user database.User) {
	helpers.RespondWithJSON(w, http.StatusOK, user)

}
