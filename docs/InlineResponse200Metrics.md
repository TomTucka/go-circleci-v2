# InlineResponse200Metrics

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuccessRate** | **float32** | The ratio of successful runs / total runs. | [default to null]
**TotalRuns** | **int64** | The total number of runs. | [default to null]
**FailedRuns** | **int64** | The number of failed runs. | [default to null]
**SuccessfulRuns** | **int64** | The number of successful runs. | [default to null]
**Throughput** | **float32** | The average number of workflow runs per day. | [default to null]
**Mttr** | **int64** | The mean time to recovery (mean time between failures and their next success) in seconds. | [default to null]
**TotalCreditsUsed** | **int64** | The total credits consumed by the workflow in the aggregation window. | [default to null]
**DurationMetrics** | [***InlineResponse200MetricsDurationMetrics**](inline_response_200_metrics_duration_metrics.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

