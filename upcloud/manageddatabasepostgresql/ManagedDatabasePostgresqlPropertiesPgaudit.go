// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabasepostgresql


type ManagedDatabasePostgresqlPropertiesPgaudit struct {
	// Enable pgaudit extension.
	//
	// Enable pgaudit extension. When enabled, pgaudit extension will be automatically installed.Otherwise, extension will be uninstalled but auditing configurations will be preserved.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_postgresql#feature_enabled ManagedDatabasePostgresql#feature_enabled}
	FeatureEnabled interface{} `field:"optional" json:"featureEnabled" yaml:"featureEnabled"`
	// Specifies which classes of statements will be logged by session audit logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_postgresql#log ManagedDatabasePostgresql#log}
	Log *[]*string `field:"optional" json:"log" yaml:"log"`
	// Specifies that session logging should be enabled in the casewhere all relations in a statement are in pg_catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_postgresql#log_catalog ManagedDatabasePostgresql#log_catalog}
	LogCatalog interface{} `field:"optional" json:"logCatalog" yaml:"logCatalog"`
	// Specifies whether log messages will be visible to a client process such as psql.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_postgresql#log_client ManagedDatabasePostgresql#log_client}
	LogClient interface{} `field:"optional" json:"logClient" yaml:"logClient"`
	// Specifies the log level that will be used for log entries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_postgresql#log_level ManagedDatabasePostgresql#log_level}
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// Crop parameters representation and whole statements if they exceed this threshold. A (default) value of -1 disable the truncation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_postgresql#log_max_string_length ManagedDatabasePostgresql#log_max_string_length}
	LogMaxStringLength *float64 `field:"optional" json:"logMaxStringLength" yaml:"logMaxStringLength"`
	// This GUC allows to turn off logging nested statements, that is, statements that are executed as part of another ExecutorRun.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_postgresql#log_nested_statements ManagedDatabasePostgresql#log_nested_statements}
	LogNestedStatements interface{} `field:"optional" json:"logNestedStatements" yaml:"logNestedStatements"`
	// Specifies that audit logging should include the parameters that were passed with the statement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_postgresql#log_parameter ManagedDatabasePostgresql#log_parameter}
	LogParameter interface{} `field:"optional" json:"logParameter" yaml:"logParameter"`
	// Specifies that parameter values longer than this setting (in bytes) should not be logged, but replaced with <long param suppressed>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_postgresql#log_parameter_max_size ManagedDatabasePostgresql#log_parameter_max_size}
	LogParameterMaxSize *float64 `field:"optional" json:"logParameterMaxSize" yaml:"logParameterMaxSize"`
	// Specifies whether session audit logging should create a separate log entry for each relation (TABLE, VIEW, etc.) referenced in a SELECT or DML statement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_postgresql#log_relation ManagedDatabasePostgresql#log_relation}
	LogRelation interface{} `field:"optional" json:"logRelation" yaml:"logRelation"`
	// Specifies that audit logging should include the rows retrieved or affected by a statement.
	//
	// When enabled the rows field will be included after the parameter field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_postgresql#log_rows ManagedDatabasePostgresql#log_rows}
	LogRows interface{} `field:"optional" json:"logRows" yaml:"logRows"`
	// Specifies whether logging will include the statement text and parameters (if enabled).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_postgresql#log_statement ManagedDatabasePostgresql#log_statement}
	LogStatement interface{} `field:"optional" json:"logStatement" yaml:"logStatement"`
	// Specifies whether logging will include the statement text and parameters with the first log entry for a statement/substatement combination or with every entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_postgresql#log_statement_once ManagedDatabasePostgresql#log_statement_once}
	LogStatementOnce interface{} `field:"optional" json:"logStatementOnce" yaml:"logStatementOnce"`
	// Specifies the master role to use for object audit logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.2/docs/resources/managed_database_postgresql#role ManagedDatabasePostgresql#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
}

