package go_nsx

type CliSettings struct {
	Password           *Password           `xml:" password,omitempty" json:"password,omitempty"`
	PasswordExpiry     *PasswordExpiry     `xml:" passwordExpiry,omitempty" json:"passwordExpiry,omitempty"`
	RemoteAccess       *RemoteAccess       `xml:" remoteAccess,omitempty" json:"remoteAccess,omitempty"`
	SshLoginBannerText *SshLoginBannerText `xml:" sshLoginBannerText,omitempty" json:"sshLoginBannerText,omitempty"`
	UserName           *UserName           `xml:" userName,omitempty" json:"userName,omitempty"`
}

type Password struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PasswordExpiry struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type RemoteAccess struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	CliSettings *CliSettings `xml:" cliSettings,omitempty" json:"cliSettings,omitempty"`
}

type SshLoginBannerText struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type UserName struct {
	Text string `xml:",chardata" json:",omitempty"`
}
