package main

import (
	"context"
	"fmt"
	"os"
)

func handlerReset(s *state, cmd command) error {

	err := s.db.DeleteAllUsers(context.Background())
	if err != nil && err.Error() != `pq: table "users" does not exist`{
		fmt.Printf("dropping users table has failed %v\n", err)
		os.Exit(1)
	}

	fmt.Println("All users deleted successfully.")

	return nil
}