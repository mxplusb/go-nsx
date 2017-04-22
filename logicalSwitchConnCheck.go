package go_nsx

type DestinationHost struct {
	HostId *HostId `xml:" hostId,omitempty" json:"hostId,omitempty"`
	Switd  *Switd  `xml:" switd,omitempty" json:"switd,omitempty"`
	VlanId *VlanId `xml:" vlanId,omitempty" json:"vlanId,omitempty"`
}

type Gateway struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type HostId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PacketSize struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PacketSizeMode struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	TestParameters *TestParameters `xml:" testParameters,omitempty" json:"testParameters,omitempty"`
}

type SourceHost struct {
	HostId *HostId `xml:" hostId,omitempty" json:"hostId,omitempty"`
	Switd  *Switd  `xml:" switd,omitempty" json:"switd,omitempty"`
	VlanId *VlanId `xml:" vlanId,omitempty" json:"vlanId,omitempty"`
}

type Switd struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type TestParameters struct {
	DestinationHost *DestinationHost `xml:" destinationHost,omitempty" json:"destinationHost,omitempty"`
	Gateway         *Gateway         `xml:" gateway,omitempty" json:"gateway,omitempty"`
	PacketSize      *PacketSize      `xml:" packetSize,omitempty" json:"packetSize,omitempty"`
	PacketSizeMode  *PacketSizeMode  `xml:" packetSizeMode,omitempty" json:"packetSizeMode,omitempty"`
	SourceHost      *SourceHost      `xml:" sourceHost,omitempty" json:"sourceHost,omitempty"`
}

type VlanId struct {
	Text string `xml:",chardata" json:",omitempty"`
}
