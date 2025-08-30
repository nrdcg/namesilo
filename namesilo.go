// Package namesilo A Go client library for accessing the Namesilo API.
package namesilo

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	querystring "github.com/google/go-querystring/query"
)

const (
	// DefaultAPIEndpoint The default API endpoint.
	DefaultAPIEndpoint = "https://www.namesilo.com/api"

	// SandboxAPIEndpoint The sandbox API endpoint.
	SandboxAPIEndpoint = "https://sandbox.namesilo.com/api"

	// OTESandboxAPIEndpoint The OTE sandbox API endpoint.
	OTESandboxAPIEndpoint = "https://ote.namesilo.com/api"
)

// Response Codes.
const (
	SuccessfulAPIOperation = "300"
	SuccessfulRegistration = "301"
	SuccessfulOrder        = "302"
)

// Client the Namesilo client.
type Client struct {
	Endpoint   string
	HTTPClient *http.Client
	apiKey     string
}

type Config struct {
	APIKey       string
	IsProduction bool
	IsOTE        bool
}

// NewClient Creates a Namesilo client.
func NewClient(httpClient *http.Client, cfg Config) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	// Set API endpoint based on environment
	endpoint := DefaultAPIEndpoint
	if !cfg.IsProduction {
		endpoint = SandboxAPIEndpoint
		if cfg.IsOTE {
			endpoint = OTESandboxAPIEndpoint
		}
	}

	return &Client{
		Endpoint:   endpoint,
		HTTPClient: httpClient,
		apiKey:     cfg.APIKey,
	}
}

func (c *Client) get(ctx context.Context, name string, params any) (*http.Response, error) {
	uri, err := url.Parse(fmt.Sprintf("%s/%s", c.Endpoint, name))
	if err != nil {
		return nil, err
	}

	if params != nil {
		var v url.Values
		v, err = querystring.Values(params)
		if err != nil {
			return nil, err
		}
		uri.RawQuery = v.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri.String(), http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("failed creating get request: %w", err)
	}

	query := req.URL.Query()
	query.Add("version", "1")
	query.Add("type", "xml")
	query.Add("key", c.apiKey)
	req.URL.RawQuery = query.Encode()

	return c.HTTPClient.Do(req)
}
