package go_nsx

type ApiPassword struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ControllerCredential struct {
	ApiPassword *ApiPassword `xml:" apiPassword,omitempty" json:"apiPassword,omitempty"`
}

type Root struct {
	ControllerCredential *ControllerCredential `xml:" controllerCredential,omitempty" json:"controllerCredential,omitempty"`
}
