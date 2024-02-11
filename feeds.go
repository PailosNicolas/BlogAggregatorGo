package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/PailosNicolas/BlogAggregatorGo/helpers"
	"github.com/PailosNicolas/BlogAggregatorGo/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (cfg *apiConfig) HandlerCreateNewFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	type createNewFeedDTO struct {
		Feed       database.Feed      `json:"feed"`
		FeedFollow database.FeedsUser `json:"feed_follow"`
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

	follow, err := cfg.DB.FollowByFeedId(r.Context(), database.FollowByFeedIdParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    newFeed.ID,
		UserID:    user.ID,
	})

	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error creating feed")
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, createNewFeedDTO{
		Feed:       newFeed,
		FeedFollow: follow,
	})
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

func (cfg *apiConfig) HandlerDeleteFollowFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	idStr := chi.URLParam(r, "feedFollowID")

	idFeedFollow, err := uuid.Parse(idStr)

	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error getting parameter")
	}

	feedFollow, err := cfg.DB.GetFeedFollowById(r.Context(), idFeedFollow)

	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error")
		return
	}

	if feedFollow.UserID != user.ID {
		helpers.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	err = cfg.DB.DeleteFeedFollowByID(r.Context(), database.DeleteFeedFollowByIDParams{
		ID:     feedFollow.ID,
		UserID: user.ID,
	})

	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error deleting follow")
		return
	}

	helpers.RespondWithOK(w)
}

func (cfg *apiConfig) HandlerGetFollowFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	allFeddFollows, err := cfg.DB.GetAllFeedFollowByUserId(r.Context(), user.ID)

	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error getting feed")
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, allFeddFollows)
}
