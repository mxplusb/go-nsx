package go_nsx

type Passphrase struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PemEncoding struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PrivateKey struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	TrustObject *TrustObject `xml:" trustObject,omitempty" json:"trustObject,omitempty"`
}

type TrustObject struct {
	Passphrase  *Passphrase  `xml:" passphrase,omitempty" json:"passphrase,omitempty"`
	PemEncoding *PemEncoding `xml:" pemEncoding,omitempty" json:"pemEncoding,omitempty"`
	PrivateKey  *PrivateKey  `xml:" privateKey,omitempty" json:"privateKey,omitempty"`
}
