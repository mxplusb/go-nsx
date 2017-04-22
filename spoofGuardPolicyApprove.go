package go_nsx

type ApprovedBy struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ApprovedIpAddress struct {
	IpAddress *IpAddress `xml:" ipAddress,omitempty" json:"ipAddress,omitempty"`
}

type ApprovedMacAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ApprovedOn struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Id struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IpAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PublishedBy struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PublishedIpAddress struct {
	IpAddress *IpAddress `xml:" ipAddress,omitempty" json:"ipAddress,omitempty"`
}

type PublishedMacAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PublishedOn struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	SpoofguardList *SpoofguardList `xml:" spoofguardList,omitempty" json:"spoofguardList,omitempty"`
}

type Spoofguard struct {
	ApprovedBy          *ApprovedBy          `xml:" approvedBy,omitempty" json:"approvedBy,omitempty"`
	ApprovedIpAddress   *ApprovedIpAddress   `xml:" approvedIpAddress,omitempty" json:"approvedIpAddress,omitempty"`
	ApprovedMacAddress  *ApprovedMacAddress  `xml:" approvedMacAddress,omitempty" json:"approvedMacAddress,omitempty"`
	ApprovedOn          *ApprovedOn          `xml:" approvedOn,omitempty" json:"approvedOn,omitempty"`
	Id                  *Id                  `xml:" id,omitempty" json:"id,omitempty"`
	PublishedBy         *PublishedBy         `xml:" publishedBy,omitempty" json:"publishedBy,omitempty"`
	PublishedIpAddress  *PublishedIpAddress  `xml:" publishedIpAddress,omitempty" json:"publishedIpAddress,omitempty"`
	PublishedMacAddress *PublishedMacAddress `xml:" publishedMacAddress,omitempty" json:"publishedMacAddress,omitempty"`
	PublishedOn         *PublishedOn         `xml:" publishedOn,omitempty" json:"publishedOn,omitempty"`
	VnicUuid            *VnicUuid            `xml:" vnicUuid,omitempty" json:"vnicUuid,omitempty"`
}

type SpoofguardList struct {
	Spoofguard *Spoofguard `xml:" spoofguard,omitempty" json:"spoofguard,omitempty"`
}

type VnicUuid struct {
	Text string `xml:",chardata" json:",omitempty"`
}
