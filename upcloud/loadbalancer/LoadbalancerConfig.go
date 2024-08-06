// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LoadbalancerConfig struct {
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
	// The name of the service must be unique within customer account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.9.1/docs/resources/loadbalancer#name Loadbalancer#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Plan which the service will have. You can list available load balancer plans with `upctl loadbalancer plans`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.9.1/docs/resources/loadbalancer#plan Loadbalancer#plan}
	Plan *string `field:"required" json:"plan" yaml:"plan"`
	// Zone in which the service will be hosted, e.g. `fi-hel1`. You can list available zones with `upctl zone list`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.9.1/docs/resources/loadbalancer#zone Loadbalancer#zone}
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// The service configured status indicates the service's current intended status. Managed by the customer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.9.1/docs/resources/loadbalancer#configured_status Loadbalancer#configured_status}
	ConfiguredStatus *string `field:"optional" json:"configuredStatus" yaml:"configuredStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.9.1/docs/resources/loadbalancer#id Loadbalancer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Key-value pairs to classify the load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.9.1/docs/resources/loadbalancer#labels Loadbalancer#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The day of the week on which maintenance will be performed.
	//
	// If not provided, we will randomly select a weekend day. Valid values `monday|tuesday|wednesday|thursday|friday|saturday|sunday`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.9.1/docs/resources/loadbalancer#maintenance_dow Loadbalancer#maintenance_dow}
	MaintenanceDow *string `field:"optional" json:"maintenanceDow" yaml:"maintenanceDow"`
	// The time at which the maintenance will begin in UTC.
	//
	// A 2-hour timeframe has been allocated for maintenance. During this period, the multi-node production plans will not experience any downtime, while the one-node plans will have a downtime of 1-2 minutes. If not provided, we will randomly select an off-peak time. Needs to be a valid time format in UTC HH:MM:SSZ, for example `20:01:01Z`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.9.1/docs/resources/loadbalancer#maintenance_time Loadbalancer#maintenance_time}
	MaintenanceTime *string `field:"optional" json:"maintenanceTime" yaml:"maintenanceTime"`
	// Private network UUID where traffic will be routed. Must reside in load balancer zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.9.1/docs/resources/loadbalancer#network Loadbalancer#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// networks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.9.1/docs/resources/loadbalancer#networks Loadbalancer#networks}
	Networks interface{} `field:"optional" json:"networks" yaml:"networks"`
}

