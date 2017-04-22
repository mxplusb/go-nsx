package go_nsx

type AutoDraftDisabled struct {
	Text bool `xml:",chardata" json:",omitempty"`
}

type GlobalConfiguration struct {
	AutoDraftDisabled  *AutoDraftDisabled  `xml:" autoDraftDisabled,omitempty" json:"autoDraftDisabled,omitempty"`
	Layer2RuleOptimize *Layer2RuleOptimize `xml:" layer2RuleOptimize,omitempty" json:"layer2RuleOptimize,omitempty"`
	Layer3RuleOptimize *Layer3RuleOptimize `xml:" layer3RuleOptimize,omitempty" json:"layer3RuleOptimize,omitempty"`
	TcpStrictOption    *TcpStrictOption    `xml:" tcpStrictOption,omitempty" json:"tcpStrictOption,omitempty"`
}

type Layer2RuleOptimize struct {
	Text bool `xml:",chardata" json:",omitempty"`
}

type Layer3RuleOptimize struct {
	Text bool `xml:",chardata" json:",omitempty"`
}

type Root struct {
	GlobalConfiguration *GlobalConfiguration `xml:" globalConfiguration,omitempty" json:"globalConfiguration,omitempty"`
}

type TcpStrictOption struct {
	Text bool `xml:",chardata" json:",omitempty"`
}
