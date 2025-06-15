package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("no username provided. Example usage: login <username>")
	}

	username := cmd.args[0]
	_, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("query for user with username '%s' failed: %v", username, err)
	}

	err = s.cfg.SetUser(username)
	if err != nil { return err }

	fmt.Printf("Usermane successfully set to: %s\n", username)
	return nil
}