package controller

type ControllerSyslogServer struct {
	Level        *Level        `xml:" level,omitempty" json:"level,omitempty"`
	Port         *Port         `xml:" port,omitempty" json:"port,omitempty"`
	Protocol     *Protocol     `xml:" protocol,omitempty" json:"protocol,omitempty"`
	SyslogServer *SyslogServer `xml:" syslogServer,omitempty" json:"syslogServer,omitempty"`
}

type Level struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Port struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Protocol struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	ControllerSyslogServer *ControllerSyslogServer `xml:" controllerSyslogServer,omitempty" json:"controllerSyslogServer,omitempty"`
}

type SyslogServer struct {
	Text string `xml:",chardata" json:",omitempty"`
}
