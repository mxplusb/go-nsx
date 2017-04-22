package go_nsx

type EventlogServer struct {
	DomainId *DomainId `xml:" domainId,omitempty" json:"domainId,omitempty"`
	Enabled  *Enabled  `xml:" enabled,omitempty" json:"enabled,omitempty"`
	HostName *HostName `xml:" hostName,omitempty" json:"hostName,omitempty"`
	Id       *Id       `xml:" id,omitempty" json:"id,omitempty"`
}

type DomainId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Enabled struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type HostName struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Id struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	EventlogServer *EventlogServer `xml:" EventlogServer,omitempty" json:"EventlogServer,omitempty"`
}
