# PipelineListResponseVcs

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderName** | **string** | Name of the VCS provider (e.g. GitHub, Bitbucket). | [default to null]
**OriginRepositoryUrl** | **string** | URL for the repository where the trigger originated. For fork-PR pipelines, this is the URL to the fork. For other pipelines the &#x60;origin_&#x60; and &#x60;target_repository_url&#x60;s will be the same. | [default to null]
**TargetRepositoryUrl** | **string** | URL for the repository the trigger targets (i.e. the repository where the PR will be merged). For fork-PR pipelines, this is the URL to the parent repo. For other pipelines, the &#x60;origin_&#x60; and &#x60;target_repository_url&#x60;s will be the same. | [default to null]
**Revision** | **string** | The code revision the pipeline ran. | [default to null]
**Branch** | **string** | The branch where the pipeline ran. The HEAD commit on this branch was used for the pipeline. Note that &#x60;branch&#x60; and &#x60;tag&#x60; are mutually exclusive. | [optional] [default to null]
**Tag** | **string** | The tag used by the pipeline. The commit that this tag points to was used for the pipeline. Note that &#x60;branch&#x60; and &#x60;tag&#x60; are mutually exclusive. | [optional] [default to null]
**Commit** | [***PipelineListResponseVcsCommit**](PipelineListResponse_vcs_commit.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

