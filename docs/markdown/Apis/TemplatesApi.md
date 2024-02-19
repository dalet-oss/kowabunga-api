# TemplatesApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**getPoolTemplates**](TemplatesApi.md#getPoolTemplates) | **GET** /pool/{poolId}/templates |  |


<a name="getPoolTemplates"></a>
# **getPoolTemplates**
> List getPoolTemplates(poolId)



    Returns the IDs of the volume templates existing in the storage pool.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **poolId** | **String**| The ID of the storage pool. | [default to null] |

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

