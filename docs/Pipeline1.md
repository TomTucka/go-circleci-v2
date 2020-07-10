# Pipeline1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID of the pipeline. | [default to null]
**Errors** | [**[]PipelineListResponseErrors**](PipelineListResponse_errors.md) | A sequence of errors that have occurred within the pipeline. | [default to null]
**ProjectSlug** | **string** | The project-slug for the pipeline. | [default to null]
**UpdatedAt** | [**time.Time**](time.Time.md) | The date and time the pipeline was last updated. | [optional] [default to null]
**Number** | **int64** | The number of the pipeline. | [default to null]
**State** | **string** | The current state of the pipeline. | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | The date and time the pipeline was created. | [default to null]
**Trigger** | [***PipelineListResponseTrigger**](PipelineListResponse_trigger.md) |  | [default to null]
**Vcs** | [***PipelineListResponseVcs**](PipelineListResponse_vcs.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

