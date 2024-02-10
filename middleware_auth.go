package main

import (
	"net/http"
	"strings"

	"github.com/PailosNicolas/BlogAggregatorGo/helpers"
	"github.com/PailosNicolas/BlogAggregatorGo/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		api_key := r.Header.Get("Authorization")
		api_key = strings.ReplaceAll(api_key, "ApiKey ", "")

		if api_key == "" {
			helpers.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		user, err := cfg.DB.GetUserByApiKey(r.Context(), api_key)

		if err != nil {
			helpers.RespondWithError(w, http.StatusInternalServerError, "Error getting user")
			return
		}

		handler(w, r, user)
	}
}
