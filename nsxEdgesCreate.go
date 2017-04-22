package go_nsx

type AddressGroup struct {
	PrimaryAddress     *PrimaryAddress     `xml:" primaryAddress,omitempty" json:"primaryAddress,omitempty"`
	SecondaryAddresses *SecondaryAddresses `xml:" secondaryAddresses,omitempty" json:"secondaryAddresses,omitempty"`
	SubnetMask         *SubnetMask         `xml:" subnetMask,omitempty" json:"subnetMask,omitempty"`
}

type AddressGroups struct {
	AddressGroup *AddressGroup `xml:" addressGroup,omitempty" json:"addressGroup,omitempty"`
}

type Appliance struct {
	CpuReservation    *CpuReservation    `xml:" cpuReservation,omitempty" json:"cpuReservation,omitempty"`
	CustomField       *CustomField       `xml:" customField,omitempty" json:"customField,omitempty"`
	DatastoreId       *DatastoreId       `xml:" datastoreId,omitempty" json:"datastoreId,omitempty"`
	HostId            *HostId            `xml:" hostId,omitempty" json:"hostId,omitempty"`
	MemoryReservation *MemoryReservation `xml:" memoryReservation,omitempty" json:"memoryReservation,omitempty"`
	ResourcePoolId    *ResourcePoolId    `xml:" resourcePoolId,omitempty" json:"resourcePoolId,omitempty"`
	VmFolderId        *VmFolderId        `xml:" vmFolderId,omitempty" json:"vmFolderId,omitempty"`
}

type ApplianceSize struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Appliances struct {
	Appliance      *Appliance      `xml:" appliance,omitempty" json:"appliance,omitempty"`
	ApplianceSize  *ApplianceSize  `xml:" applianceSize,omitempty" json:"applianceSize,omitempty"`
	EnableCoreDump *EnableCoreDump `xml:" enableCoreDump,omitempty" json:"enableCoreDump,omitempty"`
}

type AutoConfiguration struct {
	Enabled      *Enabled      `xml:" enabled,omitempty" json:"enabled,omitempty"`
	RulePriority *RulePriority `xml:" rulePriority,omitempty" json:"rulePriority,omitempty"`
}

type AverageBandwidth struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type BurstSize struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type CliSettings struct {
	Password     *Password     `xml:" password,omitempty" json:"password,omitempty"`
	RemoteAccess *RemoteAccess `xml:" remoteAccess,omitempty" json:"remoteAccess,omitempty"`
	UserName     *UserName     `xml:" userName,omitempty" json:"userName,omitempty"`
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

type DatacenterMoid struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DatastoreId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Description struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DnsClient struct {
	DomainName   *DomainName   `xml:" domainName,omitempty" json:"domainName,omitempty"`
	PrimaryDns   *PrimaryDns   `xml:" primaryDns,omitempty" json:"primaryDns,omitempty"`
	SecondaryDns *SecondaryDns `xml:" secondaryDns,omitempty" json:"secondaryDns,omitempty"`
}

type DomainName struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Edge struct {
	Appliances        *Appliances        `xml:" appliances,omitempty" json:"appliances,omitempty"`
	AutoConfiguration *AutoConfiguration `xml:" autoConfiguration,omitempty" json:"autoConfiguration,omitempty"`
	CliSettings       *CliSettings       `xml:" cliSettings,omitempty" json:"cliSettings,omitempty"`
	DatacenterMoid    *DatacenterMoid    `xml:" datacenterMoid,omitempty" json:"datacenterMoid,omitempty"`
	Description       *Description       `xml:" description,omitempty" json:"description,omitempty"`
	DnsClient         *DnsClient         `xml:" dnsClient,omitempty" json:"dnsClient,omitempty"`
	EnableAesni       *EnableAesni       `xml:" enableAesni,omitempty" json:"enableAesni,omitempty"`
	EnableFips        *EnableFips        `xml:" enableFips,omitempty" json:"enableFips,omitempty"`
	Fqdn              *Fqdn              `xml:" fqdn,omitempty" json:"fqdn,omitempty"`
	Name              *Name              `xml:" name,omitempty" json:"name,omitempty"`
	QueryDaemon       *QueryDaemon       `xml:" queryDaemon,omitempty" json:"queryDaemon,omitempty"`
	Tenant            *Tenant            `xml:" tenant,omitempty" json:"tenant,omitempty"`
	Type              *Type              `xml:" type,omitempty" json:"type,omitempty"`
	Vnics             *Vnics             `xml:" vnics,omitempty" json:"vnics,omitempty"`
	VseLogLevel       *VseLogLevel       `xml:" vseLogLevel,omitempty" json:"vseLogLevel,omitempty"`
}

type EdgeVmHaIndex struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type EnableAesni struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type EnableCoreDump struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type EnableFips struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type EnableProxyArp struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type EnableSendRedirects struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Enabled struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type FenceParameter struct {
	Key   *Key   `xml:" key,omitempty" json:"key,omitempty"`
	Value *Value `xml:" value,omitempty" json:"value,omitempty"`
}

type Fqdn struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type HostId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type InShapingPolicy struct {
	AverageBandwidth *AverageBandwidth `xml:" averageBandwidth,omitempty" json:"averageBandwidth,omitempty"`
	BurstSize        *BurstSize        `xml:" burstSize,omitempty" json:"burstSize,omitempty"`
	Enabled          *Enabled          `xml:" enabled,omitempty" json:"enabled,omitempty"`
	Inherited        *Inherited        `xml:" inherited,omitempty" json:"inherited,omitempty"`
	PeakBandwidth    *PeakBandwidth    `xml:" peakBandwidth,omitempty" json:"peakBandwidth,omitempty"`
}

type Index struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Inherited struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IpAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IsConnected struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Key struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Limit struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type MacAddress struct {
	EdgeVmHaIndex *EdgeVmHaIndex `xml:" edgeVmHaIndex,omitempty" json:"edgeVmHaIndex,omitempty"`
	Value         *Value         `xml:" value,omitempty" json:"value,omitempty"`
}

type MemoryReservation struct {
	Limit       *Limit       `xml:" limit,omitempty" json:"limit,omitempty"`
	Reservation *Reservation `xml:" reservation,omitempty" json:"reservation,omitempty"`
	Shares      *Shares      `xml:" shares,omitempty" json:"shares,omitempty"`
}

type Mtu struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type OutShapingPolicy struct {
	AverageBandwidth *AverageBandwidth `xml:" averageBandwidth,omitempty" json:"averageBandwidth,omitempty"`
	BurstSize        *BurstSize        `xml:" burstSize,omitempty" json:"burstSize,omitempty"`
	Enabled          *Enabled          `xml:" enabled,omitempty" json:"enabled,omitempty"`
	Inherited        *Inherited        `xml:" inherited,omitempty" json:"inherited,omitempty"`
	PeakBandwidth    *PeakBandwidth    `xml:" peakBandwidth,omitempty" json:"peakBandwidth,omitempty"`
}

type Password struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PeakBandwidth struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Port struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PortgroupId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PrimaryAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PrimaryDns struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type QueryDaemon struct {
	Enabled *Enabled `xml:" enabled,omitempty" json:"enabled,omitempty"`
	Port    *Port    `xml:" port,omitempty" json:"port,omitempty"`
}

type RemoteAccess struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Reservation struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ResourcePoolId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	Edge *Edge `xml:" edge,omitempty" json:"edge,omitempty"`
}

type RulePriority struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type SecondaryAddresses struct {
	IpAddress *IpAddress `xml:" ipAddress,omitempty" json:"ipAddress,omitempty"`
}

type SecondaryDns struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Shares struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type SubnetMask struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Tenant struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Type struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type UserName struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Value struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type VmFolderId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Vnic struct {
	AddressGroups       *AddressGroups       `xml:" addressGroups,omitempty" json:"addressGroups,omitempty"`
	EnableProxyArp      *EnableProxyArp      `xml:" enableProxyArp,omitempty" json:"enableProxyArp,omitempty"`
	EnableSendRedirects *EnableSendRedirects `xml:" enableSendRedirects,omitempty" json:"enableSendRedirects,omitempty"`
	FenceParameter      *FenceParameter      `xml:" fenceParameter,omitempty" json:"fenceParameter,omitempty"`
	InShapingPolicy     *InShapingPolicy     `xml:" inShapingPolicy,omitempty" json:"inShapingPolicy,omitempty"`
	Index               *Index               `xml:" index,omitempty" json:"index,omitempty"`
	IsConnected         *IsConnected         `xml:" isConnected,omitempty" json:"isConnected,omitempty"`
	MacAddress          *MacAddress          `xml:" macAddress,omitempty" json:"macAddress,omitempty"`
	Mtu                 *Mtu                 `xml:" mtu,omitempty" json:"mtu,omitempty"`
	Name                *Name                `xml:" name,omitempty" json:"name,omitempty"`
	OutShapingPolicy    *OutShapingPolicy    `xml:" outShapingPolicy,omitempty" json:"outShapingPolicy,omitempty"`
	PortgroupId         *PortgroupId         `xml:" portgroupId,omitempty" json:"portgroupId,omitempty"`
	Type                *Type                `xml:" type,omitempty" json:"type,omitempty"`
}

type Vnics struct {
	Vnic *Vnic `xml:" vnic,omitempty" json:"vnic,omitempty"`
}

type VseLogLevel struct {
	Text string `xml:",chardata" json:",omitempty"`
}
