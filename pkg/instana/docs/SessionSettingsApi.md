# \SessionSettingsApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSessionSettings**](SessionSettingsApi.md#DeleteSessionSettings) | **Delete** /api/settings/session | Delete session settings
[**GetSessionSettings**](SessionSettingsApi.md#GetSessionSettings) | **Get** /api/settings/session | Session settings
[**SetSessionSettings**](SessionSettingsApi.md#SetSessionSettings) | **Put** /api/settings/session | Configure session settings



## DeleteSessionSettings

> DeleteSessionSettings(ctx, )

Delete session settings

### Required Parameters

This endpoint does not need any parameter.

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


## GetSessionSettings

> SessionSettings GetSessionSettings(ctx, )

Session settings

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**SessionSettings**](SessionSettings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSessionSettings

> SessionSettings SetSessionSettings(ctx, optional)

Configure session settings

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SetSessionSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SetSessionSettingsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sessionSettings** | [**optional.Interface of SessionSettings**](SessionSettings.md)|  | 

### Return type

[**SessionSettings**](SessionSettings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

