package go_nsx

type ClassificationValue struct {
	Classification *Classification `xml:" classification,omitempty" json:"classification,omitempty"`
	Id             *Id             `xml:" id,omitempty" json:"id,omitempty"`
	Value          *Value          `xml:" value,omitempty" json:"value,omitempty"`
}

type Classification struct {
	Customizable *Customizable `xml:" customizable,omitempty" json:"customizable,omitempty"`
	Description  *Description  `xml:" description,omitempty" json:"description,omitempty"`
	Id           *Id           `xml:" id,omitempty" json:"id,omitempty"`
	Name         *Name         `xml:" name,omitempty" json:"name,omitempty"`
	ProviderName *ProviderName `xml:" providerName,omitempty" json:"providerName,omitempty"`
}

type Customizable struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Description struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Id struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ProviderName struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	Set *Set `xml:" set,omitempty" json:"set,omitempty"`
}

type Set struct {
	ClassificationValue *ClassificationValue `xml:" ClassificationValue,omitempty" json:"ClassificationValue,omitempty"`
}

type Value struct {
	Text string `xml:",chardata" json:",omitempty"`
}
