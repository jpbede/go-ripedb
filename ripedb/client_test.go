package ripedb

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient(APIUrlLive, "")

	assert.Equal(t, client, &Client{URL: APIUrlLive})
}

func TestGetResource(t *testing.T) {
	client := NewClient(APIUrlLive, os.Getenv("RIPE_PASSWORD"))
	resp, _ := client.GetResource("2001:67c:16c8::/48", ResourceTypeInet6num)

	assert.Equal(t, resp, []Resource([]Resource{Resource{ResourceType: "inet6num", PrimaryKey: "2001:67c:16c8::/48", Attributes: map[string]string{"admin-c": "JB16811-RIPE", "country": "DE", "created": "2017-01-12T15:30:40Z", "descr": "FastEthernet Infra", "geoloc": "50.1211277 8.4964814", "inet6num": "2001:67c:16c8::/48", "last-modified": "2020-02-09T13:53:51Z", "mnt-by": "FASTETHERNET-MNT", "netname": "DE-FE-INFRA", "org": "ORG-JB71-RIPE", "source": "RIPE", "sponsoring-org": "ORG-LIGC3-RIPE", "status": "ASSIGNED PI", "tech-c": "JB16811-RIPE"}}}))
}

func TestGetResourceRoute(t *testing.T) {
	client := NewClient(APIUrlLive, os.Getenv("RIPE_PASSWORD"))
	resp, _ := client.GetResource("185.120.22.0/24AS206479", ResourceTypeRoute)

	assert.Equal(t, resp, []Resource{Resource{ResourceType: "route", PrimaryKey: "185.120.22.0/24", Attributes: map[string]string{"created": "2017-03-12T16:31:58Z", "descr": "FastEthernet Infra IPv4", "last-modified": "2019-08-28T08:38:18Z", "mnt-by": "FASTETHERNET-MNT", "origin": "AS206479", "route": "185.120.22.0/24", "source": "RIPE"}}})
}
