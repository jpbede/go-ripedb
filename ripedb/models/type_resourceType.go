package models

type ResourceType string

const (
	ResourceTypeInetnum  ResourceType = "inetnum"
	ResourceTypeInet6num ResourceType = "inet6num"
	ResourceTypeRoute    ResourceType = "route"
	ResourceTypeRoute6   ResourceType = "route6"
	ResourceTypeDomain   ResourceType = "domain"
	ResourceTypeAutNum   ResourceType = "aut-num"
	ResourceTypeAsSet    ResourceType = "as-set"
	ResourceTypePerson   ResourceType = "person"
)
