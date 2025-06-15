package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/wbartholomay/gatorcli/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("not enough arguments. Example usage: addfeed <name> <url>")

	}

	name := cmd.args[0]
	url := cmd.args[1]

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUsername)
	if err != nil {
		return fmt.Errorf("no user found in DB for username, are you logged in? err: %w", err)
	}

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

	fmt.Println(feed)
	return nil
}