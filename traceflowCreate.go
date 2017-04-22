package go_nsx

type DstIp struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DstMac struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type EthHeader struct {
	DstMac  *DstMac  `xml:" dstMac,omitempty" json:"dstMac,omitempty"`
	EthType *EthType `xml:" ethType,omitempty" json:"ethType,omitempty"`
	SrcMac  *SrcMac  `xml:" srcMac,omitempty" json:"srcMac,omitempty"`
}

type EthType struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IpHeader struct {
	DstIp *DstIp `xml:" dstIp,omitempty" json:"dstIp,omitempty"`
	SrcIp *SrcIp `xml:" srcIp,omitempty" json:"srcIp,omitempty"`
	Ttl   *Ttl   `xml:" ttl,omitempty" json:"ttl,omitempty"`
}

type Packet struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ResourceType struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	TraceflowRequest *TraceflowRequest `xml:" traceflowRequest,omitempty" json:"traceflowRequest,omitempty"`
}

type Routed struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type SrcIp struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type SrcMac struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Timeout struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type TraceflowRequest struct {
	EthHeader    *EthHeader    `xml:" ethHeader,omitempty" json:"ethHeader,omitempty"`
	IpHeader     *IpHeader     `xml:" ipHeader,omitempty" json:"ipHeader,omitempty"`
	Packet       *Packet       `xml:" packet,omitempty" json:"packet,omitempty"`
	ResourceType *ResourceType `xml:" resourceType,omitempty" json:"resourceType,omitempty"`
	Routed       *Routed       `xml:" routed,omitempty" json:"routed,omitempty"`
	Timeout      *Timeout      `xml:" timeout,omitempty" json:"timeout,omitempty"`
	VmnicId      *VmnicId      `xml:" vmnicId,omitempty" json:"vmnicId,omitempty"`
}

type Ttl struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type VmnicId struct {
	Text string `xml:",chardata" json:",omitempty"`
}
