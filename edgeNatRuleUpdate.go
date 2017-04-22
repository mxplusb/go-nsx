package go_nsx

type Action struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Description struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DnatMatchSourceAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DnatMatchSourcePort struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Enabled struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type LoggingEnabled struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type NatRule struct {
	Action                      *Action                      `xml:" action,omitempty" json:"action,omitempty"`
	Description                 *Description                 `xml:" description,omitempty" json:"description,omitempty"`
	DnatMatchSourceAddress      *DnatMatchSourceAddress      `xml:" dnatMatchSourceAddress,omitempty" json:"dnatMatchSourceAddress,omitempty"`
	DnatMatchSourcePort         *DnatMatchSourcePort         `xml:" dnatMatchSourcePort,omitempty" json:"dnatMatchSourcePort,omitempty"`
	Enabled                     *Enabled                     `xml:" enabled,omitempty" json:"enabled,omitempty"`
	LoggingEnabled              *LoggingEnabled              `xml:" loggingEnabled,omitempty" json:"loggingEnabled,omitempty"`
	OriginalAddress             *OriginalAddress             `xml:" originalAddress,omitempty" json:"originalAddress,omitempty"`
	OriginalPort                *OriginalPort                `xml:" originalPort,omitempty" json:"originalPort,omitempty"`
	Protocol                    *Protocol                    `xml:" protocol,omitempty" json:"protocol,omitempty"`
	SnatMatchDestinationAddress *SnatMatchDestinationAddress `xml:" snatMatchDestinationAddress,omitempty" json:"snatMatchDestinationAddress,omitempty"`
	SnatMatchDestinationPort    *SnatMatchDestinationPort    `xml:" snatMatchDestinationPort,omitempty" json:"snatMatchDestinationPort,omitempty"`
	TranslatedAddress           *TranslatedAddress           `xml:" translatedAddress,omitempty" json:"translatedAddress,omitempty"`
	TranslatedPort              *TranslatedPort              `xml:" translatedPort,omitempty" json:"translatedPort,omitempty"`
	Vnic                        *Vnic                        `xml:" vnic,omitempty" json:"vnic,omitempty"`
}

type OriginalAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type OriginalPort struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Protocol struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	NatRule *NatRule `xml:" natRule,omitempty" json:"natRule,omitempty"`
}

type SnatMatchDestinationAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type SnatMatchDestinationPort struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type TranslatedAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type TranslatedPort struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Vnic struct {
	Text string `xml:",chardata" json:",omitempty"`
}
