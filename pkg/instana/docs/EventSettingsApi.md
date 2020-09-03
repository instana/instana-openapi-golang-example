# \EventSettingsApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebsiteAlertConfig**](EventSettingsApi.md#CreateWebsiteAlertConfig) | **Post** /api/events/settings/website-alert-configs | Create Website Alert Config
[**DeleteAlert**](EventSettingsApi.md#DeleteAlert) | **Delete** /api/events/settings/alerts/{id} | Delete alerting
[**DeleteAlertingChannel**](EventSettingsApi.md#DeleteAlertingChannel) | **Delete** /api/events/settings/alertingChannels/{id} | Delete alerting channel
[**DeleteBuiltInEventSpecification**](EventSettingsApi.md#DeleteBuiltInEventSpecification) | **Delete** /api/events/settings/event-specifications/built-in/{eventSpecificationId} | Delete built-in event specification
[**DeleteCustomEventSpecification**](EventSettingsApi.md#DeleteCustomEventSpecification) | **Delete** /api/events/settings/event-specifications/custom/{eventSpecificationId} | Delete custom event specification
[**DeleteWebsiteAlertConfig**](EventSettingsApi.md#DeleteWebsiteAlertConfig) | **Delete** /api/events/settings/website-alert-configs/{id} | Delete Website Alert Config
[**DisableBuiltInEventSpecification**](EventSettingsApi.md#DisableBuiltInEventSpecification) | **Post** /api/events/settings/event-specifications/built-in/{eventSpecificationId}/disable | Disable built-in event specification
[**DisableCustomEventSpecification**](EventSettingsApi.md#DisableCustomEventSpecification) | **Post** /api/events/settings/event-specifications/custom/{eventSpecificationId}/disable | Disable custom event specification
[**DisableWebsiteAlertConfig**](EventSettingsApi.md#DisableWebsiteAlertConfig) | **Put** /api/events/settings/website-alert-configs/{id}/disable | Disable Website Alert Config
[**EnableBuiltInEventSpecification**](EventSettingsApi.md#EnableBuiltInEventSpecification) | **Post** /api/events/settings/event-specifications/built-in/{eventSpecificationId}/enable | Enable built-in event specification
[**EnableCustomEventSpecification**](EventSettingsApi.md#EnableCustomEventSpecification) | **Post** /api/events/settings/event-specifications/custom/{eventSpecificationId}/enable | Enable custom event specification
[**EnableWebsiteAlertConfig**](EventSettingsApi.md#EnableWebsiteAlertConfig) | **Put** /api/events/settings/website-alert-configs/{id}/enable | Enable Website Alert Config
[**FindActiveWebsiteAlertConfigs**](EventSettingsApi.md#FindActiveWebsiteAlertConfigs) | **Get** /api/events/settings/website-alert-configs | All Website Alert Configs
[**FindWebsiteAlertConfig**](EventSettingsApi.md#FindWebsiteAlertConfig) | **Get** /api/events/settings/website-alert-configs/{id} | Get Website Alert Config
[**FindWebsiteAlertConfigVersions**](EventSettingsApi.md#FindWebsiteAlertConfigVersions) | **Get** /api/events/settings/website-alert-configs/{id}/versions | Get versions of Website Alert Config
[**GetAlert**](EventSettingsApi.md#GetAlert) | **Get** /api/events/settings/alerts/{id} | Alerting
[**GetAlertingChannel**](EventSettingsApi.md#GetAlertingChannel) | **Get** /api/events/settings/alertingChannels/{id} | Alerting channel
[**GetAlertingChannels**](EventSettingsApi.md#GetAlertingChannels) | **Get** /api/events/settings/alertingChannels | All alerting channels
[**GetAlertingChannelsOverview**](EventSettingsApi.md#GetAlertingChannelsOverview) | **Get** /api/events/settings/alertingChannels/infos | Overview over all alerting channels
[**GetAlertingConfigurationInfos**](EventSettingsApi.md#GetAlertingConfigurationInfos) | **Get** /api/events/settings/alerts/infos | All alerting configuration info
[**GetAlerts**](EventSettingsApi.md#GetAlerts) | **Get** /api/events/settings/alerts | All Alerting
[**GetBuiltInEventSpecification**](EventSettingsApi.md#GetBuiltInEventSpecification) | **Get** /api/events/settings/event-specifications/built-in/{eventSpecificationId} | Built-in event specifications
[**GetBuiltInEventSpecifications**](EventSettingsApi.md#GetBuiltInEventSpecifications) | **Get** /api/events/settings/event-specifications/built-in | All built-in event specification
[**GetCustomEventSpecification**](EventSettingsApi.md#GetCustomEventSpecification) | **Get** /api/events/settings/event-specifications/custom/{eventSpecificationId} | Custom event specification
[**GetCustomEventSpecifications**](EventSettingsApi.md#GetCustomEventSpecifications) | **Get** /api/events/settings/event-specifications/custom | All custom event specifications
[**GetEventSpecificationInfos**](EventSettingsApi.md#GetEventSpecificationInfos) | **Get** /api/events/settings/event-specifications/infos | Summary of all built-in and custom event specifications
[**GetEventSpecificationInfosByIds**](EventSettingsApi.md#GetEventSpecificationInfosByIds) | **Post** /api/events/settings/event-specifications/infos | All built-in and custom event specifications
[**GetSystemRules**](EventSettingsApi.md#GetSystemRules) | **Get** /api/events/settings/event-specifications/custom/systemRules | All system rules for custom event specifications
[**PutAlert**](EventSettingsApi.md#PutAlert) | **Put** /api/events/settings/alerts/{id} | Update alerting
[**PutAlertingChannel**](EventSettingsApi.md#PutAlertingChannel) | **Put** /api/events/settings/alertingChannels/{id} | Update alerting channel
[**PutCustomEventSpecification**](EventSettingsApi.md#PutCustomEventSpecification) | **Put** /api/events/settings/event-specifications/custom/{eventSpecificationId} | Create or Update custom event specification
[**SendTestAlerting**](EventSettingsApi.md#SendTestAlerting) | **Put** /api/events/settings/alertingChannels/test | Test alerting channel
[**UpdateWebsiteAlertConfig**](EventSettingsApi.md#UpdateWebsiteAlertConfig) | **Post** /api/events/settings/website-alert-configs/{id} | Update Website Alert Config



## CreateWebsiteAlertConfig

> []WebsiteAlertConfigWithMetadata CreateWebsiteAlertConfig(ctx, websiteAlertConfig)

Create Website Alert Config

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**websiteAlertConfig** | [**WebsiteAlertConfig**](WebsiteAlertConfig.md)|  | 

### Return type

[**[]WebsiteAlertConfigWithMetadata**](WebsiteAlertConfigWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAlert

> DeleteAlert(ctx, id)

Delete alerting

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


## DeleteAlertingChannel

> DeleteAlertingChannel(ctx, id)

Delete alerting channel

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


## DeleteBuiltInEventSpecification

> DeleteBuiltInEventSpecification(ctx, eventSpecificationId)

Delete built-in event specification

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventSpecificationId** | **string**|  | 

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


## DeleteCustomEventSpecification

> DeleteCustomEventSpecification(ctx, eventSpecificationId)

Delete custom event specification

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventSpecificationId** | **string**|  | 

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


## DeleteWebsiteAlertConfig

> DeleteWebsiteAlertConfig(ctx, id)

Delete Website Alert Config

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


## DisableBuiltInEventSpecification

> BuiltInEventSpecificationWithLastUpdated DisableBuiltInEventSpecification(ctx, eventSpecificationId, optional)

Disable built-in event specification

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventSpecificationId** | **string**|  | 
 **optional** | ***DisableBuiltInEventSpecificationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DisableBuiltInEventSpecificationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **optional.String**|  | 

### Return type

[**BuiltInEventSpecificationWithLastUpdated**](BuiltInEventSpecificationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableCustomEventSpecification

> CustomEventSpecificationWithLastUpdated DisableCustomEventSpecification(ctx, eventSpecificationId, optional)

Disable custom event specification

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventSpecificationId** | **string**|  | 
 **optional** | ***DisableCustomEventSpecificationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DisableCustomEventSpecificationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **optional.String**|  | 

### Return type

[**CustomEventSpecificationWithLastUpdated**](CustomEventSpecificationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableWebsiteAlertConfig

> []WebsiteAlertConfigWithMetadata DisableWebsiteAlertConfig(ctx, id, optional)

Disable Website Alert Config

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***DisableWebsiteAlertConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DisableWebsiteAlertConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **optional.String**|  | 

### Return type

[**[]WebsiteAlertConfigWithMetadata**](WebsiteAlertConfigWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableBuiltInEventSpecification

> BuiltInEventSpecificationWithLastUpdated EnableBuiltInEventSpecification(ctx, eventSpecificationId, optional)

Enable built-in event specification

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventSpecificationId** | **string**|  | 
 **optional** | ***EnableBuiltInEventSpecificationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EnableBuiltInEventSpecificationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **optional.String**|  | 

### Return type

[**BuiltInEventSpecificationWithLastUpdated**](BuiltInEventSpecificationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableCustomEventSpecification

> CustomEventSpecificationWithLastUpdated EnableCustomEventSpecification(ctx, eventSpecificationId, optional)

Enable custom event specification

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventSpecificationId** | **string**|  | 
 **optional** | ***EnableCustomEventSpecificationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EnableCustomEventSpecificationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **optional.String**|  | 

### Return type

[**CustomEventSpecificationWithLastUpdated**](CustomEventSpecificationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableWebsiteAlertConfig

> []WebsiteAlertConfigWithMetadata EnableWebsiteAlertConfig(ctx, id, optional)

Enable Website Alert Config

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***EnableWebsiteAlertConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a EnableWebsiteAlertConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **optional.String**|  | 

### Return type

[**[]WebsiteAlertConfigWithMetadata**](WebsiteAlertConfigWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindActiveWebsiteAlertConfigs

> []WebsiteAlertConfigWithMetadata FindActiveWebsiteAlertConfigs(ctx, optional)

All Website Alert Configs

Configs are sorted descending by their created date.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FindActiveWebsiteAlertConfigsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FindActiveWebsiteAlertConfigsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **websiteId** | **optional.String**|  | 

### Return type

[**[]WebsiteAlertConfigWithMetadata**](WebsiteAlertConfigWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindWebsiteAlertConfig

> []WebsiteAlertConfigWithMetadata FindWebsiteAlertConfig(ctx, id, optional)

Get Website Alert Config

Find a Website Alert Config by ID. This will deliver deleted configs too.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
 **optional** | ***FindWebsiteAlertConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FindWebsiteAlertConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **validOn** | **optional.Int64**|  | 

### Return type

[**[]WebsiteAlertConfigWithMetadata**](WebsiteAlertConfigWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindWebsiteAlertConfigVersions

> []ConfigVersion FindWebsiteAlertConfigVersions(ctx, id)

Get versions of Website Alert Config

Find all versions of a Website Alert Config by ID. This will deliver deleted configs too. Configs are sorted descending by their created date.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**[]ConfigVersion**](ConfigVersion.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlert

> AlertingConfigurationWithLastUpdated GetAlert(ctx, id)

Alerting

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**AlertingConfigurationWithLastUpdated**](AlertingConfigurationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertingChannel

> AbstractIntegration GetAlertingChannel(ctx, id)

Alerting channel

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 

### Return type

[**AbstractIntegration**](AbstractIntegration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertingChannels

> []AbstractIntegration GetAlertingChannels(ctx, optional)

All alerting channels

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAlertingChannelsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAlertingChannelsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**[]AbstractIntegration**](AbstractIntegration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertingChannelsOverview

> []IntegrationOverview GetAlertingChannelsOverview(ctx, optional)

Overview over all alerting channels

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAlertingChannelsOverviewOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAlertingChannelsOverviewOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**[]IntegrationOverview**](IntegrationOverview.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertingConfigurationInfos

> []ValidatedAlertingChannelInputInfo GetAlertingConfigurationInfos(ctx, optional)

All alerting configuration info

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAlertingConfigurationInfosOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAlertingConfigurationInfosOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **integrationId** | **optional.String**|  | 

### Return type

[**[]ValidatedAlertingChannelInputInfo**](ValidatedAlertingChannelInputInfo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlerts

> []ValidatedAlertingConfiguration GetAlerts(ctx, )

All Alerting

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]ValidatedAlertingConfiguration**](ValidatedAlertingConfiguration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBuiltInEventSpecification

> BuiltInEventSpecification GetBuiltInEventSpecification(ctx, eventSpecificationId)

Built-in event specifications

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventSpecificationId** | **string**|  | 

### Return type

[**BuiltInEventSpecification**](BuiltInEventSpecification.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBuiltInEventSpecifications

> []BuiltInEventSpecificationWithLastUpdated GetBuiltInEventSpecifications(ctx, optional)

All built-in event specification

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetBuiltInEventSpecificationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetBuiltInEventSpecificationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**[]BuiltInEventSpecificationWithLastUpdated**](BuiltInEventSpecificationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomEventSpecification

> CustomEventSpecificationWithLastUpdated GetCustomEventSpecification(ctx, eventSpecificationId)

Custom event specification

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventSpecificationId** | **string**|  | 

### Return type

[**CustomEventSpecificationWithLastUpdated**](CustomEventSpecificationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomEventSpecifications

> []CustomEventSpecificationWithLastUpdated GetCustomEventSpecifications(ctx, )

All custom event specifications

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]CustomEventSpecificationWithLastUpdated**](CustomEventSpecificationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventSpecificationInfos

> []EventSpecificationInfo GetEventSpecificationInfos(ctx, )

Summary of all built-in and custom event specifications

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]EventSpecificationInfo**](EventSpecificationInfo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventSpecificationInfosByIds

> []EventSpecificationInfo GetEventSpecificationInfosByIds(ctx, requestBody)

All built-in and custom event specifications

Summary of all built-in and custom event specifications by IDs

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestBody** | [**[]string**](string.md)|  | 

### Return type

[**[]EventSpecificationInfo**](EventSpecificationInfo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSystemRules

> []SystemRuleLabel GetSystemRules(ctx, )

All system rules for custom event specifications

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]SystemRuleLabel**](SystemRuleLabel.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAlert

> AlertingConfigurationWithLastUpdated PutAlert(ctx, id, alertingConfiguration)

Update alerting

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**alertingConfiguration** | [**AlertingConfiguration**](AlertingConfiguration.md)|  | 

### Return type

[**AlertingConfigurationWithLastUpdated**](AlertingConfigurationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAlertingChannel

> PutAlertingChannel(ctx, id, abstractIntegration)

Update alerting channel

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**abstractIntegration** | [**AbstractIntegration**](AbstractIntegration.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCustomEventSpecification

> CustomEventSpecificationWithLastUpdated PutCustomEventSpecification(ctx, eventSpecificationId, customEventSpecification)

Create or Update custom event specification

This endpoint creates or updates the Custom Event Specification   ## Mandatory Parameters:  - **eventSpecificationId(Path Parameter):** A unique identifier for each custom event  - **id:** Same as the eventSpecificationId  - **name:** Name for the custom event  - **entityType:** Name of tha available plugins for the selected source  - **rules.ruleType:** Type of the rule being set for the custom event  ### Rule-type specific parameters  Depending on the chosen `ruleType`, there are further required parameters:  #### Threshold Rule using a dynamic built-in metric by pattern :  - **rules.conditionOperator:** Conditional operator for the aggregation for the provided time window  - **rules.metricPattern.prefix:** Prefix pattern for the metric  - **rules.metricPattern.operator:** Operator for matching the metric  ``` curl --request PUT 'https://<HOST>/api/events/settings/event-specifications/custom/09876543225' \\ --header 'Authorization: apiToken <Token>' \\ --header 'Content-Type: application/json' \\ --data-raw '{ \"id\" :\"09876543225\", \"description\":\"Event for OpenAPI documentation\", \"enabled\":true,\"entityType\":\"host\",\"expirationTime\":\"60000\",\"name\":\"Event for OpenAPI documentation\", \"query\":<Query>,  \"rules\":[{\"aggregation\":\"sum\",\"conditionOperator\":\">\", \"conditionValue\":0.1, \"metricName\":null, \"metricPattern\":{\"prefix\":\"fs\", \"postfix\":\"free\", \"operator\":\"endsWith\", \"placeholder\":\"/xvda1\"}, \"rollup\":null, \"ruleType\":\"threshold\", \"severity\":10, \"window\":30000}], \"triggering\":false }' ``` The above example creates a custom event that matches disk devices that end with \"/xvda1\" for the metric \"fs.{device}.free\" for any host in scope.  #### Threshold Rule using fixed metric :  - **rules.conditionOperator:** Conditional operator for the aggregation for the provided time window  - **rules.metricName:** Metric name for the event  ``` curl --request PUT 'https://<Host>/api/events/settings/event-specifications/custom/09876543226' \\ --header 'Authorization: apiToken <Token>' \\ --header 'Content-Type: application/json' \\ --data-raw '{ \"id\" :\"09876543226\", \"description\":\"Event for OpenAPI documentation fixed Metric\", \"enabled\":true,\"entityType\":\"host\",\"expirationTime\":\"60000\", \"name\":\"Event for OpenAPI documentation fixed metric\",\"rules\":[{\"aggregation\":\"sum\",\"conditionOperator\":\">\", \"conditionValue\":0.1, \"metricName\":\"fs./dev/xvda1.free\",  \"rollup\":null, \"ruleType\":\"threshold\", \"severity\":10, \"window\":30000}], \"triggering\":false }' ```  #### System Rule:  - **rules.systemRuleId:** Id of the System Rule being set   ``` curl --request PUT 'https://<Host>/api/events/settings/event-specifications/custom/09876543227' \\ --header 'Authorization: apiToken <Token>' \\ --header 'Content-Type: application/json' \\ --data-raw '{ \"id\" :\"09876543227\", \"description\":\"Event for OpenAPI documentation System Rule\", \"enabled\":true,\"entityType\":\"any\",\"expirationTime\":\"60000\", \"name\":\"Event for OpenAPI documentation System Rule\", \"rules\":[{\"ruleType\":\"system\", \"systemRuleId\":\"entity.offline\",\"severity\":10}], \"triggering\":false }' ```  #### Entity Verification Rule:  - **rules.matchingEntityType:** Type of the Entity - **rules.matchingOperator:** Operator for matching the Entity name - **rules.matchingEntityLabel:** Name Pattern for the Entity  ``` curl --request PUT 'https://<Host>/api/events/settings/event-specifications/custom/09876543228' \\ --header 'Authorization: apiToken <Token>' \\ --header 'Content-Type: application/json' \\ --data-raw '{ \"id\" :\"09876543228\", \"description\":\"Event for OpenAPI Entity Verification Rule\", \"enabled\":true,\"entityType\":\"host\",\"expirationTime\":\"60000\", \"name\":\"Event for OpenAPI Entity Verification Rule\", \"rules\":[{\"matchingEntityLabel\":\"test\", \"matchingEntityType\":\"jvmRuntimePlatform\",\"matchingOperator\":\"startsWith\",\"offlineDuration\":1800000,  \"ruleType\":\"entity_verification\",\"severity\": 5}], \"triggering\":false }' `` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventSpecificationId** | **string**|  | 
**customEventSpecification** | [**CustomEventSpecification**](CustomEventSpecification.md)|  | 

### Return type

[**CustomEventSpecificationWithLastUpdated**](CustomEventSpecificationWithLastUpdated.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendTestAlerting

> SendTestAlerting(ctx, abstractIntegration)

Test alerting channel

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**abstractIntegration** | [**AbstractIntegration**](AbstractIntegration.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWebsiteAlertConfig

> []WebsiteAlertConfigWithMetadata UpdateWebsiteAlertConfig(ctx, id, websiteAlertConfig)

Update Website Alert Config

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**|  | 
**websiteAlertConfig** | [**WebsiteAlertConfig**](WebsiteAlertConfig.md)|  | 

### Return type

[**[]WebsiteAlertConfigWithMetadata**](WebsiteAlertConfigWithMetadata.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

