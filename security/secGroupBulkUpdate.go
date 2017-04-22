package security

type ClientHandle struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Criteria struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DynamicCriteria struct {
	Criteria *Criteria `xml:" criteria,omitempty" json:"criteria,omitempty"`
	Key      *Key      `xml:" key,omitempty" json:"key,omitempty"`
	Operator *Operator `xml:" operator,omitempty" json:"operator,omitempty"`
	Value    *Value    `xml:" value,omitempty" json:"value,omitempty"`
}

type DynamicMemberDefinition struct {
	DynamicSet *DynamicSet `xml:" dynamicSet,omitempty" json:"dynamicSet,omitempty"`
}

type DynamicSet struct {
	DynamicCriteria *DynamicCriteria `xml:" dynamicCriteria,omitempty" json:"dynamicCriteria,omitempty"`
	Operator        *Operator        `xml:" operator,omitempty" json:"operator,omitempty"`
}

type ExcludeMember struct {
	ClientHandle       *ClientHandle       `xml:" clientHandle,omitempty" json:"clientHandle,omitempty"`
	ExtendedAttributes *ExtendedAttributes `xml:" extendedAttributes,omitempty" json:"extendedAttributes,omitempty"`
	Name               *Name               `xml:" name,omitempty" json:"name,omitempty"`
	ObjectId           *ObjectId           `xml:" objectId,omitempty" json:"objectId,omitempty"`
	ObjectTypeName     *ObjectTypeName     `xml:" objectTypeName,omitempty" json:"objectTypeName,omitempty"`
	Revision           *Revision           `xml:" revision,omitempty" json:"revision,omitempty"`
	Scope              *Scope              `xml:" scope,omitempty" json:"scope,omitempty"`
	Type               *Type               `xml:" type,omitempty" json:"type,omitempty"`
	VsmUuid            *VsmUuid            `xml:" vsmUuid,omitempty" json:"vsmUuid,omitempty"`
}

type ExtendedAttributes struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Id struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Key struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Member struct {
	ClientHandle       *ClientHandle       `xml:" clientHandle,omitempty" json:"clientHandle,omitempty"`
	ExtendedAttributes *ExtendedAttributes `xml:" extendedAttributes,omitempty" json:"extendedAttributes,omitempty"`
	Name               *Name               `xml:" name,omitempty" json:"name,omitempty"`
	ObjectId           *ObjectId           `xml:" objectId,omitempty" json:"objectId,omitempty"`
	ObjectTypeName     *ObjectTypeName     `xml:" objectTypeName,omitempty" json:"objectTypeName,omitempty"`
	Revision           *Revision           `xml:" revision,omitempty" json:"revision,omitempty"`
	Scope              *Scope              `xml:" scope,omitempty" json:"scope,omitempty"`
	Type               *Type               `xml:" type,omitempty" json:"type,omitempty"`
	VsmUuid            *VsmUuid            `xml:" vsmUuid,omitempty" json:"vsmUuid,omitempty"`
}

type Name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ObjectId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ObjectTypeName struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Operator struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Revision struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	Securitygroup *Securitygroup `xml:" securitygroup,omitempty" json:"securitygroup,omitempty"`
}

type Scope struct {
	Id             *Id             `xml:" id,omitempty" json:"id,omitempty"`
	Name           *Name           `xml:" name,omitempty" json:"name,omitempty"`
	ObjectTypeName *ObjectTypeName `xml:" objectTypeName,omitempty" json:"objectTypeName,omitempty"`
	Revision       *Revision       `xml:" revision,omitempty" json:"revision,omitempty"`
	VsmUuid        *VsmUuid        `xml:" vsmUuid,omitempty" json:"vsmUuid,omitempty"`
}

type Securitygroup struct {
	ClientHandle            *ClientHandle            `xml:" clientHandle,omitempty" json:"clientHandle,omitempty"`
	DynamicMemberDefinition *DynamicMemberDefinition `xml:" dynamicMemberDefinition,omitempty" json:"dynamicMemberDefinition,omitempty"`
	ExcludeMember           *ExcludeMember           `xml:" excludeMember,omitempty" json:"excludeMember,omitempty"`
	ExtendedAttributes      *ExtendedAttributes      `xml:" extendedAttributes,omitempty" json:"extendedAttributes,omitempty"`
	Member                  *Member                  `xml:" member,omitempty" json:"member,omitempty"`
	Name                    *Name                    `xml:" name,omitempty" json:"name,omitempty"`
	ObjectId                *ObjectId                `xml:" objectId,omitempty" json:"objectId,omitempty"`
	ObjectTypeName          *ObjectTypeName          `xml:" objectTypeName,omitempty" json:"objectTypeName,omitempty"`
	Revision                *Revision                `xml:" revision,omitempty" json:"revision,omitempty"`
	Scope                   *Scope                   `xml:" scope,omitempty" json:"scope,omitempty"`
	Type                    *Type                    `xml:" type,omitempty" json:"type,omitempty"`
	VsmUuid                 *VsmUuid                 `xml:" vsmUuid,omitempty" json:"vsmUuid,omitempty"`
}

type Type struct {
	TypeName *TypeName `xml:" typeName,omitempty" json:"typeName,omitempty"`
}

type TypeName struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Value struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type VsmUuid struct {
	Text string `xml:",chardata" json:",omitempty"`
}
