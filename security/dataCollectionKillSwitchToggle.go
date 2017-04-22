package security

type Action struct {
	Type  *Type  `xml:" type,omitempty" json:"type,omitempty"`
	Value *Value `xml:" value,omitempty" json:"value,omitempty"`
}

type Actions struct {
	Action *Action `xml:" action,omitempty" json:"action,omitempty"`
}

type Request struct {
	Actions *Actions `xml:" actions,omitempty" json:"actions,omitempty"`
}

type Root struct {
	Request *Request `xml:" request,omitempty" json:"request,omitempty"`
}

type Type struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Value struct {
	Text string `xml:",chardata" json:",omitempty"`
}
