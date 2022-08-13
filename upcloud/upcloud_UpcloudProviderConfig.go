// Prebuilt upcloud Provider for Terraform CDK (cdktf)
package upcloud


type UpcloudProviderConfig struct {
	// Password for UpCloud API user.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud#password UpcloudProvider#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// UpCloud username with API access.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud#username UpcloudProvider#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud#alias UpcloudProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
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
}

