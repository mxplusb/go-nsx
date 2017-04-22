package go_nsx

type ControllerConfig struct {
	SslEnabled *SslEnabled `xml:" sslEnabled,omitempty" json:"sslEnabled,omitempty"`
}

type Root struct {
	ControllerConfig *ControllerConfig `xml:" controllerConfig,omitempty" json:"controllerConfig,omitempty"`
}

type SslEnabled struct {
	Text string `xml:",chardata" json:",omitempty"`
}
