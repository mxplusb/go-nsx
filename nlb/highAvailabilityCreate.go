package go_nsx

type DeclareDeadTime struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Enabled struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type HighAvailability struct {
	DeclareDeadTime *DeclareDeadTime `xml:" declareDeadTime,omitempty" json:"declareDeadTime,omitempty"`
	Enabled         *Enabled         `xml:" enabled,omitempty" json:"enabled,omitempty"`
	IpAddresses     *IpAddresses     `xml:" ipAddresses,omitempty" json:"ipAddresses,omitempty"`
	Vnic            *Vnic            `xml:" vnic,omitempty" json:"vnic,omitempty"`
}

type IpAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IpAddresses struct {
	IpAddress *IpAddress `xml:" ipAddress,omitempty" json:"ipAddress,omitempty"`
}

type Root struct {
	HighAvailability *HighAvailability `xml:" highAvailability,omitempty" json:"highAvailability,omitempty"`
}

type Vnic struct {
	Text string `xml:",chardata" json:",omitempty"`
}
