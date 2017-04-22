package go_nsx

type Action struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Application struct {
	ApplicationId *ApplicationId `xml:" applicationId,omitempty" json:"applicationId,omitempty"`
	Service       *Service       `xml:" service,omitempty" json:"service,omitempty"`
}

type ApplicationId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DefaultPolicy struct {
	Action         *Action         `xml:" action,omitempty" json:"action,omitempty"`
	LoggingEnabled *LoggingEnabled `xml:" loggingEnabled,omitempty" json:"loggingEnabled,omitempty"`
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

type DropInvalidTraffic struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type EnableSynFloodProtection struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Enabled struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Exclude struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Firewall struct {
	DefaultPolicy *DefaultPolicy `xml:" defaultPolicy,omitempty" json:"defaultPolicy,omitempty"`
	Enabled       *Enabled       `xml:" enabled,omitempty" json:"enabled,omitempty"`
	FirewallRules *FirewallRules `xml:" firewallRules,omitempty" json:"firewallRules,omitempty"`
	GlobalConfig  *GlobalConfig  `xml:" globalConfig,omitempty" json:"globalConfig,omitempty"`
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

type FirewallRules struct {
	FirewallRule *FirewallRule `xml:" firewallRule,omitempty" json:"firewallRule,omitempty"`
}

type GlobalConfig struct {
	DropInvalidTraffic            *DropInvalidTraffic            `xml:" dropInvalidTraffic,omitempty" json:"dropInvalidTraffic,omitempty"`
	EnableSynFloodProtection      *EnableSynFloodProtection      `xml:" enableSynFloodProtection,omitempty" json:"enableSynFloodProtection,omitempty"`
	Icmp6Timeout                  *Icmp6Timeout                  `xml:" icmp6Timeout,omitempty" json:"icmp6Timeout,omitempty"`
	IcmpTimeout                   *IcmpTimeout                   `xml:" icmpTimeout,omitempty" json:"icmpTimeout,omitempty"`
	IpGenericTimeout              *IpGenericTimeout              `xml:" ipGenericTimeout,omitempty" json:"ipGenericTimeout,omitempty"`
	LogInvalidTraffic             *LogInvalidTraffic             `xml:" logInvalidTraffic,omitempty" json:"logInvalidTraffic,omitempty"`
	TcpAllowOutOfWindowPackets    *TcpAllowOutOfWindowPackets    `xml:" tcpAllowOutOfWindowPackets,omitempty" json:"tcpAllowOutOfWindowPackets,omitempty"`
	TcpPickOngoingConnections     *TcpPickOngoingConnections     `xml:" tcpPickOngoingConnections,omitempty" json:"tcpPickOngoingConnections,omitempty"`
	TcpSendResetForClosedVsePorts *TcpSendResetForClosedVsePorts `xml:" tcpSendResetForClosedVsePorts,omitempty" json:"tcpSendResetForClosedVsePorts,omitempty"`
	TcpTimeoutClose               *TcpTimeoutClose               `xml:" tcpTimeoutClose,omitempty" json:"tcpTimeoutClose,omitempty"`
	TcpTimeoutEstablished         *TcpTimeoutEstablished         `xml:" tcpTimeoutEstablished,omitempty" json:"tcpTimeoutEstablished,omitempty"`
	TcpTimeoutOpen                *TcpTimeoutOpen                `xml:" tcpTimeoutOpen,omitempty" json:"tcpTimeoutOpen,omitempty"`
	UdpTimeout                    *UdpTimeout                    `xml:" udpTimeout,omitempty" json:"udpTimeout,omitempty"`
}

type GroupingObjectId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Icmp6Timeout struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IcmpTimeout struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IpAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IpGenericTimeout struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type LogInvalidTraffic struct {
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

type Port struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Protocol struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	Firewall *Firewall `xml:" firewall,omitempty" json:"firewall,omitempty"`
}

type RuleTag struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Service struct {
	Port       *Port       `xml:" port,omitempty" json:"port,omitempty"`
	Protocol   *Protocol   `xml:" protocol,omitempty" json:"protocol,omitempty"`
	SourcePort *SourcePort `xml:" sourcePort,omitempty" json:"sourcePort,omitempty"`
}

type Source struct {
	Exclude          *Exclude          `xml:" exclude,omitempty" json:"exclude,omitempty"`
	GroupingObjectId *GroupingObjectId `xml:" groupingObjectId,omitempty" json:"groupingObjectId,omitempty"`
	IpAddress        *IpAddress        `xml:" ipAddress,omitempty" json:"ipAddress,omitempty"`
	VnicGroupId      *VnicGroupId      `xml:" vnicGroupId,omitempty" json:"vnicGroupId,omitempty"`
}

type SourcePort struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type TcpAllowOutOfWindowPackets struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type TcpPickOngoingConnections struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type TcpSendResetForClosedVsePorts struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type TcpTimeoutClose struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type TcpTimeoutEstablished struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type TcpTimeoutOpen struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type UdpTimeout struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type VnicGroupId struct {
	Text string `xml:",chardata" json:",omitempty"`
}
