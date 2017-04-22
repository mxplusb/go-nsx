package go_nsx

type Action struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Config struct {
	Timestamp      string               `xml:" timestamp,attr"  json:",omitempty"`
	ContextId        *ContextId        `xml:" contextId,omitempty" json:"contextId,omitempty"`
	GenerationNumber *GenerationNumber `xml:" generationNumber,omitempty" json:"generationNumber,omitempty"`
	Layer2Sections   *Layer2Sections   `xml:" layer2Sections,omitempty" json:"layer2Sections,omitempty"`
	Layer3Sections   *Layer3Sections   `xml:" layer3Sections,omitempty" json:"layer3Sections,omitempty"`
}

type ContextId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Description struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type FirewallDraft struct {
	Id        string          `xml:" id,attr"  json:",omitempty"`
	Name      string          `xml:" name,attr"  json:",omitempty"`
	Timestamp string          `xml:" timestamp,attr"  json:",omitempty"`
	Config      *Config      `xml:" config,omitempty" json:"config,omitempty"`
	Description *Description `xml:" description,omitempty" json:"description,omitempty"`
	Mode        *Mode        `xml:" mode,omitempty" json:"mode,omitempty"`
	Preserve    *Preserve    `xml:" preserve,omitempty" json:"preserve,omitempty"`
	User        *User        `xml:" user,omitempty" json:"user,omitempty"`
}

type GenerationNumber struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Layer2Sections struct {
	Section *Section `xml:" section,omitempty" json:"section,omitempty"`
}

type Layer3Sections struct {
	Section *Section `xml:" section,omitempty" json:"section,omitempty"`
}

type Mode struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Precedence struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Preserve struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	FirewallDraft *FirewallDraft `xml:" firewallDraft,omitempty" json:"firewallDraft,omitempty"`
}

type Rule struct {
	Disabled string         `xml:" disabled,attr"  json:",omitempty"`
	Id       string         `xml:" id,attr"  json:",omitempty"`
	Logged   string         `xml:" logged,attr"  json:",omitempty"`
	Action     *Action     `xml:" action,omitempty" json:"action,omitempty"`
	Name       *Name       `xml:" name,omitempty" json:"name,omitempty"`
	Precedence *Precedence `xml:" precedence,omitempty" json:"precedence,omitempty"`
}

type Section struct {
	Name      string   `xml:" name,attr"  json:",omitempty"`
	Timestamp string   `xml:" timestamp,attr"  json:",omitempty"`
	Rule        *Rule `xml:" rule,omitempty" json:"rule,omitempty"`
}

type User struct {
	Text string `xml:",chardata" json:",omitempty"`
}
