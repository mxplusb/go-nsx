package controller

type ControllerSpec struct {
	DatastoreId    *DatastoreId    `xml:" datastoreId,omitempty" json:"datastoreId,omitempty"`
	DeployType     *DeployType     `xml:" deployType,omitempty" json:"deployType,omitempty"`
	Description    *Description    `xml:" description,omitempty" json:"description,omitempty"`
	HostId         *HostId         `xml:" hostId,omitempty" json:"hostId,omitempty"`
	IpPoolId       *IpPoolId       `xml:" ipPoolId,omitempty" json:"ipPoolId,omitempty"`
	Name           *Name           `xml:" name,omitempty" json:"name,omitempty"`
	NetworkId      *NetworkId      `xml:" networkId,omitempty" json:"networkId,omitempty"`
	Password       *Password       `xml:" password,omitempty" json:"password,omitempty"`
	ResourcePoolId *ResourcePoolId `xml:" resourcePoolId,omitempty" json:"resourcePoolId,omitempty"`
}

type DatastoreId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DeployType struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Description struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type HostId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type IpPoolId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type NetworkId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Password struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type ResourcePoolId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	ControllerSpec *ControllerSpec `xml:" controllerSpec,omitempty" json:"controllerSpec,omitempty"`
}
