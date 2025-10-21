// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedobjectstorage


type ManagedObjectStorageNetwork struct {
	// Network family. IPv6 currently not supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.29.0/docs/resources/managed_object_storage#family ManagedObjectStorage#family}
	Family *string `field:"required" json:"family" yaml:"family"`
	// Network name. Must be unique within the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.29.0/docs/resources/managed_object_storage#name ManagedObjectStorage#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Network type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.29.0/docs/resources/managed_object_storage#type ManagedObjectStorage#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Private network uuid. For public networks the field should be omitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.29.0/docs/resources/managed_object_storage#uuid ManagedObjectStorage#uuid}
	Uuid *string `field:"optional" json:"uuid" yaml:"uuid"`
}

