package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/wbartholomay/gatorcli/internal/database"
)

func handlerFollow(s *state, cmd command) error {

	if len(cmd.args) == 0 {
		return fmt.Errorf("no url provided. Example usage: follow <url>")
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("no feed found for given url %w", err)
	}

	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUsername)
	if err != nil {
		return fmt.Errorf("user not found, are you logged in? %w", err)
	}

	params := database.CreateFeedFollowParams{
		ID : uuid.New(),
		Name : feed.Name,
		Url: feed.Url,
		UserID: currentUser.ID,
		FeedID: feed.ID,
	}

	feed_follow, err := s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error following feed: %w", err)
	}

	fmt.Printf("Feed name: %s\nCurrent user: %s\n", feed_follow.FeedName, feed_follow.UserName)

	return nil
}