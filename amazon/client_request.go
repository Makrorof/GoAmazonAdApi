package amazon

import (
	"context"
	"net/http"
	"time"
)

type requestClient[T any] struct {
	client *Client

	retryCount    int
	retryDuration time.Duration

	endpoint AMAZON_ENDPOINT
}

func newRequestClient[T any](client *Client, endpoint AMAZON_ENDPOINT, retryCount int, retryDuration time.Duration) *requestClient[T] {
	return &requestClient[T]{
		client:        client,
		retryCount:    retryCount,
		retryDuration: retryDuration,
		endpoint:      endpoint,
	}
}

func (c *requestClient[T]) GET(ctx context.Context, url string) (T, error) {
	var result T
	var lastErr error

	for i := 0; i <= c.retryCount; i++ {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return result, err
		}

		resp, err := c.client.client.Do(req)
		if err != nil {
			lastErr = err
			continue
		}

		defer resp.Body.Close()

		time.Sleep(c.retryDuration)
	}

	return result, lastErr
}
