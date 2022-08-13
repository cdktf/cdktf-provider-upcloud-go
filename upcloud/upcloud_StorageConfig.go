// Prebuilt upcloud Provider for Terraform CDK (cdktf)
package upcloud

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The size of the storage in gigabytes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/storage#size Storage#size}
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// A short, informative description.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/storage#title Storage#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// The zone in which the storage will be created.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/storage#zone Storage#zone}
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// backup_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/storage#backup_rule Storage#backup_rule}
	BackupRule *StorageBackupRule `field:"optional" json:"backupRule" yaml:"backupRule"`
	// clone block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/storage#clone Storage#clone}
	Clone *StorageClone `field:"optional" json:"clone" yaml:"clone"`
	// If set to true, the backup taken before the partition and filesystem resize attempt will be deleted immediately after success.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/storage#delete_autoresize_backup Storage#delete_autoresize_backup}
	DeleteAutoresizeBackup interface{} `field:"optional" json:"deleteAutoresizeBackup" yaml:"deleteAutoresizeBackup"`
	// If set to true, provider will attempt to resize partition and filesystem when the size of the storage changes.
	//
	// Please note that before the resize attempt is made, backup of the storage will be taken. If the resize attempt fails, the backup will be used
	// to restore the storage and then deleted. If the resize attempt succeeds, backup will be kept (unless delete_autoresize_backup option is set to true).
	// Taking and keeping backups incure costs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/storage#filesystem_autoresize Storage#filesystem_autoresize}
	FilesystemAutoresize interface{} `field:"optional" json:"filesystemAutoresize" yaml:"filesystemAutoresize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/storage#id Storage#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// import block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/storage#import Storage#import}
	Import *StorageImport `field:"optional" json:"import" yaml:"import"`
	// The storage tier to use.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/storage#tier Storage#tier}
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
}

