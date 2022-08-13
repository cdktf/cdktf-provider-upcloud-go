// Prebuilt upcloud Provider for Terraform CDK (cdktf)
package upcloud


type StorageImport struct {
	// The mode of the import task. One of `http_import` or `direct_upload`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/storage#source Storage#source}
	Source *string `field:"required" json:"source" yaml:"source"`
	// The location of the file to import. For `http_import` an accessible URL for `direct_upload` a local file.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/storage#source_location Storage#source_location}
	SourceLocation *string `field:"required" json:"sourceLocation" yaml:"sourceLocation"`
	// For `direct_upload`; an optional hash of the file to upload.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/storage#source_hash Storage#source_hash}
	SourceHash *string `field:"optional" json:"sourceHash" yaml:"sourceHash"`
}

