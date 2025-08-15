// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loadbalancerbackend


type LoadbalancerBackendProperties struct {
	// Expected HTTP status code returned by the customer application to mark server as healthy. Ignored for `tcp` `health_check_type`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.0/docs/resources/loadbalancer_backend#health_check_expected_status LoadbalancerBackend#health_check_expected_status}
	HealthCheckExpectedStatus *float64 `field:"optional" json:"healthCheckExpectedStatus" yaml:"healthCheckExpectedStatus"`
	// Sets how many failed health checks are allowed until the backend member is taken off from the rotation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.0/docs/resources/loadbalancer_backend#health_check_fall LoadbalancerBackend#health_check_fall}
	HealthCheckFall *float64 `field:"optional" json:"healthCheckFall" yaml:"healthCheckFall"`
	// Interval between health checks in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.0/docs/resources/loadbalancer_backend#health_check_interval LoadbalancerBackend#health_check_interval}
	HealthCheckInterval *float64 `field:"optional" json:"healthCheckInterval" yaml:"healthCheckInterval"`
	// Sets how many successful health checks are required to put the backend member back into rotation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.0/docs/resources/loadbalancer_backend#health_check_rise LoadbalancerBackend#health_check_rise}
	HealthCheckRise *float64 `field:"optional" json:"healthCheckRise" yaml:"healthCheckRise"`
	// Enables certificate verification with the system CA certificate bundle. Works with https scheme in health_check_url, otherwise ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.0/docs/resources/loadbalancer_backend#health_check_tls_verify LoadbalancerBackend#health_check_tls_verify}
	HealthCheckTlsVerify interface{} `field:"optional" json:"healthCheckTlsVerify" yaml:"healthCheckTlsVerify"`
	// Health check type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.0/docs/resources/loadbalancer_backend#health_check_type LoadbalancerBackend#health_check_type}
	HealthCheckType *string `field:"optional" json:"healthCheckType" yaml:"healthCheckType"`
	// Target path for health check HTTP GET requests. Ignored for `tcp` `health_check_type`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.0/docs/resources/loadbalancer_backend#health_check_url LoadbalancerBackend#health_check_url}
	HealthCheckUrl *string `field:"optional" json:"healthCheckUrl" yaml:"healthCheckUrl"`
	// Allow HTTP/2 connections to backend members by utilizing ALPN extension of TLS protocol, therefore it can only be enabled when tls_enabled is set to true.
	//
	// Note: members should support HTTP/2 for this setting to work.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.0/docs/resources/loadbalancer_backend#http2_enabled LoadbalancerBackend#http2_enabled}
	Http2Enabled interface{} `field:"optional" json:"http2Enabled" yaml:"http2Enabled"`
	// Enable outbound proxy protocol by setting the desired version. Defaults to empty string. Empty string disables proxy protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.0/docs/resources/loadbalancer_backend#outbound_proxy_protocol LoadbalancerBackend#outbound_proxy_protocol}
	OutboundProxyProtocol *string `field:"optional" json:"outboundProxyProtocol" yaml:"outboundProxyProtocol"`
	// Sets sticky session cookie name. Empty string disables sticky session.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.0/docs/resources/loadbalancer_backend#sticky_session_cookie_name LoadbalancerBackend#sticky_session_cookie_name}
	StickySessionCookieName *string `field:"optional" json:"stickySessionCookieName" yaml:"stickySessionCookieName"`
	// Backend server timeout in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.0/docs/resources/loadbalancer_backend#timeout_server LoadbalancerBackend#timeout_server}
	TimeoutServer *float64 `field:"optional" json:"timeoutServer" yaml:"timeoutServer"`
	// Maximum inactivity time on the client and server side for tunnels in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.0/docs/resources/loadbalancer_backend#timeout_tunnel LoadbalancerBackend#timeout_tunnel}
	TimeoutTunnel *float64 `field:"optional" json:"timeoutTunnel" yaml:"timeoutTunnel"`
	// Enables TLS connection from the load balancer to backend servers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.0/docs/resources/loadbalancer_backend#tls_enabled LoadbalancerBackend#tls_enabled}
	TlsEnabled interface{} `field:"optional" json:"tlsEnabled" yaml:"tlsEnabled"`
	// If enabled, then the system CA certificate bundle will be used for the certificate verification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.0/docs/resources/loadbalancer_backend#tls_use_system_ca LoadbalancerBackend#tls_use_system_ca}
	TlsUseSystemCa interface{} `field:"optional" json:"tlsUseSystemCa" yaml:"tlsUseSystemCa"`
	// Enables backend servers certificate verification.
	//
	// Please make sure that TLS config with the certificate bundle of type authority attached to the backend or `tls_use_system_ca` enabled. Note: `tls_verify` has preference over `health_check_tls_verify` when `tls_enabled` in true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.24.0/docs/resources/loadbalancer_backend#tls_verify LoadbalancerBackend#tls_verify}
	TlsVerify interface{} `field:"optional" json:"tlsVerify" yaml:"tlsVerify"`
}

