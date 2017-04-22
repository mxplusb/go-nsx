package go_nsx

type Destination struct {
	Exclude          *Exclude          `xml:" exclude,omitempty" json:"exclude,omitempty"`
	GroupingObjectId *GroupingObjectId `xml:" groupingObjectId,omitempty" json:"groupingObjectId,omitempty"`
	IpAddress        *IpAddress        `xml:" ipAddress,omitempty" json:"ipAddress,omitempty"`
	VnicGroupId      *VnicGroupId      `xml:" vnicGroupId,omitempty" json:"vnicGroupId,omitempty"`
}

type Exclude struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type FirewallRule struct {
	Destination *Destination `xml:" destination,omitempty" json:"destination,omitempty"`
	Name        *Name        `xml:" name,omitempty" json:"name,omitempty"`
	RuleTag     *RuleTag     `xml:" ruleTag,omitempty" json:"ruleTag,omitempty"`
	Source      *Source      `xml:" source,omitempty" json:"source,omitempty"`
}

type FirewallRules struct {
	FirewallRule *FirewallRule `xml:" firewallRule,omitempty" json:"firewallRule,omitempty"`
}

type GroupingObjectId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IpAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	FirewallRules *FirewallRules `xml:" firewallRules,omitempty" json:"firewallRules,omitempty"`
}

type RuleTag struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Source struct {
	Exclude          *Exclude          `xml:" exclude,omitempty" json:"exclude,omitempty"`
	GroupingObjectId *GroupingObjectId `xml:" groupingObjectId,omitempty" json:"groupingObjectId,omitempty"`
	IpAddress        *IpAddress        `xml:" ipAddress,omitempty" json:"ipAddress,omitempty"`
	VnicGroupId      *VnicGroupId      `xml:" vnicGroupId,omitempty" json:"vnicGroupId,omitempty"`
}

type VnicGroupId struct {
	Text string `xml:",chardata" json:",omitempty"`
}
