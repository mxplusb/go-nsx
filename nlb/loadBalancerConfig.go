package go_nsx

type AccelerationEnabled struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ApplicationProfile struct {
	ApplicationProfileId *ApplicationProfileId `xml:" applicationProfileId,omitempty" json:"applicationProfileId,omitempty"`
	ClientSsl            *ClientSsl            `xml:" clientSsl,omitempty" json:"clientSsl,omitempty"`
	InsertXForwardedFor  *InsertXForwardedFor  `xml:" insertXForwardedFor,omitempty" json:"insertXForwardedFor,omitempty"`
	Name                 *Name                 `xml:" name,omitempty" json:"name,omitempty"`
	Persistence          *Persistence          `xml:" persistence,omitempty" json:"persistence,omitempty"`
	ServerSslEnabled     *ServerSslEnabled     `xml:" serverSslEnabled,omitempty" json:"serverSslEnabled,omitempty"`
	SslPassthrough       *SslPassthrough       `xml:" sslPassthrough,omitempty" json:"sslPassthrough,omitempty"`
	Template             *Template             `xml:" template,omitempty" json:"template,omitempty"`
}

type ApplicationProfileId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ClientSsl struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ConnectionLimit struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ConnectionRateLimit struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type CookieMode struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type CookieName struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DefaultPoolId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Description struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Enable struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type EnableServiceInsertion struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Enabled struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type InsertXForwardedFor struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IpAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type LoadBalancer struct {
	AccelerationEnabled    *AccelerationEnabled    `xml:" accelerationEnabled,omitempty" json:"accelerationEnabled,omitempty"`
	ApplicationProfile     *ApplicationProfile     `xml:" applicationProfile,omitempty" json:"applicationProfile,omitempty"`
	EnableServiceInsertion *EnableServiceInsertion `xml:" enableServiceInsertion,omitempty" json:"enableServiceInsertion,omitempty"`
	Enabled                *Enabled                `xml:" enabled,omitempty" json:"enabled,omitempty"`
	Logging                *Logging                `xml:" logging,omitempty" json:"logging,omitempty"`
	VirtualServer          *VirtualServer          `xml:" virtualServer,omitempty" json:"virtualServer,omitempty"`
}

type LogLevel struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Logging struct {
	Enable   *Enable   `xml:" enable,omitempty" json:"enable,omitempty"`
	LogLevel *LogLevel `xml:" logLevel,omitempty" json:"logLevel,omitempty"`
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

type Port struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Protocol struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	LoadBalancer *LoadBalancer `xml:" loadBalancer,omitempty" json:"loadBalancer,omitempty"`
}

type ServerSslEnabled struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type SslPassthrough struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Template struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type VirtualServer struct {
	AccelerationEnabled    *AccelerationEnabled    `xml:" accelerationEnabled,omitempty" json:"accelerationEnabled,omitempty"`
	ApplicationProfileId   *ApplicationProfileId   `xml:" applicationProfileId,omitempty" json:"applicationProfileId,omitempty"`
	ConnectionLimit        *ConnectionLimit        `xml:" connectionLimit,omitempty" json:"connectionLimit,omitempty"`
	ConnectionRateLimit    *ConnectionRateLimit    `xml:" connectionRateLimit,omitempty" json:"connectionRateLimit,omitempty"`
	DefaultPoolId          *DefaultPoolId          `xml:" defaultPoolId,omitempty" json:"defaultPoolId,omitempty"`
	Description            *Description            `xml:" description,omitempty" json:"description,omitempty"`
	EnableServiceInsertion *EnableServiceInsertion `xml:" enableServiceInsertion,omitempty" json:"enableServiceInsertion,omitempty"`
	Enabled                *Enabled                `xml:" enabled,omitempty" json:"enabled,omitempty"`
	IpAddress              *IpAddress              `xml:" ipAddress,omitempty" json:"ipAddress,omitempty"`
	Name                   *Name                   `xml:" name,omitempty" json:"name,omitempty"`
	Port                   *Port                   `xml:" port,omitempty" json:"port,omitempty"`
	Protocol               *Protocol               `xml:" protocol,omitempty" json:"protocol,omitempty"`
	VirtualServerId        *VirtualServerId        `xml:" virtualServerId,omitempty" json:"virtualServerId,omitempty"`
}

type VirtualServerId struct {
	Text string `xml:",chardata" json:",omitempty"`
}
