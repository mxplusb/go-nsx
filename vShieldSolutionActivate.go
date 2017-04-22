package go_nsx

type ActivationInfo struct {
	Moid *Moid `xml:" moid,omitempty" json:"moid,omitempty"`
}

type Moid struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	ActivationInfo *ActivationInfo `xml:" ActivationInfo,omitempty" json:"ActivationInfo,omitempty"`
}
