package ripedb

// autogenerated ripe api object
type RipeAPIObject struct {
	Objects struct {
		Object []struct {
			Type string `json:"type"`
			Link struct {
				Type string `json:"type"`
				Href string `json:"href"`
			} `json:"link"`
			Source struct {
				ID string `json:"id"`
			} `json:"source"`
			PrimaryKey struct {
				Attribute []struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"attribute"`
			} `json:"primary-key"`
			Attributes struct {
				Attribute []struct {
					Name  string `json:"name"`
					Value string `json:"value"`
					Link  struct {
						Type string `json:"type"`
						Href string `json:"href"`
					} `json:"link,omitempty"`
					ReferencedType string `json:"referenced-type,omitempty"`
				} `json:"attribute"`
			} `json:"attributes"`
		} `json:"object"`
	} `json:"objects"`
	TermsAndConditions struct {
		Type string `json:"type"`
		Href string `json:"href"`
	} `json:"terms-and-conditions"`
}
