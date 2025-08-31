// Package namesilo A Go client library for accessing the Namesilo API.
package namesilo

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"time"

	querystring "github.com/google/go-querystring/query"
)

const (
	// DefaultAPIEndpoint The default API endpoint.
	DefaultAPIEndpoint = "https://www.namesilo.com/api"

	// SandboxAPIEndpoint The sandbox API endpoint.
	SandboxAPIEndpoint = "https://sandbox.namesilo.com/api"

	// OTEAPIEndpoint The OTE sandbox API endpoint.
	OTEAPIEndpoint = "https://ote.namesilo.com/api"
)

// Response Codes.
const (
	SuccessfulAPIOperation = "300"
	SuccessfulRegistration = "301"
	SuccessfulOrder        = "302"
)

// Client the Namesilo client.
type Client struct {
	apiKey string

	Endpoint   *url.URL
	HTTPClient *http.Client
}

// NewClient Creates a Namesilo client.
func NewClient(apiKey string) *Client {
	endpoint, _ := url.Parse(DefaultAPIEndpoint)

	return &Client{
		apiKey:     apiKey,
		Endpoint:   endpoint,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *Client) get(ctx context.Context, name string, params any) (*http.Response, error) {
	endpoint := c.Endpoint.JoinPath(name)

	if params != nil {
		v, err := querystring.Values(params)
		if err != nil {
			return nil, err
		}

		endpoint.RawQuery = v.Encode()
	}

	query := endpoint.Query()
	query.Set("version", "1")
	query.Set("type", "xml")
	query.Set("key", c.apiKey)
	endpoint.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), http.NoBody)
	if err != nil {
		return nil, err
	}

	return c.HTTPClient.Do(req)
}

func GetEndpoint(prod, ote bool) (*url.URL, error) {
	if prod && ote {
		return nil, errors.New("prod and ote are mutually exclusive")
	}

	if prod {
		return url.Parse(DefaultAPIEndpoint)
	}

	if ote {
		return url.Parse(OTEAPIEndpoint)
	}

	return url.Parse(SandboxAPIEndpoint)
}
