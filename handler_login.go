package main

import (
	"context"
	"fmt"
	"os"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		fmt.Println("no username provided. Example usage: login <username>")
		os.Exit(1)
	}

	username := cmd.args[0]
	_, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		fmt.Printf("query for user with username '%s' failed: %v\n", username, err)
		os.Exit(1)
	}

	err = s.cfg.SetUser(username)
	if err != nil { return err }

	fmt.Printf("Usermane successfully set to: %s\n", username)
	return nil
}