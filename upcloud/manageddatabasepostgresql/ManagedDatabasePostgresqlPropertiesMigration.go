// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabasepostgresql


type ManagedDatabasePostgresqlPropertiesMigration struct {
	// Database name for bootstrapping the initial connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.0/docs/resources/managed_database_postgresql#dbname ManagedDatabasePostgresql#dbname}
	Dbname *string `field:"optional" json:"dbname" yaml:"dbname"`
	// Hostname or IP address of the server where to migrate data from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.0/docs/resources/managed_database_postgresql#host ManagedDatabasePostgresql#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Comma-separated list of databases, which should be ignored during migration (supported by MySQL and PostgreSQL only at the moment).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.0/docs/resources/managed_database_postgresql#ignore_dbs ManagedDatabasePostgresql#ignore_dbs}
	IgnoreDbs *string `field:"optional" json:"ignoreDbs" yaml:"ignoreDbs"`
	// Comma-separated list of database roles, which should be ignored during migration (supported by PostgreSQL only at the moment).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.0/docs/resources/managed_database_postgresql#ignore_roles ManagedDatabasePostgresql#ignore_roles}
	IgnoreRoles *string `field:"optional" json:"ignoreRoles" yaml:"ignoreRoles"`
	// The migration method to be used (currently supported only by Redis, Dragonfly, MySQL and PostgreSQL service types).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.0/docs/resources/managed_database_postgresql#method ManagedDatabasePostgresql#method}
	Method *string `field:"optional" json:"method" yaml:"method"`
	// Password for authentication with the server where to migrate data from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.0/docs/resources/managed_database_postgresql#password ManagedDatabasePostgresql#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Port number of the server where to migrate data from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.0/docs/resources/managed_database_postgresql#port ManagedDatabasePostgresql#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The server where to migrate data from is secured with SSL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.0/docs/resources/managed_database_postgresql#ssl ManagedDatabasePostgresql#ssl}
	Ssl interface{} `field:"optional" json:"ssl" yaml:"ssl"`
	// User name for authentication with the server where to migrate data from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.0/docs/resources/managed_database_postgresql#username ManagedDatabasePostgresql#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

