// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseredis

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedDatabaseRedisConfig struct {
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
	// Name of the service.
	//
	// The name is used as a prefix for the logical hostname. Must be unique within an account
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.0.1/docs/resources/managed_database_redis#name ManagedDatabaseRedis#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Service plan to use.
	//
	// This determines how much resources the instance will have. You can list available plans with `upctl database plans <type>`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.0.1/docs/resources/managed_database_redis#plan ManagedDatabaseRedis#plan}
	Plan *string `field:"required" json:"plan" yaml:"plan"`
	// Title of a managed database instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.0.1/docs/resources/managed_database_redis#title ManagedDatabaseRedis#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// Zone where the instance resides, e.g. `de-fra1`. You can list available zones with `upctl zone list`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.0.1/docs/resources/managed_database_redis#zone ManagedDatabaseRedis#zone}
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.0.1/docs/resources/managed_database_redis#id ManagedDatabaseRedis#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Maintenance window day of week. Lower case weekday name (monday, tuesday, ...).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.0.1/docs/resources/managed_database_redis#maintenance_window_dow ManagedDatabaseRedis#maintenance_window_dow}
	MaintenanceWindowDow *string `field:"optional" json:"maintenanceWindowDow" yaml:"maintenanceWindowDow"`
	// Maintenance window UTC time in hh:mm:ss format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.0.1/docs/resources/managed_database_redis#maintenance_window_time ManagedDatabaseRedis#maintenance_window_time}
	MaintenanceWindowTime *string `field:"optional" json:"maintenanceWindowTime" yaml:"maintenanceWindowTime"`
	// The administrative power state of the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.0.1/docs/resources/managed_database_redis#powered ManagedDatabaseRedis#powered}
	Powered interface{} `field:"optional" json:"powered" yaml:"powered"`
	// properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.0.1/docs/resources/managed_database_redis#properties ManagedDatabaseRedis#properties}
	Properties *ManagedDatabaseRedisProperties `field:"optional" json:"properties" yaml:"properties"`
}

