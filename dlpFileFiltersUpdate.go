package go_nsx

type FileFilters struct {
	Extensions         *Extensions         `xml:" extensions,omitempty" json:"extensions,omitempty"`
	ExtensionsIncluded *ExtensionsIncluded `xml:" extensionsIncluded,omitempty" json:"extensionsIncluded,omitempty"`
	LastModifiedAfter  *LastModifiedAfter  `xml:" lastModifiedAfter,omitempty" json:"lastModifiedAfter,omitempty"`
	ScanAllFiles       *ScanAllFiles       `xml:" scanAllFiles,omitempty" json:"scanAllFiles,omitempty"`
}

type Extensions struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ExtensionsIncluded struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type LastModifiedAfter struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	FileFilters *FileFilters `xml:" FileFilters,omitempty" json:"FileFilters,omitempty"`
}

type ScanAllFiles struct {
	Text string `xml:",chardata" json:",omitempty"`
}
