// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch


type ManagedDatabaseOpensearchPropertiesSaml struct {
	// Enable or disable OpenSearch SAML authentication.
	//
	// Enables or disables SAML-based authentication for OpenSearch. When enabled, users can authenticate using SAML with an Identity Provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.9.1/docs/resources/managed_database_opensearch#enabled ManagedDatabaseOpensearch#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Identity Provider Entity ID.
	//
	// The unique identifier for the Identity Provider (IdP) entity that is used for SAML authentication. This value is typically provided by the IdP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.9.1/docs/resources/managed_database_opensearch#idp_entity_id ManagedDatabaseOpensearch#idp_entity_id}
	IdpEntityId *string `field:"optional" json:"idpEntityId" yaml:"idpEntityId"`
	// Identity Provider (IdP) SAML metadata URL.
	//
	// The URL of the SAML metadata for the Identity Provider (IdP). This is used to configure SAML-based authentication with the IdP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.9.1/docs/resources/managed_database_opensearch#idp_metadata_url ManagedDatabaseOpensearch#idp_metadata_url}
	IdpMetadataUrl *string `field:"optional" json:"idpMetadataUrl" yaml:"idpMetadataUrl"`
	// PEM-encoded root CA Content for SAML IdP server verification.
	//
	// This parameter specifies the PEM-encoded root certificate authority (CA) content for the SAML identity provider (IdP) server verification. The root CA content is used to verify the SSL/TLS certificate presented by the server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.9.1/docs/resources/managed_database_opensearch#idp_pemtrustedcas_content ManagedDatabaseOpensearch#idp_pemtrustedcas_content}
	IdpPemtrustedcasContent *string `field:"optional" json:"idpPemtrustedcasContent" yaml:"idpPemtrustedcasContent"`
	// SAML response role attribute.
	//
	// Optional. Specifies the attribute in the SAML response where role information is stored, if available. Role attributes are not required for SAML authentication, but can be included in SAML assertions by most Identity Providers (IdPs) to determine user access levels or permissions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.9.1/docs/resources/managed_database_opensearch#roles_key ManagedDatabaseOpensearch#roles_key}
	RolesKey *string `field:"optional" json:"rolesKey" yaml:"rolesKey"`
	// Service Provider Entity ID.
	//
	// The unique identifier for the Service Provider (SP) entity that is used for SAML authentication. This value is typically provided by the SP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.9.1/docs/resources/managed_database_opensearch#sp_entity_id ManagedDatabaseOpensearch#sp_entity_id}
	SpEntityId *string `field:"optional" json:"spEntityId" yaml:"spEntityId"`
	// SAML response subject attribute.
	//
	// Optional. Specifies the attribute in the SAML response where the subject identifier is stored. If not configured, the NameID attribute is used by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.9.1/docs/resources/managed_database_opensearch#subject_key ManagedDatabaseOpensearch#subject_key}
	SubjectKey *string `field:"optional" json:"subjectKey" yaml:"subjectKey"`
}

