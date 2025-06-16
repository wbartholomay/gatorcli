package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/wbartholomay/gatorcli/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit int32
	if len(cmd.args) == 0 {
		limit = 2
	} else {
		 limit64, err := strconv.ParseInt(cmd.args[0], 10, 32)
		if err != nil {
			return fmt.Errorf("expecting an int as an argument. Example usage: browse <limit (int)>")
		}

		limit = int32(limit64)
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID : user.ID,
		Limit: limit,
	})


	if err != nil {
		return fmt.Errorf("getting posts has failed %v", err)
	}

	for _, post := range posts {
		fmt.Printf("Title - %v\nDescription: %v\n\n", post.Title, post.Description)
	}

	return nil
}