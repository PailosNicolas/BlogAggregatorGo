package main

import (
	"net/http"
	"os"

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

	http.ListenAndServe(":"+port, r)
}
