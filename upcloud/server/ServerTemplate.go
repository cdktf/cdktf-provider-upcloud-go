// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package server


type ServerTemplate struct {
	// A valid storage UUID or template name.
	//
	// You can list available public templates with `upctl storage list --public --template` and available private templates with `upctl storage list --template`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.12.0/docs/resources/server#storage Server#storage}
	Storage *string `field:"required" json:"storage" yaml:"storage"`
	// The device address the storage will be attached to (`scsi`|`virtio`|`ide`).
	//
	// Leave `address_position` field empty to auto-select next available address from that bus.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.12.0/docs/resources/server#address Server#address}
	Address *string `field:"optional" json:"address" yaml:"address"`
	// The device position in the given bus (defined via field `address`).
	//
	// For example `0:0`, or `0`. Leave empty to auto-select next available address in the given bus.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.12.0/docs/resources/server#address_position Server#address_position}
	AddressPosition *string `field:"optional" json:"addressPosition" yaml:"addressPosition"`
	// backup_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.12.0/docs/resources/server#backup_rule Server#backup_rule}
	BackupRule *ServerTemplateBackupRule `field:"optional" json:"backupRule" yaml:"backupRule"`
	// If set to true, the backup taken before the partition and filesystem resize attempt will be deleted immediately after success.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.12.0/docs/resources/server#delete_autoresize_backup Server#delete_autoresize_backup}
	DeleteAutoresizeBackup interface{} `field:"optional" json:"deleteAutoresizeBackup" yaml:"deleteAutoresizeBackup"`
	// Sets if the storage is encrypted at rest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.12.0/docs/resources/server#encrypt Server#encrypt}
	Encrypt interface{} `field:"optional" json:"encrypt" yaml:"encrypt"`
	// If set to true, provider will attempt to resize partition and filesystem when the size of template storage changes.
	//
	// Please note that before the resize attempt is made, backup of the storage will be taken. If the resize attempt fails, the backup will be used
	// 							to restore the storage and then deleted. If the resize attempt succeeds, backup will be kept (unless delete_autoresize_backup option is set to true).
	// 							Taking and keeping backups incure costs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.12.0/docs/resources/server#filesystem_autoresize Server#filesystem_autoresize}
	FilesystemAutoresize interface{} `field:"optional" json:"filesystemAutoresize" yaml:"filesystemAutoresize"`
	// The size of the storage in gigabytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.12.0/docs/resources/server#size Server#size}
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// A short, informative description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.12.0/docs/resources/server#title Server#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
}

