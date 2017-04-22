package go_nsx

type AutoConfigureDNS struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Code struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DefaultGateway struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DestinationSubnet struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DhcpOptions struct {
	Option121 *Option121 `xml:" option121,omitempty" json:"option121,omitempty"`
	Option150 *Option150 `xml:" option150,omitempty" json:"option150,omitempty"`
	Option26  *Option26  `xml:" option26,omitempty" json:"option26,omitempty"`
	Option66  *Option66  `xml:" option66,omitempty" json:"option66,omitempty"`
	Option67  *Option67  `xml:" option67,omitempty" json:"option67,omitempty"`
	Other     *Other     `xml:" other,omitempty" json:"other,omitempty"`
}

type DomainName struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Hostname struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IpAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type LeaseTime struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type NextServer struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Option121 struct {
	StaticRoute *StaticRoute `xml:" staticRoute,omitempty" json:"staticRoute,omitempty"`
}

type Option150 struct {
	Server *Server `xml:" server,omitempty" json:"server,omitempty"`
}

type Option26 struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Option66 struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Option67 struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Other struct {
	Code  *Code  `xml:" code,omitempty" json:"code,omitempty"`
	Value *Value `xml:" value,omitempty" json:"value,omitempty"`
}

type PrimaryNameServer struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	StaticBinding *StaticBinding `xml:" staticBinding,omitempty" json:"staticBinding,omitempty"`
}

type Router struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type SecondaryNameServer struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Server struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type StaticBinding struct {
	AutoConfigureDNS    *AutoConfigureDNS    `xml:" autoConfigureDNS,omitempty" json:"autoConfigureDNS,omitempty"`
	DefaultGateway      *DefaultGateway      `xml:" defaultGateway,omitempty" json:"defaultGateway,omitempty"`
	DhcpOptions         *DhcpOptions         `xml:" dhcpOptions,omitempty" json:"dhcpOptions,omitempty"`
	DomainName          *DomainName          `xml:" domainName,omitempty" json:"domainName,omitempty"`
	Hostname            *Hostname            `xml:" hostname,omitempty" json:"hostname,omitempty"`
	IpAddress           *IpAddress           `xml:" ipAddress,omitempty" json:"ipAddress,omitempty"`
	LeaseTime           *LeaseTime           `xml:" leaseTime,omitempty" json:"leaseTime,omitempty"`
	NextServer          *NextServer          `xml:" nextServer,omitempty" json:"nextServer,omitempty"`
	PrimaryNameServer   *PrimaryNameServer   `xml:" primaryNameServer,omitempty" json:"primaryNameServer,omitempty"`
	SecondaryNameServer *SecondaryNameServer `xml:" secondaryNameServer,omitempty" json:"secondaryNameServer,omitempty"`
	VmId                *VmId                `xml:" vmId,omitempty" json:"vmId,omitempty"`
	VnicId              *VnicId              `xml:" vnicId,omitempty" json:"vnicId,omitempty"`
}

type StaticRoute struct {
	DestinationSubnet *DestinationSubnet `xml:" destinationSubnet,omitempty" json:"destinationSubnet,omitempty"`
	Router            *Router            `xml:" router,omitempty" json:"router,omitempty"`
}

type Value struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type VmId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type VnicId struct {
	Text string `xml:",chardata" json:",omitempty"`
}
