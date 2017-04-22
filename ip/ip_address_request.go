package ip

type AllocationMode struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type AddressRequest struct {
	AllocationMode *AllocationMode `xml:" allocationMode,omitempty" json:"allocationMode,omitempty"`
	IpAddress      *IPAddress      `xml:" ipAddress,omitempty" json:"ipAddress,omitempty"`
}

type AddressRequestRoot struct {
	AddressRequest *AddressRequest `xml:" ipAddressRequest,omitempty" json:"ipAddressRequest,omitempty"`
}
