package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/PailosNicolas/BlogAggregatorGo/helpers"
	"github.com/PailosNicolas/BlogAggregatorGo/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) HandlerCreateNewFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
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

	if params.URL == "" {
		helpers.RespondWithError(w, http.StatusInternalServerError, "URL can not be empty")
		return
	}

	newFeed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})

	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error creating feed")
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, newFeed)
}
