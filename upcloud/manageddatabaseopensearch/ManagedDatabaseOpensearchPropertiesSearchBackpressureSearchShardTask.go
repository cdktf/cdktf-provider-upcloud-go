// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch


type ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchShardTask struct {
	// The maximum number of search tasks to cancel in a single iteration of the observer thread.
	//
	// The maximum number of search tasks to cancel in a single iteration of the observer thread. Default is 10.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.29.0/docs/resources/managed_database_opensearch#cancellation_burst ManagedDatabaseOpensearch#cancellation_burst}
	CancellationBurst *float64 `field:"optional" json:"cancellationBurst" yaml:"cancellationBurst"`
	// The maximum number of tasks to cancel per millisecond of elapsed time.
	//
	// The maximum number of tasks to cancel per millisecond of elapsed time. Default is 0.003.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.29.0/docs/resources/managed_database_opensearch#cancellation_rate ManagedDatabaseOpensearch#cancellation_rate}
	CancellationRate *float64 `field:"optional" json:"cancellationRate" yaml:"cancellationRate"`
	// The maximum number of tasks to cancel.
	//
	// The maximum number of tasks to cancel, as a percentage of successful task completions. Default is 0.1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.29.0/docs/resources/managed_database_opensearch#cancellation_ratio ManagedDatabaseOpensearch#cancellation_ratio}
	CancellationRatio *float64 `field:"optional" json:"cancellationRatio" yaml:"cancellationRatio"`
	// The CPU usage threshold (in milliseconds) required for a single search shard task before it is considered for cancellation.
	//
	// The CPU usage threshold (in milliseconds) required for a single search shard task before it is considered for cancellation. Default is 15000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.29.0/docs/resources/managed_database_opensearch#cpu_time_millis_threshold ManagedDatabaseOpensearch#cpu_time_millis_threshold}
	CpuTimeMillisThreshold *float64 `field:"optional" json:"cpuTimeMillisThreshold" yaml:"cpuTimeMillisThreshold"`
	// The elapsed time threshold (in milliseconds) required for a single search shard task before it is considered for cancellation.
	//
	// The elapsed time threshold (in milliseconds) required for a single search shard task before it is considered for cancellation. Default is 30000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.29.0/docs/resources/managed_database_opensearch#elapsed_time_millis_threshold ManagedDatabaseOpensearch#elapsed_time_millis_threshold}
	ElapsedTimeMillisThreshold *float64 `field:"optional" json:"elapsedTimeMillisThreshold" yaml:"elapsedTimeMillisThreshold"`
	// The number of previously completed search shard tasks to consider when calculating the rolling average of heap usage.
	//
	// The number of previously completed search shard tasks to consider when calculating the rolling average of heap usage. Default is 100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.29.0/docs/resources/managed_database_opensearch#heap_moving_average_window_size ManagedDatabaseOpensearch#heap_moving_average_window_size}
	HeapMovingAverageWindowSize *float64 `field:"optional" json:"heapMovingAverageWindowSize" yaml:"heapMovingAverageWindowSize"`
	// The heap usage threshold (as a percentage) required for a single search shard task before it is considered for cancellation.
	//
	// The heap usage threshold (as a percentage) required for a single search shard task before it is considered for cancellation. Default is 0.5.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.29.0/docs/resources/managed_database_opensearch#heap_percent_threshold ManagedDatabaseOpensearch#heap_percent_threshold}
	HeapPercentThreshold *float64 `field:"optional" json:"heapPercentThreshold" yaml:"heapPercentThreshold"`
	// The minimum variance required for a single search shard task’s heap usage compared to the rolling average of previously completed tasks before it is considered for cancellation.
	//
	// The minimum variance required for a single search shard task’s heap usage compared to the rolling average of previously completed tasks before it is considered for cancellation. Default is 2.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.29.0/docs/resources/managed_database_opensearch#heap_variance ManagedDatabaseOpensearch#heap_variance}
	HeapVariance *float64 `field:"optional" json:"heapVariance" yaml:"heapVariance"`
	// The heap usage threshold (as a percentage) required for the sum of heap usages of all search shard tasks before cancellation is applied.
	//
	// The heap usage threshold (as a percentage) required for the sum of heap usages of all search shard tasks before cancellation is applied. Default is 0.5.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.29.0/docs/resources/managed_database_opensearch#total_heap_percent_threshold ManagedDatabaseOpensearch#total_heap_percent_threshold}
	TotalHeapPercentThreshold *float64 `field:"optional" json:"totalHeapPercentThreshold" yaml:"totalHeapPercentThreshold"`
}

