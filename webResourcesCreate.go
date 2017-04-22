package go_nsx

type Data struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Description struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Enabled struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Method struct {
	Attr_name string `xml:" name,attr"  json:",omitempty"`
	Data      *Data  `xml:" data,omitempty" json:"data,omitempty"`
}

type Name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	WebResource *WebResource `xml:" webResource,omitempty" json:"webResource,omitempty"`
}

type Url struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type WebResource struct {
	Description *Description `xml:" description,omitempty" json:"description,omitempty"`
	Enabled     *Enabled     `xml:" enabled,omitempty" json:"enabled,omitempty"`
	Method      *Method      `xml:" method,omitempty" json:"method,omitempty"`
	Name        *Name        `xml:" name,omitempty" json:"name,omitempty"`
	Url         *Url         `xml:" url,omitempty" json:"url,omitempty"`
}
