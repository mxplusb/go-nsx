package go_nsx

type Command struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Nsxcli struct {
	Command *Command `xml:" command,omitempty" json:"command,omitempty"`
}

type Root struct {
	Nsxcli *Nsxcli `xml:" nsxcli,omitempty" json:"nsxcli,omitempty"`
}
