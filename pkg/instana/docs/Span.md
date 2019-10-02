# Span

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] 
**Name** | **string** |  | [optional] 
**Label** | **string** |  | [optional] 
**Start** | **int64** |  | [optional] 
**Duration** | **int64** |  | [optional] 
**ErrorCount** | **int32** |  | [optional] 
**BatchSize** | **int32** |  | [optional] 
**BatchSelfTime** | **int64** |  | [optional] 
**Kind** | **string** |  | 
**Source** | [**SpanRelation**](SpanRelation.md) |  | [optional] 
**Destination** | [**SpanRelation**](SpanRelation.md) |  | [optional] 
**StackTrace** | [**[]StackTraceItem**](StackTraceItem.md) |  | 
**ChildSpans** | [**[]Span**](Span.md) |  | 
**Data** | [**map[string]interface{}**](interface{}.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


