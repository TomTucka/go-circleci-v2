# {{classname}}

All URIs are relative to *https://circleci.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelJob**](PreviewApi.md#CancelJob) | **Post** /project/{project-slug}/job/{job-number}/cancel | Cancel job
[**CreateCheckoutKey**](PreviewApi.md#CreateCheckoutKey) | **Post** /project/{project-slug}/checkout-key | Create a new checkout key
[**CreateEnvVar**](PreviewApi.md#CreateEnvVar) | **Post** /project/{project-slug}/envvar | Create an environment variable
[**DeleteCheckoutKey**](PreviewApi.md#DeleteCheckoutKey) | **Delete** /project/{project-slug}/checkout-key/{fingerprint} | Delete a checkout key
[**DeleteEnvVar**](PreviewApi.md#DeleteEnvVar) | **Delete** /project/{project-slug}/envvar/{name} | Delete an environment variable
[**GetCheckoutKey**](PreviewApi.md#GetCheckoutKey) | **Get** /project/{project-slug}/checkout-key/{fingerprint} | Get a checkout key
[**GetCollaborations**](PreviewApi.md#GetCollaborations) | **Get** /me/collaborations | Collaborations
[**GetCurrentUser**](PreviewApi.md#GetCurrentUser) | **Get** /me | User Information
[**GetEnvVar**](PreviewApi.md#GetEnvVar) | **Get** /project/{project-slug}/envvar/{name} | Get a masked environment variable
[**GetJobArtifacts**](PreviewApi.md#GetJobArtifacts) | **Get** /project/{project-slug}/{job-number}/artifacts | Get a job&#x27;s artifacts
[**GetJobDetails**](PreviewApi.md#GetJobDetails) | **Get** /project/{project-slug}/job/{job-number} | Get job details
[**GetProjectBySlug**](PreviewApi.md#GetProjectBySlug) | **Get** /project/{project-slug} | Get a project
[**GetProjectJobRuns**](PreviewApi.md#GetProjectJobRuns) | **Get** /insights/{project-slug}/workflows/{workflow-name}/jobs/{job-name} | Get recent runs of a workflow job
[**GetProjectWorkflowJobMetrics**](PreviewApi.md#GetProjectWorkflowJobMetrics) | **Get** /insights/{project-slug}/workflows/{workflow-name}/jobs | Get summary metrics for a project workflow&#x27;s jobs.
[**GetProjectWorkflowMetrics**](PreviewApi.md#GetProjectWorkflowMetrics) | **Get** /insights/{project-slug}/workflows | Get summary metrics for a project&#x27;s workflows
[**GetProjectWorkflowRuns**](PreviewApi.md#GetProjectWorkflowRuns) | **Get** /insights/{project-slug}/workflows/{workflow-name} | Get recent runs of a workflow
[**GetTests**](PreviewApi.md#GetTests) | **Get** /project/{project-slug}/{job-number}/tests | Get test metadata
[**GetUser**](PreviewApi.md#GetUser) | **Get** /user/{id} | User Information
[**ListCheckoutKeys**](PreviewApi.md#ListCheckoutKeys) | **Get** /project/{project-slug}/checkout-key | Get all checkout keys
[**ListEnvVars**](PreviewApi.md#ListEnvVars) | **Get** /project/{project-slug}/envvar | List all environment variables
[**ListPipelines**](PreviewApi.md#ListPipelines) | **Get** /pipeline | Get a list of pipelines

# **CancelJob**
> MessageResponse CancelJob(ctx, jobNumber, projectSlug)
Cancel job

Cancel job with a given job number.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobNumber** | [**Object**](.md)| The number of the job. | 
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCheckoutKey**
> CheckoutKey CreateCheckoutKey(ctx, projectSlug, optional)
Create a new checkout key

Creates a new checkout key. This API request is only usable with a user API token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 
 **optional** | ***PreviewApiCreateCheckoutKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PreviewApiCreateCheckoutKeyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CheckoutKeyInput**](CheckoutKeyInput.md)|  | 

### Return type

[**CheckoutKey**](CheckoutKey.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateEnvVar**
> EnvironmentVariablePair CreateEnvVar(ctx, projectSlug, optional)
Create an environment variable

Creates a new environment variable.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 
 **optional** | ***PreviewApiCreateEnvVarOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PreviewApiCreateEnvVarOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of EnvironmentVariablePair**](EnvironmentVariablePair.md)|  | 

### Return type

[**EnvironmentVariablePair**](EnvironmentVariablePair.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCheckoutKey**
> MessageResponse DeleteCheckoutKey(ctx, projectSlug, fingerprint)
Delete a checkout key

Deletes the checkout key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 
  **fingerprint** | **string**| An SSH key fingerprint. | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEnvVar**
> MessageResponse DeleteEnvVar(ctx, projectSlug, name)
Delete an environment variable

Deletes the environment variable named :name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 
  **name** | **string**| The name of the environment variable. | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCheckoutKey**
> CheckoutKey GetCheckoutKey(ctx, projectSlug, fingerprint)
Get a checkout key

Returns an individual checkout key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 
  **fingerprint** | **string**| An SSH key fingerprint. | 

### Return type

[**CheckoutKey**](CheckoutKey.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCollaborations**
> []Collaboration GetCollaborations(ctx, )
Collaborations

Provides the set of organizations of which a user is a member or a collaborator.  The set of organizations that a user can collaborate on is composed of:  * Organizations that the current user belongs to across VCS types (e.g. BitBucket, GitHub) * The parent organization of repository that the user can collaborate on, but is not necessarily a member of * The organization of the current user's account

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Collaboration**](Collaboration.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCurrentUser**
> User GetCurrentUser(ctx, )
User Information

Provides information about the user that is currently signed in.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**User**](User.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnvVar**
> EnvironmentVariablePair GetEnvVar(ctx, projectSlug, name)
Get a masked environment variable

Returns the masked value of environment variable :name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 
  **name** | **string**| The name of the environment variable. | 

### Return type

[**EnvironmentVariablePair**](EnvironmentVariablePair.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJobArtifacts**
> ArtifactListResponse GetJobArtifacts(ctx, jobNumber, projectSlug)
Get a job's artifacts

Returns a job's artifacts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobNumber** | [**Object**](.md)| The number of the job. | 
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 

### Return type

[**ArtifactListResponse**](ArtifactListResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetJobDetails**
> JobDetails GetJobDetails(ctx, jobNumber, projectSlug)
Get job details

Returns job details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobNumber** | [**Object**](.md)| The number of the job. | 
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 

### Return type

[**JobDetails**](Job Details.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectBySlug**
> Project GetProjectBySlug(ctx, projectSlug)
Get a project

Retrieves a project by project slug.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 

### Return type

[**Project**](Project.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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
 **optional** | ***PreviewApiGetProjectJobRunsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PreviewApiGetProjectJobRunsOpts struct
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
 **optional** | ***PreviewApiGetProjectWorkflowJobMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PreviewApiGetProjectWorkflowJobMetricsOpts struct
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
 **optional** | ***PreviewApiGetProjectWorkflowMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PreviewApiGetProjectWorkflowMetricsOpts struct
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
 **optional** | ***PreviewApiGetProjectWorkflowRunsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PreviewApiGetProjectWorkflowRunsOpts struct
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

# **GetTests**
> TestsResponse GetTests(ctx, jobNumber, projectSlug)
Get test metadata

Get test metadata for a build.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobNumber** | [**Object**](.md)| The number of the job. | 
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 

### Return type

[**TestsResponse**](TestsResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUser**
> User GetUser(ctx, id)
User Information

Provides information about the user with the given ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | [**string**](.md)| The unique ID of the user. | 

### Return type

[**User**](User.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCheckoutKeys**
> CheckoutKeyListResponse ListCheckoutKeys(ctx, projectSlug)
Get all checkout keys

Returns a sequence of checkout keys for `:project`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 

### Return type

[**CheckoutKeyListResponse**](CheckoutKeyListResponse.md)

### Authorization

[api_key_header](../README.md#api_key_header), [api_key_query](../README.md#api_key_query), [basic_auth](../README.md#basic_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEnvVars**
> EnvironmentVariableListResponse ListEnvVars(ctx, projectSlug)
List all environment variables

Returns four 'x' characters, in addition to the last four ASCII characters of the value, consistent with the display of environment variable values on the CircleCI website.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 

### Return type

[**EnvironmentVariableListResponse**](EnvironmentVariableListResponse.md)

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
 **optional** | ***PreviewApiListPipelinesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PreviewApiListPipelinesOpts struct
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

