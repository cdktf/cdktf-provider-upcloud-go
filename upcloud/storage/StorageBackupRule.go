// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storage


type StorageBackupRule struct {
	// The weekday when the backup is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.9.1/docs/resources/storage#interval Storage#interval}
	Interval *string `field:"required" json:"interval" yaml:"interval"`
	// The number of days before a backup is automatically deleted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.9.1/docs/resources/storage#retention Storage#retention}
	Retention *float64 `field:"required" json:"retention" yaml:"retention"`
	// The time of day when the backup is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.9.1/docs/resources/storage#time Storage#time}
	Time *string `field:"required" json:"time" yaml:"time"`
}

