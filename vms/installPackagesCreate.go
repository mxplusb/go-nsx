package vms

type ClientInstallPackage struct {
	CreateDesktopIcon                   *CreateDesktopIcon                   `xml:" createDesktopIcon,omitempty" json:"createDesktopIcon,omitempty"`
	CreateLinuxClient                   *CreateLinuxClient                   `xml:" createLinuxClient,omitempty" json:"createLinuxClient,omitempty"`
	CreateMacClient                     *CreateMacClient                     `xml:" createMacClient,omitempty" json:"createMacClient,omitempty"`
	Description                         *Description                         `xml:" description,omitempty" json:"description,omitempty"`
	Enabled                             *Enabled                             `xml:" enabled,omitempty" json:"enabled,omitempty"`
	EnforceServerSecurityCertValidation *EnforceServerSecurityCertValidation `xml:" enforceServerSecurityCertValidation,omitempty" json:"enforceServerSecurityCertValidation,omitempty"`
	GatewayList                         *GatewayList                         `xml:" gatewayList,omitempty" json:"gatewayList,omitempty"`
	HideNetworkAdaptor                  *HideNetworkAdaptor                  `xml:" hideNetworkAdaptor,omitempty" json:"hideNetworkAdaptor,omitempty"`
	HideSystrayIcon                     *HideSystrayIcon                     `xml:" hideSystrayIcon,omitempty" json:"hideSystrayIcon,omitempty"`
	ProfileName                         *ProfileName                         `xml:" profileName,omitempty" json:"profileName,omitempty"`
	RememberPassword                    *RememberPassword                    `xml:" rememberPassword,omitempty" json:"rememberPassword,omitempty"`
	SilentModeInstallation              *SilentModeInstallation              `xml:" silentModeInstallation,omitempty" json:"silentModeInstallation,omitempty"`
	SilentModeOperation                 *SilentModeOperation                 `xml:" silentModeOperation,omitempty" json:"silentModeOperation,omitempty"`
	StartClientOnLogon                  *StartClientOnLogon                  `xml:" startClientOnLogon,omitempty" json:"startClientOnLogon,omitempty"`
}

type CreateDesktopIcon struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type CreateLinuxClient struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type CreateMacClient struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Description struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Enabled struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type EnforceServerSecurityCertValidation struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Gateway struct {
	HostName *HostName `xml:" hostName,omitempty" json:"hostName,omitempty"`
	Port     *Port     `xml:" port,omitempty" json:"port,omitempty"`
}

type GatewayList struct {
	Gateway *Gateway `xml:" gateway,omitempty" json:"gateway,omitempty"`
}

type HideNetworkAdaptor struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type HideSystrayIcon struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type HostName struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Port struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ProfileName struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type RememberPassword struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	ClientInstallPackage *ClientInstallPackage `xml:" clientInstallPackage,omitempty" json:"clientInstallPackage,omitempty"`
}

type SilentModeInstallation struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type SilentModeOperation struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type StartClientOnLogon struct {
	Text string `xml:",chardata" json:",omitempty"`
}
