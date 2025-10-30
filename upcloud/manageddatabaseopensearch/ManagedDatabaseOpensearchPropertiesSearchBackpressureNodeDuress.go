// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch


type ManagedDatabaseOpensearchPropertiesSearchBackpressureNodeDuress struct {
	// The CPU usage threshold (as a percentage) required for a node to be considered to be under duress.
	//
	// The CPU usage threshold (as a percentage) required for a node to be considered to be under duress. Default is 0.9.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.30.0/docs/resources/managed_database_opensearch#cpu_threshold ManagedDatabaseOpensearch#cpu_threshold}
	CpuThreshold *float64 `field:"optional" json:"cpuThreshold" yaml:"cpuThreshold"`
	// The heap usage threshold (as a percentage) required for a node to be considered to be under duress.
	//
	// The heap usage threshold (as a percentage) required for a node to be considered to be under duress. Default is 0.7.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.30.0/docs/resources/managed_database_opensearch#heap_threshold ManagedDatabaseOpensearch#heap_threshold}
	HeapThreshold *float64 `field:"optional" json:"heapThreshold" yaml:"heapThreshold"`
	// The number of successive limit breaches after which the node is considered to be under duress.
	//
	// The number of successive limit breaches after which the node is considered to be under duress. Default is 3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.30.0/docs/resources/managed_database_opensearch#num_successive_breaches ManagedDatabaseOpensearch#num_successive_breaches}
	NumSuccessiveBreaches *float64 `field:"optional" json:"numSuccessiveBreaches" yaml:"numSuccessiveBreaches"`
}

