package ripedb

import (
	"context"
	"github.com/jpbede/go-ripedb/ripedb/internal/transport"
	"github.com/jpbede/go-ripedb/ripedb/models"
)

// Client represents the main client
type Client struct {
	transport *transport.Client
}

// GetResource returns an object from the RIPE Database by key and resource type
func (c *Client) GetResource(ctx context.Context, key string, resourceType models.ResourceType) ([]*models.Resource, error) {
	var out models.WhoisResourceResponse
	if err := c.transport.Get(ctx, "/ripe/"+string(resourceType)+"/"+key, &out); err != nil {
		return nil, err
	}
	return out.Objects.Object, nil
}
