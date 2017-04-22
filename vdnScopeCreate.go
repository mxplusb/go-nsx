package go_nsx

type Cluster struct {
	Cluster  *Cluster  `xml:" cluster,omitempty" json:"cluster,omitempty"`
	ObjectId *ObjectId `xml:" objectId,omitempty" json:"objectId,omitempty"`
}

type Clusters struct {
	Cluster *Cluster `xml:" cluster,omitempty" json:"cluster,omitempty"`
}

type ControlPlaneMode struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Description struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ObjectId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	VdnScope *VdnScope `xml:" vdnScope,omitempty" json:"vdnScope,omitempty"`
}

type VdnScope struct {
	Clusters         *Clusters         `xml:" clusters,omitempty" json:"clusters,omitempty"`
	ControlPlaneMode *ControlPlaneMode `xml:" controlPlaneMode,omitempty" json:"controlPlaneMode,omitempty"`
	Description      *Description      `xml:" description,omitempty" json:"description,omitempty"`
	Name             *Name             `xml:" name,omitempty" json:"name,omitempty"`
}
