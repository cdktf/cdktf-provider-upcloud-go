// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch


type ManagedDatabaseOpensearchPropertiesJwt struct {
	// Enable or disable OpenSearch JWT authentication.
	//
	// Enables or disables JWT-based authentication for OpenSearch. When enabled, users can authenticate using JWT tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.0/docs/resources/managed_database_opensearch#enabled ManagedDatabaseOpensearch#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// JWT clock skew tolerance in seconds.
	//
	// The maximum allowed time difference in seconds between the JWT issuer's clock and the OpenSearch server's clock. This helps prevent token validation failures due to minor time synchronization issues.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.0/docs/resources/managed_database_opensearch#jwt_clock_skew_tolerance_seconds ManagedDatabaseOpensearch#jwt_clock_skew_tolerance_seconds}
	JwtClockSkewToleranceSeconds *float64 `field:"optional" json:"jwtClockSkewToleranceSeconds" yaml:"jwtClockSkewToleranceSeconds"`
	// HTTP header name for JWT token.
	//
	// The HTTP header name where the JWT token is transmitted. Typically 'Authorization' for Bearer tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.0/docs/resources/managed_database_opensearch#jwt_header ManagedDatabaseOpensearch#jwt_header}
	JwtHeader *string `field:"optional" json:"jwtHeader" yaml:"jwtHeader"`
	// URL parameter name for JWT token.
	//
	// If the JWT token is transmitted as a URL parameter instead of an HTTP header, specify the parameter name here.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.0/docs/resources/managed_database_opensearch#jwt_url_parameter ManagedDatabaseOpensearch#jwt_url_parameter}
	JwtUrlParameter *string `field:"optional" json:"jwtUrlParameter" yaml:"jwtUrlParameter"`
	// Required JWT audience.
	//
	// If specified, the JWT must contain an 'aud' claim that matches this value. This provides additional security by ensuring the JWT was issued for the expected audience.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.0/docs/resources/managed_database_opensearch#required_audience ManagedDatabaseOpensearch#required_audience}
	RequiredAudience *string `field:"optional" json:"requiredAudience" yaml:"requiredAudience"`
	// Required JWT issuer.
	//
	// If specified, the JWT must contain an 'iss' claim that matches this value. This provides additional security by ensuring the JWT was issued by the expected issuer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.0/docs/resources/managed_database_opensearch#required_issuer ManagedDatabaseOpensearch#required_issuer}
	RequiredIssuer *string `field:"optional" json:"requiredIssuer" yaml:"requiredIssuer"`
	// JWT claim key for roles.
	//
	// The key in the JWT payload that contains the user's roles. If specified, roles will be extracted from the JWT for authorization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.0/docs/resources/managed_database_opensearch#roles_key ManagedDatabaseOpensearch#roles_key}
	RolesKey *string `field:"optional" json:"rolesKey" yaml:"rolesKey"`
	// JWT signing key.
	//
	// The secret key used to sign and verify JWT tokens. This should be a secure, randomly generated key HMAC key or public RSA/ECDSA key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.0/docs/resources/managed_database_opensearch#signing_key ManagedDatabaseOpensearch#signing_key}
	SigningKey *string `field:"optional" json:"signingKey" yaml:"signingKey"`
	// JWT claim key for subject.
	//
	// The key in the JWT payload that contains the user's subject identifier. If not specified, the 'sub' claim is used by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.31.0/docs/resources/managed_database_opensearch#subject_key ManagedDatabaseOpensearch#subject_key}
	SubjectKey *string `field:"optional" json:"subjectKey" yaml:"subjectKey"`
}

