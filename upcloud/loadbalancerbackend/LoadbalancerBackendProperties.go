package loadbalancerbackend


type LoadbalancerBackendProperties struct {
	// Expected HTTP status code returned by the customer application to mark server as healthy. Ignored for tcp type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.11.0/docs/resources/loadbalancer_backend#health_check_expected_status LoadbalancerBackend#health_check_expected_status}
	HealthCheckExpectedStatus *float64 `field:"optional" json:"healthCheckExpectedStatus" yaml:"healthCheckExpectedStatus"`
	// Sets how many failed health checks are allowed until the backend member is taken off from the rotation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.11.0/docs/resources/loadbalancer_backend#health_check_fall LoadbalancerBackend#health_check_fall}
	HealthCheckFall *float64 `field:"optional" json:"healthCheckFall" yaml:"healthCheckFall"`
	// Interval between health checks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.11.0/docs/resources/loadbalancer_backend#health_check_interval LoadbalancerBackend#health_check_interval}
	HealthCheckInterval *float64 `field:"optional" json:"healthCheckInterval" yaml:"healthCheckInterval"`
	// Sets how many passing checks there must be before returning the backend member to the rotation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.11.0/docs/resources/loadbalancer_backend#health_check_rise LoadbalancerBackend#health_check_rise}
	HealthCheckRise *float64 `field:"optional" json:"healthCheckRise" yaml:"healthCheckRise"`
	// Health check type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.11.0/docs/resources/loadbalancer_backend#health_check_type LoadbalancerBackend#health_check_type}
	HealthCheckType *string `field:"optional" json:"healthCheckType" yaml:"healthCheckType"`
	// Target path for health check HTTP GET requests. Ignored for tcp type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.11.0/docs/resources/loadbalancer_backend#health_check_url LoadbalancerBackend#health_check_url}
	HealthCheckUrl *string `field:"optional" json:"healthCheckUrl" yaml:"healthCheckUrl"`
	// Enable outbound proxy protocol by setting the desired version. Empty string disables proxy protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.11.0/docs/resources/loadbalancer_backend#outbound_proxy_protocol LoadbalancerBackend#outbound_proxy_protocol}
	OutboundProxyProtocol *string `field:"optional" json:"outboundProxyProtocol" yaml:"outboundProxyProtocol"`
	// Sets sticky session cookie name. Empty string disables sticky session.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.11.0/docs/resources/loadbalancer_backend#sticky_session_cookie_name LoadbalancerBackend#sticky_session_cookie_name}
	StickySessionCookieName *string `field:"optional" json:"stickySessionCookieName" yaml:"stickySessionCookieName"`
	// Backend server timeout in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.11.0/docs/resources/loadbalancer_backend#timeout_server LoadbalancerBackend#timeout_server}
	TimeoutServer *float64 `field:"optional" json:"timeoutServer" yaml:"timeoutServer"`
	// Maximum inactivity time on the client and server side for tunnels in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/2.11.0/docs/resources/loadbalancer_backend#timeout_tunnel LoadbalancerBackend#timeout_tunnel}
	TimeoutTunnel *float64 `field:"optional" json:"timeoutTunnel" yaml:"timeoutTunnel"`
}

