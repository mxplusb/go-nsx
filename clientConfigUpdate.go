package go_nsx

type AutoReconnect struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ClientConfiguration struct {
	AutoReconnect       *AutoReconnect       `xml:" autoReconnect,omitempty" json:"autoReconnect,omitempty"`
	FullTunnel          *FullTunnel          `xml:" fullTunnel,omitempty" json:"fullTunnel,omitempty"`
	UpgradeNotification *UpgradeNotification `xml:" upgradeNotification,omitempty" json:"upgradeNotification,omitempty"`
}

type ExcludeLocalSubnets struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type FullTunnel struct {
	ExcludeLocalSubnets *ExcludeLocalSubnets `xml:" excludeLocalSubnets,omitempty" json:"excludeLocalSubnets,omitempty"`
	GatewayIp           *GatewayIp           `xml:" gatewayIp,omitempty" json:"gatewayIp,omitempty"`
}

type GatewayIp struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	ClientConfiguration *ClientConfiguration `xml:" clientConfiguration,omitempty" json:"clientConfiguration,omitempty"`
}

type UpgradeNotification struct {
	Text string `xml:",chardata" json:",omitempty"`
}
