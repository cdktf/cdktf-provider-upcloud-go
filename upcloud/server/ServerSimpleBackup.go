// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package server


type ServerSimpleBackup struct {
	// Simple backup plan. Accepted values: daily, dailies, weeklies, monthlies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.1/docs/resources/server#plan Server#plan}
	Plan *string `field:"optional" json:"plan" yaml:"plan"`
	// Time of the day at which backup will be taken. Should be provided in a hhmm format (e.g. 2230).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.1/docs/resources/server#time Server#time}
	Time *string `field:"optional" json:"time" yaml:"time"`
}

