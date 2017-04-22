package layer_7

type ApplicationRule struct {
	Name   *Name   `xml:" name,omitempty" json:"name,omitempty"`
	Script *Script `xml:" script,omitempty" json:"script,omitempty"`
}

type Name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	ApplicationRule *ApplicationRule `xml:" applicationRule,omitempty" json:"applicationRule,omitempty"`
}

type Script struct {
	Text string `xml:",chardata" json:",omitempty"`
}
