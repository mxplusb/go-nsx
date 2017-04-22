package go_nsx

type Action struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type AreaId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Authentication struct {
	Type  *Type  `xml:" type,omitempty" json:"type,omitempty"`
	Value *Value `xml:" value,omitempty" json:"value,omitempty"`
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

type Cost struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DeadInterval struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DefaultRoute struct {
	Description    *Description    `xml:" description,omitempty" json:"description,omitempty"`
	GatewayAddress *GatewayAddress `xml:" gatewayAddress,omitempty" json:"gatewayAddress,omitempty"`
	Mtu            *Mtu            `xml:" mtu,omitempty" json:"mtu,omitempty"`
	Vnic           *Vnic           `xml:" vnic,omitempty" json:"vnic,omitempty"`
}

type Description struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Direction struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Ecmp struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Enable struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Enabled struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ForwardingAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type From struct {
	Bgp       *Bgp       `xml:" bgp,omitempty" json:"bgp,omitempty"`
	Connected *Connected `xml:" connected,omitempty" json:"connected,omitempty"`
	Ospf      *Ospf      `xml:" ospf,omitempty" json:"ospf,omitempty"`
	Static    *Static    `xml:" static,omitempty" json:"static,omitempty"`
}

type GatewayAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type GracefulRestart struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type HelloInterval struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type HoldDownTimer struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IpAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IpPrefix struct {
	IpAddress *IpAddress `xml:" ipAddress,omitempty" json:"ipAddress,omitempty"`
	Name      *Name      `xml:" name,omitempty" json:"name,omitempty"`
}

type IpPrefixGe struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IpPrefixLe struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IpPrefixes struct {
	IpPrefix *IpPrefix `xml:" ipPrefix,omitempty" json:"ipPrefix,omitempty"`
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

type LogLevel struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Logging struct {
	Enable   *Enable   `xml:" enable,omitempty" json:"enable,omitempty"`
	LogLevel *LogLevel `xml:" logLevel,omitempty" json:"logLevel,omitempty"`
}

type Mtu struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Network struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type NextHop struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Ospf struct {
	Enabled           *Enabled           `xml:" enabled,omitempty" json:"enabled,omitempty"`
	ForwardingAddress *ForwardingAddress `xml:" forwardingAddress,omitempty" json:"forwardingAddress,omitempty"`
	GracefulRestart   *GracefulRestart   `xml:" gracefulRestart,omitempty" json:"gracefulRestart,omitempty"`
	OspfAreas         *OspfAreas         `xml:" ospfAreas,omitempty" json:"ospfAreas,omitempty"`
	OspfInterfaces    *OspfInterfaces    `xml:" ospfInterfaces,omitempty" json:"ospfInterfaces,omitempty"`
	ProtocolAddress   *ProtocolAddress   `xml:" protocolAddress,omitempty" json:"protocolAddress,omitempty"`
	Redistribution    *Redistribution    `xml:" redistribution,omitempty" json:"redistribution,omitempty"`
	Text              string             `xml:",chardata" json:",omitempty"`
}

type OspfArea struct {
	AreaId                *AreaId                `xml:" areaId,omitempty" json:"areaId,omitempty"`
	Authentication        *Authentication        `xml:" authentication,omitempty" json:"authentication,omitempty"`
	TranslateType7ToType5 *TranslateType7ToType5 `xml:" translateType7ToType5,omitempty" json:"translateType7ToType5,omitempty"`
	Type                  *Type                  `xml:" type,omitempty" json:"type,omitempty"`
}

type OspfAreas struct {
	OspfArea *OspfArea `xml:" ospfArea,omitempty" json:"ospfArea,omitempty"`
}

type OspfInterface struct {
	AreaId        *AreaId        `xml:" areaId,omitempty" json:"areaId,omitempty"`
	Cost          *Cost          `xml:" cost,omitempty" json:"cost,omitempty"`
	DeadInterval  *DeadInterval  `xml:" deadInterval,omitempty" json:"deadInterval,omitempty"`
	HelloInterval *HelloInterval `xml:" helloInterval,omitempty" json:"helloInterval,omitempty"`
	Priority      *Priority      `xml:" priority,omitempty" json:"priority,omitempty"`
	Vnic          *Vnic          `xml:" vnic,omitempty" json:"vnic,omitempty"`
}

type OspfInterfaces struct {
	OspfInterface *OspfInterface `xml:" ospfInterface,omitempty" json:"ospfInterface,omitempty"`
}

type Password struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PrefixName struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Priority struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ProtocolAddress struct {
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
	Routing *Routing `xml:" routing,omitempty" json:"routing,omitempty"`
}

type Route struct {
	Description *Description `xml:" description,omitempty" json:"description,omitempty"`
	Mtu         *Mtu         `xml:" mtu,omitempty" json:"mtu,omitempty"`
	Network     *Network     `xml:" network,omitempty" json:"network,omitempty"`
	NextHop     *NextHop     `xml:" nextHop,omitempty" json:"nextHop,omitempty"`
	Vnic        *Vnic        `xml:" vnic,omitempty" json:"vnic,omitempty"`
}

type RouterId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Routing struct {
	Bgp                 *Bgp                 `xml:" bgp,omitempty" json:"bgp,omitempty"`
	Ospf                *Ospf                `xml:" ospf,omitempty" json:"ospf,omitempty"`
	RoutingGlobalConfig *RoutingGlobalConfig `xml:" routingGlobalConfig,omitempty" json:"routingGlobalConfig,omitempty"`
	StaticRouting       *StaticRouting       `xml:" staticRouting,omitempty" json:"staticRouting,omitempty"`
}

type RoutingGlobalConfig struct {
	Ecmp       *Ecmp       `xml:" ecmp,omitempty" json:"ecmp,omitempty"`
	IpPrefixes *IpPrefixes `xml:" ipPrefixes,omitempty" json:"ipPrefixes,omitempty"`
	Logging    *Logging    `xml:" logging,omitempty" json:"logging,omitempty"`
	RouterId   *RouterId   `xml:" routerId,omitempty" json:"routerId,omitempty"`
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

type StaticRoutes struct {
	Route *Route `xml:" route,omitempty" json:"route,omitempty"`
}

type StaticRouting struct {
	DefaultRoute *DefaultRoute `xml:" defaultRoute,omitempty" json:"defaultRoute,omitempty"`
	StaticRoutes *StaticRoutes `xml:" staticRoutes,omitempty" json:"staticRoutes,omitempty"`
}

type TranslateType7ToType5 struct {
	Text bool `xml:",chardata" json:",omitempty"`
}

type Type struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Value struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Vnic struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Weight struct {
	Text string `xml:",chardata" json:",omitempty"`
}
