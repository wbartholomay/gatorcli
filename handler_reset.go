package main

import (
	"context"
	"fmt"
	"os"
)

func handlerReset(s *state, cmd command) error {

	err := s.db.DeleteAllUsers(context.Background())
	if err != nil{
		fmt.Printf("deleting all users has failed %v\n", err)
		os.Exit(1)
	}

	err = s.db.DeleteAllFeeds(context.Background())
	if err != nil {
		fmt.Printf("deleting all feeds has failed %v\n", err)
		os.Exit(1)
	}

	fmt.Println("All users and feeds deleted successfully.")

	return nil
}