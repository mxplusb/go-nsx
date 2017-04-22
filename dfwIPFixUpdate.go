package go_nsx

type Collector struct {
	Ip   *Ip   `xml:" ip,omitempty" json:"ip,omitempty"`
	Port *Port `xml:" port,omitempty" json:"port,omitempty"`
}

type ContextId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type FlowTimeout struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Ip struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IpfixConfiguration struct {
	Collector           *Collector           `xml:" collector,omitempty" json:"collector,omitempty"`
	ContextId           *ContextId           `xml:" contextId,omitempty" json:"contextId,omitempty"`
	FlowTimeout         *FlowTimeout         `xml:" flowTimeout,omitempty" json:"flowTimeout,omitempty"`
	IpfixEnabled        *IpfixEnabled        `xml:" ipfixEnabled,omitempty" json:"ipfixEnabled,omitempty"`
	ObservationDomainId *ObservationDomainId `xml:" observationDomainId,omitempty" json:"observationDomainId,omitempty"`
}

type IpfixEnabled struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ObservationDomainId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Port struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	IpfixConfiguration *IpfixConfiguration `xml:" ipfixConfiguration,omitempty" json:"ipfixConfiguration,omitempty"`
}
