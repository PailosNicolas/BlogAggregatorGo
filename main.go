package main

import (
	"net/http"
	"os"

	v1 "github.com/PailosNicolas/BlogAggregatorGo/v1"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{}))
	rV1 := chi.NewRouter()
	r.Mount("/v1", rV1)

	rV1.Get("/readiness", v1.ReadinessHandler)

	http.ListenAndServe(":"+port, r)
}
