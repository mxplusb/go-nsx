package ip

type Description struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type InheritanceAllowed struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Ipset struct {
	Description        *Description        `xml:" description,omitempty" json:"description,omitempty"`
	InheritanceAllowed *InheritanceAllowed `xml:" inheritanceAllowed,omitempty" json:"inheritanceAllowed,omitempty"`
	Name               *Name               `xml:" name,omitempty" json:"name,omitempty"`
	ObjectId           *ObjectId           `xml:" objectId,omitempty" json:"objectId,omitempty"`
	ObjectTypeName     *ObjectTypeName     `xml:" objectTypeName,omitempty" json:"objectTypeName,omitempty"`
	Revision           *Revision           `xml:" revision,omitempty" json:"revision,omitempty"`
	Type               *Type               `xml:" type,omitempty" json:"type,omitempty"`
	Value              *Value              `xml:" value,omitempty" json:"value,omitempty"`
}

type Name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ObjectId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ObjectTypeName struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Revision struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	Ipset *Ipset `xml:" ipset,omitempty" json:"ipset,omitempty"`
}

type Type struct {
	TypeName *TypeName `xml:" typeName,omitempty" json:"typeName,omitempty"`
}

type TypeName struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Value struct {
	Text string `xml:",chardata" json:",omitempty"`
}
