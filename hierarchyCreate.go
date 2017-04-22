package go_nsx

type Description struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	SecurityPolicyHierarchy *SecurityPolicyHierarchy `xml:" securityPolicyHierarchy,omitempty" json:"securityPolicyHierarchy,omitempty"`
}

type SecurityGroup struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type SecurityPolicy struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type SecurityPolicyHierarchy struct {
	Description    *Description    `xml:" description,omitempty" json:"description,omitempty"`
	Name           *Name           `xml:" name,omitempty" json:"name,omitempty"`
	SecurityGroup  *SecurityGroup  `xml:" securityGroup,omitempty" json:"securityGroup,omitempty"`
	SecurityPolicy *SecurityPolicy `xml:" securityPolicy,omitempty" json:"securityPolicy,omitempty"`
}
