package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/wbartholomay/gatorcli/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("no username provided. Example usage: login <username>")
	}

	username := cmd.args[0]

	userParams := database.CreateUserParams{
		ID : uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: username,
	}

	user, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		return fmt.Errorf("an error has occurred: %v", err)
	}

	s.cfg.SetUser(username)
	fmt.Printf("User has been created: \n%v\n", user)

	return nil
}