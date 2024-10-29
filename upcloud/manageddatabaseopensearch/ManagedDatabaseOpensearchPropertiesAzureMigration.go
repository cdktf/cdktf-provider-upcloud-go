// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch


type ManagedDatabaseOpensearchPropertiesAzureMigration struct {
	// Account name. Azure account name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.14.0/docs/resources/managed_database_opensearch#account ManagedDatabaseOpensearch#account}
	Account *string `field:"optional" json:"account" yaml:"account"`
	// The path to the repository data within its container.
	//
	// The path to the repository data within its container. The value of this setting should not start or end with a /.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.14.0/docs/resources/managed_database_opensearch#base_path ManagedDatabaseOpensearch#base_path}
	BasePath *string `field:"optional" json:"basePath" yaml:"basePath"`
	// Chunk size.
	//
	// Big files can be broken down into chunks during snapshotting if needed. Should be the same as for the 3rd party repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.14.0/docs/resources/managed_database_opensearch#chunk_size ManagedDatabaseOpensearch#chunk_size}
	ChunkSize *string `field:"optional" json:"chunkSize" yaml:"chunkSize"`
	// Metadata files are stored in compressed format. when set to true metadata files are stored in compressed format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.14.0/docs/resources/managed_database_opensearch#compress ManagedDatabaseOpensearch#compress}
	Compress interface{} `field:"optional" json:"compress" yaml:"compress"`
	// Azure container name. Azure container name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.14.0/docs/resources/managed_database_opensearch#container ManagedDatabaseOpensearch#container}
	Container *string `field:"optional" json:"container" yaml:"container"`
	// Endpoint suffix. Defines the DNS suffix for Azure Storage endpoints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.14.0/docs/resources/managed_database_opensearch#endpoint_suffix ManagedDatabaseOpensearch#endpoint_suffix}
	EndpointSuffix *string `field:"optional" json:"endpointSuffix" yaml:"endpointSuffix"`
	// Account secret key. Azure account secret key. One of key or sas_token should be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.14.0/docs/resources/managed_database_opensearch#key ManagedDatabaseOpensearch#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// SAS token. A shared access signatures (SAS) token. One of key or sas_token should be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.14.0/docs/resources/managed_database_opensearch#sas_token ManagedDatabaseOpensearch#sas_token}
	SasToken *string `field:"optional" json:"sasToken" yaml:"sasToken"`
	// The snapshot name to restore from. The snapshot name to restore from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.14.0/docs/resources/managed_database_opensearch#snapshot_name ManagedDatabaseOpensearch#snapshot_name}
	SnapshotName *string `field:"optional" json:"snapshotName" yaml:"snapshotName"`
}

