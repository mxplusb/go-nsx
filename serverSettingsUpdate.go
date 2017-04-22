package go_nsx

type CertificateId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IpAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Port struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	ServerSettings *ServerSettings `xml:" serverSettings,omitempty" json:"serverSettings,omitempty"`
}

type ServerAddresses struct {
	IpAddress *IpAddress `xml:" ipAddress,omitempty" json:"ipAddress,omitempty"`
}

type ServerSettings struct {
	CertificateId   *CertificateId   `xml:" certificateId,omitempty" json:"certificateId,omitempty"`
	Port            *Port            `xml:" port,omitempty" json:"port,omitempty"`
	ServerAddresses *ServerAddresses `xml:" serverAddresses,omitempty" json:"serverAddresses,omitempty"`
	SslVersionList  *SslVersionList  `xml:" sslVersionList,omitempty" json:"sslVersionList,omitempty"`
}

type SslVersionList struct {
	Version *Version `xml:" version,omitempty" json:"version,omitempty"`
}

type Version struct {
	Text string `xml:",chardata" json:",omitempty"`
}
