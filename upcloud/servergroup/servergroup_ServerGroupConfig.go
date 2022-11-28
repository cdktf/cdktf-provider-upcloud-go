package servergroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServerGroupConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Title of your server group.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/server_group#title ServerGroup#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// Is group an anti-affinity group.
	//
	// Setting this to true will result in all servers in the group being placed on separate compute hosts.
	// NOTE: this is an experimental feature. The anti-affinity policy is "best-effort" and it is not
	// guaranteed that all the servers will end up on a separate compute hosts. You can verify if the
	// anti-affinity policies are met by requesting a server group details from API. For more information
	// please see UpCloud API documentation on server groups
	// Plese also note that anti-affinity policies are only applied on server start. This means that if anti-affinity
	// policies in server group are not met, you need to manually restart the servers in said group,
	// for example via API, UpCloud Control Panel or upctl (UpCloud CLI)
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/server_group#anti_affinity ServerGroup#anti_affinity}
	AntiAffinity interface{} `field:"optional" json:"antiAffinity" yaml:"antiAffinity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/server_group#id ServerGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels for your server group.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/server_group#labels ServerGroup#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// UUIDs of the servers that are members of this group.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/server_group#members ServerGroup#members}
	Members *[]*string `field:"optional" json:"members" yaml:"members"`
}

