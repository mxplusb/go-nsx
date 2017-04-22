package application

type ApplicationGroup struct {
	Description        *Description        `xml:" description,omitempty" json:"description,omitempty"`
	InheritanceAllowed *InheritanceAllowed `xml:" inheritanceAllowed,omitempty" json:"inheritanceAllowed,omitempty"`
	Name               *Name               `xml:" name,omitempty" json:"name,omitempty"`
	Revision           *Revision           `xml:" revision,omitempty" json:"revision,omitempty"`
}

type Description struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type InheritanceAllowed struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Revision struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	ApplicationGroup *ApplicationGroup `xml:" applicationGroup,omitempty" json:"applicationGroup,omitempty"`
}
