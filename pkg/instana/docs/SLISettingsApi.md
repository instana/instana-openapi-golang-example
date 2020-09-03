# \SLISettingsApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSli**](SLISettingsApi.md#CreateSli) | **Post** /api/settings/sli | Create SLI Config
[**Delete**](SLISettingsApi.md#Delete) | **Delete** /api/settings/sli/{id} | Delete SLI Config
[**GetSli1**](SLISettingsApi.md#GetSli1) | **Get** /api/settings/sli/{id} | Get SLI Config
[**GetSli2**](SLISettingsApi.md#GetSli2) | **Get** /api/settings/sli | Get All SLI Configs
[**PutSli**](SLISettingsApi.md#PutSli) | **Put** /api/settings/sli/{id} | Update SLI Config



## CreateSli

> []SliConfigurationWithLastUpdated CreateSli(ctx, sliConfiguration)

Create SLI Config

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sliConfiguration** | [**SliConfiguration**](SliConfiguration.md)|  | 

### Return type

[**[]SliConfigurationWithLastUpdated**](SliConfigurationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, id)

Delete SLI Config

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSli1

> SliConfigurationWithLastUpdated GetSli1(ctx, id)

Get SLI Config

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**SliConfigurationWithLastUpdated**](SliConfigurationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSli2

> []SliConfigurationWithLastUpdated GetSli2(ctx, )

Get All SLI Configs

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]SliConfigurationWithLastUpdated**](SliConfigurationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSli

> []SliConfigurationWithLastUpdated PutSli(ctx, id, sliConfigurationWithLastUpdated)

Update SLI Config

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**sliConfigurationWithLastUpdated** | [**SliConfigurationWithLastUpdated**](SliConfigurationWithLastUpdated.md)|  | 

### Return type

[**[]SliConfigurationWithLastUpdated**](SliConfigurationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

