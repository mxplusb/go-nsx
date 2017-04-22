package dataloss

type Root struct {
	Set *Set `xml:" set,omitempty" json:"set,omitempty"`
}

type Set struct {
	String *String `xml:" string,omitempty" json:"string,omitempty"`
}

type String struct {
	Text string `xml:",chardata" json:",omitempty"`
}
