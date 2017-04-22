package ip

type AllocationMode struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IpAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IpAddressRequest struct {
	AllocationMode *AllocationMode `xml:" allocationMode,omitempty" json:"allocationMode,omitempty"`
	IpAddress      *IpAddress      `xml:" ipAddress,omitempty" json:"ipAddress,omitempty"`
}

type Root struct {
	IpAddressRequest *IpAddressRequest `xml:" ipAddressRequest,omitempty" json:"ipAddressRequest,omitempty"`
}
