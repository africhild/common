package http

import (
	"net/http"

	"github.com/africhild/common/pkg/auth"
)

// AuthenticatedHTTPClient is an interface for an http client which can automatically
// authenticate itself. This is useful for OAuth authentication, where the access token
// needs to be refreshed automatically. The signatures are a subset of http.Client,
// so it can be used as a (mostly) drop-in replacement.
type AuthenticatedHTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
	CloseIdleConnections()
}

type authenticatedClient struct {
	client   *http.Client
	strategy auth.Strategy
}

func NewAuthenticatedClient(strategy auth.Strategy) AuthenticatedHTTPClient {
	return &authenticatedClient{
		client:   &http.Client{},
		strategy: strategy,
	}
}

// func NewClient() AuthenticatedHTTPClient {
//     return &authenticatedClient{
//         client:   &http.Client{},
//         strategy: nil,
//     }
// }

func (c *authenticatedClient) Do(req *http.Request) (*http.Response, error) {
	if err := c.strategy.ApplyAuth(req); err != nil {
		return nil, err
	}
	return c.client.Do(req)
}

func (c *authenticatedClient) CloseIdleConnections() {
	c.client.CloseIdleConnections()
}
