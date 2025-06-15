package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {

	err := s.db.DeleteAllUsers(context.Background())
	if err != nil{
		return fmt.Errorf("deleting all users has failed %v", err)
	}

	//not necessary due to cascading
	err = s.db.DeleteAllFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("deleting all feeds has failed %v", err)
	}

	fmt.Println("All users and feeds deleted successfully.")

	return nil
}