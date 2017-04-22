package go_nsx

type Action struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Bgp struct {
	BgpNeighbours  *BgpNeighbours  `xml:" bgpNeighbours,omitempty" json:"bgpNeighbours,omitempty"`
	Enabled        *Enabled        `xml:" enabled,omitempty" json:"enabled,omitempty"`
	LocalAS        *LocalAS        `xml:" localAS,omitempty" json:"localAS,omitempty"`
	LocalASNumber  *LocalASNumber  `xml:" localASNumber,omitempty" json:"localASNumber,omitempty"`
	Redistribution *Redistribution `xml:" redistribution,omitempty" json:"redistribution,omitempty"`
	Text           string          `xml:",chardata" json:",omitempty"`
}

type BgpFilter struct {
	Action     *Action     `xml:" action,omitempty" json:"action,omitempty"`
	Direction  *Direction  `xml:" direction,omitempty" json:"direction,omitempty"`
	IpPrefixGe *IpPrefixGe `xml:" ipPrefixGe,omitempty" json:"ipPrefixGe,omitempty"`
	IpPrefixLe *IpPrefixLe `xml:" ipPrefixLe,omitempty" json:"ipPrefixLe,omitempty"`
	Network    *Network    `xml:" network,omitempty" json:"network,omitempty"`
}

type BgpFilters struct {
	BgpFilter *BgpFilter `xml:" bgpFilter,omitempty" json:"bgpFilter,omitempty"`
}

type BgpNeighbour struct {
	BgpFilters     *BgpFilters     `xml:" bgpFilters,omitempty" json:"bgpFilters,omitempty"`
	HoldDownTimer  *HoldDownTimer  `xml:" holdDownTimer,omitempty" json:"holdDownTimer,omitempty"`
	IpAddress      *IpAddress      `xml:" ipAddress,omitempty" json:"ipAddress,omitempty"`
	KeepAliveTimer *KeepAliveTimer `xml:" keepAliveTimer,omitempty" json:"keepAliveTimer,omitempty"`
	Password       *Password       `xml:" password,omitempty" json:"password,omitempty"`
	RemoteAS       *RemoteAS       `xml:" remoteAS,omitempty" json:"remoteAS,omitempty"`
	RemoteASNumber *RemoteASNumber `xml:" remoteASNumber,omitempty" json:"remoteASNumber,omitempty"`
	Weight         *Weight         `xml:" weight,omitempty" json:"weight,omitempty"`
}

type BgpNeighbours struct {
	BgpNeighbour *BgpNeighbour `xml:" bgpNeighbour,omitempty" json:"bgpNeighbour,omitempty"`
}

type Connected struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Direction struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Enabled struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type From struct {
	Bgp       *Bgp       `xml:" bgp,omitempty" json:"bgp,omitempty"`
	Connected *Connected `xml:" connected,omitempty" json:"connected,omitempty"`
	Ospf      *Ospf      `xml:" ospf,omitempty" json:"ospf,omitempty"`
	Static    *Static    `xml:" static,omitempty" json:"static,omitempty"`
}

type HoldDownTimer struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IpAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IpPrefixGe struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IpPrefixLe struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type KeepAliveTimer struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type LocalAS struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type LocalASNumber struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Network struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Ospf struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Password struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PrefixName struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Redistribution struct {
	Enabled *Enabled `xml:" enabled,omitempty" json:"enabled,omitempty"`
	Rules   *Rules   `xml:" rules,omitempty" json:"rules,omitempty"`
}

type RemoteAS struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type RemoteASNumber struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	Bgp *Bgp `xml:" bgp,omitempty" json:"bgp,omitempty"`
}

type Rule struct {
	Action     *Action     `xml:" action,omitempty" json:"action,omitempty"`
	From       *From       `xml:" from,omitempty" json:"from,omitempty"`
	PrefixName *PrefixName `xml:" prefixName,omitempty" json:"prefixName,omitempty"`
}

type Rules struct {
	Rule *Rule `xml:" rule,omitempty" json:"rule,omitempty"`
}

type Static struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Weight struct {
	Text string `xml:",chardata" json:",omitempty"`
}
