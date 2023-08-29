// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch


type ManagedDatabaseOpensearchProperties struct {
	// Automatic utility network IP Filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.12.0/docs/resources/managed_database_opensearch#automatic_utility_network_ip_filter ManagedDatabaseOpensearch#automatic_utility_network_ip_filter}
	AutomaticUtilityNetworkIpFilter interface{} `field:"optional" json:"automaticUtilityNetworkIpFilter" yaml:"automaticUtilityNetworkIpFilter"`
	// IP filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.12.0/docs/resources/managed_database_opensearch#ip_filter ManagedDatabaseOpensearch#ip_filter}
	IpFilter *[]*string `field:"optional" json:"ipFilter" yaml:"ipFilter"`
	// migration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.12.0/docs/resources/managed_database_opensearch#migration ManagedDatabaseOpensearch#migration}
	Migration *ManagedDatabaseOpensearchPropertiesMigration `field:"optional" json:"migration" yaml:"migration"`
	// Public access allows connections to your Managed Database services via the public internet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.12.0/docs/resources/managed_database_opensearch#public_access ManagedDatabaseOpensearch#public_access}
	PublicAccess interface{} `field:"optional" json:"publicAccess" yaml:"publicAccess"`
	// OpenSearch major version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.12.0/docs/resources/managed_database_opensearch#version ManagedDatabaseOpensearch#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

