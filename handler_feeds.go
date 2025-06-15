package main

import (
	"context"
	"fmt"
	"os"
)

func handlerFeeds(s *state, cmd command) error {

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		fmt.Printf("getting all feeds has failed %v\n", err)
		os.Exit(1)
	}

	for _, feed := range feeds {
		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return err
		}

		fmt.Printf("Name: %s\nURL: %s\nCreated by: %s\n", feed.Name, feed.Url, user.Name)
	}

	return nil
}