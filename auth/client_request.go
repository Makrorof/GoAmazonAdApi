package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"

	"github.com/Makrorof/GoAmazonAdApi"
)

type RequestClient struct {
	client *Client

	retryCount    int
	retryDuration time.Duration

	endpoint AMAZON_ENDPOINT
}

func NewRequestClient(client *Client, endpoint AMAZON_ENDPOINT, retryCount int, retryDuration time.Duration) *RequestClient {
	return &RequestClient{
		client:        client,
		retryCount:    retryCount,
		retryDuration: retryDuration,
		endpoint:      endpoint,
	}
}

func (c *RequestClient) request(ctx context.Context, method string, apiPath string, header map[string][]string, body io.Reader, outBody any) error {
	rv := reflect.ValueOf(outBody)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return errors.New("outBody must be a non-nil pointer")
	}

	var lastErr error
	var onceUnauthorizedError bool

	for i := 0; i <= c.retryCount; i++ {
		req, err := http.NewRequestWithContext(ctx, method, c.endpoint.Join(apiPath), body)
		if err != nil {
			return err
		}

		h := map[string][]string{
			"Accept":                          {"*/*"},
			"Content-Type":                    {"application/json"}, //
			"Amazon-Advertising-API-ClientId": {c.client.clientId},
			"User-Agent":                      {c.client.UserAgent},
		}

		for k, v := range header {
			h[k] = v
		}

		req.Header = h

		resp, err := c.client.client.Do(req)
		if err != nil {
			lastErr = err
			time.Sleep(500 * time.Millisecond)
			continue
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body) //Performans sorunu yaratabilir: json.NewDecoder(resp.Body).Decode(x) kullanilabilir.
		if err != nil {
			return err
		}

		if GoAmazonAdApi.DEBUG_PRINT_API_BODY {
			fmt.Printf("DEBUG | endpoint: %s, apiPath: %s, body: %s \n", c.endpoint, apiPath, string(body))
		}

		//success message parse
		//100 - 399
		if resp.StatusCode < 400 {
			if err := json.Unmarshal(body, outBody); err != nil {
				return err
			}
			return nil
		}

		//AMAZON API: This is the universal format for error response objects.
		//her turlu error datasi veriyor.
		amazonErrorData := new(GoAmazonAdApi.AmazonError)
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

// Returns an error, which can be either a standard error or an GoAmazonAdApi.AmazonError.
func (c *RequestClient) GET(ctx context.Context, apiPath string, header map[string][]string, outBody any) error {
	return c.request(ctx, http.MethodGet, apiPath, header, nil, outBody)
}

// Returns an error, which can be either a standard error or an GoAmazonAdApi.AmazonError.
func (c *RequestClient) POST(ctx context.Context, apiPath string, header map[string][]string, inBody any, outBody any) error {
	body, err := json.Marshal(inBody)
	if err != nil {
		return err
	}

	return c.request(ctx, http.MethodPost, apiPath, header, bytes.NewBuffer(body), outBody)
}

func (c *RequestClient) checkRequestStatus(statusCode int) bool {
	//Amazon yogunluktan dolayi 500 - 502 - 504 - 429 verebilecegini soyluyor. Digerleri request hatasi veya izin hatasi olarak yorumluyoruz.
	//Ek olarak gecmis token olabilecegini dusunerek 401 bir seferlik kabul ediyoruz.
	return (statusCode >= 500 && statusCode <= 599 || statusCode == http.StatusTooManyRequests)
}
