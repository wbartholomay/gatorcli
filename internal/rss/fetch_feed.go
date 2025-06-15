package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
)

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	data, err := makeGetReqWithContext(ctx, feedURL)
	if err != nil {
		return nil, err
	}

	var feed RSSFeed
	err = xml.Unmarshal(data, &feed)
	if err != nil {
		return nil, fmt.Errorf("error converting XML to struct: %w", err)
	}

	cleanFeed(&feed)

	return &feed, nil
}

func cleanFeed(feed *RSSFeed) {
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	for i := range feed.Channel.Item {
		feed.Channel.Item[i].Title = html.UnescapeString(feed.Channel.Item[i].Title)
		feed.Channel.Item[i].Description = html.UnescapeString(feed.Channel.Item[i].Description)
	}
}