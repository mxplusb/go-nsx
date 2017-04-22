package go_nsx

type Algorithm struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Description struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IpAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type MaxConn struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Member struct {
	IpAddress   *IpAddress   `xml:" ipAddress,omitempty" json:"ipAddress,omitempty"`
	MaxConn     *MaxConn     `xml:" maxConn,omitempty" json:"maxConn,omitempty"`
	MinConn     *MinConn     `xml:" minConn,omitempty" json:"minConn,omitempty"`
	MonitorPort *MonitorPort `xml:" monitorPort,omitempty" json:"monitorPort,omitempty"`
	Name        *Name        `xml:" name,omitempty" json:"name,omitempty"`
	Port        *Port        `xml:" port,omitempty" json:"port,omitempty"`
	Weight      *Weight      `xml:" weight,omitempty" json:"weight,omitempty"`
}

type MinConn struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type MonitorId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type MonitorPort struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Pool struct {
	Algorithm   *Algorithm   `xml:" algorithm,omitempty" json:"algorithm,omitempty"`
	Description *Description `xml:" description,omitempty" json:"description,omitempty"`
	Member      *Member      `xml:" member,omitempty" json:"member,omitempty"`
	MonitorId   *MonitorId   `xml:" monitorId,omitempty" json:"monitorId,omitempty"`
	Name        *Name        `xml:" name,omitempty" json:"name,omitempty"`
	Transparent *Transparent `xml:" transparent,omitempty" json:"transparent,omitempty"`
}

type Port struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	Pool *Pool `xml:" pool,omitempty" json:"pool,omitempty"`
}

type Transparent struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Weight struct {
	Text string `xml:",chardata" json:",omitempty"`
}
