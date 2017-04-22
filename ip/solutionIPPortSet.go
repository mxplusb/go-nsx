package go_nsx

type LocationInfo struct {
	Ip   *Ip   `xml:" ip,omitempty" json:"ip,omitempty"`
	Port *Port `xml:" port,omitempty" json:"port,omitempty"`
}

type Ip struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Port struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	LocationInfo *LocationInfo `xml:" LocationInfo,omitempty" json:"LocationInfo,omitempty"`
}
