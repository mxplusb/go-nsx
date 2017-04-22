package ip

type Description struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Enabled struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Network struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Optimize struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Ports struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PrivateNetwork struct {
	Description    *Description    `xml:" description,omitempty" json:"description,omitempty"`
	Enabled        *Enabled        `xml:" enabled,omitempty" json:"enabled,omitempty"`
	Network        *Network        `xml:" network,omitempty" json:"network,omitempty"`
	SendOverTunnel *SendOverTunnel `xml:" sendOverTunnel,omitempty" json:"sendOverTunnel,omitempty"`
}

type Root struct {
	PrivateNetwork *PrivateNetwork `xml:" privateNetwork,omitempty" json:"privateNetwork,omitempty"`
}

type SendOverTunnel struct {
	Optimize *Optimize `xml:" optimize,omitempty" json:"optimize,omitempty"`
	Ports    *Ports    `xml:" ports,omitempty" json:"ports,omitempty"`
}
