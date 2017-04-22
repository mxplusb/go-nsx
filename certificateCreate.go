package go_nsx

type PemEncoding struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	TrustObject *TrustObject `xml:" trustObject,omitempty" json:"trustObject,omitempty"`
}

type TrustObject struct {
	PemEncoding *PemEncoding `xml:" pemEncoding,omitempty" json:"pemEncoding,omitempty"`
}
