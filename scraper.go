package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/PailosNicolas/BlogAggregatorGo/internal/database"
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
		log.Print("error getting rrs", err)
		return
	}

	for _, item := range rss.Channel.Item {
		log.Println("Found post", item)
	}

}
