package go_nsx

type DefaultRoute struct {
	Description    *Description    `xml:" description,omitempty" json:"description,omitempty"`
	GatewayAddress *GatewayAddress `xml:" gatewayAddress,omitempty" json:"gatewayAddress,omitempty"`
	Mtu            *Mtu            `xml:" mtu,omitempty" json:"mtu,omitempty"`
	Vnic           *Vnic           `xml:" vnic,omitempty" json:"vnic,omitempty"`
}

type Description struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type GatewayAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Mtu struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Network struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type NextHop struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	StaticRouting *StaticRouting `xml:" staticRouting,omitempty" json:"staticRouting,omitempty"`
}

type Route struct {
	Description *Description `xml:" description,omitempty" json:"description,omitempty"`
	Mtu         *Mtu         `xml:" mtu,omitempty" json:"mtu,omitempty"`
	Network     *Network     `xml:" network,omitempty" json:"network,omitempty"`
	NextHop     *NextHop     `xml:" nextHop,omitempty" json:"nextHop,omitempty"`
	Vnic        *Vnic        `xml:" vnic,omitempty" json:"vnic,omitempty"`
}

type StaticRoutes struct {
	Route *Route `xml:" route,omitempty" json:"route,omitempty"`
}

type StaticRouting struct {
	DefaultRoute *DefaultRoute `xml:" defaultRoute,omitempty" json:"defaultRoute,omitempty"`
	StaticRoutes *StaticRoutes `xml:" staticRoutes,omitempty" json:"staticRoutes,omitempty"`
}

type Vnic struct {
	Text string `xml:",chardata" json:",omitempty"`
}
