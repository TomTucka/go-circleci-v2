# {{classname}}

All URIs are relative to *https://circleci.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApprovePendingApprovalJobById**](WorkflowApi.md#ApprovePendingApprovalJobById) | **Post** /workflow/{id}/approve/{approval_request_id} | Approve a job
[**CancelWorkflow**](WorkflowApi.md#CancelWorkflow) | **Post** /workflow/{id}/cancel | Cancel a workflow
[**GetWorkflowById**](WorkflowApi.md#GetWorkflowById) | **Get** /workflow/{id} | Get a workflow
[**ListWorkflowJobs**](WorkflowApi.md#ListWorkflowJobs) | **Get** /workflow/{id}/job | Get a workflow&#x27;s jobs
[**RerunWorkflow**](WorkflowApi.md#RerunWorkflow) | **Post** /workflow/{id}/rerun | Rerun a workflow

# **ApprovePendingApprovalJobById**
> MessageResponse ApprovePendingApprovalJobById(ctx, approvalRequestId, id)
Approve a job

Approves a pending approval job in a workflow.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **approvalRequestId** | [**string**](.md)| The ID of the job being approved. | 
  **id** | [**string**](.md)| The unique ID of the workflow. | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelWorkflow**
> MessageResponse CancelWorkflow(ctx, id)
Cancel a workflow

Cancels a running workflow.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The unique ID of the workflow. | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWorkflowById**
> Workflow GetWorkflowById(ctx, id)
Get a workflow

Returns summary fields of a workflow by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The unique ID of the workflow. | 

### Return type

[**Workflow**](Workflow.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListWorkflowJobs**
> WorkflowJobListResponse ListWorkflowJobs(ctx, id)
Get a workflow's jobs

Returns a sequence of jobs for a workflow.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The unique ID of the workflow. | 

### Return type

[**WorkflowJobListResponse**](WorkflowJobListResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RerunWorkflow**
> MessageResponse RerunWorkflow(ctx, id, optional)
Rerun a workflow

Reruns a workflow.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The unique ID of the workflow. | 
 **optional** | ***WorkflowApiRerunWorkflowOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowApiRerunWorkflowOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of RerunWorkflowParameters**](RerunWorkflowParameters.md)|  | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

