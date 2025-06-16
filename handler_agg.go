package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/wbartholomay/gatorcli/internal/database"
	"github.com/wbartholomay/gatorcli/internal/rss"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("not enough arguments provided. Example usage: agg <time_interval>")
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("expecting a time interval: %w", err)
	}

	fmt.Printf("Collecting feeds every %v", timeBetweenReqs)

	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
			return err
		}
	}
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	nullTime := sql.NullTime {
		Time: time.Now().UTC(),
		Valid: true,
	}

	markFetchedParams := database.MarkFeedFetchedParams{
		ID: nextFeed.ID,
		LastFetchedAt: nullTime,
		UpdatedAt: time.Now().UTC(),
	}
	err = s.db.MarkFeedFetched(context.Background(), markFetchedParams)
	if err != nil {
		return fmt.Errorf("error marking feed fetched: %w", err)
	}

	feed, err := rss.FetchFeed(context.Background(), nextFeed.Url)
	if err != nil { return err }

	for _, item := range feed.Channel.Item {
		fmt.Println(item.Title)
	}

	return nil

}