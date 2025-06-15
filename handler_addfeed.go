package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/wbartholomay/gatorcli/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("not enough arguments. Example usage: addfeed <name> <url>")

	}

	name := cmd.args[0]
	url := cmd.args[1]

	feedParams := database.CreateFeedParams{
		ID : uuid.New(),
		Name : name,
		Url : url,
		UserID: user.ID,
	}


	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return fmt.Errorf("error occurred while adding feed to DB: %w", err)
	}
	fmt.Printf("Successfully created feed: %v\n",feed)

	feedFollowParams := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: feed.ID,
	}
	feed_follow, err := s.db.CreateFeedFollow(context.Background(), feedFollowParams)
	if err != nil {
		return fmt.Errorf("error creating feed_follow: %w", err)
	}

	fmt.Printf("Successfully created feed_follow %v\n", feed_follow)
	return nil
}