package manageddatabaseuser


type ManagedDatabaseUserPgAccessControl struct {
	// Grant replication privilege.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/managed_database_user#allow_replication ManagedDatabaseUser#allow_replication}
	AllowReplication interface{} `field:"optional" json:"allowReplication" yaml:"allowReplication"`
}

