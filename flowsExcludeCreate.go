package go_nsx

type FlowConfiguration struct {
	CollectFlows         *CollectFlows         `xml:" collectFlows,omitempty" json:"collectFlows,omitempty"`
	DestinationContainer *DestinationContainer `xml:" destinationContainer,omitempty" json:"destinationContainer,omitempty"`
	DestinationIPs       *DestinationIPs       `xml:" destinationIPs,omitempty" json:"destinationIPs,omitempty"`
	DestinationPorts     *DestinationPorts     `xml:" destinationPorts,omitempty" json:"destinationPorts,omitempty"`
	IgnoreBlockedFlows   *IgnoreBlockedFlows   `xml:" ignoreBlockedFlows,omitempty" json:"ignoreBlockedFlows,omitempty"`
	IgnoreLayer2Flows    *IgnoreLayer2Flows    `xml:" ignoreLayer2Flows,omitempty" json:"ignoreLayer2Flows,omitempty"`
	Service              *Service              `xml:" service,omitempty" json:"service,omitempty"`
	SourceContainer      *SourceContainer      `xml:" sourceContainer,omitempty" json:"sourceContainer,omitempty"`
	SourceIPs            *SourceIPs            `xml:" sourceIPs,omitempty" json:"sourceIPs,omitempty"`
}

type CollectFlows struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DestinationContainer struct {
	Id   *Id   `xml:" id,omitempty" json:"id,omitempty"`
	Name *Name `xml:" name,omitempty" json:"name,omitempty"`
	Type *Type `xml:" type,omitempty" json:"type,omitempty"`
}

type DestinationIPs struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DestinationPorts struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Id struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IgnoreBlockedFlows struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IgnoreLayer2Flows struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	FlowConfiguration *FlowConfiguration `xml:" FlowConfiguration,omitempty" json:"FlowConfiguration,omitempty"`
}

type Service struct {
	Id   *Id   `xml:" id,omitempty" json:"id,omitempty"`
	Name *Name `xml:" name,omitempty" json:"name,omitempty"`
}

type SourceContainer struct {
	Id   *Id   `xml:" id,omitempty" json:"id,omitempty"`
	Name *Name `xml:" name,omitempty" json:"name,omitempty"`
	Type *Type `xml:" type,omitempty" json:"type,omitempty"`
}

type SourceIPs struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Type struct {
	Text string `xml:",chardata" json:",omitempty"`
}
