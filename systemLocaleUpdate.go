package go_nsx

type Country struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Language struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Locale struct {
	Country  *Country  `xml:" country,omitempty" json:"country,omitempty"`
	Language *Language `xml:" language,omitempty" json:"language,omitempty"`
}

type Root struct {
	Locale *Locale `xml:" locale,omitempty" json:"locale,omitempty"`
}
