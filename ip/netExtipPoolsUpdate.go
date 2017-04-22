package ip

type Description struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DnsSuffix struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Enabled struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Gateway struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IpAddressPool struct {
	Description  *Description  `xml:" description,omitempty" json:"description,omitempty"`
	DnsSuffix    *DnsSuffix    `xml:" dnsSuffix,omitempty" json:"dnsSuffix,omitempty"`
	Enabled      *Enabled      `xml:" enabled,omitempty" json:"enabled,omitempty"`
	Gateway      *Gateway      `xml:" gateway,omitempty" json:"gateway,omitempty"`
	IpRange      *IpRange      `xml:" ipRange,omitempty" json:"ipRange,omitempty"`
	Netmask      *Netmask      `xml:" netmask,omitempty" json:"netmask,omitempty"`
	PrimaryDns   *PrimaryDns   `xml:" primaryDns,omitempty" json:"primaryDns,omitempty"`
	SecondaryDns *SecondaryDns `xml:" secondaryDns,omitempty" json:"secondaryDns,omitempty"`
	WinsServer   *WinsServer   `xml:" winsServer,omitempty" json:"winsServer,omitempty"`
}

type IpRange struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Netmask struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PrimaryDns struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	IpAddressPool *IpAddressPool `xml:" ipAddressPool,omitempty" json:"ipAddressPool,omitempty"`
}

type SecondaryDns struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type WinsServer struct {
	Text string `xml:",chardata" json:",omitempty"`
}
