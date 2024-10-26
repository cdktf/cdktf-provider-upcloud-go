// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gatewayconnectiontunnel


type GatewayConnectionTunnelIpsecProperties struct {
	// IKE child SA rekey time in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.13.2/docs/resources/gateway_connection_tunnel#child_rekey_time GatewayConnectionTunnel#child_rekey_time}
	ChildRekeyTime *float64 `field:"optional" json:"childRekeyTime" yaml:"childRekeyTime"`
	// Delay before sending Dead Peer Detection packets if no traffic is detected, in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.13.2/docs/resources/gateway_connection_tunnel#dpd_delay GatewayConnectionTunnel#dpd_delay}
	DpdDelay *float64 `field:"optional" json:"dpdDelay" yaml:"dpdDelay"`
	// Timeout period for DPD reply before considering the peer to be dead, in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.13.2/docs/resources/gateway_connection_tunnel#dpd_timeout GatewayConnectionTunnel#dpd_timeout}
	DpdTimeout *float64 `field:"optional" json:"dpdTimeout" yaml:"dpdTimeout"`
	// Maximum IKE SA lifetime in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.13.2/docs/resources/gateway_connection_tunnel#ike_lifetime GatewayConnectionTunnel#ike_lifetime}
	IkeLifetime *float64 `field:"optional" json:"ikeLifetime" yaml:"ikeLifetime"`
	// List of Phase 1: Proposal algorithms.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.13.2/docs/resources/gateway_connection_tunnel#phase1_algorithms GatewayConnectionTunnel#phase1_algorithms}
	Phase1Algorithms *[]*string `field:"optional" json:"phase1Algorithms" yaml:"phase1Algorithms"`
	// List of Phase 1 Diffie-Hellman group numbers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.13.2/docs/resources/gateway_connection_tunnel#phase1_dh_group_numbers GatewayConnectionTunnel#phase1_dh_group_numbers}
	Phase1DhGroupNumbers *[]*float64 `field:"optional" json:"phase1DhGroupNumbers" yaml:"phase1DhGroupNumbers"`
	// List of Phase 1 integrity algorithms.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.13.2/docs/resources/gateway_connection_tunnel#phase1_integrity_algorithms GatewayConnectionTunnel#phase1_integrity_algorithms}
	Phase1IntegrityAlgorithms *[]*string `field:"optional" json:"phase1IntegrityAlgorithms" yaml:"phase1IntegrityAlgorithms"`
	// List of Phase 2: Security Association algorithms.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.13.2/docs/resources/gateway_connection_tunnel#phase2_algorithms GatewayConnectionTunnel#phase2_algorithms}
	Phase2Algorithms *[]*string `field:"optional" json:"phase2Algorithms" yaml:"phase2Algorithms"`
	// List of Phase 2 Diffie-Hellman group numbers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.13.2/docs/resources/gateway_connection_tunnel#phase2_dh_group_numbers GatewayConnectionTunnel#phase2_dh_group_numbers}
	Phase2DhGroupNumbers *[]*float64 `field:"optional" json:"phase2DhGroupNumbers" yaml:"phase2DhGroupNumbers"`
	// List of Phase 2 integrity algorithms.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.13.2/docs/resources/gateway_connection_tunnel#phase2_integrity_algorithms GatewayConnectionTunnel#phase2_integrity_algorithms}
	Phase2IntegrityAlgorithms *[]*string `field:"optional" json:"phase2IntegrityAlgorithms" yaml:"phase2IntegrityAlgorithms"`
	// IKE SA rekey time in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.13.2/docs/resources/gateway_connection_tunnel#rekey_time GatewayConnectionTunnel#rekey_time}
	RekeyTime *float64 `field:"optional" json:"rekeyTime" yaml:"rekeyTime"`
}

