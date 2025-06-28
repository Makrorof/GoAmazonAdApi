package amazon

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"
)

type requestClient struct {
	client *Client

	retryCount    int
	retryDuration time.Duration

	endpoint AMAZON_ENDPOINT
}

func newRequestClient(client *Client, endpoint AMAZON_ENDPOINT, retryCount int, retryDuration time.Duration) *requestClient {
	return &requestClient{
		client:        client,
		retryCount:    retryCount,
		retryDuration: retryDuration,
		endpoint:      endpoint,
	}
}

// Returns an error, which can be either a standard error or an AmazonError.
func (c *requestClient) GET(ctx context.Context, apiPath string, target any) error {
	rv := reflect.ValueOf(target)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return errors.New("target must be a non-nil pointer")
	}

	var lastErr error
	var onceUnauthorizedError bool

	for i := 0; i <= c.retryCount; i++ {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.endpoint.Join(apiPath), nil)
		if err != nil {
			return err
		}

		req.Header = map[string][]string{
			"Accept":                          {"application/json"},
			"Amazon-Advertising-API-ClientId": {c.client.clientId},
			"User-Agent":                      {c.client.UserAgent},
		}
		resp, err := c.client.client.Do(req)
		if err != nil {
			lastErr = err
			time.Sleep(500 * time.Millisecond)
			continue
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		if DEBUG_PRINT_API_BODY {
			fmt.Printf("DEBUG | endpoint: %s, apiPath: %s, body: %s \n", c.endpoint, apiPath, string(body))
		}

		//success message parse
		//100 - 399
		if resp.StatusCode < 400 {
			if err := json.Unmarshal(body, target); err != nil {
				return err
			}
			return nil
		}

		//AMAZON API: This is the universal format for error response objects.
		//her turlu error datasi veriyor.
		amazonErrorData := new(AmazonError)
		if err := json.Unmarshal(body, amazonErrorData); err != nil {
			return err
		}

		if resp.StatusCode == http.StatusUnauthorized && !onceUnauthorizedError {
			//Tek seferlik token guncellemesi denenecek.
			onceUnauthorizedError = true
			//Update token.
			if terr := c.client.updateToken(); terr != nil {
				return amazonErrorData //amazon hatasi daha onemli o yuzden terr kullanilmadi.
			}
		} else if !c.checkRequestStatus(resp.StatusCode) {
			//Eger hata tekrar denenmeyecek bir sekildeyse amazonError donucek.
			return amazonErrorData
		}

		lastErr = amazonErrorData
		time.Sleep(time.Duration(i+1) * 500 * time.Millisecond) // exponential-ish backoff
	}

	return lastErr
}

func (c *requestClient) checkRequestStatus(statusCode int) bool {
	//Amazon yogunluktan dolayi 500 - 502 - 504 - 429 verebilecegini soyluyor. Digerleri request hatasi veya izin hatasi olarak yorumluyoruz.
	//Ek olarak gecmis token olabilecegini dusunerek 401 bir seferlik kabul ediyoruz.
	return (statusCode >= 500 && statusCode <= 599 || statusCode == http.StatusTooManyRequests)
}
