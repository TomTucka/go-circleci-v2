# JobDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebUrl** | **string** | URL of the job in CircleCI Web UI. | [default to null]
**Project** | [***JobDetailsProject**](Job Details_project.md) |  | [default to null]
**ParallelRuns** | [**[]JobDetailsParallelRuns**](Job Details_parallel_runs.md) | Info about parallels runs and their status. | [default to null]
**StartedAt** | [**time.Time**](time.Time.md) | The date and time the job started. | [default to null]
**LatestWorkflow** | [***JobDetailsLatestWorkflow**](Job Details_latest_workflow.md) |  | [default to null]
**Name** | **string** | The name of the job. | [default to null]
**Executor** | [***JobDetailsExecutor**](Job Details_executor.md) |  | [default to null]
**Parallelism** | **int64** | A number of parallel runs the job has. | [default to null]
**Status** | [***Object**](.md) | The current status of the job. | [default to null]
**Number** | **int64** | The number of the job. | [default to null]
**Pipeline** | [***JobDetailsPipeline**](Job Details_pipeline.md) |  | [default to null]
**Duration** | **int64** | Duration of a job in milliseconds. | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The time when the job was created. | [default to null]
**Messages** | [**[]JobDetailsMessages**](Job Details_messages.md) | Messages from CircleCI execution platform. | [default to null]
**Contexts** | [**[]JobDetailsContexts**](Job Details_contexts.md) | List of contexts used by the job. | [default to null]
**Organization** | [***JobDetailsOrganization**](Job Details_organization.md) |  | [default to null]
**QueuedAt** | [**time.Time**](time.Time.md) | The time when the job was placed in a queue. | [default to null]
**StoppedAt** | [**time.Time**](time.Time.md) | The time when the job stopped. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

