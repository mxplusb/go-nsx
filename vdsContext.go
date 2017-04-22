package go_nsx

type Mtu struct {
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
	VdsContext *VdsContext `xml:" vdsContext,omitempty" json:"vdsContext,omitempty"`
}

type Switch struct {
	Name           *Name           `xml:" name,omitempty" json:"name,omitempty"`
	ObjectId       *ObjectId       `xml:" objectId,omitempty" json:"objectId,omitempty"`
	ObjectTypeName *ObjectTypeName `xml:" objectTypeName,omitempty" json:"objectTypeName,omitempty"`
	Revision       *Revision       `xml:" revision,omitempty" json:"revision,omitempty"`
	Type           *Type           `xml:" type,omitempty" json:"type,omitempty"`
}

type Teaming struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Type struct {
	TypeName *TypeName `xml:" typeName,omitempty" json:"typeName,omitempty"`
}

type TypeName struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type VdsContext struct {
	Mtu     *Mtu     `xml:" mtu,omitempty" json:"mtu,omitempty"`
	Switch  *Switch  `xml:" switch,omitempty" json:"switch,omitempty"`
	Teaming *Teaming `xml:" teaming,omitempty" json:"teaming,omitempty"`
}
