// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch


type ManagedDatabaseOpensearchPropertiesOpenid struct {
	// The ID of the OpenID Connect client. The ID of the OpenID Connect client configured in your IdP. Required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.5/docs/resources/managed_database_opensearch#client_id ManagedDatabaseOpensearch#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The client secret of the OpenID Connect.
	//
	// The client secret of the OpenID Connect client configured in your IdP. Required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.5/docs/resources/managed_database_opensearch#client_secret ManagedDatabaseOpensearch#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// OpenID Connect metadata/configuration URL.
	//
	// The URL of your IdP where the Security plugin can find the OpenID Connect metadata/configuration settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.5/docs/resources/managed_database_opensearch#connect_url ManagedDatabaseOpensearch#connect_url}
	ConnectUrl *string `field:"optional" json:"connectUrl" yaml:"connectUrl"`
	// Enable or disable OpenSearch OpenID Connect authentication.
	//
	// Enables or disables OpenID Connect authentication for OpenSearch. When enabled, users can authenticate using OpenID Connect with an Identity Provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.5/docs/resources/managed_database_opensearch#enabled ManagedDatabaseOpensearch#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// HTTP header name of the JWT token. HTTP header name of the JWT token. Optional. Default is Authorization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.5/docs/resources/managed_database_opensearch#header ManagedDatabaseOpensearch#header}
	Header *string `field:"optional" json:"header" yaml:"header"`
	// The HTTP header that stores the token.
	//
	// The HTTP header that stores the token. Typically the Authorization header with the Bearer schema: Authorization: Bearer <token>. Optional. Default is Authorization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.5/docs/resources/managed_database_opensearch#jwt_header ManagedDatabaseOpensearch#jwt_header}
	JwtHeader *string `field:"optional" json:"jwtHeader" yaml:"jwtHeader"`
	// URL JWT token.
	//
	// If the token is not transmitted in the HTTP header, but as an URL parameter, define the name of the parameter here. Optional.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.5/docs/resources/managed_database_opensearch#jwt_url_parameter ManagedDatabaseOpensearch#jwt_url_parameter}
	JwtUrlParameter *string `field:"optional" json:"jwtUrlParameter" yaml:"jwtUrlParameter"`
	// The maximum number of unknown key IDs in the time frame.
	//
	// The maximum number of unknown key IDs in the time frame. Default is 10. Optional.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.5/docs/resources/managed_database_opensearch#refresh_rate_limit_count ManagedDatabaseOpensearch#refresh_rate_limit_count}
	RefreshRateLimitCount *float64 `field:"optional" json:"refreshRateLimitCount" yaml:"refreshRateLimitCount"`
	// The time frame to use when checking the maximum number of unknown key IDs, in milliseconds.
	//
	// The time frame to use when checking the maximum number of unknown key IDs, in milliseconds. Optional.Default is 10000 (10 seconds).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.5/docs/resources/managed_database_opensearch#refresh_rate_limit_time_window_ms ManagedDatabaseOpensearch#refresh_rate_limit_time_window_ms}
	RefreshRateLimitTimeWindowMs *float64 `field:"optional" json:"refreshRateLimitTimeWindowMs" yaml:"refreshRateLimitTimeWindowMs"`
	// The key in the JSON payload that stores the user’s roles.
	//
	// The key in the JSON payload that stores the user’s roles. The value of this key must be a comma-separated list of roles. Required only if you want to use roles in the JWT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.5/docs/resources/managed_database_opensearch#roles_key ManagedDatabaseOpensearch#roles_key}
	RolesKey *string `field:"optional" json:"rolesKey" yaml:"rolesKey"`
	// The scope of the identity token issued by the IdP.
	//
	// The scope of the identity token issued by the IdP. Optional. Default is openid profile email address phone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.5/docs/resources/managed_database_opensearch#scope ManagedDatabaseOpensearch#scope}
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// The key in the JSON payload that stores the user’s name.
	//
	// The key in the JSON payload that stores the user’s name. If not defined, the subject registered claim is used. Most IdP providers use the preferred_username claim. Optional.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.20.5/docs/resources/managed_database_opensearch#subject_key ManagedDatabaseOpensearch#subject_key}
	SubjectKey *string `field:"optional" json:"subjectKey" yaml:"subjectKey"`
}

