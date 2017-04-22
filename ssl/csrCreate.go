package go_nsx

type Algorithm struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Attribute struct {
	Key   *Key   `xml:" key,omitempty" json:"key,omitempty"`
	Value *Value `xml:" value,omitempty" json:"value,omitempty"`
}

type Csr struct {
	Algorithm *Algorithm `xml:" algorithm,omitempty" json:"algorithm,omitempty"`
	KeySize   *KeySize   `xml:" keySize,omitempty" json:"keySize,omitempty"`
	Subject   *Subject   `xml:" subject,omitempty" json:"subject,omitempty"`
}

type Key struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type KeySize struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	Csr *Csr `xml:" csr,omitempty" json:"csr,omitempty"`
}

type Subject struct {
	Attribute *Attribute `xml:" attribute,omitempty" json:"attribute,omitempty"`
}

type Value struct {
	Text string `xml:",chardata" json:",omitempty"`
}
