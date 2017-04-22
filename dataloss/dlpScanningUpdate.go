package go_nsx

type ScanOp struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	ScanOp *ScanOp `xml:" ScanOp,omitempty" json:"ScanOp,omitempty"`
}
