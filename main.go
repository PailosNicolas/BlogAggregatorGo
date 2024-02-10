package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/PailosNicolas/BlogAggregatorGo/internal/database"
	v1 "github.com/PailosNicolas/BlogAggregatorGo/v1"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	dbURL := os.Getenv("CONN")
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		println("error starting DB")
		return
	}

	config := apiConfig{
		DB: database.New(db),
	}

	port := os.Getenv("PORT")

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{}))
	rV1 := chi.NewRouter()
	r.Mount("/v1", rV1)

	// these two are just for testing purpuses
	rV1.Get("/readiness", v1.ReadinessHandler)
	rV1.Get("/err", v1.ErrHandler)

	// User
	rV1.Post("/users", config.HandlerCreateNewUser)
	rV1.Get("/users", config.middlewareAuth(config.HandlerGetUserByApiKey))

	//Feeds
	rV1.Post("/feeds", config.middlewareAuth(config.HandlerCreateNewFeed))

	log.Fatal(http.ListenAndServe(":"+port, r))
}
