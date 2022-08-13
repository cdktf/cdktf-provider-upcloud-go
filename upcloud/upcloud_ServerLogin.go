// Prebuilt upcloud Provider for Terraform CDK (cdktf)
package upcloud


type ServerLogin struct {
	// Indicates a password should be create to allow access.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/server#create_password Server#create_password}
	CreatePassword interface{} `field:"optional" json:"createPassword" yaml:"createPassword"`
	// A list of ssh keys to access the server.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/server#keys Server#keys}
	Keys *[]*string `field:"optional" json:"keys" yaml:"keys"`
	// The delivery method for the server’s root password.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/server#password_delivery Server#password_delivery}
	PasswordDelivery *string `field:"optional" json:"passwordDelivery" yaml:"passwordDelivery"`
	// Username to be create to access the server.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/upcloud/r/server#user Server#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

