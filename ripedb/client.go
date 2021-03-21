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

// SetDryRun enables/disabled the dry run mode
func (c *Client) SetDryRun(enable bool) {
	c.transport.DryRun = enable
}

// GetResource returns an object from the RIPE Database by key and resource type
func (c *Client) GetResource(ctx context.Context, key string, resourceType models.ResourceType) ([]*models.Resource, error) {
	var out models.WhoisResource
	if err := c.transport.Get(ctx, "/"+string(resourceType)+"/"+key, &out); err != nil {
		return nil, err
	}
	return out.Objects.Object, nil
}

// CreateResource creates a resource in RIPE database
func (c *Client) CreateResource(ctx context.Context, resourceType models.ResourceType, object models.WhoisResource) ([]*models.Resource, error) {
	var out models.WhoisResource
	if err := c.transport.Post(ctx, "/"+string(resourceType), &out, transport.WithJSONRequestBody(object)); err != nil {
		return nil, err
	}
	return out.Objects.Object, nil
}

// DeleteResource deletes a resource in RIPE database
func (c *Client) DeleteResource(ctx context.Context, key string, resourceType models.ResourceType) ([]*models.Resource, error) {
	var out models.WhoisResource
	if err := c.transport.Delete(ctx, "/"+string(resourceType)+"/"+key, &out); err != nil {
		return nil, err
	}
	return out.Objects.Object, nil
}
