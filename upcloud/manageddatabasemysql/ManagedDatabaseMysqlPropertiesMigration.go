// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabasemysql


type ManagedDatabaseMysqlPropertiesMigration struct {
	// Database name for bootstrapping the initial connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.4.0/docs/resources/managed_database_mysql#dbname ManagedDatabaseMysql#dbname}
	Dbname *string `field:"optional" json:"dbname" yaml:"dbname"`
	// Hostname or IP address of the server where to migrate data from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.4.0/docs/resources/managed_database_mysql#host ManagedDatabaseMysql#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Comma-separated list of databases, which should be ignored during migration (supported by MySQL and PostgreSQL only at the moment).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.4.0/docs/resources/managed_database_mysql#ignore_dbs ManagedDatabaseMysql#ignore_dbs}
	IgnoreDbs *string `field:"optional" json:"ignoreDbs" yaml:"ignoreDbs"`
	// The migration method to be used (currently supported only by Redis, Dragonfly, MySQL and PostgreSQL service types).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.4.0/docs/resources/managed_database_mysql#method ManagedDatabaseMysql#method}
	Method *string `field:"optional" json:"method" yaml:"method"`
	// Password for authentication with the server where to migrate data from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.4.0/docs/resources/managed_database_mysql#password ManagedDatabaseMysql#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Port number of the server where to migrate data from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.4.0/docs/resources/managed_database_mysql#port ManagedDatabaseMysql#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The server where to migrate data from is secured with SSL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.4.0/docs/resources/managed_database_mysql#ssl ManagedDatabaseMysql#ssl}
	Ssl interface{} `field:"optional" json:"ssl" yaml:"ssl"`
	// User name for authentication with the server where to migrate data from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.4.0/docs/resources/managed_database_mysql#username ManagedDatabaseMysql#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

