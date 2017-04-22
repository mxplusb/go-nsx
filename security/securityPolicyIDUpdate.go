package security

type Action struct {
	Class                string                     `xml:" class,attr"  json:",omitempty"`
	ActionType             *ActionType             `xml:" actionType,omitempty" json:"actionType,omitempty"`
	Applications           *Applications           `xml:" applications,omitempty" json:"applications,omitempty"`
	Category               *Category               `xml:" category,omitempty" json:"category,omitempty"`
	Description            *Description            `xml:" description,omitempty" json:"description,omitempty"`
	IsActionEnforced       *IsActionEnforced       `xml:" isActionEnforced,omitempty" json:"isActionEnforced,omitempty"`
	IsActive               *IsActive               `xml:" isActive,omitempty" json:"isActive,omitempty"`
	IsEnabled              *IsEnabled              `xml:" isEnabled,omitempty" json:"isEnabled,omitempty"`
	Logged                 *Logged                 `xml:" logged,omitempty" json:"logged,omitempty"`
	Name                   *Name                   `xml:" name,omitempty" json:"name,omitempty"`
	Scope                  *Scope                  `xml:" scope,omitempty" json:"scope,omitempty"`
	SecondarySecurityGroup *SecondarySecurityGroup `xml:" secondarySecurityGroup,omitempty" json:"secondarySecurityGroup,omitempty"`
}

type ActionType struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ActionsByCategory struct {
	Action                    *Action                    `xml:" action,omitempty" json:"action,omitempty"`
	Category                  *Category                  `xml:" category,omitempty" json:"category,omitempty"`
	Direction                 *Direction                 `xml:" direction,omitempty" json:"direction,omitempty"`
	OutsideSecondaryContainer *OutsideSecondaryContainer `xml:" outsideSecondaryContainer,omitempty" json:"outsideSecondaryContainer,omitempty"`
}

type Application struct {
	ObjectId *ObjectId `xml:" objectId,omitempty" json:"objectId,omitempty"`
}

type ApplicationGroup struct {
	ObjectId *ObjectId `xml:" objectId,omitempty" json:"objectId,omitempty"`
}

type Applications struct {
	Application      *Application      `xml:" application,omitempty" json:"application,omitempty"`
	ApplicationGroup *ApplicationGroup `xml:" applicationGroup,omitempty" json:"applicationGroup,omitempty"`
}

type Category struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Description struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Direction struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Id struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IsActionEnforced struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IsActive struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IsEnabled struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Logged struct {
	Text string `xml:",chardata" json:",omitempty"`
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

type OutsideSecondaryContainer struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Parent struct {
	ObjectId *ObjectId `xml:" objectId,omitempty" json:"objectId,omitempty"`
}

type Precedence struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	SecurityPolicy *SecurityPolicy `xml:" securityPolicy,omitempty" json:"securityPolicy,omitempty"`
}

type Scope struct {
	Id             *Id             `xml:" id,omitempty" json:"id,omitempty"`
	Name           *Name           `xml:" name,omitempty" json:"name,omitempty"`
	ObjectTypeName *ObjectTypeName `xml:" objectTypeName,omitempty" json:"objectTypeName,omitempty"`
}

type SecondarySecurityGroup struct {
	ObjectId *ObjectId `xml:" objectId,omitempty" json:"objectId,omitempty"`
}

type SecurityGroupBinding struct {
	ObjectId *ObjectId `xml:" objectId,omitempty" json:"objectId,omitempty"`
}

type SecurityPolicy struct {
	ActionsByCategory    *ActionsByCategory    `xml:" actionsByCategory,omitempty" json:"actionsByCategory,omitempty"`
	Description          *Description          `xml:" description,omitempty" json:"description,omitempty"`
	Name                 *Name                 `xml:" name,omitempty" json:"name,omitempty"`
	Parent               *Parent               `xml:" parent,omitempty" json:"parent,omitempty"`
	Precedence           *Precedence           `xml:" precedence,omitempty" json:"precedence,omitempty"`
	SecurityGroupBinding *SecurityGroupBinding `xml:" securityGroupBinding,omitempty" json:"securityGroupBinding,omitempty"`
	SecurityPolicy       *SecurityPolicy       `xml:" securityPolicy,omitempty" json:"securityPolicy,omitempty"`
}
