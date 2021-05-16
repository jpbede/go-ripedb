package models

type Resource struct {
	Type       ResourceType `json:"type"`
	Attributes map[string]string
}

func (r *Resource) TypeString() string {
	return string(r.Type)
}

func (r *Resource) RIPEWhoisResource() *WhoisResource {
	return nil
}
