package namesilo

import (
	"fmt"
	"net/http"
	"net/url"

	querystring "github.com/google/go-querystring/query"
)

// DefaultAPIEndpoint The default API endpoint.
const DefaultAPIEndpoint = "https://www.namesilo.com/api"

// Response Codes
const (
	SuccessfulAPIOperation = "300"
	SuccessfulRegistration = "301"
	SuccessfulOrder        = "302"
)

type Client struct {
	Endpoint   string
	HTTPClient *http.Client
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{
		Endpoint:   DefaultAPIEndpoint,
		HTTPClient: httpClient,
	}
}

func (c *Client) get(name string, params interface{}) (*http.Response, error) {
	uri, err := url.Parse(fmt.Sprintf("%s/%s", c.Endpoint, name))
	if err != nil {
		return nil, err
	}

	if params != nil {
		v, _ := querystring.Values(params)
		uri.RawQuery = v.Encode()

	}

	return c.HTTPClient.Get(uri.String())
}
