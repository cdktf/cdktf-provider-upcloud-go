package provider


type UpcloudProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud#alias UpcloudProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Password for UpCloud API user. Can also be configured using the `UPCLOUD_PASSWORD` environment variable.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud#password UpcloudProvider#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Maximum number of retries.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud#retry_max UpcloudProvider#retry_max}
	RetryMax *float64 `field:"optional" json:"retryMax" yaml:"retryMax"`
	// Maximum time to wait between retries.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud#retry_wait_max_sec UpcloudProvider#retry_wait_max_sec}
	RetryWaitMaxSec *float64 `field:"optional" json:"retryWaitMaxSec" yaml:"retryWaitMaxSec"`
	// Minimum time to wait between retries.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud#retry_wait_min_sec UpcloudProvider#retry_wait_min_sec}
	RetryWaitMinSec *float64 `field:"optional" json:"retryWaitMinSec" yaml:"retryWaitMinSec"`
	// UpCloud username with API access. Can also be configured using the `UPCLOUD_USERNAME` environment variable.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud#username UpcloudProvider#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

