package go_nsx

type Long struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	Set *Set `xml:" set,omitempty" json:"set,omitempty"`
}

type Set struct {
	Long *Long `xml:" long,omitempty" json:"long,omitempty"`
}
