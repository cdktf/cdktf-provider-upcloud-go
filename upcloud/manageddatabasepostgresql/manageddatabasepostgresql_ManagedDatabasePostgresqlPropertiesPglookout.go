package manageddatabasepostgresql


type ManagedDatabasePostgresqlPropertiesPglookout struct {
	// max_failover_replication_time_lag.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_postgresql#max_failover_replication_time_lag ManagedDatabasePostgresql#max_failover_replication_time_lag}
	MaxFailoverReplicationTimeLag *float64 `field:"optional" json:"maxFailoverReplicationTimeLag" yaml:"maxFailoverReplicationTimeLag"`
}

