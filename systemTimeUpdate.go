package go_nsx

type Datetime struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type NtpServer struct {
	String *String `xml:" string,omitempty" json:"string,omitempty"`
}

type Root struct {
	TimeSettings *TimeSettings `xml:" timeSettings,omitempty" json:"timeSettings,omitempty"`
}

type String struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type TimeSettings struct {
	Datetime  *Datetime  `xml:" datetime,omitempty" json:"datetime,omitempty"`
	NtpServer *NtpServer `xml:" ntpServer,omitempty" json:"ntpServer,omitempty"`
	Timezone  *Timezone  `xml:" timezone,omitempty" json:"timezone,omitempty"`
}

type Timezone struct {
	Text string `xml:",chardata" json:",omitempty"`
}
