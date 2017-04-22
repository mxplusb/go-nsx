package ip

// this is for singletons and things I don't really care about.

type IPAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DnsServer1 struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DnsServer2 struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DNSSuffix struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Gateway struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Description struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Port struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Id struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ObjectId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ObjectTypeName struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PrefixLength struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Revision struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Type struct {
	TypeName *TypeName `xml:" typeName,omitempty" json:"typeName,omitempty"`
}

type Enabled struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PrimaryDns struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Netmask struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type SecondaryDns struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type WinsServer struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DnsSuffix struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IpRange struct {
	Text string `xml:",chardata" json:",omitempty"`
}