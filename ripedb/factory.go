package ripedb

import (
	"errors"
	"github.com/jpbede/go-ripedb/ripedb/internal/transport"
	"net/http"
)

// APIEndpoint represents a api endpoint
type APIEndpoint string

const (
	APIEndpointLive APIEndpoint = "https://rest.db.ripe.net/ripe"
	APIEndpointTest APIEndpoint = "https://rest-test.db.ripe.net/test"
)

// New creates a new Client with password and context
func New(password string) (*Client, error) {
	return NewWithOptions(WithAPIEndpoint(APIEndpointLive), WithCredentials(password))
}

// NewWithClient creates a new Client with a given http.Client
func NewWithClient(httpClient *http.Client, password string) (*Client, error) {
	return NewWithOptions(WithAPIEndpoint(APIEndpointLive), WithHTTPClient(httpClient), WithCredentials(password))
}

// NewWithOptions creates a new Client with given options
func NewWithOptions(options ...ClientOption) (*Client, error) {
	c := &Client{}

	// always create a base transport it can be overwritten with options
	c.transport = transport.New(string(APIEndpointLive))

	// run given options
	for _, option := range options {
		option(c)
	}

	// check if there are credentials
	if !c.transport.HasCredentials() {
		return nil, errors.New("no api credentials supplied")
	}

	return c, nil
}
