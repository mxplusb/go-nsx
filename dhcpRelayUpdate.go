package go_nsx

type Fqdn struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type GiAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type GroupingObjectId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IpAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Relay struct {
	RelayAgents *RelayAgents `xml:" relayAgents,omitempty" json:"relayAgents,omitempty"`
	RelayServer *RelayServer `xml:" relayServer,omitempty" json:"relayServer,omitempty"`
}

type RelayAgent struct {
	GiAddress *GiAddress `xml:" giAddress,omitempty" json:"giAddress,omitempty"`
	VnicIndex *VnicIndex `xml:" vnicIndex,omitempty" json:"vnicIndex,omitempty"`
}

type RelayAgents struct {
	RelayAgent *RelayAgent `xml:" relayAgent,omitempty" json:"relayAgent,omitempty"`
}

type RelayServer struct {
	Fqdn             *Fqdn             `xml:" fqdn,omitempty" json:"fqdn,omitempty"`
	GroupingObjectId *GroupingObjectId `xml:" groupingObjectId,omitempty" json:"groupingObjectId,omitempty"`
	IpAddress        *IpAddress        `xml:" ipAddress,omitempty" json:"ipAddress,omitempty"`
}

type Root struct {
	Relay *Relay `xml:" relay,omitempty" json:"relay,omitempty"`
}

type VnicIndex struct {
	Text string `xml:",chardata" json:",omitempty"`
}
