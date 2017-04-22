package interfaces

type AddressGroup struct {
	PrimaryAddress *PrimaryAddress `xml:" primaryAddress,omitempty" json:"primaryAddress,omitempty"`
	SubnetMask     *SubnetMask     `xml:" subnetMask,omitempty" json:"subnetMask,omitempty"`
}

type AddressGroups struct {
	AddressGroup *AddressGroup `xml:" addressGroup,omitempty" json:"addressGroup,omitempty"`
}

type ConnectedToId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type MgmtInterface struct {
	AddressGroups *AddressGroups `xml:" addressGroups,omitempty" json:"addressGroups,omitempty"`
	ConnectedToId *ConnectedToId `xml:" connectedToId,omitempty" json:"connectedToId,omitempty"`
	Mtu           *Mtu           `xml:" mtu,omitempty" json:"mtu,omitempty"`
}

type Mtu struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PrimaryAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	MgmtInterface *MgmtInterface `xml:" mgmtInterface,omitempty" json:"mgmtInterface,omitempty"`
}

type SubnetMask struct {
	Text string `xml:",chardata" json:",omitempty"`
}
