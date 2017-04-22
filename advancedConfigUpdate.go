package go_nsx

type AdvancedConfig struct {
	ClientNotification   *ClientNotification   `xml:" clientNotification,omitempty" json:"clientNotification,omitempty"`
	EnableCompression    *EnableCompression    `xml:" enableCompression,omitempty" json:"enableCompression,omitempty"`
	EnableLogging        *EnableLogging        `xml:" enableLogging,omitempty" json:"enableLogging,omitempty"`
	ForceVirtualKeyboard *ForceVirtualKeyboard `xml:" forceVirtualKeyboard,omitempty" json:"forceVirtualKeyboard,omitempty"`
	PreventMultipleLogon *PreventMultipleLogon `xml:" preventMultipleLogon,omitempty" json:"preventMultipleLogon,omitempty"`
	RandomizeVirtualkeys *RandomizeVirtualkeys `xml:" randomizeVirtualkeys,omitempty" json:"randomizeVirtualkeys,omitempty"`
	Timeout              *Timeout              `xml:" timeout,omitempty" json:"timeout,omitempty"`
}

type ClientNotification struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type EnableCompression struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type EnableLogging struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ForceVirtualKeyboard struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ForcedTimeout struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PreventMultipleLogon struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type RandomizeVirtualkeys struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type AdvancedConfigRoot struct {
	AdvancedConfig *AdvancedConfig `xml:" advancedConfig,omitempty" json:"advancedConfig,omitempty"`
}

type SessionIdleTimeout struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type AdvancedConfigTimeout struct {
	ForcedTimeout      *ForcedTimeout      `xml:" forcedTimeout,omitempty" json:"forcedTimeout,omitempty"`
	SessionIdleTimeout *SessionIdleTimeout `xml:" sessionIdleTimeout,omitempty" json:"sessionIdleTimeout,omitempty"`
}
