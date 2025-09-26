// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch


type ManagedDatabaseOpensearchPropertiesDiskWatermarks struct {
	// Flood stage watermark (percentage). The flood stage watermark for disk usage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.27.0/docs/resources/managed_database_opensearch#flood_stage ManagedDatabaseOpensearch#flood_stage}
	FloodStage *float64 `field:"optional" json:"floodStage" yaml:"floodStage"`
	// High watermark (percentage). The high watermark for disk usage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.27.0/docs/resources/managed_database_opensearch#high ManagedDatabaseOpensearch#high}
	High *float64 `field:"optional" json:"high" yaml:"high"`
	// Low watermark (percentage). The low watermark for disk usage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.27.0/docs/resources/managed_database_opensearch#low ManagedDatabaseOpensearch#low}
	Low *float64 `field:"optional" json:"low" yaml:"low"`
}

