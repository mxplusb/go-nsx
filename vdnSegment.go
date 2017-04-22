package go_nsx

type Begin struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type End struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	SegmentRange *SegmentRange `xml:" segmentRange,omitempty" json:"segmentRange,omitempty"`
}

type SegmentRange struct {
	Begin *Begin `xml:" begin,omitempty" json:"begin,omitempty"`
	End   *End   `xml:" end,omitempty" json:"end,omitempty"`
	Name  *Name  `xml:" name,omitempty" json:"name,omitempty"`
}
