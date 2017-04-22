package go_nsx

type Action struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type AddressGroup struct {
	PrimaryAddress     *PrimaryAddress     `xml:" primaryAddress,omitempty" json:"primaryAddress,omitempty"`
	SecondaryAddresses *SecondaryAddresses `xml:" secondaryAddresses,omitempty" json:"secondaryAddresses,omitempty"`
	SubnetMask         *SubnetMask         `xml:" subnetMask,omitempty" json:"subnetMask,omitempty"`
}

type AddressGroups struct {
	AddressGroup *AddressGroup `xml:" addressGroup,omitempty" json:"addressGroup,omitempty"`
}

type Appliance struct {
	DatastoreId    *DatastoreId    `xml:" datastoreId,omitempty" json:"datastoreId,omitempty"`
	HaAdminState   *HaAdminState   `xml:" haAdminState,omitempty" json:"haAdminState,omitempty"`
	ResourcePoolId *ResourcePoolId `xml:" resourcePoolId,omitempty" json:"resourcePoolId,omitempty"`
	VmFolderId     *VmFolderId     `xml:" vmFolderId,omitempty" json:"vmFolderId,omitempty"`
}

type ApplianceSize struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Appliances struct {
	Appliance     *Appliance     `xml:" appliance,omitempty" json:"appliance,omitempty"`
	ApplianceSize *ApplianceSize `xml:" applianceSize,omitempty" json:"applianceSize,omitempty"`
}

type Application struct {
	ApplicationId *ApplicationId `xml:" applicationId,omitempty" json:"applicationId,omitempty"`
	Service       *Service       `xml:" service,omitempty" json:"service,omitempty"`
}

type ApplicationId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type AuthenticationMode struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type AutoConfiguration struct {
	Enabled      *Enabled      `xml:" enabled,omitempty" json:"enabled,omitempty"`
	RulePriority *RulePriority `xml:" rulePriority,omitempty" json:"rulePriority,omitempty"`
}

type AutoConfigureDNS struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type AverageBandwidth struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type BindingId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type BurstSize struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type CaCertificates struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type CliSettings struct {
	RemoteAccess *RemoteAccess `xml:" remoteAccess,omitempty" json:"remoteAccess,omitempty"`
	UserName     *UserName     `xml:" userName,omitempty" json:"userName,omitempty"`
}

type CrlCertificates struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DatacenterMoid struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DatastoreId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DeclareDeadTime struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DefaultGateway struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DefaultPolicy struct {
	Action         *Action         `xml:" action,omitempty" json:"action,omitempty"`
	LoggingEnabled *LoggingEnabled `xml:" loggingEnabled,omitempty" json:"loggingEnabled,omitempty"`
}

type DefaultRoute struct {
	Description    *Description    `xml:" description,omitempty" json:"description,omitempty"`
	GatewayAddress *GatewayAddress `xml:" gatewayAddress,omitempty" json:"gatewayAddress,omitempty"`
	Vnic           *Vnic           `xml:" vnic,omitempty" json:"vnic,omitempty"`
}

type Description struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Destination struct {
	Exclude          *Exclude          `xml:" exclude,omitempty" json:"exclude,omitempty"`
	GroupingObjectId *GroupingObjectId `xml:" groupingObjectId,omitempty" json:"groupingObjectId,omitempty"`
	IpAddress        *IpAddress        `xml:" ipAddress,omitempty" json:"ipAddress,omitempty"`
	VnicGroupId      *VnicGroupId      `xml:" vnicGroupId,omitempty" json:"vnicGroupId,omitempty"`
}

type DhGroup struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Dhcp struct {
	Enabled        *Enabled        `xml:" enabled,omitempty" json:"enabled,omitempty"`
	IpPools        *IpPools        `xml:" ipPools,omitempty" json:"ipPools,omitempty"`
	Logging        *Logging        `xml:" logging,omitempty" json:"logging,omitempty"`
	StaticBindings *StaticBindings `xml:" staticBindings,omitempty" json:"staticBindings,omitempty"`
}

type DnatMatchSourceAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DnatMatchSourcePort struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Edge struct {
	Appliances        *Appliances        `xml:" appliances,omitempty" json:"appliances,omitempty"`
	AutoConfiguration *AutoConfiguration `xml:" autoConfiguration,omitempty" json:"autoConfiguration,omitempty"`
	CliSettings       *CliSettings       `xml:" cliSettings,omitempty" json:"cliSettings,omitempty"`
	DatacenterMoid    *DatacenterMoid    `xml:" datacenterMoid,omitempty" json:"datacenterMoid,omitempty"`
	Description       *Description       `xml:" description,omitempty" json:"description,omitempty"`
	EnableAesni       *EnableAesni       `xml:" enableAesni,omitempty" json:"enableAesni,omitempty"`
	EnableFips        *EnableFips        `xml:" enableFips,omitempty" json:"enableFips,omitempty"`
	Features          *Features          `xml:" features,omitempty" json:"features,omitempty"`
	Fqdn              *Fqdn              `xml:" fqdn,omitempty" json:"fqdn,omitempty"`
	Id                *Id                `xml:" id,omitempty" json:"id,omitempty"`
	Name              *Name              `xml:" name,omitempty" json:"name,omitempty"`
	Vnics             *Vnics             `xml:" vnics,omitempty" json:"vnics,omitempty"`
	VseLogLevel       *VseLogLevel       `xml:" vseLogLevel,omitempty" json:"vseLogLevel,omitempty"`
}

type Enable struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type EnableAesni struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type EnableFips struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type EnablePfs struct {
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

type EncryptionAlgorithm struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Exclude struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Features struct {
	Dhcp             *Dhcp             `xml:" dhcp,omitempty" json:"dhcp,omitempty"`
	Firewall         *Firewall         `xml:" firewall,omitempty" json:"firewall,omitempty"`
	HighAvailability *HighAvailability `xml:" highAvailability,omitempty" json:"highAvailability,omitempty"`
	Ipsec            *Ipsec            `xml:" ipsec,omitempty" json:"ipsec,omitempty"`
	Nat              *Nat              `xml:" nat,omitempty" json:"nat,omitempty"`
	Routing          *Routing          `xml:" routing,omitempty" json:"routing,omitempty"`
	Syslog           *Syslog           `xml:" syslog,omitempty" json:"syslog,omitempty"`
}

type Firewall struct {
	DefaultPolicy *DefaultPolicy `xml:" defaultPolicy,omitempty" json:"defaultPolicy,omitempty"`
	FirewallRules *FirewallRules `xml:" firewallRules,omitempty" json:"firewallRules,omitempty"`
}

type FirewallRule struct {
	Action          *Action          `xml:" action,omitempty" json:"action,omitempty"`
	Application     *Application     `xml:" application,omitempty" json:"application,omitempty"`
	Destination     *Destination     `xml:" destination,omitempty" json:"destination,omitempty"`
	Enabled         *Enabled         `xml:" enabled,omitempty" json:"enabled,omitempty"`
	Id              *Id              `xml:" id,omitempty" json:"id,omitempty"`
	LoggingEnabled  *LoggingEnabled  `xml:" loggingEnabled,omitempty" json:"loggingEnabled,omitempty"`
	MatchTranslated *MatchTranslated `xml:" matchTranslated,omitempty" json:"matchTranslated,omitempty"`
	Name            *Name            `xml:" name,omitempty" json:"name,omitempty"`
	RuleTag         *RuleTag         `xml:" ruleTag,omitempty" json:"ruleTag,omitempty"`
	RuleType        *RuleType        `xml:" ruleType,omitempty" json:"ruleType,omitempty"`
	Source          *Source          `xml:" source,omitempty" json:"source,omitempty"`
}

type FirewallRules struct {
	FirewallRule *FirewallRule `xml:" firewallRule,omitempty" json:"firewallRule,omitempty"`
}

type Fqdn struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type GatewayAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Global struct {
	CaCertificates  *CaCertificates  `xml:" caCertificates,omitempty" json:"caCertificates,omitempty"`
	CrlCertificates *CrlCertificates `xml:" crlCertificates,omitempty" json:"crlCertificates,omitempty"`
}

type GroupingObjectId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type HaAdminState struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type HighAvailability struct {
	DeclareDeadTime *DeclareDeadTime `xml:" declareDeadTime,omitempty" json:"declareDeadTime,omitempty"`
	Enabled         *Enabled         `xml:" enabled,omitempty" json:"enabled,omitempty"`
	Logging         *Logging         `xml:" logging,omitempty" json:"logging,omitempty"`
}

type Hostname struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Id struct {
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

type IpPool struct {
	AutoConfigureDNS *AutoConfigureDNS `xml:" autoConfigureDNS,omitempty" json:"autoConfigureDNS,omitempty"`
	DefaultGateway   *DefaultGateway   `xml:" defaultGateway,omitempty" json:"defaultGateway,omitempty"`
	IpRange          *IpRange          `xml:" ipRange,omitempty" json:"ipRange,omitempty"`
	LeaseTime        *LeaseTime        `xml:" leaseTime,omitempty" json:"leaseTime,omitempty"`
	PoolId           *PoolId           `xml:" poolId,omitempty" json:"poolId,omitempty"`
}

type IpPools struct {
	IpPool *IpPool `xml:" ipPool,omitempty" json:"ipPool,omitempty"`
}

type IpRange struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Ipsec struct {
	Enabled *Enabled `xml:" enabled,omitempty" json:"enabled,omitempty"`
	Global  *Global  `xml:" global,omitempty" json:"global,omitempty"`
	Logging *Logging `xml:" logging,omitempty" json:"logging,omitempty"`
	Sites   *Sites   `xml:" sites,omitempty" json:"sites,omitempty"`
}

type IsConnected struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type LeaseTime struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type LocalId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type LocalIp struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type LocalSubnets struct {
	Subnet *Subnet `xml:" subnet,omitempty" json:"subnet,omitempty"`
}

type LogLevel struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Logging struct {
	Enable   *Enable   `xml:" enable,omitempty" json:"enable,omitempty"`
	LogLevel *LogLevel `xml:" logLevel,omitempty" json:"logLevel,omitempty"`
}

type LoggingEnabled struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type MatchTranslated struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Mtu struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Nat struct {
	NatRules *NatRules `xml:" natRules,omitempty" json:"natRules,omitempty"`
}

type NatRule struct {
	Action                      *Action                      `xml:" action,omitempty" json:"action,omitempty"`
	DnatMatchSourceAddress      *DnatMatchSourceAddress      `xml:" dnatMatchSourceAddress,omitempty" json:"dnatMatchSourceAddress,omitempty"`
	DnatMatchSourcePort         *DnatMatchSourcePort         `xml:" dnatMatchSourcePort,omitempty" json:"dnatMatchSourcePort,omitempty"`
	Enabled                     *Enabled                     `xml:" enabled,omitempty" json:"enabled,omitempty"`
	LoggingEnabled              *LoggingEnabled              `xml:" loggingEnabled,omitempty" json:"loggingEnabled,omitempty"`
	OriginalAddress             *OriginalAddress             `xml:" originalAddress,omitempty" json:"originalAddress,omitempty"`
	OriginalPort                *OriginalPort                `xml:" originalPort,omitempty" json:"originalPort,omitempty"`
	Protocol                    *Protocol                    `xml:" protocol,omitempty" json:"protocol,omitempty"`
	RuleId                      *RuleId                      `xml:" ruleId,omitempty" json:"ruleId,omitempty"`
	RuleTag                     *RuleTag                     `xml:" ruleTag,omitempty" json:"ruleTag,omitempty"`
	RuleType                    *RuleType                    `xml:" ruleType,omitempty" json:"ruleType,omitempty"`
	SnatMatchDestinationAddress *SnatMatchDestinationAddress `xml:" snatMatchDestinationAddress,omitempty" json:"snatMatchDestinationAddress,omitempty"`
	SnatMatchDestinationPort    *SnatMatchDestinationPort    `xml:" snatMatchDestinationPort,omitempty" json:"snatMatchDestinationPort,omitempty"`
	TranslatedAddress           *TranslatedAddress           `xml:" translatedAddress,omitempty" json:"translatedAddress,omitempty"`
	TranslatedPort              *TranslatedPort              `xml:" translatedPort,omitempty" json:"translatedPort,omitempty"`
	Vnic                        *Vnic                        `xml:" vnic,omitempty" json:"vnic,omitempty"`
}

type NatRules struct {
	NatRule *NatRule `xml:" natRule,omitempty" json:"natRule,omitempty"`
}

type Network struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type NextHop struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type OriginalAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type OriginalPort struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Ospf struct {
	Enabled *Enabled `xml:" enabled,omitempty" json:"enabled,omitempty"`
}

type OutShapingPolicy struct {
	AverageBandwidth *AverageBandwidth `xml:" averageBandwidth,omitempty" json:"averageBandwidth,omitempty"`
	BurstSize        *BurstSize        `xml:" burstSize,omitempty" json:"burstSize,omitempty"`
	Enabled          *Enabled          `xml:" enabled,omitempty" json:"enabled,omitempty"`
	Inherited        *Inherited        `xml:" inherited,omitempty" json:"inherited,omitempty"`
	PeakBandwidth    *PeakBandwidth    `xml:" peakBandwidth,omitempty" json:"peakBandwidth,omitempty"`
}

type PeakBandwidth struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PeerId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PeerSubnets struct {
	Subnet *Subnet `xml:" subnet,omitempty" json:"subnet,omitempty"`
}

type PoolId struct {
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

type Protocol struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Psk struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type RemoteAccess struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ResourcePoolId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	Edge *Edge `xml:" edge,omitempty" json:"edge,omitempty"`
}

type Route struct {
	Network *Network `xml:" network,omitempty" json:"network,omitempty"`
	NextHop *NextHop `xml:" nextHop,omitempty" json:"nextHop,omitempty"`
	Type    *Type    `xml:" type,omitempty" json:"type,omitempty"`
	Vnic    *Vnic    `xml:" vnic,omitempty" json:"vnic,omitempty"`
}

type Routing struct {
	Ospf          *Ospf          `xml:" ospf,omitempty" json:"ospf,omitempty"`
	StaticRouting *StaticRouting `xml:" staticRouting,omitempty" json:"staticRouting,omitempty"`
}

type RuleId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type RulePriority struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type RuleTag struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type RuleType struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type SecondaryAddresses struct {
	IpAddress *IpAddress `xml:" ipAddress,omitempty" json:"ipAddress,omitempty"`
}

type ServerAddresses struct {
	IpAddress *IpAddress `xml:" ipAddress,omitempty" json:"ipAddress,omitempty"`
}

type Service struct {
	Port       *Port       `xml:" port,omitempty" json:"port,omitempty"`
	Protocol   *Protocol   `xml:" protocol,omitempty" json:"protocol,omitempty"`
	SourcePort *SourcePort `xml:" sourcePort,omitempty" json:"sourcePort,omitempty"`
}

type Site struct {
	AuthenticationMode  *AuthenticationMode  `xml:" authenticationMode,omitempty" json:"authenticationMode,omitempty"`
	DhGroup             *DhGroup             `xml:" dhGroup,omitempty" json:"dhGroup,omitempty"`
	EnablePfs           *EnablePfs           `xml:" enablePfs,omitempty" json:"enablePfs,omitempty"`
	Enabled             *Enabled             `xml:" enabled,omitempty" json:"enabled,omitempty"`
	EncryptionAlgorithm *EncryptionAlgorithm `xml:" encryptionAlgorithm,omitempty" json:"encryptionAlgorithm,omitempty"`
	LocalId             *LocalId             `xml:" localId,omitempty" json:"localId,omitempty"`
	LocalIp             *LocalIp             `xml:" localIp,omitempty" json:"localIp,omitempty"`
	LocalSubnets        *LocalSubnets        `xml:" localSubnets,omitempty" json:"localSubnets,omitempty"`
	Mtu                 *Mtu                 `xml:" mtu,omitempty" json:"mtu,omitempty"`
	Name                *Name                `xml:" name,omitempty" json:"name,omitempty"`
	PeerId              *PeerId              `xml:" peerId,omitempty" json:"peerId,omitempty"`
	PeerSubnets         *PeerSubnets         `xml:" peerSubnets,omitempty" json:"peerSubnets,omitempty"`
	Psk                 *Psk                 `xml:" psk,omitempty" json:"psk,omitempty"`
}

type Sites struct {
	Site *Site `xml:" site,omitempty" json:"site,omitempty"`
}

type SnatMatchDestinationAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type SnatMatchDestinationPort struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Source struct {
	Exclude          *Exclude          `xml:" exclude,omitempty" json:"exclude,omitempty"`
	GroupingObjectId *GroupingObjectId `xml:" groupingObjectId,omitempty" json:"groupingObjectId,omitempty"`
	IpAddress        *IpAddress        `xml:" ipAddress,omitempty" json:"ipAddress,omitempty"`
	VnicGroupId      *VnicGroupId      `xml:" vnicGroupId,omitempty" json:"vnicGroupId,omitempty"`
}

type SourcePort struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type StaticBinding struct {
	AutoConfigureDNS *AutoConfigureDNS `xml:" autoConfigureDNS,omitempty" json:"autoConfigureDNS,omitempty"`
	BindingId        *BindingId        `xml:" bindingId,omitempty" json:"bindingId,omitempty"`
	DefaultGateway   *DefaultGateway   `xml:" defaultGateway,omitempty" json:"defaultGateway,omitempty"`
	Hostname         *Hostname         `xml:" hostname,omitempty" json:"hostname,omitempty"`
	IpAddress        *IpAddress        `xml:" ipAddress,omitempty" json:"ipAddress,omitempty"`
	LeaseTime        *LeaseTime        `xml:" leaseTime,omitempty" json:"leaseTime,omitempty"`
	VmId             *VmId             `xml:" vmId,omitempty" json:"vmId,omitempty"`
	VnicId           *VnicId           `xml:" vnicId,omitempty" json:"vnicId,omitempty"`
}

type StaticBindings struct {
	StaticBinding *StaticBinding `xml:" staticBinding,omitempty" json:"staticBinding,omitempty"`
}

type StaticRoutes struct {
	Route *Route `xml:" route,omitempty" json:"route,omitempty"`
}

type StaticRouting struct {
	DefaultRoute *DefaultRoute `xml:" defaultRoute,omitempty" json:"defaultRoute,omitempty"`
	StaticRoutes *StaticRoutes `xml:" staticRoutes,omitempty" json:"staticRoutes,omitempty"`
}

type Subnet struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type SubnetMask struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Syslog struct {
	Protocol        *Protocol        `xml:" protocol,omitempty" json:"protocol,omitempty"`
	ServerAddresses *ServerAddresses `xml:" serverAddresses,omitempty" json:"serverAddresses,omitempty"`
}

type TranslatedAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type TranslatedPort struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Type struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type UserName struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type VmFolderId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type VmId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Vnic struct {
	AddressGroups       *AddressGroups       `xml:" addressGroups,omitempty" json:"addressGroups,omitempty"`
	EnableProxyArp      *EnableProxyArp      `xml:" enableProxyArp,omitempty" json:"enableProxyArp,omitempty"`
	EnableSendRedirects *EnableSendRedirects `xml:" enableSendRedirects,omitempty" json:"enableSendRedirects,omitempty"`
	InShapingPolicy     *InShapingPolicy     `xml:" inShapingPolicy,omitempty" json:"inShapingPolicy,omitempty"`
	Index               *Index               `xml:" index,omitempty" json:"index,omitempty"`
	IsConnected         *IsConnected         `xml:" isConnected,omitempty" json:"isConnected,omitempty"`
	Mtu                 *Mtu                 `xml:" mtu,omitempty" json:"mtu,omitempty"`
	Name                *Name                `xml:" name,omitempty" json:"name,omitempty"`
	OutShapingPolicy    *OutShapingPolicy    `xml:" outShapingPolicy,omitempty" json:"outShapingPolicy,omitempty"`
	PortgroupId         *PortgroupId         `xml:" portgroupId,omitempty" json:"portgroupId,omitempty"`
	Type                *Type                `xml:" type,omitempty" json:"type,omitempty"`
	Text                string               `xml:",chardata" json:",omitempty"`
}

type VnicGroupId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type VnicId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Vnics struct {
	Vnic *Vnic `xml:" vnic,omitempty" json:"vnic,omitempty"`
}

type VseLogLevel struct {
	Text string `xml:",chardata" json:",omitempty"`
}
