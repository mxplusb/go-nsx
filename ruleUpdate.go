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

type IsValid struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	Rule *Rule `xml:" rule,omitempty" json:"rule,omitempty"`
}

type Rule struct {
	Attr_disabled string         `xml:" disabled,attr"  json:",omitempty"`
	Attr_id       string         `xml:" id,attr"  json:",omitempty"`
	Attr_logged   string         `xml:" logged,attr"  json:",omitempty"`
	Action        *Action        `xml:" action,omitempty" json:"action,omitempty"`
	AppliedToList *AppliedToList `xml:" appliedToList,omitempty" json:"appliedToList,omitempty"`
	Name          *Name          `xml:" name,omitempty" json:"name,omitempty"`
}

type Type struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Value struct {
	Text string `xml:",chardata" json:",omitempty"`
}
