package manageddatabaseuser


type ManagedDatabaseUserOpensearchAccessControl struct {
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.12.0/docs/resources/managed_database_user#rules ManagedDatabaseUser#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

