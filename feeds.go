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

func (cfg *apiConfig) HandlerGetAllFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetAllFeeds(r.Context())

	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error getting feeds")
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, feeds)
}

func (cfg *apiConfig) HandlerFollowFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error decoding parameters")
		return
	}

	feed, err := cfg.DB.GetFeedById(r.Context(), params.FeedId)

	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error getting feed")
		return
	}

	follow, err := cfg.DB.FollowByFeedId(r.Context(), database.FollowByFeedIdParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    feed.ID,
		UserID:    user.ID,
	})

	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error creating follow")
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, follow)
}
