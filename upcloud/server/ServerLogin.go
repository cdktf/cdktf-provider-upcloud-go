// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package server


type ServerLogin struct {
	// Indicates a password should be create to allow access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.1/docs/resources/server#create_password Server#create_password}
	CreatePassword interface{} `field:"optional" json:"createPassword" yaml:"createPassword"`
	// A list of ssh keys to access the server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.1/docs/resources/server#keys Server#keys}
	Keys *[]*string `field:"optional" json:"keys" yaml:"keys"`
	// The delivery method for the server's root password (one of `none`, `email` or `sms`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.1/docs/resources/server#password_delivery Server#password_delivery}
	PasswordDelivery *string `field:"optional" json:"passwordDelivery" yaml:"passwordDelivery"`
	// Username to be create to access the server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.1/docs/resources/server#user Server#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

