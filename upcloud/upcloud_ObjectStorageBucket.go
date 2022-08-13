// Prebuilt upcloud Provider for Terraform CDK (cdktf)
package upcloud


type ObjectStorageBucket struct {
	// The name of the bucket.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/object_storage#name ObjectStorage#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

