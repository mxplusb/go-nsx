package go_nsx

type Application struct {
	Description    *Description    `xml:" description,omitempty" json:"description,omitempty"`
	Element        *Element        `xml:" element,omitempty" json:"element,omitempty"`
	Name           *Name           `xml:" name,omitempty" json:"name,omitempty"`
	ObjectId       *ObjectId       `xml:" objectId,omitempty" json:"objectId,omitempty"`
	ObjectTypeName *ObjectTypeName `xml:" objectTypeName,omitempty" json:"objectTypeName,omitempty"`
	Revision       *Revision       `xml:" revision,omitempty" json:"revision,omitempty"`
	Type           *Type           `xml:" type,omitempty" json:"type,omitempty"`
}

type ApplicationProtocol struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Description struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Element struct {
	ApplicationProtocol *ApplicationProtocol `xml:" applicationProtocol,omitempty" json:"applicationProtocol,omitempty"`
	Value               *Value               `xml:" value,omitempty" json:"value,omitempty"`
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
	Application *Application `xml:" application,omitempty" json:"application,omitempty"`
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
