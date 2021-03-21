package ripedb

import (
	"github.com/jpbede/go-ripedb/ripedb/internal/transport"
	"github.com/rs/zerolog"
	"net/http"
)

// ClientOption options for creating the client
type ClientOption func(*Client)

// WithAPIEndpoint supplies a API Endpoint to be used
func WithAPIEndpoint(api APIEndpoint) ClientOption {
	return func(c *Client) {
		c.transport.BaseURL = string(api)
	}
}

// WithHTTPClient supplies a already created http.Client
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.transport.HTTPClient = httpClient
	}
}

// WithLogger supplies a logger for the transport client
func WithLogger(logger *zerolog.Logger) ClientOption {
	return func(c *Client) {
		c.transport.SetLogger(logger)
	}
}

// WithCredentials supplies the api credentials
func WithCredentials(password string) ClientOption {
	return func(c *Client) {
		c.transport.Credentials = &transport.APICredentials{
			Password: password,
		}
	}
}
