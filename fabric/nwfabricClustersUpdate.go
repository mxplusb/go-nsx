package fabric

type LocaleId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type NwFabricClusterConfig struct {
	LocaleId *LocaleId `xml:" localeId,omitempty" json:"localeId,omitempty"`
}

type Root struct {
	NwFabricClusterConfig *NwFabricClusterConfig `xml:" nwFabricClusterConfig,omitempty" json:"nwFabricClusterConfig,omitempty"`
}
