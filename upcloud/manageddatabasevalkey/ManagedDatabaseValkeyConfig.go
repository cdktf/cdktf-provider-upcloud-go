// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabasevalkey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedDatabaseValkeyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.1/docs/resources/managed_database_valkey#name ManagedDatabaseValkey#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Service plan to use.
	//
	// This determines how much resources the instance will have. You can list available plans with `upctl database plans <type>`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.1/docs/resources/managed_database_valkey#plan ManagedDatabaseValkey#plan}
	Plan *string `field:"required" json:"plan" yaml:"plan"`
	// Title of a managed database instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.1/docs/resources/managed_database_valkey#title ManagedDatabaseValkey#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// Zone where the instance resides, e.g. `de-fra1`. You can list available zones with `upctl zone list`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.1/docs/resources/managed_database_valkey#zone ManagedDatabaseValkey#zone}
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.1/docs/resources/managed_database_valkey#id ManagedDatabaseValkey#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// User defined key-value pairs to classify the managed database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.1/docs/resources/managed_database_valkey#labels ManagedDatabaseValkey#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Maintenance window day of week. Lower case weekday name (monday, tuesday, ...).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.1/docs/resources/managed_database_valkey#maintenance_window_dow ManagedDatabaseValkey#maintenance_window_dow}
	MaintenanceWindowDow *string `field:"optional" json:"maintenanceWindowDow" yaml:"maintenanceWindowDow"`
	// Maintenance window UTC time in hh:mm:ss format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.1/docs/resources/managed_database_valkey#maintenance_window_time ManagedDatabaseValkey#maintenance_window_time}
	MaintenanceWindowTime *string `field:"optional" json:"maintenanceWindowTime" yaml:"maintenanceWindowTime"`
	// network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.1/docs/resources/managed_database_valkey#network ManagedDatabaseValkey#network}
	Network interface{} `field:"optional" json:"network" yaml:"network"`
	// The administrative power state of the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.1/docs/resources/managed_database_valkey#powered ManagedDatabaseValkey#powered}
	Powered interface{} `field:"optional" json:"powered" yaml:"powered"`
	// properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.1/docs/resources/managed_database_valkey#properties ManagedDatabaseValkey#properties}
	Properties *ManagedDatabaseValkeyProperties `field:"optional" json:"properties" yaml:"properties"`
	// If set to true, prevents the managed service from being powered off, or deleted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.1/docs/resources/managed_database_valkey#termination_protection ManagedDatabaseValkey#termination_protection}
	TerminationProtection interface{} `field:"optional" json:"terminationProtection" yaml:"terminationProtection"`
}

