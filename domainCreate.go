package go_nsx

type DirectoryDomain struct {
	Name        *Name        `xml:" name,omitempty" json:"name,omitempty"`
	NetbiosName *NetbiosName `xml:" netbiosName,omitempty" json:"netbiosName,omitempty"`
	Password    *Password    `xml:" password,omitempty" json:"password,omitempty"`
	Type        *Type        `xml:" type,omitempty" json:"type,omitempty"`
	Username    *Username    `xml:" username,omitempty" json:"username,omitempty"`
}

type Name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type NetbiosName struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Password struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	DirectoryDomain *DirectoryDomain `xml:" DirectoryDomain,omitempty" json:"DirectoryDomain,omitempty"`
}

type Type struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Username struct {
	Text string `xml:",chardata" json:",omitempty"`
}
