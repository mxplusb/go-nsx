package go_nsx

type Appliance struct {
	CpuReservation    *CpuReservation    `xml:" cpuReservation,omitempty" json:"cpuReservation,omitempty"`
	CustomField       *CustomField       `xml:" customField,omitempty" json:"customField,omitempty"`
	DatastoreId       *DatastoreId       `xml:" datastoreId,omitempty" json:"datastoreId,omitempty"`
	HaAdminState      *HaAdminState      `xml:" haAdminState,omitempty" json:"haAdminState,omitempty"`
	HostId            *HostId            `xml:" hostId,omitempty" json:"hostId,omitempty"`
	MemoryReservation *MemoryReservation `xml:" memoryReservation,omitempty" json:"memoryReservation,omitempty"`
	ResourcePoolId    *ResourcePoolId    `xml:" resourcePoolId,omitempty" json:"resourcePoolId,omitempty"`
	VmFolderId        *VmFolderId        `xml:" vmFolderId,omitempty" json:"vmFolderId,omitempty"`
}

type CpuReservation struct {
	Limit       *Limit       `xml:" limit,omitempty" json:"limit,omitempty"`
	Reservation *Reservation `xml:" reservation,omitempty" json:"reservation,omitempty"`
	Shares      *Shares      `xml:" shares,omitempty" json:"shares,omitempty"`
}

type CustomField struct {
	Key   *Key   `xml:" key,omitempty" json:"key,omitempty"`
	Value *Value `xml:" value,omitempty" json:"value,omitempty"`
}

type DatastoreId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type HaAdminState struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type HostId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Key struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Limit struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type MemoryReservation struct {
	Limit       *Limit       `xml:" limit,omitempty" json:"limit,omitempty"`
	Reservation *Reservation `xml:" reservation,omitempty" json:"reservation,omitempty"`
	Shares      *Shares      `xml:" shares,omitempty" json:"shares,omitempty"`
}

type Reservation struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ResourcePoolId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	Appliance *Appliance `xml:" utilities,omitempty" json:"utilities,omitempty"`
}

type Shares struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Value struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type VmFolderId struct {
	Text string `xml:",chardata" json:",omitempty"`
}
