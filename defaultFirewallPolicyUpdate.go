package go_nsx

type Action struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type FirewallDefaultPolicy struct {
	Action         *Action         `xml:" action,omitempty" json:"action,omitempty"`
	LoggingEnabled *LoggingEnabled `xml:" loggingEnabled,omitempty" json:"loggingEnabled,omitempty"`
}

type LoggingEnabled struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	FirewallDefaultPolicy *FirewallDefaultPolicy `xml:" firewallDefaultPolicy,omitempty" json:"firewallDefaultPolicy,omitempty"`
}
