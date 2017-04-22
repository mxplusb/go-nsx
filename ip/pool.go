package ip

type Algorithm struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type MaxConn struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Member struct {
	IpAddress   *IPAddress   `xml:" ipAddress,omitempty" json:"ipAddress,omitempty"`
	MaxConn     *MaxConn     `xml:" maxConn,omitempty" json:"maxConn,omitempty"`
	MinConn     *MinConn     `xml:" minConn,omitempty" json:"minConn,omitempty"`
	MonitorPort *MonitorPort `xml:" monitorPort,omitempty" json:"monitorPort,omitempty"`
	Name        *Name        `xml:" name,omitempty" json:"name,omitempty"`
	Port        *Port        `xml:" port,omitempty" json:"port,omitempty"`
	Weight      *Weight      `xml:" weight,omitempty" json:"weight,omitempty"`
}

type MinConn struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type MonitorId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type MonitorPort struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Pool struct {
	Algorithm   *Algorithm   `xml:" algorithm,omitempty" json:"algorithm,omitempty"`
	Description *Description `xml:" description,omitempty" json:"description,omitempty"`
	Member      *Member      `xml:" member,omitempty" json:"member,omitempty"`
	MonitorId   *MonitorId   `xml:" monitorId,omitempty" json:"monitorId,omitempty"`
	Name        *Name        `xml:" name,omitempty" json:"name,omitempty"`
	Transparent *Transparent `xml:" transparent,omitempty" json:"transparent,omitempty"`
}

type PoolRoot struct {
	Pool *Pool `xml:" pool,omitempty" json:"pool,omitempty"`
}

type Transparent struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Weight struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ExtendedAttributes struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IPRanges struct {
	IpRangeDto *IPRange `xml:" ipRangeDto,omitempty" json:"ipRangeDto,omitempty"`
}

type IpamAddressPool struct {
	DnsServer1         *DnsServer1         `xml:" dnsServer1,omitempty" json:"dnsServer1,omitempty"`
	DnsServer2         *DnsServer2         `xml:" dnsServer2,omitempty" json:"dnsServer2,omitempty"`
	DnsSuffix          *DNSSuffix          `xml:" dnsSuffix,omitempty" json:"dnsSuffix,omitempty"`
	ExtendedAttributes *ExtendedAttributes `xml:" extendedAttributes,omitempty" json:"extendedAttributes,omitempty"`
	Gateway            *Gateway            `xml:" gateway,omitempty" json:"gateway,omitempty"`
	IpRanges           *IPRanges           `xml:" ipRanges,omitempty" json:"ipRanges,omitempty"`
	Name               *Name               `xml:" name,omitempty" json:"name,omitempty"`
	ObjectId           *ObjectId           `xml:" objectId,omitempty" json:"objectId,omitempty"`
	ObjectTypeName     *ObjectTypeName     `xml:" objectTypeName,omitempty" json:"objectTypeName,omitempty"`
	PrefixLength       *PrefixLength       `xml:" prefixLength,omitempty" json:"prefixLength,omitempty"`
	Revision           *Revision           `xml:" revision,omitempty" json:"revision,omitempty"`
	Type               *Type               `xml:" type,omitempty" json:"type,omitempty"`
	VsmUuid            *VsmUuid            `xml:" vsmUuid,omitempty" json:"vsmUuid,omitempty"`
}

type IPRange struct {
	EndAddress   *EndAddress   `xml:" endAddress,omitempty" json:"endAddress,omitempty"`
	StartAddress *StartAddress `xml:" startAddress,omitempty" json:"startAddress,omitempty"`
}

type IPAMAddressPoolRoot struct {
	IpamAddressPool *IpamAddressPool `xml:" ipamAddressPool,omitempty" json:"ipamAddressPool,omitempty"`
}

type VsmUuid struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type StartAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type EndAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IPAddressPoolRoot struct {
	IpAddressPool *IPAddressPool `xml:" ipAddressPool,omitempty" json:"ipAddressPool,omitempty"`
}

type IPAddressPool struct {
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
