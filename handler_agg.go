package main

import (
	"context"
	"fmt"

	"github.com/wbartholomay/gatorcli/internal/rss"
)

func handlerAgg(s *state, cmd command) error {
	feedUrl := "https://www.wagslane.dev/index.xml"

	feed, err := rss.FetchFeed(context.Background(), feedUrl)
	if err != nil { return err }

	fmt.Println(*feed)

	return nil
}