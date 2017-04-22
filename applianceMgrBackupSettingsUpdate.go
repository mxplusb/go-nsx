package go_nsx

type BackupDirectory struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type BackupFrequency struct {
	DayOfWeek    *DayOfWeek    `xml:" dayOfWeek,omitempty" json:"dayOfWeek,omitempty"`
	Frequency    *Frequency    `xml:" frequency,omitempty" json:"frequency,omitempty"`
	HourOfDay    *HourOfDay    `xml:" hourOfDay,omitempty" json:"hourOfDay,omitempty"`
	MinuteOfHour *MinuteOfHour `xml:" minuteOfHour,omitempty" json:"minuteOfHour,omitempty"`
}

type BackupRestoreSettings struct {
	BackupFrequency *BackupFrequency `xml:" backupFrequency,omitempty" json:"backupFrequency,omitempty"`
	ExcludeTables   *ExcludeTables   `xml:" excludeTables,omitempty" json:"excludeTables,omitempty"`
	FtpSettings     *FtpSettings     `xml:" ftpSettings,omitempty" json:"ftpSettings,omitempty"`
}

type DayOfWeek struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ExcludeTable struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ExcludeTables struct {
	ExcludeTable *ExcludeTable `xml:" excludeTable,omitempty" json:"excludeTable,omitempty"`
}

type FilenamePrefix struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Frequency struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type FtpSettings struct {
	BackupDirectory   *BackupDirectory   `xml:" backupDirectory,omitempty" json:"backupDirectory,omitempty"`
	FilenamePrefix    *FilenamePrefix    `xml:" filenamePrefix,omitempty" json:"filenamePrefix,omitempty"`
	HostNameIPAddress *HostNameIPAddress `xml:" hostNameIPAddress,omitempty" json:"hostNameIPAddress,omitempty"`
	PassPhrase        *PassPhrase        `xml:" passPhrase,omitempty" json:"passPhrase,omitempty"`
	PassiveMode       *PassiveMode       `xml:" passiveMode,omitempty" json:"passiveMode,omitempty"`
	Password          *Password          `xml:" password,omitempty" json:"password,omitempty"`
	Port              *Port              `xml:" port,omitempty" json:"port,omitempty"`
	TransferProtocol  *TransferProtocol  `xml:" transferProtocol,omitempty" json:"transferProtocol,omitempty"`
	UseEPRT           *UseEPRT           `xml:" useEPRT,omitempty" json:"useEPRT,omitempty"`
	UseEPSV           *UseEPSV           `xml:" useEPSV,omitempty" json:"useEPSV,omitempty"`
	UserName          *UserName          `xml:" userName,omitempty" json:"userName,omitempty"`
}

type HostNameIPAddress struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type HourOfDay struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type MinuteOfHour struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PassPhrase struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PassiveMode struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Password struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Port struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	BackupRestoreSettings *BackupRestoreSettings `xml:" backupRestoreSettings,omitempty" json:"backupRestoreSettings,omitempty"`
}

type TransferProtocol struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type UseEPRT struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type UseEPSV struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type UserName struct {
	Text string `xml:",chardata" json:",omitempty"`
}
