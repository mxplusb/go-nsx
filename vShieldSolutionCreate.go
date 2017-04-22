package go_nsx

type SolutionInfo struct {
	Altitude    *Altitude    `xml:" altitude,omitempty" json:"altitude,omitempty"`
	Description *Description `xml:" description,omitempty" json:"description,omitempty"`
	Title       *Title       `xml:" title,omitempty" json:"title,omitempty"`
}

type Altitude struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Description struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	SolutionInfo *SolutionInfo `xml:" SolutionInfo,omitempty" json:"SolutionInfo,omitempty"`
}

type Title struct {
	Text string `xml:",chardata" json:",omitempty"`
}
