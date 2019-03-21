package namesilo

import (
	"fmt"
	"net/http"
)

// TokenTransport HTTP transport for API authentication
type TokenTransport struct {
	apiKey string

	// Transport is the underlying HTTP transport to use when making requests.
	// It will default to http.DefaultTransport if nil.
	Transport http.RoundTripper
}

func NewTokenTransport(apiKey string) (*TokenTransport, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("credentials missing: API key")
	}

	return &TokenTransport{apiKey: apiKey}, nil
}

// RoundTrip executes a single HTTP transaction
func (t *TokenTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	enrichedReq := &http.Request{}
	*enrichedReq = *req

	enrichedReq.Header = make(http.Header, len(req.Header))
	for k, s := range req.Header {
		enrichedReq.Header[k] = append([]string(nil), s...)
	}

	if t.apiKey != "" {
		query := req.URL.Query()
		query.Add("version", "1")
		query.Add("type", "xml")
		query.Add("key", t.apiKey)
		req.URL.RawQuery = query.Encode()
	}

	fmt.Println(req.URL.RawQuery)
	return t.transport().RoundTrip(enrichedReq)
}

// Wrap Wrap a HTTP client Transport with the TokenTransport
func (t *TokenTransport) Wrap(client *http.Client) *http.Client {
	backup := client.Transport
	t.Transport = backup
	client.Transport = t
	return client
}

// Client Creates a new HTTP client
func (t *TokenTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}

func (t *TokenTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}
