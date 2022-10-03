package server


type ServerTemplateBackupRule struct {
	// The weekday when the backup is created.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/server#interval Server#interval}
	Interval *string `field:"required" json:"interval" yaml:"interval"`
	// The number of days before a backup is automatically deleted.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/server#retention Server#retention}
	Retention *float64 `field:"required" json:"retention" yaml:"retention"`
	// The time of day when the backup is created.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/server#time Server#time}
	Time *string `field:"required" json:"time" yaml:"time"`
}

