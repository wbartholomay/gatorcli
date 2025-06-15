package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/wbartholomay/gatorcli/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {

	if len(cmd.args) == 0 {
		return fmt.Errorf("no url provided. Example usage: follow <url>")
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("no feed found for given url %w", err)
	}

	params := database.CreateFeedFollowParams{
		ID : uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: feed.ID,
	}

	feed_follow, err := s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error following feed: %w", err)
	}

	fmt.Printf("Feed name: %s\nCurrent user: %s\n", feed_follow.FeedName, feed_follow.UserName)

	return nil
}