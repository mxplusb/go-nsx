package go_nsx

type Action struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Application struct {
	ApplicationId *ApplicationId `xml:" applicationId,omitempty" json:"applicationId,omitempty"`
}

type ApplicationId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Description struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Destination struct {
	Exclude          *Exclude          `xml:" exclude,omitempty" json:"exclude,omitempty"`
	GroupingObjectId *GroupingObjectId `xml:" groupingObjectId,omitempty" json:"groupingObjectId,omitempty"`
	IpAddress        *IpAddress        `xml:" ipAddress,omitempty" json:"ipAddress,omitempty"`
	VnicGroupId      *VnicGroupId      `xml:" vnicGroupId,omitempty" json:"vnicGroupId,omitempty"`
}

type Direction struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Enabled struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Exclude struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type FirewallRule struct {
	Action          *Action          `xml:" action,omitempty" json:"action,omitempty"`
	Application     *Application     `xml:" application,omitempty" json:"application,omitempty"`
	Description     *Description     `xml:" description,omitempty" json:"description,omitempty"`
	Destination     *Destination     `xml:" destination,omitempty" json:"destination,omitempty"`
	Direction       *Direction       `xml:" direction,omitempty" json:"direction,omitempty"`
	Enabled         *Enabled         `xml:" enabled,omitempty" json:"enabled,omitempty"`
	LoggingEnabled  *LoggingEnabled  `xml:" loggingEnabled,omitempty" json:"loggingEnabled,omitempty"`
	MatchTranslated *MatchTranslated `xml:" matchTranslated,omitempty" json:"matchTranslated,omitempty"`
	Name            *Name            `xml:" name,omitempty" json:"name,omitempty"`
	RuleTag         *RuleTag         `xml:" ruleTag,omitempty" json:"ruleTag,omitempty"`
	Source          *Source          `xml:" source,omitempty" json:"source,omitempty"`
}

type GroupingObjectId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IpAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type LoggingEnabled struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type MatchTranslated struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	FirewallRule *FirewallRule `xml:" firewallRule,omitempty" json:"firewallRule,omitempty"`
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
