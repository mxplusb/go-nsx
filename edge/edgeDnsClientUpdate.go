package edge

type DnsClient struct {
	DomainName   *DomainName   `xml:" domainName,omitempty" json:"domainName,omitempty"`
	PrimaryDns   *PrimaryDns   `xml:" primaryDns,omitempty" json:"primaryDns,omitempty"`
	SecondaryDns *SecondaryDns `xml:" secondaryDns,omitempty" json:"secondaryDns,omitempty"`
}

type DomainName struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PrimaryDns struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	DnsClient *DnsClient `xml:" dnsClient,omitempty" json:"dnsClient,omitempty"`
}

type SecondaryDns struct {
	Text string `xml:",chardata" json:",omitempty"`
}
