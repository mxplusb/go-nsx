package go_nsx

type Description struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ExtendedAttributes struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IsUniversal struct {
	Text bool `xml:",chardata" json:",omitempty"`
}

type Name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ObjectTypeName struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	SecurityTag *SecurityTag `xml:" securityTag,omitempty" json:"securityTag,omitempty"`
}

type SecurityTag struct {
	Description        *Description        `xml:" description,omitempty" json:"description,omitempty"`
	ExtendedAttributes *ExtendedAttributes `xml:" extendedAttributes,omitempty" json:"extendedAttributes,omitempty"`
	IsUniversal        *IsUniversal        `xml:" isUniversal,omitempty" json:"isUniversal,omitempty"`
	Name               *Name               `xml:" name,omitempty" json:"name,omitempty"`
	ObjectTypeName     *ObjectTypeName     `xml:" objectTypeName,omitempty" json:"objectTypeName,omitempty"`
	Type               *Type               `xml:" type,omitempty" json:"type,omitempty"`
}

type Type struct {
	TypeName *TypeName `xml:" typeName,omitempty" json:"typeName,omitempty"`
}

type TypeName struct {
	Text string `xml:",chardata" json:",omitempty"`
}
