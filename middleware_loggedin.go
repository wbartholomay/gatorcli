package main

import (
	"context"
	"fmt"

	"github.com/wbartholomay/gatorcli/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUsername)
		if err != nil {
			return fmt.Errorf("getting current user has failed, are you logged in? err: %w", err)
		}

		return handler(s, cmd, user)
	}
}