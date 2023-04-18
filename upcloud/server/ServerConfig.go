package server

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServerConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// A valid domain name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.9.1/docs/resources/server#hostname Server#hostname}
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// network_interface block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.9.1/docs/resources/server#network_interface Server#network_interface}
	NetworkInterface interface{} `field:"required" json:"networkInterface" yaml:"networkInterface"`
	// The zone in which the server will be hosted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.9.1/docs/resources/server#zone Server#zone}
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// The number of CPU for the server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.9.1/docs/resources/server#cpu Server#cpu}
	Cpu *float64 `field:"optional" json:"cpu" yaml:"cpu"`
	// Are firewall rules active for the server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.9.1/docs/resources/server#firewall Server#firewall}
	Firewall interface{} `field:"optional" json:"firewall" yaml:"firewall"`
	// Use this to start the VM on a specific host.
	//
	// Refers to value from host -attribute. Only available for private cloud hosts
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.9.1/docs/resources/server#host Server#host}
	Host *float64 `field:"optional" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.9.1/docs/resources/server#id Server#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Key-value pairs to classify the server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.9.1/docs/resources/server#labels Server#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// login block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.9.1/docs/resources/server#login Server#login}
	Login *ServerLogin `field:"optional" json:"login" yaml:"login"`
	// The size of memory for the server (in megabytes).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.9.1/docs/resources/server#mem Server#mem}
	Mem *float64 `field:"optional" json:"mem" yaml:"mem"`
	// Is the metadata service active for the server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.9.1/docs/resources/server#metadata Server#metadata}
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// The pricing plan used for the server. You can list available server plans with `upctl server plans`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.9.1/docs/resources/server#plan Server#plan}
	Plan *string `field:"optional" json:"plan" yaml:"plan"`
	// simple_backup block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.9.1/docs/resources/server#simple_backup Server#simple_backup}
	SimpleBackup *ServerSimpleBackup `field:"optional" json:"simpleBackup" yaml:"simpleBackup"`
	// storage_devices block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.9.1/docs/resources/server#storage_devices Server#storage_devices}
	StorageDevices interface{} `field:"optional" json:"storageDevices" yaml:"storageDevices"`
	// The server related tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.9.1/docs/resources/server#tags Server#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.9.1/docs/resources/server#template Server#template}
	Template *ServerTemplate `field:"optional" json:"template" yaml:"template"`
	// A short, informational description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.9.1/docs/resources/server#title Server#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
	// Defines URL for a server setup script, or the script body itself.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.9.1/docs/resources/server#user_data Server#user_data}
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
}

