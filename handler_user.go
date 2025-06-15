package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("getting all users has failed %v", err)
	}

	for _, user := range users {
		fmt.Printf("* %s", user.Name)
		if user.Name == s.cfg.CurrentUsername {
			fmt.Print(" (current)")
		}
		fmt.Println()
	}

	return nil
}