# {{classname}}

All URIs are relative to *https://circleci.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPipelineById**](PipelineApi.md#GetPipelineById) | **Get** /pipeline/{pipeline-id} | Get a pipeline
[**GetPipelineByNumber**](PipelineApi.md#GetPipelineByNumber) | **Get** /project/{project-slug}/pipeline/{pipeline-number} | Get a pipeline
[**GetPipelineConfigById**](PipelineApi.md#GetPipelineConfigById) | **Get** /pipeline/{pipeline-id}/config | Get a pipeline&#x27;s configuration
[**ListMyPipelines**](PipelineApi.md#ListMyPipelines) | **Get** /project/{project-slug}/pipeline/mine | Get your pipelines
[**ListPipelines**](PipelineApi.md#ListPipelines) | **Get** /pipeline | Get a list of pipelines
[**ListPipelinesForProject**](PipelineApi.md#ListPipelinesForProject) | **Get** /project/{project-slug}/pipeline | Get all pipelines
[**ListWorkflowsByPipelineId**](PipelineApi.md#ListWorkflowsByPipelineId) | **Get** /pipeline/{pipeline-id}/workflow | Get a pipeline&#x27;s workflows
[**TriggerPipeline**](PipelineApi.md#TriggerPipeline) | **Post** /project/{project-slug}/pipeline | Trigger a new pipeline

# **GetPipelineById**
> Pipeline GetPipelineById(ctx, pipelineId)
Get a pipeline

Returns a pipeline by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pipelineId** | [**string**](.md)| The unique ID of the pipeline. | 

### Return type

[**Pipeline**](Pipeline.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipelineByNumber**
> Pipeline GetPipelineByNumber(ctx, projectSlug, pipelineNumber)
Get a pipeline

Returns a pipeline by number.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 
  **pipelineNumber** | [**Object**](.md)| The number of the pipeline. | 

### Return type

[**Pipeline**](Pipeline.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipelineConfigById**
> PipelineConfig GetPipelineConfigById(ctx, pipelineId)
Get a pipeline's configuration

Returns a pipeline's configuration by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pipelineId** | [**string**](.md)| The unique ID of the pipeline. | 

### Return type

[**PipelineConfig**](PipelineConfig.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMyPipelines**
> PipelineListResponse ListMyPipelines(ctx, projectSlug, optional)
Get your pipelines

Returns a sequence of all pipelines for this project triggered by the user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 
 **optional** | ***PipelineApiListMyPipelinesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineApiListMyPipelinesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **optional.String**| A token to retrieve the next page of results. | 

### Return type

[**PipelineListResponse**](PipelineListResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPipelines**
> PipelineListResponse ListPipelines(ctx, orgSlug, mine, optional)
Get a list of pipelines

Returns all pipelines for the most recently built projects (max 250) you follow in an organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgSlug** | **string**| Org slug in the form &#x60;vcs-slug/org-name&#x60; | 
  **mine** | **bool**| Only include entries created by your user. | 
 **optional** | ***PipelineApiListPipelinesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineApiListPipelinesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageToken** | **optional.String**| A token to retrieve the next page of results. | 

### Return type

[**PipelineListResponse**](PipelineListResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPipelinesForProject**
> PipelineListResponse ListPipelinesForProject(ctx, projectSlug, optional)
Get all pipelines

Returns all pipelines for this project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 
 **optional** | ***PipelineApiListPipelinesForProjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineApiListPipelinesForProjectOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **branch** | **optional.String**| The name of a vcs branch. | 
 **pageToken** | **optional.String**| A token to retrieve the next page of results. | 

### Return type

[**PipelineListResponse**](PipelineListResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListWorkflowsByPipelineId**
> WorkflowListResponse ListWorkflowsByPipelineId(ctx, pipelineId, optional)
Get a pipeline's workflows

Returns a paginated list of workflows by pipeline ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pipelineId** | [**string**](.md)| The unique ID of the pipeline. | 
 **optional** | ***PipelineApiListWorkflowsByPipelineIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineApiListWorkflowsByPipelineIdOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **optional.String**| A token to retrieve the next page of results. | 

### Return type

[**WorkflowListResponse**](WorkflowListResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TriggerPipeline**
> PipelineCreation TriggerPipeline(ctx, projectSlug, optional)
Trigger a new pipeline

Triggers a new pipeline on the project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 
 **optional** | ***PipelineApiTriggerPipelineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PipelineApiTriggerPipelineOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of TriggerPipelineParameters**](TriggerPipelineParameters.md)|  | 
 **xAttributionLogin** | **optional.**| The login or user-readable identifier for the pipeline&#x27;s triggerer. Internal use only. | 
 **xAttributionActorId** | **optional.**| The id the integration uses to identify the pipeline&#x27;s triggerer. Internal use only. | 

### Return type

[**PipelineCreation**](PipelineCreation.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

