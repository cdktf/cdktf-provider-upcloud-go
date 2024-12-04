// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gateway


type GatewayAddress struct {
	// Name of the IP address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.16.0/docs/resources/gateway#name Gateway#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

