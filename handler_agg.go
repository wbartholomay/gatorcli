package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
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

	_, err = saveFeed(s, feed, nextFeed)
	if err != nil {
		return fmt.Errorf("error saving post: %w", err)
	}

	return nil

}

func saveFeed(s *state, feed *rss.RSSFeed, dbFeed database.Feed) (database.Post, error) {
	createPostParams := database.CreatePostParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Title: feed.Channel.Title,
		Url: dbFeed.Url,
		Description: sql.NullString{
			String: feed.Channel.Description,
			Valid: true,
		},
		PublishedAt: dbFeed.CreatedAt,
		FeedID: dbFeed.ID,
	}

	post, err := s.db.CreatePost(context.Background(), createPostParams)
	if err != nil && err.Error() != "post with that URL already exists" {
		return database.Post{}, err
	}

	return post, nil
}