package rss

import (
	"context"
	"fmt"
	"io"
	"net/http"
)


func makeGetReqWithContext(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return []byte{}, err
	}
	req.Header.Set("User-Agent", "gator")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, fmt.Errorf("request to %s failed: %w", url, err)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{},fmt.Errorf("error reading response: %w", err)
	}

	return data, nil
}