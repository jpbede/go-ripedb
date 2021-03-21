package models

type WhoisResourceResponse struct {
	Link    *Link    `json:"link"`
	Objects *Objects `json:"objects"`
	TaC     *Link    `json:"terms-and-conditions"`
}

type Link struct {
	Type string `json:"type"`
	HRef string `json:"href"`
}

type Objects struct {
	Object []*Resource `json:"object"`
}

type Resource struct {
	ResourceType   ResourceType
	PrimaryKey     *Attributes `json:"primary-key"`
	Attributes     *Attributes
	Source         *Source
	Link           *Link  `json:"link,omitempty"`
	ReferencedType string `json:"referenced-type,omitempty"`
}

type Source struct {
	ID string
}

type Attributes struct {
	Attribute []*Attribute
}

type Attribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
