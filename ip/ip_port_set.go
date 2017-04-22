package ip

type LocationInfo struct {
	Ip   *IPAddr   `xml:" ip,omitempty" json:"ip,omitempty"`
	Port *Port `xml:" port,omitempty" json:"port,omitempty"`
}

type IPAddr struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type LocationInfoRoot struct {
	LocationInfo *LocationInfo `xml:" LocationInfo,omitempty" json:"LocationInfo,omitempty"`
}
