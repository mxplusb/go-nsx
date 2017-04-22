package go_nsx

type ConnectionsPerSecond struct {
	Value *Value `xml:" value,omitempty" json:"value,omitempty"`
}

type Cpu struct {
	PercentValue *PercentValue `xml:" percentValue,omitempty" json:"percentValue,omitempty"`
}

type EventThresholds struct {
	ConnectionsPerSecond *ConnectionsPerSecond `xml:" connectionsPerSecond,omitempty" json:"connectionsPerSecond,omitempty"`
	Cpu                  *Cpu                  `xml:" cpu,omitempty" json:"cpu,omitempty"`
	Memory               *Memory               `xml:" memory,omitempty" json:"memory,omitempty"`
}

type Memory struct {
	PercentValue *PercentValue `xml:" percentValue,omitempty" json:"percentValue,omitempty"`
}

type PercentValue struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	EventThresholds *EventThresholds `xml:" eventThresholds,omitempty" json:"eventThresholds,omitempty"`
}

type Value struct {
	Text string `xml:",chardata" json:",omitempty"`
}
