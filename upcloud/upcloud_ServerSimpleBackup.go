// Prebuilt upcloud Provider for Terraform CDK (cdktf)
package upcloud


type ServerSimpleBackup struct {
	// Simple backup plan. Accepted values: dailies, weeklies, monthlies.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/server#plan Server#plan}
	Plan *string `field:"required" json:"plan" yaml:"plan"`
	// Time of the day at which backup will be taken. Should be provided in a hhmm format (e.g. 2230).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/server#time Server#time}
	Time *string `field:"required" json:"time" yaml:"time"`
}

