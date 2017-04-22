package users

type LDAPServer struct {
	DomainId *DomainId `xml:" domainId,omitempty" json:"domainId,omitempty"`
	Enabled  *Enabled  `xml:" enabled,omitempty" json:"enabled,omitempty"`
	HostName *HostName `xml:" hostName,omitempty" json:"hostName,omitempty"`
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

type Root struct {
	LDAPServer *LDAPServer `xml:" LDAPServer,omitempty" json:"LDAPServer,omitempty"`
}
