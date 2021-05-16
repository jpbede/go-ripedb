package ripedb_test

import (
	"github.com/stretchr/testify/assert"
	"go.bnck.me/ripedb/ripedb"
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	c, err := ripedb.New("abc")
	assert.NoError(t, err)
	assert.NotNil(t, c)
}

func TestNewWithClient(t *testing.T) {
	c, err := ripedb.NewWithClient(&http.Client{}, "abc")
	assert.NoError(t, err)
	assert.NotNil(t, c)
}

func TestNewWithOptions(t *testing.T) {
	c, err := ripedb.NewWithOptions(ripedb.WithAPIEndpoint(ripedb.APIEndpointTest), ripedb.WithCredentials("abc"))
	assert.NoError(t, err)
	assert.NotNil(t, c)
}
