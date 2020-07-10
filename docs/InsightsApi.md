# {{classname}}

All URIs are relative to *https://circleci.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProjectJobRuns**](InsightsApi.md#GetProjectJobRuns) | **Get** /insights/{project-slug}/workflows/{workflow-name}/jobs/{job-name} | Get recent runs of a workflow job
[**GetProjectWorkflowJobMetrics**](InsightsApi.md#GetProjectWorkflowJobMetrics) | **Get** /insights/{project-slug}/workflows/{workflow-name}/jobs | Get summary metrics for a project workflow&#x27;s jobs.
[**GetProjectWorkflowMetrics**](InsightsApi.md#GetProjectWorkflowMetrics) | **Get** /insights/{project-slug}/workflows | Get summary metrics for a project&#x27;s workflows
[**GetProjectWorkflowRuns**](InsightsApi.md#GetProjectWorkflowRuns) | **Get** /insights/{project-slug}/workflows/{workflow-name} | Get recent runs of a workflow

# **GetProjectJobRuns**
> InlineResponse2003 GetProjectJobRuns(ctx, projectSlug, workflowName, jobName, startDate, endDate, optional)
Get recent runs of a workflow job

Get recent runs of a job within a workflow. Runs going back at most 90 days are returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 
  **workflowName** | **string**| The name of the workflow. | 
  **jobName** | **string**| The name of the job. | 
  **startDate** | **time.Time**| Include only executions that started at or after this date. This must be specified if an end-date is provided. | 
  **endDate** | **time.Time**| Include only executions that started before this date. This date can be at most 90 days after the start-date. | 
 **optional** | ***InsightsApiGetProjectJobRunsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightsApiGetProjectJobRunsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **branch** | **optional.String**| The name of a vcs branch. | 
 **pageToken** | **optional.String**| A token to retrieve the next page of results. | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectWorkflowJobMetrics**
> InlineResponse2002 GetProjectWorkflowJobMetrics(ctx, projectSlug, workflowName, optional)
Get summary metrics for a project workflow's jobs.

Get summary metrics for a project workflow's jobs. Job runs going back at most 90 days are included in the aggregation window. Metrics are refreshed daily, and thus may not include executions from the last 24 hours.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 
  **workflowName** | **string**| The name of the workflow. | 
 **optional** | ***InsightsApiGetProjectWorkflowJobMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightsApiGetProjectWorkflowJobMetricsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageToken** | **optional.String**| A token to retrieve the next page of results. | 
 **branch** | **optional.String**| The name of a vcs branch. | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectWorkflowMetrics**
> InlineResponse200 GetProjectWorkflowMetrics(ctx, projectSlug, optional)
Get summary metrics for a project's workflows

Get summary metrics for a project's workflows. Workflow runs going back at most 90 days are included in the aggregation window. Metrics are refreshed daily, and thus may not include executions from the last 24 hours.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 
 **optional** | ***InsightsApiGetProjectWorkflowMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightsApiGetProjectWorkflowMetricsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageToken** | **optional.String**| A token to retrieve the next page of results. | 
 **branch** | **optional.String**| The name of a vcs branch. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectWorkflowRuns**
> InlineResponse2001 GetProjectWorkflowRuns(ctx, projectSlug, workflowName, startDate, endDate, optional)
Get recent runs of a workflow

Get recent runs of a workflow. Runs going back at most 90 days are returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 
  **workflowName** | **string**| The name of the workflow. | 
  **startDate** | **time.Time**| Include only executions that started at or after this date. This must be specified if an end-date is provided. | 
  **endDate** | **time.Time**| Include only executions that started before this date. This date can be at most 90 days after the start-date. | 
 **optional** | ***InsightsApiGetProjectWorkflowRunsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InsightsApiGetProjectWorkflowRunsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **branch** | **optional.String**| The name of a vcs branch. | 
 **pageToken** | **optional.String**| A token to retrieve the next page of results. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

