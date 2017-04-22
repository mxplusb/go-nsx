package go_nsx

type Action struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type AppliedTo struct {
	Type  *Type  `xml:" type,omitempty" json:"type,omitempty"`
	Value *Value `xml:" value,omitempty" json:"value,omitempty"`
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

type Destinations struct {
	Attr_excluded string       `xml:" excluded,attr"  json:",omitempty"`
	Destination   *Destination `xml:" destination,omitempty" json:"destination,omitempty"`
}

type IsValid struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Notes struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	Rule *Rule `xml:" rule,omitempty" json:"rule,omitempty"`
}

type Rule struct {
	Attr_disabled string         `xml:" disabled,attr"  json:",omitempty"`
	Attr_logged   string         `xml:" logged,attr"  json:",omitempty"`
	Action        *Action        `xml:" action,omitempty" json:"action,omitempty"`
	AppliedToList *AppliedToList `xml:" appliedToList,omitempty" json:"appliedToList,omitempty"`
	Destinations  *Destinations  `xml:" destinations,omitempty" json:"destinations,omitempty"`
	Name          *Name          `xml:" name,omitempty" json:"name,omitempty"`
	Notes         *Notes         `xml:" notes,omitempty" json:"notes,omitempty"`
	Services      *Services      `xml:" services,omitempty" json:"services,omitempty"`
	Sources       *Sources       `xml:" sources,omitempty" json:"sources,omitempty"`
}

type Service struct {
	Value *Value `xml:" value,omitempty" json:"value,omitempty"`
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
	Attr_excluded string  `xml:" excluded,attr"  json:",omitempty"`
	Source        *Source `xml:" source,omitempty" json:"source,omitempty"`
}

type Type struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Value struct {
	Text string `xml:",chardata" json:",omitempty"`
}
