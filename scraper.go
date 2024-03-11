package main

import (
	"log"
	"time"

	"github.com/myk-francis/learning-go/internal/database"
)

func startScrapint(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Printf("Scraping on %v goroutines every %s duration", concurrency, timeBetweenRequest)
}