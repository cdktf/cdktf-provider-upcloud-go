// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servergroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServerGroupConfig struct {
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
	// Title of your server group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.0/docs/resources/server_group#title ServerGroup#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// Defines if a server group is an anti-affinity group.
	//
	// Setting this to `strict` or `yes` will
	// 	result in all servers in the group being placed on separate compute hosts. The value can be `strict`, `yes`, or `no`.
	//
	// 	* `strict` policy doesn't allow servers in the same server group to be on the same host
	// 	* `yes` refers to best-effort policy and tries to put servers on different hosts, but this is not guaranteed
	// 	* `no` refers to having no policy and thus no effect on server host affinity
	//
	// 	To verify if the anti-affinity policies are met by requesting a server group details from API. For more information
	// 	please see UpCloud API documentation on server groups.
	//
	// 	Plese also note that anti-affinity policies are only applied on server start. This means that if anti-affinity
	// 	policies in server group are not met, you need to manually restart the servers in said group,
	// 	for example via API, UpCloud Control Panel or upctl (UpCloud CLI)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.0/docs/resources/server_group#anti_affinity_policy ServerGroup#anti_affinity_policy}
	AntiAffinityPolicy *string `field:"optional" json:"antiAffinityPolicy" yaml:"antiAffinityPolicy"`
	// User defined key-value pairs to classify the server group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.0/docs/resources/server_group#labels ServerGroup#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// UUIDs of the servers that are members of this group.
	//
	// Servers can also be attached to the server group via `server_group` property of `upcloud_server`. See also `track_members` property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.0/docs/resources/server_group#members ServerGroup#members}
	Members *[]*string `field:"optional" json:"members" yaml:"members"`
	// Controls if members of the server group are being tracked in this resource.
	//
	// Set to `false` when using `server_group` property of `upcloud_server` to attach servers to the server group to avoid delayed state updates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.0/docs/resources/server_group#track_members ServerGroup#track_members}
	TrackMembers interface{} `field:"optional" json:"trackMembers" yaml:"trackMembers"`
}

