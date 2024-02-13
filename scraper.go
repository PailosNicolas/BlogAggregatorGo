package main

import (
	"context"
	"database/sql"
	"log"
	"sync"
	"time"

	"github.com/PailosNicolas/BlogAggregatorGo/internal/database"
	"github.com/google/uuid"
)

func startScraping(
	db *database.Queries,
	concurrency int,
	timeBetweenRequests time.Duration,
) {
	log.Printf("Scraping started.")
	ticker := time.NewTicker(timeBetweenRequests)

	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))

		if err != nil {
			log.Print("error getting feeds", err)
			continue
		}

		wg := &sync.WaitGroup{}

		for _, feed := range feeds {
			wg.Add(1)
			go scrappeFeed(db, wg, feed)
		}

		wg.Wait()

	}
}

func scrappeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.UpdateLastUpdatedAt(context.Background(), feed.ID)

	if err != nil {
		log.Print("error updating feed", err)
		return
	}

	rss, err := urlToFeed(feed.Url)

	if err != nil {
		log.Print("error parsing rss", err)
		return
	}

	// quick and ugly fix por parser
	layout := "Mon, 02 Jan 2006 15:04:05 -0700"

	for _, item := range rss.Channel.Item {
		parsedTime, err := time.Parse(layout, item.PubDate)
		if err != nil {
			log.Println("Error parsing time:", err)
			continue
		}
		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title:     item.Title,
			Url:       item.Link,
			Description: sql.NullString{
				String: item.Text,
				Valid:  true,
			},
			PublishedAt: sql.NullTime{
				Time:  parsedTime,
				Valid: true,
			},
			FeedID: feed.ID,
		})

		if err != nil {
			log.Println("Error creating post:", err)
			continue
		}
	}

}
