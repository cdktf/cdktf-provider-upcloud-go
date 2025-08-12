// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package manageddatabaseopensearch


type ManagedDatabaseOpensearchPropertiesSearchBackpressureSearchTask struct {
	// The maximum number of search tasks to cancel in a single iteration of the observer thread.
	//
	// The maximum number of search tasks to cancel in a single iteration of the observer thread. Default is 5.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.4/docs/resources/managed_database_opensearch#cancellation_burst ManagedDatabaseOpensearch#cancellation_burst}
	CancellationBurst *float64 `field:"optional" json:"cancellationBurst" yaml:"cancellationBurst"`
	// The maximum number of search tasks to cancel per millisecond of elapsed time.
	//
	// The maximum number of search tasks to cancel per millisecond of elapsed time. Default is 0.003.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.4/docs/resources/managed_database_opensearch#cancellation_rate ManagedDatabaseOpensearch#cancellation_rate}
	CancellationRate *float64 `field:"optional" json:"cancellationRate" yaml:"cancellationRate"`
	// The maximum number of search tasks to cancel, as a percentage of successful search task completions.
	//
	// The maximum number of search tasks to cancel, as a percentage of successful search task completions. Default is 0.1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.4/docs/resources/managed_database_opensearch#cancellation_ratio ManagedDatabaseOpensearch#cancellation_ratio}
	CancellationRatio *float64 `field:"optional" json:"cancellationRatio" yaml:"cancellationRatio"`
	// The CPU usage threshold (in milliseconds) required for an individual parent task before it is considered for cancellation.
	//
	// The CPU usage threshold (in milliseconds) required for an individual parent task before it is considered for cancellation. Default is 30000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.4/docs/resources/managed_database_opensearch#cpu_time_millis_threshold ManagedDatabaseOpensearch#cpu_time_millis_threshold}
	CpuTimeMillisThreshold *float64 `field:"optional" json:"cpuTimeMillisThreshold" yaml:"cpuTimeMillisThreshold"`
	// The elapsed time threshold (in milliseconds) required for an individual parent task before it is considered for cancellation.
	//
	// The elapsed time threshold (in milliseconds) required for an individual parent task before it is considered for cancellation. Default is 45000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.4/docs/resources/managed_database_opensearch#elapsed_time_millis_threshold ManagedDatabaseOpensearch#elapsed_time_millis_threshold}
	ElapsedTimeMillisThreshold *float64 `field:"optional" json:"elapsedTimeMillisThreshold" yaml:"elapsedTimeMillisThreshold"`
	// The window size used to calculate the rolling average of the heap usage for the completed parent tasks.
	//
	// The window size used to calculate the rolling average of the heap usage for the completed parent tasks. Default is 10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.4/docs/resources/managed_database_opensearch#heap_moving_average_window_size ManagedDatabaseOpensearch#heap_moving_average_window_size}
	HeapMovingAverageWindowSize *float64 `field:"optional" json:"heapMovingAverageWindowSize" yaml:"heapMovingAverageWindowSize"`
	// The heap usage threshold (as a percentage) required for an individual parent task before it is considered for cancellation.
	//
	// The heap usage threshold (as a percentage) required for an individual parent task before it is considered for cancellation. Default is 0.2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.4/docs/resources/managed_database_opensearch#heap_percent_threshold ManagedDatabaseOpensearch#heap_percent_threshold}
	HeapPercentThreshold *float64 `field:"optional" json:"heapPercentThreshold" yaml:"heapPercentThreshold"`
	// The heap usage variance required for an individual parent task before it is considered for cancellation.
	//
	// The heap usage variance required for an individual parent task before it is considered for cancellation. A task is considered for cancellation when taskHeapUsage is greater than or equal to heapUsageMovingAverage * variance. Default is 2.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.4/docs/resources/managed_database_opensearch#heap_variance ManagedDatabaseOpensearch#heap_variance}
	HeapVariance *float64 `field:"optional" json:"heapVariance" yaml:"heapVariance"`
	// The heap usage threshold (as a percentage) required for the sum of heap usages of all search tasks before cancellation is applied.
	//
	// The heap usage threshold (as a percentage) required for the sum of heap usages of all search tasks before cancellation is applied. Default is 0.5.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/upcloudltd/upcloud/5.23.4/docs/resources/managed_database_opensearch#total_heap_percent_threshold ManagedDatabaseOpensearch#total_heap_percent_threshold}
	TotalHeapPercentThreshold *float64 `field:"optional" json:"totalHeapPercentThreshold" yaml:"totalHeapPercentThreshold"`
}

