package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/wbartholomay/gatorcli/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		fmt.Println("no username provided. Example usage: login <username>")
		os.Exit(1)
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
		fmt.Printf("An error has occurred: %v\n", err)
		os.Exit(1)
	}

	s.cfg.SetUser(username)
	fmt.Printf("User has been created: \n%v\n", user)

	return nil
}