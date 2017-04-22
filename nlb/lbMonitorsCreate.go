package nlb

type Interval struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type MaxRetries struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Method struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Monitor struct {
	Interval   *Interval   `xml:" interval,omitempty" json:"interval,omitempty"`
	MaxRetries *MaxRetries `xml:" maxRetries,omitempty" json:"maxRetries,omitempty"`
	Method     *Method     `xml:" method,omitempty" json:"method,omitempty"`
	Name       *Name       `xml:" name,omitempty" json:"name,omitempty"`
	Timeout    *Timeout    `xml:" timeout,omitempty" json:"timeout,omitempty"`
	Type       *Type       `xml:" type,omitempty" json:"type,omitempty"`
	Url        *Url        `xml:" url,omitempty" json:"url,omitempty"`
}

type Name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	Monitor *Monitor `xml:" monitor,omitempty" json:"monitor,omitempty"`
}

type Timeout struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Type struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Url struct {
	Text string `xml:",chardata" json:",omitempty"`
}
