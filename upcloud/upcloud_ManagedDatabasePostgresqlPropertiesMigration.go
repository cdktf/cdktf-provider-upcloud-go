// Prebuilt upcloud Provider for Terraform CDK (cdktf)
package upcloud


type ManagedDatabasePostgresqlPropertiesMigration struct {
	// Database name for bootstrapping the initial connection.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#dbname ManagedDatabasePostgresql#dbname}
	Dbname *string `field:"optional" json:"dbname" yaml:"dbname"`
	// Hostname or IP address of the server where to migrate data from.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#host ManagedDatabasePostgresql#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Comma-separated list of databases, which should be ignored during migration (supported by MySQL only at the moment).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#ignore_dbs ManagedDatabasePostgresql#ignore_dbs}
	IgnoreDbs *string `field:"optional" json:"ignoreDbs" yaml:"ignoreDbs"`
	// Password for authentication with the server where to migrate data from.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#password ManagedDatabasePostgresql#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Port number of the server where to migrate data from.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#port ManagedDatabasePostgresql#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The server where to migrate data from is secured with SSL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#ssl ManagedDatabasePostgresql#ssl}
	Ssl interface{} `field:"optional" json:"ssl" yaml:"ssl"`
	// User name for authentication with the server where to migrate data from.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#username ManagedDatabasePostgresql#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

