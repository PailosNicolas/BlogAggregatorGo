package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

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

	go startScraping(
		config.DB,
		5,
		time.Minute,
	)

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
	rV1.Get("/feeds", config.HandlerGetAllFeeds)
	rV1.Post("/feed_follows", config.middlewareAuth(config.HandlerFollowFeed))
	rV1.Delete("/feed_follows/{feedFollowID}", config.middlewareAuth(config.HandlerDeleteFollowFeed))
	rV1.Get("/feed_follows", config.middlewareAuth(config.HandlerGetFollowFeed))

	//Posts
	rV1.Get("/posts", config.middlewareAuth(config.HandlerGetPostsByUser))

	log.Fatal(http.ListenAndServe(":"+port, r))
}
