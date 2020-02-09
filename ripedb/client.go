package ripedb

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client struct {
	URL      APIUrl
	Password string
}

type APIUrl string

const (
	APIUrlLive APIUrl = "https://rest.db.ripe.net"
	APIUrlTest APIUrl = "https://rest-test.db.ripe.net"
)

type ResourceType string

const (
	ResourceTypeInetnum  ResourceType = "inetnum"
	ResourceTypeInet6num ResourceType = "inet6num"
	ResourceTypeRoute    ResourceType = "route"
	ResourceTypeRoute6   ResourceType = "route6"
	ResourceTypeDomain   ResourceType = "domain"
	ResourceTypeAutNum   ResourceType = "aut-num"
	ResourceTypeAsSet    ResourceType = "as-set"
)

func NewClient(apiURL APIUrl, password string) *Client {
	return &Client{URL: apiURL, Password: password}
}

func (c *Client) ResourcesFromAPIObject(object *RipeAPIObject) []Resource {
	result := make([]Resource, len(object.Objects.Object)-1)

	for _, resourceObject := range object.Objects.Object {
		res := Resource{}
		res.ResourceType = ResourceType(resourceObject.Type)
		res.PrimaryKey = resourceObject.PrimaryKey.Attribute[0].Value

		res.Attributes = make(map[string]string)

		for _, attr := range resourceObject.Attributes.Attribute {
			res.Attributes[attr.Name] = attr.Value
		}

		result = append(result, res)
	}

	return result
}

func (c *Client) GetResource(key string, resourceType ResourceType) ([]Resource, error) {
	resp, _ := http.Get(string(c.URL) + "/ripe/" + string(resourceType) + "/" + key + ".json")
	body, _ := ioutil.ReadAll(resp.Body)

	var f RipeAPIObject
	err := json.Unmarshal(body, &f)
	if err != nil {
		return nil, err
	}
	resources := c.ResourcesFromAPIObject(&f)

	return resources, nil
}

func (c *Client) CreateResource(resource Resource) {
}
