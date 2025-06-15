package main

import (
	"context"
	"fmt"
	"os"
)

func handlerUsers(s *state, cmd command) error {

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		fmt.Printf("getting all users has failed %v\n", err)
		os.Exit(1)
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