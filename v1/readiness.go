package v1

import (
	"net/http"

	"github.com/PailosNicolas/BlogAggregatorGo/helpers"
)

type readinessDTO struct {
	Status string `json:"status"`
}

func ReadinessHandler(w http.ResponseWriter, request *http.Request) {
	helpers.RespondWithJSON(w, 200, readinessDTO{
		Status: "ok",
	})
}
