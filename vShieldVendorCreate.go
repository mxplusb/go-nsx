package go_nsx

type VendorInfo struct {
	Description *Description `xml:" description,omitempty" json:"description,omitempty"`
	Id          *Id          `xml:" id,omitempty" json:"id,omitempty"`
	Title       *Title       `xml:" title,omitempty" json:"title,omitempty"`
}

type Description struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Id struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	VendorInfo *VendorInfo `xml:" VendorInfo,omitempty" json:"VendorInfo,omitempty"`
}

type Title struct {
	Text string `xml:",chardata" json:",omitempty"`
}
