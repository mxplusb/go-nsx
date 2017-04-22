package utilities

type Dns struct {
	DomainList  *DomainList  `xml:" domainList,omitempty" json:"domainList,omitempty"`
	Ipv4Address *Ipv4Address `xml:" ipv4Address,omitempty" json:"ipv4Address,omitempty"`
	Ipv6Address *Ipv6Address `xml:" ipv6Address,omitempty" json:"ipv6Address,omitempty"`
}

type DomainList struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Ipv4Address struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Ipv6Address struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	Dns *Dns `xml:" dns,omitempty" json:"dns,omitempty"`
}
