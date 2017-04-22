package go_nsx

type FeatureId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type NwFabricFeatureConfig struct {
	FeatureId      *FeatureId      `xml:" featureId,omitempty" json:"featureId,omitempty"`
	ResourceConfig *ResourceConfig `xml:" resourceConfig,omitempty" json:"resourceConfig,omitempty"`
}

type ResourceConfig struct {
	ResourceId *ResourceId `xml:" resourceId,omitempty" json:"resourceId,omitempty"`
}

type ResourceId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	NwFabricFeatureConfig *NwFabricFeatureConfig `xml:" nwFabricFeatureConfig,omitempty" json:"nwFabricFeatureConfig,omitempty"`
}
