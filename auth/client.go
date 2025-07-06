package auth

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"golang.org/x/oauth2"
)

type Client struct {
	//PRIVATE
	mu     sync.Mutex
	client *http.Client

	clientId     string
	clientSecret string
	tokenSource  oauth2.TokenSource

	//PUBLIC
	UserAgent string
}

func createOauthConfig(clientId, clientSecret, redirectURL string, scopes ...string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.amazon.com/ap/oa",
			TokenURL: "https://api.amazon.com/auth/o2/token",
		},
	}
}

func New(clientId string, clientSecret string) *Client {
	return NewClient(nil, clientId, clientSecret)
}

func NewClient(httpClient *http.Client, clientId string, clientSecret string) *Client {
	if httpClient == nil {
		httpClient = &http.Client{
			Transport: http.DefaultTransport,
		}
	}
	httpClient2 := *httpClient
	c := &Client{client: &httpClient2, clientId: clientId, clientSecret: clientSecret, UserAgent: "Makrorof/GoAmazonAdAPI"}
	c.init()
	return c
}

func NewClientWithEnv(ctx context.Context, httpClient *http.Client) *Client {
	clientId := os.Getenv("AMAZON_CLIENT_ID")
	clientSecret := os.Getenv("AMAZON_CLIENT_SECRET")
	//redirectURL := os.Getenv("AMAZON_REDIRECT_URL")
	accessToken := os.Getenv("AMAZON_ACCESS_TOKEN")
	refreshToken := os.Getenv("AMAZON_REFRESH_TOKEN")

	if len(strings.TrimSpace(clientId)) == 0 || len(strings.TrimSpace(clientSecret)) == 0 {
		log.Fatalln("Missing Client ID or Client Secret.")
		return nil
	}

	client := NewClient(httpClient, clientId, clientSecret)
	client = client.WithToken(ctx, accessToken, refreshToken)
	return client
}

func (c *Client) clone() *Client {
	c.mu.Lock()

	clone := Client{
		client:       &http.Client{},
		UserAgent:    c.UserAgent,
		clientId:     c.clientId,
		clientSecret: c.clientSecret,
	}
	c.mu.Unlock()
	if c.client != nil {
		clone.client.Transport = c.client.Transport
		clone.client.CheckRedirect = c.client.CheckRedirect
		clone.client.Jar = c.client.Jar
		clone.client.Timeout = c.client.Timeout
	}

	return &clone
}

func (c *Client) init() {

}

func (c *Client) WithTokenSource(ctx context.Context, tokenSource oauth2.TokenSource) *Client {
	c2 := c.clone()
	defer c2.init()

	c2.tokenSource = tokenSource
	c2.client.Transport = &oauth2.Transport{
		Source: tokenSource,
		Base:   c2.client.Transport,
	}

	return c2
}

func (c *Client) WithToken(ctx context.Context, accessToken string, refreshToken string) *Client {
	config := createOauthConfig(c.clientId, c.clientSecret, "")

	token := &oauth2.Token{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}

	return c.WithTokenSource(ctx, config.TokenSource(ctx, token))
}

func (c *Client) WithRefreshToken(ctx context.Context, refreshToken string) *Client {
	return c.WithToken(ctx, "", refreshToken)
}

func (c *Client) WithAccessToken(ctx context.Context, accessToken string) *Client {
	return c.WithToken(ctx, accessToken, "")
}

func (c *Client) WithCode(ctx context.Context, code string, redirectURL string, scopes ...string) (*Client, error) {
	config := createOauthConfig(c.clientId, c.clientSecret, redirectURL, scopes...)

	token, err := config.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	return c.WithTokenSource(ctx, config.TokenSource(ctx, token)), nil
}

func (c *Client) updateToken() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.tokenSource == nil {
		return errors.New("token source not found")
	}

	_, err := c.tokenSource.Token()
	return err
}

func (c *Client) RefreshToken() (string, error) {
	if c.tokenSource == nil {
		return "", nil
	}

	token, err := c.tokenSource.Token()
	if err != nil {
		return "", err
	}

	return token.RefreshToken, nil
}

func (c *Client) Token() (*AmazonToken, error) {
	if c.tokenSource == nil {
		return nil, nil
	}

	token, err := c.tokenSource.Token()
	if err != nil {
		return nil, err
	}

	return &AmazonToken{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		TokenType:    token.TokenType,
		ExpiresIn:    token.ExpiresIn,
	}, nil
}
