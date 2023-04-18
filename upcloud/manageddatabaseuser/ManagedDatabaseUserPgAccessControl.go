package manageddatabaseuser


type ManagedDatabaseUserPgAccessControl struct {
	// Grant replication privilege.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.9.1/docs/resources/managed_database_user#allow_replication ManagedDatabaseUser#allow_replication}
	AllowReplication interface{} `field:"optional" json:"allowReplication" yaml:"allowReplication"`
}

