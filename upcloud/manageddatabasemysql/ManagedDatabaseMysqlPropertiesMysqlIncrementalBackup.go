// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabasemysql


type ManagedDatabaseMysqlPropertiesMysqlIncrementalBackup struct {
	// Enable incremental backups.
	//
	// Enable periodic incremental backups. When enabled, full_backup_week_schedule must be set. Incremental backups only store changes since the last backup, making them faster and more storage-efficient than full backups. This is particularly useful for large databases where daily full backups would be too time-consuming or expensive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.1/docs/resources/managed_database_mysql#enabled ManagedDatabaseMysql#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Full backup week schedule.
	//
	// Comma-separated list of days of the week when full backups should be created. Valid values: mon, tue, wed, thu, fri, sat, sun.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.1/docs/resources/managed_database_mysql#full_backup_week_schedule ManagedDatabaseMysql#full_backup_week_schedule}
	FullBackupWeekSchedule *string `field:"optional" json:"fullBackupWeekSchedule" yaml:"fullBackupWeekSchedule"`
}

