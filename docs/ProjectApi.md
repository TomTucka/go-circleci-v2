# {{classname}}

All URIs are relative to *https://circleci.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCheckoutKey**](ProjectApi.md#CreateCheckoutKey) | **Post** /project/{project-slug}/checkout-key | Create a new checkout key
[**CreateEnvVar**](ProjectApi.md#CreateEnvVar) | **Post** /project/{project-slug}/envvar | Create an environment variable
[**DeleteCheckoutKey**](ProjectApi.md#DeleteCheckoutKey) | **Delete** /project/{project-slug}/checkout-key/{fingerprint} | Delete a checkout key
[**DeleteEnvVar**](ProjectApi.md#DeleteEnvVar) | **Delete** /project/{project-slug}/envvar/{name} | Delete an environment variable
[**GetCheckoutKey**](ProjectApi.md#GetCheckoutKey) | **Get** /project/{project-slug}/checkout-key/{fingerprint} | Get a checkout key
[**GetEnvVar**](ProjectApi.md#GetEnvVar) | **Get** /project/{project-slug}/envvar/{name} | Get a masked environment variable
[**GetProjectBySlug**](ProjectApi.md#GetProjectBySlug) | **Get** /project/{project-slug} | Get a project
[**ListCheckoutKeys**](ProjectApi.md#ListCheckoutKeys) | **Get** /project/{project-slug}/checkout-key | Get all checkout keys
[**ListEnvVars**](ProjectApi.md#ListEnvVars) | **Get** /project/{project-slug}/envvar | List all environment variables

# **CreateCheckoutKey**
> CheckoutKey CreateCheckoutKey(ctx, projectSlug, optional)
Create a new checkout key

Creates a new checkout key. This API request is only usable with a user API token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectSlug** | **string**| Project slug in the form &#x60;vcs-slug/org-name/repo-name&#x60;. The &#x60;/&#x60; characters may be URL-escaped. | 
 **optional** | ***ProjectApiCreateCheckoutKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiCreateCheckoutKeyOpts struct
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
 **optional** | ***ProjectApiCreateEnvVarOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectApiCreateEnvVarOpts struct
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

