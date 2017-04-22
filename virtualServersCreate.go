package go_nsx

type AccelerationEnabled struct {
	Text bool `xml:",chardata" json:",omitempty"`
}

type ApplicationProfileId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ApplicationRuleId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ConnectionLimit struct {
	Text int8 `xml:",chardata" json:",omitempty"`
}

type ConnectionRateLimit struct {
	Text int8 `xml:",chardata" json:",omitempty"`
}

type DefaultPoolId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Description struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type EnableServiceInsertion struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Enabled struct {
	Text bool `xml:",chardata" json:",omitempty"`
}

type IpAddress struct {
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
	VirtualServer *VirtualServer `xml:" virtualServer,omitempty" json:"virtualServer,omitempty"`
}

type VirtualServer struct {
	AccelerationEnabled    *AccelerationEnabled    `xml:" accelerationEnabled,omitempty" json:"accelerationEnabled,omitempty"`
	ApplicationProfileId   *ApplicationProfileId   `xml:" applicationProfileId,omitempty" json:"applicationProfileId,omitempty"`
	ApplicationRuleId      *ApplicationRuleId      `xml:" applicationRuleId,omitempty" json:"applicationRuleId,omitempty"`
	ConnectionLimit        *ConnectionLimit        `xml:" connectionLimit,omitempty" json:"connectionLimit,omitempty"`
	ConnectionRateLimit    *ConnectionRateLimit    `xml:" connectionRateLimit,omitempty" json:"connectionRateLimit,omitempty"`
	DefaultPoolId          *DefaultPoolId          `xml:" defaultPoolId,omitempty" json:"defaultPoolId,omitempty"`
	Description            *Description            `xml:" description,omitempty" json:"description,omitempty"`
	EnableServiceInsertion *EnableServiceInsertion `xml:" enableServiceInsertion,omitempty" json:"enableServiceInsertion,omitempty"`
	Enabled                *Enabled                `xml:" enabled,omitempty" json:"enabled,omitempty"`
	IpAddress              *IpAddress              `xml:" ipAddress,omitempty" json:"ipAddress,omitempty"`
	Name                   *Name                   `xml:" name,omitempty" json:"name,omitempty"`
	Port                   *Port                   `xml:" port,omitempty" json:"port,omitempty"`
	Protocol               *Protocol               `xml:" protocol,omitempty" json:"protocol,omitempty"`
}
