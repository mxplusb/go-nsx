package go_nsx

type AutoConfiguration struct {
	Enabled      *Enabled      `xml:" enabled,omitempty" json:"enabled,omitempty"`
	RulePriority *RulePriority `xml:" rulePriority,omitempty" json:"rulePriority,omitempty"`
}

type Enabled struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	AutoConfiguration *AutoConfiguration `xml:" autoConfiguration,omitempty" json:"autoConfiguration,omitempty"`
}

type RulePriority struct {
	Text string `xml:",chardata" json:",omitempty"`
}
