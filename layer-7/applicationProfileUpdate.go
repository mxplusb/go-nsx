package layer_7

type ApplicationProfile struct {
	InsertXForwardedFor *InsertXForwardedFor `xml:" insertXForwardedFor,omitempty" json:"insertXForwardedFor,omitempty"`
	Name                *Name                `xml:" name,omitempty" json:"name,omitempty"`
	Persistence         *Persistence         `xml:" persistence,omitempty" json:"persistence,omitempty"`
	SslPassthrough      *SslPassthrough      `xml:" sslPassthrough,omitempty" json:"sslPassthrough,omitempty"`
}

type CookieMode struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type CookieName struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type InsertXForwardedFor struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Method struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Persistence struct {
	CookieMode *CookieMode `xml:" cookieMode,omitempty" json:"cookieMode,omitempty"`
	CookieName *CookieName `xml:" cookieName,omitempty" json:"cookieName,omitempty"`
	Method     *Method     `xml:" method,omitempty" json:"method,omitempty"`
}

type Root struct {
	ApplicationProfile *ApplicationProfile `xml:" applicationProfile,omitempty" json:"applicationProfile,omitempty"`
}

type SslPassthrough struct {
	Text string `xml:",chardata" json:",omitempty"`
}
