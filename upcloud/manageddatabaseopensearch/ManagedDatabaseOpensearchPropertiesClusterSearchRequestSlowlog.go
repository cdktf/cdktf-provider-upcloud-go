// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch


type ManagedDatabaseOpensearchPropertiesClusterSearchRequestSlowlog struct {
	// Log level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.1/docs/resources/managed_database_opensearch#level ManagedDatabaseOpensearch#level}
	Level *string `field:"optional" json:"level" yaml:"level"`
	// threshold block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.1/docs/resources/managed_database_opensearch#threshold ManagedDatabaseOpensearch#threshold}
	Threshold *ManagedDatabaseOpensearchPropertiesClusterSearchRequestSlowlogThreshold `field:"optional" json:"threshold" yaml:"threshold"`
}

