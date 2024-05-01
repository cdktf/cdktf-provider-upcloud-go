// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch


type ManagedDatabaseOpensearchPropertiesAuthFailureListeners struct {
	// internal_authentication_backend_limiting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.3/docs/resources/managed_database_opensearch#internal_authentication_backend_limiting ManagedDatabaseOpensearch#internal_authentication_backend_limiting}
	InternalAuthenticationBackendLimiting *ManagedDatabaseOpensearchPropertiesAuthFailureListenersInternalAuthenticationBackendLimiting `field:"optional" json:"internalAuthenticationBackendLimiting" yaml:"internalAuthenticationBackendLimiting"`
	// ip_rate_limiting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.2.3/docs/resources/managed_database_opensearch#ip_rate_limiting ManagedDatabaseOpensearch#ip_rate_limiting}
	IpRateLimiting *ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimiting `field:"optional" json:"ipRateLimiting" yaml:"ipRateLimiting"`
}

