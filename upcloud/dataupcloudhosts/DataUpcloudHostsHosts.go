// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataupcloudhosts


type DataUpcloudHostsHosts struct {
	// statistics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.32.0/docs/data-sources/hosts#statistics DataUpcloudHosts#statistics}
	Statistics interface{} `field:"optional" json:"statistics" yaml:"statistics"`
}

