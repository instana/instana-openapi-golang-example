# \ApplicationSettingsApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddApplicationConfig**](ApplicationSettingsApi.md#AddApplicationConfig) | **Post** /api/application-monitoring/settings/application | Add application configuration
[**AddEndpointConfig**](ApplicationSettingsApi.md#AddEndpointConfig) | **Post** /api/application-monitoring/settings/service | Add service configuration
[**CreateEndpointConfig**](ApplicationSettingsApi.md#CreateEndpointConfig) | **Post** /api/application-monitoring/settings/http-endpoint | Create endpoint configuration
[**DeleteApplicationConfig**](ApplicationSettingsApi.md#DeleteApplicationConfig) | **Delete** /api/application-monitoring/settings/application/{id} | Delete application configuration
[**DeleteEndpointConfig**](ApplicationSettingsApi.md#DeleteEndpointConfig) | **Delete** /api/application-monitoring/settings/http-endpoint/{id} | Delete endpoint configuration
[**DeleteEndpointConfig1**](ApplicationSettingsApi.md#DeleteEndpointConfig1) | **Delete** /api/application-monitoring/settings/service/{id} | Delete service configuration
[**GetApplicationConfig**](ApplicationSettingsApi.md#GetApplicationConfig) | **Get** /api/application-monitoring/settings/application/{id} | Application configuration
[**GetApplicationConfigs**](ApplicationSettingsApi.md#GetApplicationConfigs) | **Get** /api/application-monitoring/settings/application | All Application configurations
[**GetEndpointConfig**](ApplicationSettingsApi.md#GetEndpointConfig) | **Get** /api/application-monitoring/settings/http-endpoint/{id} | Endpoint configuration
[**GetEndpointConfigs**](ApplicationSettingsApi.md#GetEndpointConfigs) | **Get** /api/application-monitoring/settings/http-endpoint | All Endpoint configurations
[**GetServiceConfig**](ApplicationSettingsApi.md#GetServiceConfig) | **Get** /api/application-monitoring/settings/service/{id} | Service configuration
[**GetServiceConfigs**](ApplicationSettingsApi.md#GetServiceConfigs) | **Get** /api/application-monitoring/settings/service | All service configurations
[**OrderEndpointConfig**](ApplicationSettingsApi.md#OrderEndpointConfig) | **Put** /api/application-monitoring/settings/service/order | Order of service configuration
[**PutApplicationConfig**](ApplicationSettingsApi.md#PutApplicationConfig) | **Put** /api/application-monitoring/settings/application/{id} | Update application configuration
[**PutEndpointConfig**](ApplicationSettingsApi.md#PutEndpointConfig) | **Put** /api/application-monitoring/settings/service/{id} | Update service configuration
[**UpdateEndpointConfig**](ApplicationSettingsApi.md#UpdateEndpointConfig) | **Put** /api/application-monitoring/settings/http-endpoint/{id} | Update endpoint configuration



## AddApplicationConfig

> ApplicationConfig AddApplicationConfig(ctx, applicationConfig)
Add application configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationConfig** | [**ApplicationConfig**](ApplicationConfig.md)|  | 

### Return type

[**ApplicationConfig**](ApplicationConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddEndpointConfig

> ServiceConfig AddEndpointConfig(ctx, serviceConfig)
Add service configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceConfig** | [**ServiceConfig**](ServiceConfig.md)|  | 

### Return type

[**ServiceConfig**](ServiceConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEndpointConfig

> HttpEndpointConfig CreateEndpointConfig(ctx, httpEndpointConfig)
Create endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**httpEndpointConfig** | [**HttpEndpointConfig**](HttpEndpointConfig.md)|  | 

### Return type

[**HttpEndpointConfig**](HttpEndpointConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicationConfig

> DeleteApplicationConfig(ctx, id)
Delete application configuration

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


## DeleteEndpointConfig

> DeleteEndpointConfig(ctx, id)
Delete endpoint configuration

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


## DeleteEndpointConfig1

> DeleteEndpointConfig1(ctx, id)
Delete service configuration

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


## GetApplicationConfig

> ApplicationConfig GetApplicationConfig(ctx, id)
Application configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**ApplicationConfig**](ApplicationConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationConfigs

> []ApplicationConfig GetApplicationConfigs(ctx, )
All Application configurations

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]ApplicationConfig**](ApplicationConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpointConfig

> HttpEndpointConfig GetEndpointConfig(ctx, id)
Endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**HttpEndpointConfig**](HttpEndpointConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEndpointConfigs

> []HttpEndpointConfig GetEndpointConfigs(ctx, )
All Endpoint configurations

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]HttpEndpointConfig**](HttpEndpointConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceConfig

> ServiceConfig GetServiceConfig(ctx, id)
Service configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**ServiceConfig**](ServiceConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceConfigs

> []ServiceConfig GetServiceConfigs(ctx, )
All service configurations

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]ServiceConfig**](ServiceConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderEndpointConfig

> OrderEndpointConfig(ctx, requestBody)
Order of service configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestBody** | [**[]string**](string.md)|  | 

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


## PutApplicationConfig

> ApplicationConfig PutApplicationConfig(ctx, id, applicationConfig)
Update application configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**applicationConfig** | [**ApplicationConfig**](ApplicationConfig.md)|  | 

### Return type

[**ApplicationConfig**](ApplicationConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutEndpointConfig

> ServiceConfig PutEndpointConfig(ctx, id, serviceConfig)
Update service configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**serviceConfig** | [**ServiceConfig**](ServiceConfig.md)|  | 

### Return type

[**ServiceConfig**](ServiceConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEndpointConfig

> HttpEndpointConfig UpdateEndpointConfig(ctx, id, httpEndpointConfig)
Update endpoint configuration

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**httpEndpointConfig** | [**HttpEndpointConfig**](HttpEndpointConfig.md)|  | 

### Return type

[**HttpEndpointConfig**](HttpEndpointConfig.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

