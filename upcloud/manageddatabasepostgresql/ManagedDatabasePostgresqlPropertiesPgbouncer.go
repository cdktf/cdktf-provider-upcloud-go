// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabasepostgresql


type ManagedDatabasePostgresqlPropertiesPgbouncer struct {
	// If the automatically created database pools have been unused this many seconds, they are freed.
	//
	// If 0 then timeout is disabled. [seconds].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.1/docs/resources/managed_database_postgresql#autodb_idle_timeout ManagedDatabasePostgresql#autodb_idle_timeout}
	AutodbIdleTimeout *float64 `field:"optional" json:"autodbIdleTimeout" yaml:"autodbIdleTimeout"`
	// Do not allow more than this many server connections per database (regardless of user).
	//
	// Setting it to 0 means unlimited.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.1/docs/resources/managed_database_postgresql#autodb_max_db_connections ManagedDatabasePostgresql#autodb_max_db_connections}
	AutodbMaxDbConnections *float64 `field:"optional" json:"autodbMaxDbConnections" yaml:"autodbMaxDbConnections"`
	// PGBouncer pool mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.1/docs/resources/managed_database_postgresql#autodb_pool_mode ManagedDatabasePostgresql#autodb_pool_mode}
	AutodbPoolMode *string `field:"optional" json:"autodbPoolMode" yaml:"autodbPoolMode"`
	// If non-zero then create automatically a pool of that size per user when a pool doesn't exist.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.1/docs/resources/managed_database_postgresql#autodb_pool_size ManagedDatabasePostgresql#autodb_pool_size}
	AutodbPoolSize *float64 `field:"optional" json:"autodbPoolSize" yaml:"autodbPoolSize"`
	// List of parameters to ignore when given in startup packet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.1/docs/resources/managed_database_postgresql#ignore_startup_parameters ManagedDatabasePostgresql#ignore_startup_parameters}
	IgnoreStartupParameters *[]*string `field:"optional" json:"ignoreStartupParameters" yaml:"ignoreStartupParameters"`
	// PgBouncer tracks protocol-level named prepared statements related commands sent by the client in transaction and statement pooling modes when max_prepared_statements is set to a non-zero value.
	//
	// Setting it to 0 disables prepared statements. max_prepared_statements defaults to 100, and its maximum is 3000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.1/docs/resources/managed_database_postgresql#max_prepared_statements ManagedDatabasePostgresql#max_prepared_statements}
	MaxPreparedStatements *float64 `field:"optional" json:"maxPreparedStatements" yaml:"maxPreparedStatements"`
	// Add more server connections to pool if below this number.
	//
	// Improves behavior when usual load comes suddenly back after period of total inactivity. The value is effectively capped at the pool size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.1/docs/resources/managed_database_postgresql#min_pool_size ManagedDatabasePostgresql#min_pool_size}
	MinPoolSize *float64 `field:"optional" json:"minPoolSize" yaml:"minPoolSize"`
	// If a server connection has been idle more than this many seconds it will be dropped.
	//
	// If 0 then timeout is disabled. [seconds].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.1/docs/resources/managed_database_postgresql#server_idle_timeout ManagedDatabasePostgresql#server_idle_timeout}
	ServerIdleTimeout *float64 `field:"optional" json:"serverIdleTimeout" yaml:"serverIdleTimeout"`
	// The pooler will close an unused server connection that has been connected longer than this. [seconds].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.1/docs/resources/managed_database_postgresql#server_lifetime ManagedDatabasePostgresql#server_lifetime}
	ServerLifetime *float64 `field:"optional" json:"serverLifetime" yaml:"serverLifetime"`
	// Run server_reset_query (DISCARD ALL) in all pooling modes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.1/docs/resources/managed_database_postgresql#server_reset_query_always ManagedDatabasePostgresql#server_reset_query_always}
	ServerResetQueryAlways interface{} `field:"optional" json:"serverResetQueryAlways" yaml:"serverResetQueryAlways"`
}

