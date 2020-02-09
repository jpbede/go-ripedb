package ripedb

type Resource struct {
	ResourceType ResourceType
	PrimaryKey   string
	Attributes   map[string]string
}
