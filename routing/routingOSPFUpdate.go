package routing

type Action struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type AreaId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Authentication struct {
	Type  *Type  `xml:" type,omitempty" json:"type,omitempty"`
	Value *Value `xml:" value,omitempty" json:"value,omitempty"`
}

type Bgp struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Connected struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Cost struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DeadInterval struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Enabled struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type From struct {
	Bgp       *Bgp       `xml:" bgp,omitempty" json:"bgp,omitempty"`
	Connected *Connected `xml:" connected,omitempty" json:"connected,omitempty"`
	Ospf      *Ospf      `xml:" ospf,omitempty" json:"ospf,omitempty"`
	Static    *Static    `xml:" static,omitempty" json:"static,omitempty"`
}

type HelloInterval struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Ospf struct {
	Enabled        *Enabled        `xml:" enabled,omitempty" json:"enabled,omitempty"`
	OspfAreas      *OspfAreas      `xml:" ospfAreas,omitempty" json:"ospfAreas,omitempty"`
	OspfInterfaces *OspfInterfaces `xml:" ospfInterfaces,omitempty" json:"ospfInterfaces,omitempty"`
	Redistribution *Redistribution `xml:" redistribution,omitempty" json:"redistribution,omitempty"`
	Text              string             `xml:",chardata" json:",omitempty"`
}

type OspfArea struct {
	AreaId         *AreaId         `xml:" areaId,omitempty" json:"areaId,omitempty"`
	Authentication *Authentication `xml:" authentication,omitempty" json:"authentication,omitempty"`
	Type           *Type           `xml:" type,omitempty" json:"type,omitempty"`
}

type OspfAreas struct {
	OspfArea *OspfArea `xml:" ospfArea,omitempty" json:"ospfArea,omitempty"`
}

type OspfInterface struct {
	AreaId        *AreaId        `xml:" areaId,omitempty" json:"areaId,omitempty"`
	Cost          *Cost          `xml:" cost,omitempty" json:"cost,omitempty"`
	DeadInterval  *DeadInterval  `xml:" deadInterval,omitempty" json:"deadInterval,omitempty"`
	HelloInterval *HelloInterval `xml:" helloInterval,omitempty" json:"helloInterval,omitempty"`
	Priority      *Priority      `xml:" priority,omitempty" json:"priority,omitempty"`
	Vnic          *Vnic          `xml:" vnic,omitempty" json:"vnic,omitempty"`
}

type OspfInterfaces struct {
	OspfInterface *OspfInterface `xml:" ospfInterface,omitempty" json:"ospfInterface,omitempty"`
}

type PrefixName struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Priority struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Redistribution struct {
	Enabled *Enabled `xml:" enabled,omitempty" json:"enabled,omitempty"`
	Rules   *Rules   `xml:" rules,omitempty" json:"rules,omitempty"`
}

type Root struct {
	Ospf *Ospf `xml:" ospf,omitempty" json:"ospf,omitempty"`
}

type Rule struct {
	Action     *Action     `xml:" action,omitempty" json:"action,omitempty"`
	From       *From       `xml:" from,omitempty" json:"from,omitempty"`
	PrefixName *PrefixName `xml:" prefixName,omitempty" json:"prefixName,omitempty"`
}

type Rules struct {
	Rule *Rule `xml:" rule,omitempty" json:"rule,omitempty"`
}

type Static struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Type struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Value struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Vnic struct {
	Text string `xml:",chardata" json:",omitempty"`
}
