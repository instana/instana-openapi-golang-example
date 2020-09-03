# \SLIReportApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSli**](SLIReportApi.md#GetSli) | **Get** /api/sli/report/{sliId} | Generate SLI report



## GetSli

> []SliReport GetSli(ctx, sliId, optional)

Generate SLI report

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sliId** | **string**|  | 
 **optional** | ***GetSliOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetSliOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **slo** | **optional.Float64**|  | 
 **from** | **optional.Int64**|  | 
 **to** | **optional.Int64**|  | 

### Return type

[**[]SliReport**](SliReport.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

