package go_nsx

type Action struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type AppliedTo struct {
	IsValid *IsValid `xml:" isValid,omitempty" json:"isValid,omitempty"`
	Name    *Name    `xml:" name,omitempty" json:"name,omitempty"`
	Type    *Type    `xml:" type,omitempty" json:"type,omitempty"`
	Value   *Value   `xml:" value,omitempty" json:"value,omitempty"`
}

type AppliedToList struct {
	AppliedTo *AppliedTo `xml:" appliedTo,omitempty" json:"appliedTo,omitempty"`
}

type Destination struct {
	IsValid *IsValid `xml:" isValid,omitempty" json:"isValid,omitempty"`
	Name    *Name    `xml:" name,omitempty" json:"name,omitempty"`
	Type    *Type    `xml:" type,omitempty" json:"type,omitempty"`
	Value   *Value   `xml:" value,omitempty" json:"value,omitempty"`
}

type DestinationPort struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Destinations struct {
	Excluded  string          `xml:" excluded,attr"  json:",omitempty"`
	Destination *Destination `xml:" destination,omitempty" json:"destination,omitempty"`
}

type IsValid struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Protocol struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	Section *Section `xml:" section,omitempty" json:"section,omitempty"`
}

type Rule struct {
	Disabled    string            `xml:" disabled,attr"  json:",omitempty"`
	Logged      string            `xml:" logged,attr"  json:",omitempty"`
	Action        *Action        `xml:" action,omitempty" json:"action,omitempty"`
	AppliedToList *AppliedToList `xml:" appliedToList,omitempty" json:"appliedToList,omitempty"`
	Destinations  *Destinations  `xml:" destinations,omitempty" json:"destinations,omitempty"`
	Name          *Name          `xml:" name,omitempty" json:"name,omitempty"`
	Services      *Services      `xml:" services,omitempty" json:"services,omitempty"`
	Sources       *Sources       `xml:" sources,omitempty" json:"sources,omitempty"`
	Text             string            `xml:",chardata" json:",omitempty"`
}

type Section struct {
	ManagedBy string   `xml:" managedBy,attr"  json:",omitempty"`
	Name      string   `xml:" name,attr"  json:",omitempty"`
	Type      string   `xml:" type,attr"  json:",omitempty"`
	Rule        *Rule `xml:" rule,omitempty" json:"rule,omitempty"`
}

type Service struct {
	DestinationPort *DestinationPort `xml:" destinationPort,omitempty" json:"destinationPort,omitempty"`
	Protocol        *Protocol        `xml:" protocol,omitempty" json:"protocol,omitempty"`
	SubProtocol     *SubProtocol     `xml:" subProtocol,omitempty" json:"subProtocol,omitempty"`
}

type Services struct {
	Service *Service `xml:" service,omitempty" json:"service,omitempty"`
}

type Source struct {
	IsValid *IsValid `xml:" isValid,omitempty" json:"isValid,omitempty"`
	Name    *Name    `xml:" name,omitempty" json:"name,omitempty"`
	Type    *Type    `xml:" type,omitempty" json:"type,omitempty"`
	Value   *Value   `xml:" value,omitempty" json:"value,omitempty"`
}

type Sources struct {
	Excluded string     `xml:" excluded,attr"  json:",omitempty"`
	Source     *Source `xml:" source,omitempty" json:"source,omitempty"`
}

type SubProtocol struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Type struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Value struct {
	Text string `xml:",chardata" json:",omitempty"`
}
