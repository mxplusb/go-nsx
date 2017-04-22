package security

type ClusterDeploymentConfig struct {
	ClusterId *ClusterId `xml:" clusterId,omitempty" json:"clusterId,omitempty"`
	Datastore *Datastore `xml:" datastore,omitempty" json:"datastore,omitempty"`
	Services  *Services  `xml:" services,omitempty" json:"services,omitempty"`
}

type ClusterDeploymentConfigs struct {
	ClusterDeploymentConfig *ClusterDeploymentConfig `xml:" clusterDeploymentConfig,omitempty" json:"clusterDeploymentConfig,omitempty"`
}

type ClusterId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Datastore struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DvPortGroup struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IpPool struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	ClusterDeploymentConfigs *ClusterDeploymentConfigs `xml:" clusterDeploymentConfigs,omitempty" json:"clusterDeploymentConfigs,omitempty"`
}

type ServiceDeploymentConfig struct {
	DvPortGroup *DvPortGroup `xml:" dvPortGroup,omitempty" json:"dvPortGroup,omitempty"`
	IpPool      *IpPool      `xml:" ipPool,omitempty" json:"ipPool,omitempty"`
	ServiceId   *ServiceId   `xml:" serviceId,omitempty" json:"serviceId,omitempty"`
}

type ServiceId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Services struct {
	ServiceDeploymentConfig *ServiceDeploymentConfig `xml:" serviceDeploymentConfig,omitempty" json:"serviceDeploymentConfig,omitempty"`
}
