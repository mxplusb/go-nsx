package go_nsx

type Ecmp struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Enable struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IpAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IpPrefix struct {
	IpAddress *IpAddress `xml:" ipAddress,omitempty" json:"ipAddress,omitempty"`
	Name      *Name      `xml:" name,omitempty" json:"name,omitempty"`
}

type IpPrefixes struct {
	IpPrefix *IpPrefix `xml:" ipPrefix,omitempty" json:"ipPrefix,omitempty"`
}

type LogLevel struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Logging struct {
	Enable   *Enable   `xml:" enable,omitempty" json:"enable,omitempty"`
	LogLevel *LogLevel `xml:" logLevel,omitempty" json:"logLevel,omitempty"`
}

type Name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	RoutingGlobalConfig *RoutingGlobalConfig `xml:" routingGlobalConfig,omitempty" json:"routingGlobalConfig,omitempty"`
}

type RouterId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type RoutingGlobalConfig struct {
	Ecmp       *Ecmp       `xml:" ecmp,omitempty" json:"ecmp,omitempty"`
	IpPrefixes *IpPrefixes `xml:" ipPrefixes,omitempty" json:"ipPrefixes,omitempty"`
	Logging    *Logging    `xml:" logging,omitempty" json:"logging,omitempty"`
	RouterId   *RouterId   `xml:" routerId,omitempty" json:"routerId,omitempty"`
}
