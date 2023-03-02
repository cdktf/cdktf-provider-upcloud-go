package storage


type StorageBackupRule struct {
	// The weekday when the backup is created.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/storage#interval Storage#interval}
	Interval *string `field:"required" json:"interval" yaml:"interval"`
	// The number of days before a backup is automatically deleted.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/storage#retention Storage#retention}
	Retention *float64 `field:"required" json:"retention" yaml:"retention"`
	// The time of day when the backup is created.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/storage#time Storage#time}
	Time *string `field:"required" json:"time" yaml:"time"`
}

