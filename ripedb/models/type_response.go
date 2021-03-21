package models

type WhoisResource struct {
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
	ResourceType   ResourceType `json:"resource-type,omitempty"`
	PrimaryKey     *Attributes  `json:"primary-key,omitempty"`
	Attributes     *Attributes  `json:"attributes"`
	Source         *Source      `json:"source"`
	Link           *Link        `json:"link,omitempty"`
	ReferencedType string       `json:"referenced-type,omitempty"`
}

type Source struct {
	ID string `json:"id"`
}

type Attributes struct {
	Attribute []*Attribute `json:"attribute"`
}

type Attribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
