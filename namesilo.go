// Package namesilo A Go client library for accessing the Namesilo API.
package namesilo

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
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

func (c *Client) do(ctx context.Context, name string, params, op any) error {
	endpoint := c.Endpoint.JoinPath(name)

	query, err := querystring.Values(params)
	if err != nil {
		return err
	}

	query.Set("version", "1")
	query.Set("type", "xml")
	query.Set("key", c.apiKey)

	endpoint.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), http.NoBody)
	if err != nil {
		return err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		data, _ := io.ReadAll(resp.Body)

		return fmt.Errorf("error: %d: %s", resp.StatusCode, string(data))
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(data, op)
	if err != nil {
		return fmt.Errorf("failed to decode: %w: %s", err, data)
	}

	return nil
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

type replyGetter interface {
	getCode() string
	getDetail() string
}

func checkReply(reply replyGetter) error {
	switch reply.getCode() {
	case SuccessfulAPIOperation:
		// Successful API operation
		return nil
	case SuccessfulRegistration:
		// Successful registration, but not all provided hosts were valid resulting in our nameservers being used
		return nil
	case SuccessfulOrder:
		// Successful order, but there was an error with the contact information provided so your account default contact profile was used
		return nil
	default:
		// error
		return fmt.Errorf("code: %s, details: %s", reply.getCode(), reply.getDetail())
	}
}
