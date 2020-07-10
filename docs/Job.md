# Job

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanceledBy** | **string** | The unique ID of the user. | [optional] [default to null]
**Dependencies** | **[]string** | A sequence of the unique job IDs for the jobs that this job depends upon in the workflow. | [default to null]
**JobNumber** | **int64** | The number of the job. | [optional] [default to null]
**Id** | **string** | The unique ID of the job. | [default to null]
**StartedAt** | [**time.Time**](time.Time.md) | The date and time the job started. | [default to null]
**Name** | **string** | The name of the job. | [default to null]
**ApprovedBy** | **string** | The unique ID of the user. | [optional] [default to null]
**ProjectSlug** | **string** | The project-slug for the job. | [default to null]
**Status** | [***Object**](.md) | The current status of the job. | [default to null]
**Type_** | **string** | The type of job. | [default to null]
**StoppedAt** | [**time.Time**](time.Time.md) | The time when the job stopped. | [optional] [default to null]
**ApprovalRequestId** | **string** | The unique ID of the job. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

