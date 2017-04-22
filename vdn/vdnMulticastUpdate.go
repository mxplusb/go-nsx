package go_nsx

type Desc struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type End struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type MulticastRange struct {
	Desc *Desc `xml:" desc,omitempty" json:"desc,omitempty"`
	End  *End  `xml:" end,omitempty" json:"end,omitempty"`
	Name *Name `xml:" name,omitempty" json:"name,omitempty"`
}

type Name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	MulticastRange *MulticastRange `xml:" multicastRange,omitempty" json:"multicastRange,omitempty"`
}
