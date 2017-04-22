package go_nsx

type DropInvalidTraffic struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type EnableSynFloodProtection struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type GlobalConfig struct {
	DropInvalidTraffic            *DropInvalidTraffic            `xml:" dropInvalidTraffic,omitempty" json:"dropInvalidTraffic,omitempty"`
	EnableSynFloodProtection      *EnableSynFloodProtection      `xml:" enableSynFloodProtection,omitempty" json:"enableSynFloodProtection,omitempty"`
	Icmp6Timeout                  *Icmp6Timeout                  `xml:" icmp6Timeout,omitempty" json:"icmp6Timeout,omitempty"`
	IcmpTimeout                   *IcmpTimeout                   `xml:" icmpTimeout,omitempty" json:"icmpTimeout,omitempty"`
	IpGenericTimeout              *IpGenericTimeout              `xml:" ipGenericTimeout,omitempty" json:"ipGenericTimeout,omitempty"`
	LogInvalidTraffic             *LogInvalidTraffic             `xml:" logInvalidTraffic,omitempty" json:"logInvalidTraffic,omitempty"`
	TcpAllowOutOfWindowPackets    *TcpAllowOutOfWindowPackets    `xml:" tcpAllowOutOfWindowPackets,omitempty" json:"tcpAllowOutOfWindowPackets,omitempty"`
	TcpPickOngoingConnections     *TcpPickOngoingConnections     `xml:" tcpPickOngoingConnections,omitempty" json:"tcpPickOngoingConnections,omitempty"`
	TcpSendResetForClosedVsePorts *TcpSendResetForClosedVsePorts `xml:" tcpSendResetForClosedVsePorts,omitempty" json:"tcpSendResetForClosedVsePorts,omitempty"`
	TcpTimeoutClose               *TcpTimeoutClose               `xml:" tcpTimeoutClose,omitempty" json:"tcpTimeoutClose,omitempty"`
	TcpTimeoutEstablished         *TcpTimeoutEstablished         `xml:" tcpTimeoutEstablished,omitempty" json:"tcpTimeoutEstablished,omitempty"`
	TcpTimeoutOpen                *TcpTimeoutOpen                `xml:" tcpTimeoutOpen,omitempty" json:"tcpTimeoutOpen,omitempty"`
	UdpTimeout                    *UdpTimeout                    `xml:" udpTimeout,omitempty" json:"udpTimeout,omitempty"`
}

type Icmp6Timeout struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IcmpTimeout struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IpGenericTimeout struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type LogInvalidTraffic struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	GlobalConfig *GlobalConfig `xml:" globalConfig,omitempty" json:"globalConfig,omitempty"`
}

type TcpAllowOutOfWindowPackets struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type TcpPickOngoingConnections struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type TcpSendResetForClosedVsePorts struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type TcpTimeoutClose struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type TcpTimeoutEstablished struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type TcpTimeoutOpen struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type UdpTimeout struct {
	Text string `xml:",chardata" json:",omitempty"`
}
