package go_nsx

type Enabled struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IpDiscoveryConfig struct {
	Enabled *Enabled `xml:" enabled,omitempty" json:"enabled,omitempty"`
}

type MacLearningConfig struct {
	Enabled *Enabled `xml:" enabled,omitempty" json:"enabled,omitempty"`
}

type NetworkFeatureConfig struct {
	IpDiscoveryConfig *IpDiscoveryConfig `xml:" ipDiscoveryConfig,omitempty" json:"ipDiscoveryConfig,omitempty"`
	MacLearningConfig *MacLearningConfig `xml:" macLearningConfig,omitempty" json:"macLearningConfig,omitempty"`
}

type Root struct {
	NetworkFeatureConfig *NetworkFeatureConfig `xml:" networkFeatureConfig,omitempty" json:"networkFeatureConfig,omitempty"`
}
