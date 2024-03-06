// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch


type ManagedDatabaseOpensearchPropertiesMigration struct {
	// Database name for bootstrapping the initial connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.0.3/docs/resources/managed_database_opensearch#dbname ManagedDatabaseOpensearch#dbname}
	Dbname *string `field:"optional" json:"dbname" yaml:"dbname"`
	// Hostname or IP address of the server where to migrate data from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.0.3/docs/resources/managed_database_opensearch#host ManagedDatabaseOpensearch#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Comma-separated list of databases, which should be ignored during migration (supported by MySQL only at the moment).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.0.3/docs/resources/managed_database_opensearch#ignore_dbs ManagedDatabaseOpensearch#ignore_dbs}
	IgnoreDbs *string `field:"optional" json:"ignoreDbs" yaml:"ignoreDbs"`
	// Password for authentication with the server where to migrate data from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.0.3/docs/resources/managed_database_opensearch#password ManagedDatabaseOpensearch#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Port number of the server where to migrate data from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.0.3/docs/resources/managed_database_opensearch#port ManagedDatabaseOpensearch#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The server where to migrate data from is secured with SSL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.0.3/docs/resources/managed_database_opensearch#ssl ManagedDatabaseOpensearch#ssl}
	Ssl interface{} `field:"optional" json:"ssl" yaml:"ssl"`
	// User name for authentication with the server where to migrate data from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.0.3/docs/resources/managed_database_opensearch#username ManagedDatabaseOpensearch#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

