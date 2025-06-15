package main

import (
	"fmt"
	"os"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		fmt.Println("no username provided. Example usage: login <username>")
		os.Exit(1)
	}

	username := cmd.args[0]
	err := s.cfg.SetUser(username)
	if err != nil { return err }

	fmt.Printf("Usermane successfully set to: %s\n", username)
	return nil
}