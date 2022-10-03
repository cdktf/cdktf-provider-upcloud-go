package manageddatabasepostgresql


type ManagedDatabasePostgresqlPropertiesTimescaledb struct {
	// timescaledb.max_background_workers.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#max_background_workers ManagedDatabasePostgresql#max_background_workers}
	MaxBackgroundWorkers *float64 `field:"optional" json:"maxBackgroundWorkers" yaml:"maxBackgroundWorkers"`
}

