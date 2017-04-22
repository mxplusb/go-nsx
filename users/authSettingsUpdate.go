package users

type AccountLockoutPolicy struct {
	LockoutDuration *LockoutDuration `xml:" lockoutDuration,omitempty" json:"lockoutDuration,omitempty"`
	RetryCount      *RetryCount      `xml:" retryCount,omitempty" json:"retryCount,omitempty"`
	RetryDuration   *RetryDuration   `xml:" retryDuration,omitempty" json:"retryDuration,omitempty"`
}

type AllowUserIdWithinPassword struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type AuthenticationConfig struct {
	PasswordAuthentication *PasswordAuthentication `xml:" passwordAuthentication,omitempty" json:"passwordAuthentication,omitempty"`
}

type AuthenticationTimeout struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type BindDomainName struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type BindPassword struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Com_dot_vmware_dot_vshield_dot_edge_dot_sslvpn_dot_dto_dot_AdAuthServerDto struct {
	BindDomainName              *BindDomainName              `xml:" bindDomainName,omitempty" json:"bindDomainName,omitempty"`
	BindPassword                *BindPassword                `xml:" bindPassword,omitempty" json:"bindPassword,omitempty"`
	EnableSsl                   *EnableSsl                   `xml:" enableSsl,omitempty" json:"enableSsl,omitempty"`
	Enabled                     *Enabled                     `xml:" enabled,omitempty" json:"enabled,omitempty"`
	Ip                          *Ip                          `xml:" ip,omitempty" json:"ip,omitempty"`
	LoginAttributeName          *LoginAttributeName          `xml:" loginAttributeName,omitempty" json:"loginAttributeName,omitempty"`
	Port                        *Port                        `xml:" port,omitempty" json:"port,omitempty"`
	SearchBase                  *SearchBase                  `xml:" searchBase,omitempty" json:"searchBase,omitempty"`
	SearchFilter                *SearchFilter                `xml:" searchFilter,omitempty" json:"searchFilter,omitempty"`
	TerminateSessionOnAuthFails *TerminateSessionOnAuthFails `xml:" terminateSessionOnAuthFails,omitempty" json:"terminateSessionOnAuthFails,omitempty"`
	TimeOut                     *TimeOut                     `xml:" timeOut,omitempty" json:"timeOut,omitempty"`
}

type Com_dot_vmware_dot_vshield_dot_edge_dot_sslvpn_dot_dto_dot_LdapAuthServerDto struct {
	BindDomainName     *BindDomainName     `xml:" bindDomainName,omitempty" json:"bindDomainName,omitempty"`
	BindPassword       *BindPassword       `xml:" bindPassword,omitempty" json:"bindPassword,omitempty"`
	EnableSsl          *EnableSsl          `xml:" enableSsl,omitempty" json:"enableSsl,omitempty"`
	Enabled            *Enabled            `xml:" enabled,omitempty" json:"enabled,omitempty"`
	Ip                 *Ip                 `xml:" ip,omitempty" json:"ip,omitempty"`
	LoginAttributeName *LoginAttributeName `xml:" loginAttributeName,omitempty" json:"loginAttributeName,omitempty"`
	Port               *Port               `xml:" port,omitempty" json:"port,omitempty"`
	SearchBase         *SearchBase         `xml:" searchBase,omitempty" json:"searchBase,omitempty"`
	SearchFilter       *SearchFilter       `xml:" searchFilter,omitempty" json:"searchFilter,omitempty"`
	TimeOut            *TimeOut            `xml:" timeOut,omitempty" json:"timeOut,omitempty"`
}

type Com_dot_vmware_dot_vshield_dot_edge_dot_sslvpn_dot_dto_dot_LocalAuthServerDto struct {
	AccountLockoutPolicy *AccountLockoutPolicy `xml:" accountLockoutPolicy,omitempty" json:"accountLockoutPolicy,omitempty"`
	Enabled              *Enabled              `xml:" enabled,omitempty" json:"enabled,omitempty"`
	PasswordPolicy       *PasswordPolicy       `xml:" passwordPolicy,omitempty" json:"passwordPolicy,omitempty"`
}

type RadiusAuthServerD struct {
	Ip         *Ip         `xml:" ip,omitempty" json:"ip,omitempty"`
	NasIp      *NasIp      `xml:" nasIp,omitempty" json:"nasIp,omitempty"`
	Port       *Port       `xml:" port,omitempty" json:"port,omitempty"`
	RetryCount *RetryCount `xml:" retryCount,omitempty" json:"retryCount,omitempty"`
	Secret     *Secret     `xml:" secret,omitempty" json:"secret,omitempty"`
	TimeOut    *TimeOut    `xml:" timeOut,omitempty" json:"timeOut,omitempty"`
}

type Com_dot_vmware_dot_vshield_dot_edge_dot_sslvpn_dot_dto_dot_RsaAuthServerDto struct {
	SourceIp *SourceIp `xml:" sourceIp,omitempty" json:"sourceIp,omitempty"`
	TimeOut  *TimeOut  `xml:" timeOut,omitempty" json:"timeOut,omitempty"`
}

type EnableSsl struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Enabled struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ExpiryNotification struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Ip struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type LockoutDuration struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type LoginAttributeName struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type MaxLength struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type MinAlphabets struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type MinDigits struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type MinLength struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type MinSpecialChar struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type NasIp struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PasswordAuthentication struct {
	AuthenticationTimeout *AuthenticationTimeout `xml:" authenticationTimeout,omitempty" json:"authenticationTimeout,omitempty"`
	PrimaryAuthServers    *PrimaryAuthServers    `xml:" primaryAuthServers,omitempty" json:"primaryAuthServers,omitempty"`
	SecondaryAuthServer   *SecondaryAuthServer   `xml:" secondaryAuthServer,omitempty" json:"secondaryAuthServer,omitempty"`
}

type PasswordLifeTime struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PasswordPolicy struct {
	AllowUserIdWithinPassword *AllowUserIdWithinPassword `xml:" allowUserIdWithinPassword,omitempty" json:"allowUserIdWithinPassword,omitempty"`
	ExpiryNotification        *ExpiryNotification        `xml:" expiryNotification,omitempty" json:"expiryNotification,omitempty"`
	MaxLength                 *MaxLength                 `xml:" maxLength,omitempty" json:"maxLength,omitempty"`
	MinAlphabets              *MinAlphabets              `xml:" minAlphabets,omitempty" json:"minAlphabets,omitempty"`
	MinDigits                 *MinDigits                 `xml:" minDigits,omitempty" json:"minDigits,omitempty"`
	MinLength                 *MinLength                 `xml:" minLength,omitempty" json:"minLength,omitempty"`
	MinSpecialChar            *MinSpecialChar            `xml:" minSpecialChar,omitempty" json:"minSpecialChar,omitempty"`
	PasswordLifeTime          *PasswordLifeTime          `xml:" passwordLifeTime,omitempty" json:"passwordLifeTime,omitempty"`
}

type Port struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PrimaryAuthServers struct {
	Com_dot_vmware_dot_vshield_dot_edge_dot_sslvpn_dot_dto_dot_LdapAuthServerDto   *Com_dot_vmware_dot_vshield_dot_edge_dot_sslvpn_dot_dto_dot_LdapAuthServerDto   `xml:" com.vmware.vshield.edge.sslvpn.dto.LdapAuthServerDto,omitempty" json:"com.vmware.vshield.edge.sslvpn.dto.LdapAuthServerDto,omitempty"`
	Com_dot_vmware_dot_vshield_dot_edge_dot_sslvpn_dot_dto_dot_LocalAuthServerDto  *Com_dot_vmware_dot_vshield_dot_edge_dot_sslvpn_dot_dto_dot_LocalAuthServerDto  `xml:" com.vmware.vshield.edge.sslvpn.dto.LocalAuthServerDto,omitempty" json:"com.vmware.vshield.edge.sslvpn.dto.LocalAuthServerDto,omitempty"`
	Com_dot_vmware_dot_vshield_dot_edge_dot_sslvpn_dot_dto_dot_RadiusAuthServerDto *RadiusAuthServerD `xml:" com.vmware.vshield.edge.sslvpn.dto.RadiusAuthServerDto,omitempty" json:"com.vmware.vshield.edge.sslvpn.dto.RadiusAuthServerDto,omitempty"`
	Com_dot_vmware_dot_vshield_dot_edge_dot_sslvpn_dot_dto_dot_RsaAuthServerDto    *Com_dot_vmware_dot_vshield_dot_edge_dot_sslvpn_dot_dto_dot_RsaAuthServerDto    `xml:" com.vmware.vshield.edge.sslvpn.dto.RsaAuthServerDto,omitempty" json:"com.vmware.vshield.edge.sslvpn.dto.RsaAuthServerDto,omitempty"`
}

type RetryCount struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type RetryDuration struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	AuthenticationConfig *AuthenticationConfig `xml:" authenticationConfig,omitempty" json:"authenticationConfig,omitempty"`
}

type SearchBase struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type SearchFilter struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type SecondaryAuthServer struct {
	Com_dot_vmware_dot_vshield_dot_edge_dot_sslvpn_dot_dto_dot_AdAuthServerDto *Com_dot_vmware_dot_vshield_dot_edge_dot_sslvpn_dot_dto_dot_AdAuthServerDto `xml:" com.vmware.vshield.edge.sslvpn.dto.AdAuthServerDto,omitempty" json:"com.vmware.vshield.edge.sslvpn.dto.AdAuthServerDto,omitempty"`
}

type Secret struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type SourceIp struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type TerminateSessionOnAuthFails struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type TimeOut struct {
	Text string `xml:",chardata" json:",omitempty"`
}
