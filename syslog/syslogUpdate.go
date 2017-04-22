package syslog

type IpAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Protocol struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	Syslog *Syslog `xml:" syslog,omitempty" json:"syslog,omitempty"`
}

type ServerAddresses struct {
	IpAddress *IpAddress `xml:" ipAddress,omitempty" json:"ipAddress,omitempty"`
}

type Syslog struct {
	Protocol        *Protocol        `xml:" protocol,omitempty" json:"protocol,omitempty"`
	ServerAddresses *ServerAddresses `xml:" serverAddresses,omitempty" json:"serverAddresses,omitempty"`
}
