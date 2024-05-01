// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch


type ManagedDatabaseOpensearchPropertiesIndexTemplate struct {
	// index.mapping.nested_objects.limit. The maximum number of nested JSON objects that a single document can contain across all nested types. This limit helps to prevent out of memory errors when a document contains too many nested objects. Default is 10000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.3/docs/resources/managed_database_opensearch#mapping_nested_objects_limit ManagedDatabaseOpensearch#mapping_nested_objects_limit}
	MappingNestedObjectsLimit *float64 `field:"optional" json:"mappingNestedObjectsLimit" yaml:"mappingNestedObjectsLimit"`
	// The number of replicas each primary shard has.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.3/docs/resources/managed_database_opensearch#number_of_replicas ManagedDatabaseOpensearch#number_of_replicas}
	NumberOfReplicas *float64 `field:"optional" json:"numberOfReplicas" yaml:"numberOfReplicas"`
	// The number of primary shards that an index should have.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.3/docs/resources/managed_database_opensearch#number_of_shards ManagedDatabaseOpensearch#number_of_shards}
	NumberOfShards *float64 `field:"optional" json:"numberOfShards" yaml:"numberOfShards"`
}

