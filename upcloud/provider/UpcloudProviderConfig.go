// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type UpcloudProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.27.0/docs#alias UpcloudProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Password for UpCloud API user. Can also be configured using the `UPCLOUD_PASSWORD` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.27.0/docs#password UpcloudProvider#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The duration (in seconds) that the provider waits for an HTTP request towards UpCloud API to complete.
	//
	// Defaults to 120 seconds
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.27.0/docs#request_timeout_sec UpcloudProvider#request_timeout_sec}
	RequestTimeoutSec *float64 `field:"optional" json:"requestTimeoutSec" yaml:"requestTimeoutSec"`
	// Maximum number of retries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.27.0/docs#retry_max UpcloudProvider#retry_max}
	RetryMax *float64 `field:"optional" json:"retryMax" yaml:"retryMax"`
	// Maximum time to wait between retries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.27.0/docs#retry_wait_max_sec UpcloudProvider#retry_wait_max_sec}
	RetryWaitMaxSec *float64 `field:"optional" json:"retryWaitMaxSec" yaml:"retryWaitMaxSec"`
	// Minimum time to wait between retries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.27.0/docs#retry_wait_min_sec UpcloudProvider#retry_wait_min_sec}
	RetryWaitMinSec *float64 `field:"optional" json:"retryWaitMinSec" yaml:"retryWaitMinSec"`
	// Token for authenticating to UpCloud API.
	//
	// Can also be configured using the `UPCLOUD_TOKEN` environment variable or using the system keyring. Use `upctl account login` command to save a token to the system keyring. (EXPERIMENTAL)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.27.0/docs#token UpcloudProvider#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
	// UpCloud username with API access. Can also be configured using the `UPCLOUD_USERNAME` environment variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.27.0/docs#username UpcloudProvider#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

