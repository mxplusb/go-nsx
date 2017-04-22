package go_nsx

type CacheSize struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Dns struct {
	CacheSize *CacheSize `xml:" cacheSize,omitempty" json:"cacheSize,omitempty"`
	DnsViews  *DnsViews  `xml:" dnsViews,omitempty" json:"dnsViews,omitempty"`
	Enabled   *Enabled   `xml:" enabled,omitempty" json:"enabled,omitempty"`
	Listeners *Listeners `xml:" listeners,omitempty" json:"listeners,omitempty"`
	Version   *Version   `xml:" version,omitempty" json:"version,omitempty"`
}

type DnsView struct {
	ViewId *ViewId `xml:" viewId,omitempty" json:"viewId,omitempty"`
}

type DnsViews struct {
	DnsView *DnsView `xml:" dnsView,omitempty" json:"dnsView,omitempty"`
}

type Enabled struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IpAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Listeners struct {
	IpAddress *IpAddress `xml:" ipAddress,omitempty" json:"ipAddress,omitempty"`
}

type Root struct {
	Dns *Dns `xml:" dns,omitempty" json:"dns,omitempty"`
}

type Version struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ViewId struct {
	Text string `xml:",chardata" json:",omitempty"`
}
