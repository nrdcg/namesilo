// Package namesilo A Go client library for accessing the Namesilo API.
package namesilo

import (
	"context"
	"errors"
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
}

// NewClient Creates a Namesilo client.
func NewClient(httpClient *http.Client, isProduction bool, isOTE bool) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	endpoint := DefaultAPIEndpoint
	if !isProduction {
		endpoint = SandboxAPIEndpoint
		if isOTE {
			endpoint = OTESandboxAPIEndpoint
		}
	}

	return &Client{
		Endpoint:   endpoint,
		HTTPClient: httpClient,
	}
}

// NewClientWithAPIKey Creates a Namesilo client with API Key.
func NewClientWithAPIKey(httpClient *http.Client, apiKey string, isProduction bool, isOTE bool) (*Client, error) {
	if apiKey == "" {
		return nil, errors.New("credentials missing: API key")
	}
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	// Extract existing transport if present, otherwise use DefaultTransport
	baseTransport := http.DefaultTransport
	if httpClient != nil && httpClient.Transport != nil {
		baseTransport = httpClient.Transport
	}

	// Initialize TokenTransport with API key, wrapping the existing transport
	tokenTransport := &TokenTransport{
		apiKey:    apiKey,
		Transport: baseTransport, // Preserve existing transport
	}

	// Create a new http.Client from the previous one that applies TokenTransport on top
	newHTTPClient := *httpClient
	newHTTPClient.Transport = tokenTransport

	// Set API endpoint based on environment
	endpoint := DefaultAPIEndpoint
	if !isProduction {
		endpoint = SandboxAPIEndpoint
		if isOTE {
			endpoint = OTESandboxAPIEndpoint
		}
	}

	return &Client{
		Endpoint:   endpoint,
		HTTPClient: &newHTTPClient,
	}, nil
}

func (c *Client) get(ctx context.Context, name string, params interface{}) (*http.Response, error) {
	uri, err := url.Parse(fmt.Sprintf("%s/%s", c.Endpoint, name))
	if err != nil {
		return nil, err
	}

	if params != nil {
		v, err := querystring.Values(params)
		if err != nil {
			return nil, err
		}
		uri.RawQuery = v.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri.String(), http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("failed creating get request: %w", err)
	}

	return c.HTTPClient.Do(req)
}
