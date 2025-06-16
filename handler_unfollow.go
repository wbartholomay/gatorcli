package main


import (
	"context"
	"fmt"

	"github.com/wbartholomay/gatorcli/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {

	if len(cmd.args) == 0 {
		return fmt.Errorf("no url provided. Example usage: follow <url>")
	}
	url := cmd.args[0]

	unfollowParams := database.DeleteFeedFollowParams{
		UserID: user.ID,
		Url: url,
	}

	err := s.db.DeleteFeedFollow(context.Background(), unfollowParams)
	if err != nil {
		return fmt.Errorf("failed to unfollow feed: %w", err)
	}

	fmt.Printf("Successfully unfollowed feed with url: %s", url)
	return nil
}