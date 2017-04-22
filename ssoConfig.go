package go_nsx

type CertificateThumbprint struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	SsoConfig *SsoConfig `xml:" ssoConfig,omitempty" json:"ssoConfig,omitempty"`
}

type SsoAdminUsername struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type SsoAdminUserpassword struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type SsoConfig struct {
	CertificateThumbprint *CertificateThumbprint `xml:" certificateThumbprint,omitempty" json:"certificateThumbprint,omitempty"`
	SsoAdminUsername      *SsoAdminUsername      `xml:" ssoAdminUsername,omitempty" json:"ssoAdminUsername,omitempty"`
	SsoAdminUserpassword  *SsoAdminUserpassword  `xml:" ssoAdminUserpassword,omitempty" json:"ssoAdminUserpassword,omitempty"`
	SsoLookupServiceUrl   *SsoLookupServiceUrl   `xml:" ssoLookupServiceUrl,omitempty" json:"ssoLookupServiceUrl,omitempty"`
}

type SsoLookupServiceUrl struct {
	Text string `xml:",chardata" json:",omitempty"`
}
