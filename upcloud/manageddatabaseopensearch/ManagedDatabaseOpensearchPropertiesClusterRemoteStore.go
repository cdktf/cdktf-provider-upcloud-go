// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch


type ManagedDatabaseOpensearchPropertiesClusterRemoteStore struct {
	// The amount of time to wait for the cluster state upload to complete.
	//
	// The amount of time to wait for the cluster state upload to complete. Defaults to 20s.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.27.0/docs/resources/managed_database_opensearch#state_global_metadata_upload_timeout ManagedDatabaseOpensearch#state_global_metadata_upload_timeout}
	StateGlobalMetadataUploadTimeout *string `field:"optional" json:"stateGlobalMetadataUploadTimeout" yaml:"stateGlobalMetadataUploadTimeout"`
	// The amount of time to wait for the manifest file upload to complete.
	//
	// The amount of time to wait for the manifest file upload to complete. The manifest file contains the details of each of the files uploaded for a single cluster state, both index metadata files and global metadata files. Defaults to 20s.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.27.0/docs/resources/managed_database_opensearch#state_metadata_manifest_upload_timeout ManagedDatabaseOpensearch#state_metadata_manifest_upload_timeout}
	StateMetadataManifestUploadTimeout *string `field:"optional" json:"stateMetadataManifestUploadTimeout" yaml:"stateMetadataManifestUploadTimeout"`
	// The default value of the translog buffer interval.
	//
	// The default value of the translog buffer interval used when performing periodic translog updates. This setting is only effective when the index setting `index.remote_store.translog.buffer_interval` is not present. Defaults to 650ms.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.27.0/docs/resources/managed_database_opensearch#translog_buffer_interval ManagedDatabaseOpensearch#translog_buffer_interval}
	TranslogBufferInterval *string `field:"optional" json:"translogBufferInterval" yaml:"translogBufferInterval"`
	// The maximum number of open translog files for remote-backed indexes.
	//
	// Sets the maximum number of open translog files for remote-backed indexes. This limits the total number of translog files per shard. After reaching this limit, the remote store flushes the translog files. Default is 1000. The minimum required is 100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.27.0/docs/resources/managed_database_opensearch#translog_max_readers ManagedDatabaseOpensearch#translog_max_readers}
	TranslogMaxReaders *float64 `field:"optional" json:"translogMaxReaders" yaml:"translogMaxReaders"`
}

