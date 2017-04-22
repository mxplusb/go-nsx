package go_nsx

type Port struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Protocol struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	Syslogserver *Syslogserver `xml:" syslogserver,omitempty" json:"syslogserver,omitempty"`
}

type SyslogServer struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Syslogserver struct {
	Port         *Port         `xml:" port,omitempty" json:"port,omitempty"`
	Protocol     *Protocol     `xml:" protocol,omitempty" json:"protocol,omitempty"`
	SyslogServer *SyslogServer `xml:" syslogServer,omitempty" json:"syslogServer,omitempty"`
}
