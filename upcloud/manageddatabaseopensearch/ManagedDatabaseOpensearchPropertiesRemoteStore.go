// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch


type ManagedDatabaseOpensearchPropertiesRemoteStore struct {
	// The variance factor that is used to calculate the dynamic bytes lag threshold.
	//
	// The variance factor that is used together with the moving average to calculate the dynamic bytes lag threshold for activating remote segment backpressure. Defaults to 10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.0/docs/resources/managed_database_opensearch#segment_pressure_bytes_lag_variance_factor ManagedDatabaseOpensearch#segment_pressure_bytes_lag_variance_factor}
	SegmentPressureBytesLagVarianceFactor *float64 `field:"optional" json:"segmentPressureBytesLagVarianceFactor" yaml:"segmentPressureBytesLagVarianceFactor"`
	// The minimum consecutive failure count for activating remote segment backpressure.
	//
	// The minimum consecutive failure count for activating remote segment backpressure. Defaults to 5.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.0/docs/resources/managed_database_opensearch#segment_pressure_consecutive_failures_limit ManagedDatabaseOpensearch#segment_pressure_consecutive_failures_limit}
	SegmentPressureConsecutiveFailuresLimit *float64 `field:"optional" json:"segmentPressureConsecutiveFailuresLimit" yaml:"segmentPressureConsecutiveFailuresLimit"`
	// Enables remote segment backpressure. Enables remote segment backpressure. Default is `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.0/docs/resources/managed_database_opensearch#segment_pressure_enabled ManagedDatabaseOpensearch#segment_pressure_enabled}
	SegmentPressureEnabled interface{} `field:"optional" json:"segmentPressureEnabled" yaml:"segmentPressureEnabled"`
	// The variance factor that is used to calculate the dynamic bytes lag threshold.
	//
	// The variance factor that is used together with the moving average to calculate the dynamic time lag threshold for activating remote segment backpressure. Defaults to 10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.0/docs/resources/managed_database_opensearch#segment_pressure_time_lag_variance_factor ManagedDatabaseOpensearch#segment_pressure_time_lag_variance_factor}
	SegmentPressureTimeLagVarianceFactor *float64 `field:"optional" json:"segmentPressureTimeLagVarianceFactor" yaml:"segmentPressureTimeLagVarianceFactor"`
}

