package go_nsx

type LocaleId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type NwFabricHostConfig struct {
	LocaleId *LocaleId `xml:" localeId,omitempty" json:"localeId,omitempty"`
}

type Root struct {
	NwFabricHostConfig *NwFabricHostConfig `xml:" nwFabricHostConfig,omitempty" json:"nwFabricHostConfig,omitempty"`
}
