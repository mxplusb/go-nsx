package go_nsx

type ControlPlaneMode struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Description struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	VirtualWire *VirtualWire `xml:" virtualWire,omitempty" json:"virtualWire,omitempty"`
}

type TenantId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type VirtualWire struct {
	ControlPlaneMode *ControlPlaneMode `xml:" controlPlaneMode,omitempty" json:"controlPlaneMode,omitempty"`
	Description      *Description      `xml:" description,omitempty" json:"description,omitempty"`
	Name             *Name             `xml:" name,omitempty" json:"name,omitempty"`
	TenantId         *TenantId         `xml:" tenantId,omitempty" json:"tenantId,omitempty"`
}
