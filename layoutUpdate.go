package go_nsx

type BodyColor struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type CompanyName struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Layout struct {
	BodyColor           *BodyColor           `xml:" bodyColor,omitempty" json:"bodyColor,omitempty"`
	CompanyName         *CompanyName         `xml:" companyName,omitempty" json:"companyName,omitempty"`
	LogoBackgroundColor *LogoBackgroundColor `xml:" logoBackgroundColor,omitempty" json:"logoBackgroundColor,omitempty"`
	MenuBarColor        *MenuBarColor        `xml:" menuBarColor,omitempty" json:"menuBarColor,omitempty"`
	PortalTitle         *PortalTitle         `xml:" portalTitle,omitempty" json:"portalTitle,omitempty"`
	RowAlternativeColor *RowAlternativeColor `xml:" rowAlternativeColor,omitempty" json:"rowAlternativeColor,omitempty"`
	RowColor            *RowColor            `xml:" rowColor,omitempty" json:"rowColor,omitempty"`
	TitleColor          *TitleColor          `xml:" titleColor,omitempty" json:"titleColor,omitempty"`
	TopFrameColor       *TopFrameColor       `xml:" topFrameColor,omitempty" json:"topFrameColor,omitempty"`
}

type LogoBackgroundColor struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type MenuBarColor struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PortalTitle struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	Layout *Layout `xml:" layout,omitempty" json:"layout,omitempty"`
}

type RowAlternativeColor struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type RowColor struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type TitleColor struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type TopFrameColor struct {
	Text string `xml:",chardata" json:",omitempty"`
}
