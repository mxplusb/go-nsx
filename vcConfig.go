package go_nsx

type AssignRoleToUser struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type CertificateThumbprint struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IpAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Password struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PluginDownloadPort struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PluginDownloadServer struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	VcInfo *VcInfo `xml:" vcInfo,omitempty" json:"vcInfo,omitempty"`
}

type UserName struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type VcInfo struct {
	AssignRoleToUser      *AssignRoleToUser      `xml:" assignRoleToUser,omitempty" json:"assignRoleToUser,omitempty"`
	CertificateThumbprint *CertificateThumbprint `xml:" certificateThumbprint,omitempty" json:"certificateThumbprint,omitempty"`
	IpAddress             *IpAddress             `xml:" ipAddress,omitempty" json:"ipAddress,omitempty"`
	Password              *Password              `xml:" password,omitempty" json:"password,omitempty"`
	PluginDownloadPort    *PluginDownloadPort    `xml:" pluginDownloadPort,omitempty" json:"pluginDownloadPort,omitempty"`
	PluginDownloadServer  *PluginDownloadServer  `xml:" pluginDownloadServer,omitempty" json:"pluginDownloadServer,omitempty"`
	UserName              *UserName              `xml:" userName,omitempty" json:"userName,omitempty"`
}
