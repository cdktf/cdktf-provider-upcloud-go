// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storage


type StorageClone struct {
	// The unique identifier of the storage/template to clone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.0/docs/resources/storage#id Storage#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
}

