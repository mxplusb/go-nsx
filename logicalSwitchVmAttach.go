package go_nsx

type Com_dot_vmware_dot_vshield_dot_vsm_dot_inventory_dot_dto_dot_VnicDto struct {
	ObjectId    *ObjectId    `xml:" objectId,omitempty" json:"objectId,omitempty"`
	PortgroupId *PortgroupId `xml:" portgroupId,omitempty" json:"portgroupId,omitempty"`
	VnicUuid    *VnicUuid    `xml:" vnicUuid,omitempty" json:"vnicUuid,omitempty"`
}

type ObjectId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type PortgroupId struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Root struct {
	Com_dot_vmware_dot_vshield_dot_vsm_dot_inventory_dot_dto_dot_VnicDto *Com_dot_vmware_dot_vshield_dot_vsm_dot_inventory_dot_dto_dot_VnicDto `xml:" com.vmware.vshield.vsm.inventory.dto.VnicDto,omitempty" json:"com.vmware.vshield.vsm.inventory.dto.VnicDto,omitempty"`
}

type VnicUuid struct {
	Text string `xml:",chardata" json:",omitempty"`
}
