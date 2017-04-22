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

type Interface struct {
	AddressGroups *AddressGroups `xml:" addressGroups,omitempty" json:"addressGroups,omitempty"`
	ConnectedToId *ConnectedToId `xml:" connectedToId,omitempty" json:"connectedToId,omitempty"`
	IsConnected   *IsConnected   `xml:" isConnected,omitempty" json:"isConnected,omitempty"`
	Mtu           *Mtu           `xml:" mtu,omitempty" json:"mtu,omitempty"`
	Name          *Name          `xml:" name,omitempty" json:"name,omitempty"`
	Type          *Type          `xml:" type,omitempty" json:"type,omitempty"`
}

type Interfaces struct {
	Interface *Interface `xml:" interface,omitempty" json:"interface,omitempty"`
}

type IsConnected struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Mtu struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PrimaryAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	Interfaces *Interfaces `xml:" interfaces,omitempty" json:"interfaces,omitempty"`
}

type SubnetMask struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Type struct {
	Text string `xml:",chardata" json:",omitempty"`
}
