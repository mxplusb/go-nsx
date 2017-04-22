package go_nsx

type DnsServer1 struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DnsServer2 struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DnsSuffix struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type EndAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Gateway struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IpRangeDto struct {
	EndAddress   *EndAddress   `xml:" endAddress,omitempty" json:"endAddress,omitempty"`
	StartAddress *StartAddress `xml:" startAddress,omitempty" json:"startAddress,omitempty"`
}

type IpRanges struct {
	IpRangeDto *IpRangeDto `xml:" ipRangeDto,omitempty" json:"ipRangeDto,omitempty"`
}

type IpamAddressPool struct {
	DnsServer1   *DnsServer1   `xml:" dnsServer1,omitempty" json:"dnsServer1,omitempty"`
	DnsServer2   *DnsServer2   `xml:" dnsServer2,omitempty" json:"dnsServer2,omitempty"`
	DnsSuffix    *DnsSuffix    `xml:" dnsSuffix,omitempty" json:"dnsSuffix,omitempty"`
	Gateway      *Gateway      `xml:" gateway,omitempty" json:"gateway,omitempty"`
	IpRanges     *IpRanges     `xml:" ipRanges,omitempty" json:"ipRanges,omitempty"`
	Name         *Name         `xml:" name,omitempty" json:"name,omitempty"`
	PrefixLength *PrefixLength `xml:" prefixLength,omitempty" json:"prefixLength,omitempty"`
}

type Name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PrefixLength struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	IpamAddressPool *IpamAddressPool `xml:" ipamAddressPool,omitempty" json:"ipamAddressPool,omitempty"`
}

type StartAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}
