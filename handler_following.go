package main

import (
	"context"
	"fmt"

	"github.com/wbartholomay/gatorcli/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {

	feed_follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error getting feed follows: %w", err)
	}

	for _, feed_follow := range feed_follows {
		fmt.Println(feed_follow.FeedName)
	}

	return nil
}