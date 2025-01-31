// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch


type ManagedDatabaseOpensearchPropertiesClusterSearchRequestSlowlogThreshold struct {
	// Debug threshold for total request took time.
	//
	// The value should be in the form count and unit, where unit one of (s,m,h,d,nanos,ms,micros) or -1. Default is -1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.18.0/docs/resources/managed_database_opensearch#debug ManagedDatabaseOpensearch#debug}
	Debug *string `field:"optional" json:"debug" yaml:"debug"`
	// Info threshold for total request took time.
	//
	// The value should be in the form count and unit, where unit one of (s,m,h,d,nanos,ms,micros) or -1. Default is -1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.18.0/docs/resources/managed_database_opensearch#info ManagedDatabaseOpensearch#info}
	Info *string `field:"optional" json:"info" yaml:"info"`
	// Trace threshold for total request took time.
	//
	// The value should be in the form count and unit, where unit one of (s,m,h,d,nanos,ms,micros) or -1. Default is -1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.18.0/docs/resources/managed_database_opensearch#trace ManagedDatabaseOpensearch#trace}
	Trace *string `field:"optional" json:"trace" yaml:"trace"`
	// Warning threshold for total request took time.
	//
	// The value should be in the form count and unit, where unit one of (s,m,h,d,nanos,ms,micros) or -1. Default is -1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.18.0/docs/resources/managed_database_opensearch#warn ManagedDatabaseOpensearch#warn}
	Warn *string `field:"optional" json:"warn" yaml:"warn"`
}

