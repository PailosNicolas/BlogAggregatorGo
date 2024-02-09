package v1

import (
	"net/http"

	"github.com/PailosNicolas/BlogAggregatorGo/helpers"
)

func ErrHandler(w http.ResponseWriter, request *http.Request) {
	helpers.RespondWithError(w, 500, "Internal Server Error")
}
