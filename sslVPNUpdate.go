package go_nsx

type Description struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Enabled struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type LogonLogoffScript struct {
	Description       *Description       `xml:" description,omitempty" json:"description,omitempty"`
	Enabled           *Enabled           `xml:" enabled,omitempty" json:"enabled,omitempty"`
	LogonLogoffScript *LogonLogoffScript `xml:" logonLogoffScript,omitempty" json:"logonLogoffScript,omitempty"`
	ObjectId          *ObjectId          `xml:" objectId,omitempty" json:"objectId,omitempty"`
	ScriptFileId      *ScriptFileId      `xml:" scriptFileId,omitempty" json:"scriptFileId,omitempty"`
	Type              *Type              `xml:" type,omitempty" json:"type,omitempty"`
}

type ObjectId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	LogonLogoffScript *LogonLogoffScript `xml:" logonLogoffScript,omitempty" json:"logonLogoffScript,omitempty"`
}

type ScriptFileId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Type struct {
	Text string `xml:",chardata" json:",omitempty"`
}
