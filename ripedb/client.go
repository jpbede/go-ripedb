package ripedb

import (
	"context"
	"go.bnck.me/ripedb/ripedb/internal/transport"
	"go.bnck.me/ripedb/ripedb/models"
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
func (c *Client) GetResource(ctx context.Context, key string, resourceType models.ResourceType) ([]*models.Object, error) {
	var out models.WhoisResource
	if err := c.transport.Get(ctx, "/"+string(resourceType)+"/"+key, &out); err != nil {
		return nil, err
	}
	return out.Objects.Object, nil
}

// CreateResource creates a resource in RIPE database
func (c *Client) CreateResource(ctx context.Context, resource models.Resource) ([]*models.Object, error) {
	var out models.WhoisResource
	if err := c.transport.Post(ctx, "/"+resource.TypeString(), &out, transport.WithJSONRequestBody(resource.RIPEWhoisResource())); err != nil {
		return nil, err
	}
	return out.Objects.Object, nil
}

// DeleteResource deletes a resource in RIPE database
func (c *Client) DeleteResource(ctx context.Context, key string, resourceType models.ResourceType) ([]*models.Object, error) {
	var out models.WhoisResource
	if err := c.transport.Delete(ctx, "/"+string(resourceType)+"/"+key, &out); err != nil {
		return nil, err
	}
	return out.Objects.Object, nil
}
