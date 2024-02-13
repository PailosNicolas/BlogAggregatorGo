package main

import (
	"encoding/json"
	"net/http"

	"github.com/PailosNicolas/BlogAggregatorGo/helpers"
	"github.com/PailosNicolas/BlogAggregatorGo/internal/database"
)

func (cfg *apiConfig) HandlerGetPostsByUser(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Limit int `json:"limit"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error decoding parameters")
		return
	}

	if params.Limit == 0 {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Limit can not be empty")
		return
	}

	posts, err := cfg.DB.GetPostByUser(r.Context(), database.GetPostByUserParams{
		UserID: user.ID,
		Limit:  int32(params.Limit),
	})

	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error getting posts")
		return
	}

	helpers.RespondWithJSON(w, http.StatusOK, posts)
}
