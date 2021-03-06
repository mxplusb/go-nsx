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
	VirtualWireCreateSpec *VirtualWireCreateSpec `xml:" virtualWireCreateSpec,omitempty" json:"virtualWireCreateSpec,omitempty"`
}

type TenantId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type VirtualWireCreateSpec struct {
	ControlPlaneMode *ControlPlaneMode `xml:" controlPlaneMode,omitempty" json:"controlPlaneMode,omitempty"`
	Description      *Description      `xml:" description,omitempty" json:"description,omitempty"`
	Name             *Name             `xml:" name,omitempty" json:"name,omitempty"`
	TenantId         *TenantId         `xml:" tenantId,omitempty" json:"tenantId,omitempty"`
}
