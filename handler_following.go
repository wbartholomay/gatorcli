package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {

	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUsername)
	if err != nil {
		return fmt.Errorf("getting current user has failed, are you logged in? err: %w", err)
	}

	feed_follows, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		return fmt.Errorf("error getting feed follows: %w", err)
	}

	for _, feed_follow := range feed_follows {
		fmt.Println(feed_follow.FeedName)
	}

	return nil
}