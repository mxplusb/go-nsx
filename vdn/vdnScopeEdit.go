package go_nsx

type Cluster struct {
	Cluster  *Cluster  `xml:" cluster,omitempty" json:"cluster,omitempty"`
	ObjectId *ObjectId `xml:" objectId,omitempty" json:"objectId,omitempty"`
}

type Clusters struct {
	Cluster *Cluster `xml:" cluster,omitempty" json:"cluster,omitempty"`
}

type ObjectId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	VdnScope *VdnScope `xml:" vdnScope,omitempty" json:"vdnScope,omitempty"`
}

type VdnScope struct {
	Clusters *Clusters `xml:" clusters,omitempty" json:"clusters,omitempty"`
	ObjectId *ObjectId `xml:" objectId,omitempty" json:"objectId,omitempty"`
}
