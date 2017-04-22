package go_nsx

type AllowLocalIPs struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Description struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type EnforcementPoint struct {
	Id   *Id   `xml:" id,omitempty" json:"id,omitempty"`
	Name *Name `xml:" name,omitempty" json:"name,omitempty"`
	Type *Type `xml:" type,omitempty" json:"type,omitempty"`
}

type Id struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type OperationMode struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	SpoofguardPolicy *SpoofguardPolicy `xml:" spoofguardPolicy,omitempty" json:"spoofguardPolicy,omitempty"`
}

type SpoofguardPolicy struct {
	AllowLocalIPs    *AllowLocalIPs    `xml:" allowLocalIPs,omitempty" json:"allowLocalIPs,omitempty"`
	Description      *Description      `xml:" description,omitempty" json:"description,omitempty"`
	EnforcementPoint *EnforcementPoint `xml:" enforcementPoint,omitempty" json:"enforcementPoint,omitempty"`
	Name             *Name             `xml:" name,omitempty" json:"name,omitempty"`
	OperationMode    *OperationMode    `xml:" operationMode,omitempty" json:"operationMode,omitempty"`
}

type Type struct {
	Text string `xml:",chardata" json:",omitempty"`
}
