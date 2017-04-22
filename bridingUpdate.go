package go_nsx

type Bridge struct {
	DvportGroup *DvportGroup `xml:" dvportGroup,omitempty" json:"dvportGroup,omitempty"`
	Name        *Name        `xml:" name,omitempty" json:"name,omitempty"`
	VirtualWire *VirtualWire `xml:" virtualWire,omitempty" json:"virtualWire,omitempty"`
}

type Bridges struct {
	Bridge  *Bridge  `xml:" bridge,omitempty" json:"bridge,omitempty"`
	Enabled *Enabled `xml:" enabled,omitempty" json:"enabled,omitempty"`
	Version *Version `xml:" version,omitempty" json:"version,omitempty"`
}

type DvportGroup struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Enabled struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	Bridges *Bridges `xml:" bridges,omitempty" json:"bridges,omitempty"`
}

type Version struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type VirtualWire struct {
	Text string `xml:",chardata" json:",omitempty"`
}
