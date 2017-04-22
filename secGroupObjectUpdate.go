package go_nsx

type ClientHandle struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ExtendedAttributes struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Id struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type InheritanceAllowed struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IsUniversal struct {
	Text string `xml:",chardata" json:",omitempty"`
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
	Securitygroup *Securitygroup `xml:" securitygroup,omitempty" json:"securitygroup,omitempty"`
}

type Scope struct {
	Id             *Id             `xml:" id,omitempty" json:"id,omitempty"`
	Name           *Name           `xml:" name,omitempty" json:"name,omitempty"`
	ObjectTypeName *ObjectTypeName `xml:" objectTypeName,omitempty" json:"objectTypeName,omitempty"`
	Revision       *Revision       `xml:" revision,omitempty" json:"revision,omitempty"`
}

type Securitygroup struct {
	ClientHandle       *ClientHandle       `xml:" clientHandle,omitempty" json:"clientHandle,omitempty"`
	ExtendedAttributes *ExtendedAttributes `xml:" extendedAttributes,omitempty" json:"extendedAttributes,omitempty"`
	InheritanceAllowed *InheritanceAllowed `xml:" inheritanceAllowed,omitempty" json:"inheritanceAllowed,omitempty"`
	IsUniversal        *IsUniversal        `xml:" isUniversal,omitempty" json:"isUniversal,omitempty"`
	Name               *Name               `xml:" name,omitempty" json:"name,omitempty"`
	ObjectId           *ObjectId           `xml:" objectId,omitempty" json:"objectId,omitempty"`
	ObjectTypeName     *ObjectTypeName     `xml:" objectTypeName,omitempty" json:"objectTypeName,omitempty"`
	Revision           *Revision           `xml:" revision,omitempty" json:"revision,omitempty"`
	Scope              *Scope              `xml:" scope,omitempty" json:"scope,omitempty"`
	Type               *Type               `xml:" type,omitempty" json:"type,omitempty"`
}

type Type struct {
	TypeName *TypeName `xml:" typeName,omitempty" json:"typeName,omitempty"`
}

type TypeName struct {
	Text string `xml:",chardata" json:",omitempty"`
}
