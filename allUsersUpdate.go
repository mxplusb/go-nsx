package go_nsx

type AllowChangePassword struct {
	ChangePasswordOnNextLogin *ChangePasswordOnNextLogin `xml:" changePasswordOnNextLogin,omitempty" json:"changePasswordOnNextLogin,omitempty"`
}

type ChangePasswordOnNextLogin struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Description struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DisableUserAccount struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type FirstName struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type LastName struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Password struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PasswordNeverExpires struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	Users *Users `xml:" users,omitempty" json:"users,omitempty"`
}

type User struct {
	AllowChangePassword  *AllowChangePassword  `xml:" allowChangePassword,omitempty" json:"allowChangePassword,omitempty"`
	Description          *Description          `xml:" description,omitempty" json:"description,omitempty"`
	DisableUserAccount   *DisableUserAccount   `xml:" disableUserAccount,omitempty" json:"disableUserAccount,omitempty"`
	FirstName            *FirstName            `xml:" firstName,omitempty" json:"firstName,omitempty"`
	LastName             *LastName             `xml:" lastName,omitempty" json:"lastName,omitempty"`
	Password             *Password             `xml:" password,omitempty" json:"password,omitempty"`
	PasswordNeverExpires *PasswordNeverExpires `xml:" passwordNeverExpires,omitempty" json:"passwordNeverExpires,omitempty"`
	UserId               *UserId               `xml:" userId,omitempty" json:"userId,omitempty"`
}

type UserId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Users struct {
	User *User `xml:" user,omitempty" json:"user,omitempty"`
}
