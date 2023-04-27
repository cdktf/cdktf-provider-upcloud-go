package server


type ServerStorageDevices struct {
	// A valid storage UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.10.0/docs/resources/server#storage Server#storage}
	Storage *string `field:"required" json:"storage" yaml:"storage"`
	// The device address the storage will be attached to.
	//
	// Specify only the bus name (ide/scsi/virtio) to auto-select next available address from that bus.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.10.0/docs/resources/server#address Server#address}
	Address *string `field:"optional" json:"address" yaml:"address"`
	// The device type the storage will be attached as.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.10.0/docs/resources/server#type Server#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

