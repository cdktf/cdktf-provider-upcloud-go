// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch


type ManagedDatabaseOpensearchPropertiesAuthFailureListenersIpRateLimiting struct {
	// The number of login attempts allowed before login is blocked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.4.0/docs/resources/managed_database_opensearch#allowed_tries ManagedDatabaseOpensearch#allowed_tries}
	AllowedTries *float64 `field:"optional" json:"allowedTries" yaml:"allowedTries"`
	// The duration of time that login remains blocked after a failed login.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.4.0/docs/resources/managed_database_opensearch#block_expiry_seconds ManagedDatabaseOpensearch#block_expiry_seconds}
	BlockExpirySeconds *float64 `field:"optional" json:"blockExpirySeconds" yaml:"blockExpirySeconds"`
	// The maximum number of blocked IP addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.4.0/docs/resources/managed_database_opensearch#max_blocked_clients ManagedDatabaseOpensearch#max_blocked_clients}
	MaxBlockedClients *float64 `field:"optional" json:"maxBlockedClients" yaml:"maxBlockedClients"`
	// The maximum number of tracked IP addresses that have failed login.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.4.0/docs/resources/managed_database_opensearch#max_tracked_clients ManagedDatabaseOpensearch#max_tracked_clients}
	MaxTrackedClients *float64 `field:"optional" json:"maxTrackedClients" yaml:"maxTrackedClients"`
	// The window of time in which the value for `allowed_tries` is enforced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.4.0/docs/resources/managed_database_opensearch#time_window_seconds ManagedDatabaseOpensearch#time_window_seconds}
	TimeWindowSeconds *float64 `field:"optional" json:"timeWindowSeconds" yaml:"timeWindowSeconds"`
	// The type of rate limiting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.4.0/docs/resources/managed_database_opensearch#type ManagedDatabaseOpensearch#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

