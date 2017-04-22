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
	Section *Section `xml:" section,omitempty" json:"section,omitempty"`
}

type Rule struct {
	Disabled    string            `xml:" disabled,attr"  json:",omitempty"`
	Id          string            `xml:" id,attr"  json:",omitempty"`
	Logged      string            `xml:" logged,attr"  json:",omitempty"`
	Action        *Action        `xml:" action,omitempty" json:"action,omitempty"`
	AppliedToList *AppliedToList `xml:" appliedToList,omitempty" json:"appliedToList,omitempty"`
	Name          *Name          `xml:" name,omitempty" json:"name,omitempty"`
	SectionId     *SectionId     `xml:" sectionId,omitempty" json:"sectionId,omitempty"`
}

type Section struct {
	GenerationNumber string   `xml:" generationNumber,attr"  json:",omitempty"`
	Id               string   `xml:" id,attr"  json:",omitempty"`
	Name             string   `xml:" name,attr"  json:",omitempty"`
	Timestamp        string   `xml:" timestamp,attr"  json:",omitempty"`
	Rule               *Rule `xml:" rule,omitempty" json:"rule,omitempty"`
}

type SectionId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Type struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Value struct {
	Text string `xml:",chardata" json:",omitempty"`
}
