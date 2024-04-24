# KgwApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createProjectRegionKGW**](KgwApi.md#createProjectRegionKGW) | **POST** /project/{projectId}/region/{regionId}/kgw |  |
| [**deleteKGW**](KgwApi.md#deleteKGW) | **DELETE** /kgw/{kgwId} |  |
| [**listKGWs**](KgwApi.md#listKGWs) | **GET** /kgw |  |
| [**listProjectRegionKGWs**](KgwApi.md#listProjectRegionKGWs) | **GET** /project/{projectId}/region/{regionId}/kgws |  |
| [**readKGW**](KgwApi.md#readKGW) | **GET** /kgw/{kgwId} |  |
| [**updateKGW**](KgwApi.md#updateKGW) | **PUT** /kgw/{kgwId} |  |


<a name="createProjectRegionKGW"></a>
# **createProjectRegionKGW**
> KGW createProjectRegionKGW(projectId, regionId, KGW)



    Creates a new KGW (Kowabunga Network Gateway).

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |
| **regionId** | **String**| The ID of the region. | [default to null] |
| **KGW** | [**KGW**](../Models/KGW.md)| KGW payload. | |

### Return type

[**KGW**](../Models/KGW.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteKGW"></a>
# **deleteKGW**
> deleteKGW(kgwId)



    Deletes an existing KGW (Kowabunga Network Gateway).

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kgwId** | **String**| The ID of the KGW (Kowabunga Network Gateway). | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listKGWs"></a>
# **listKGWs**
> List listKGWs()



    Returns the IDs of KGW (Kowabunga Network Gateway) objects.

### Parameters
This endpoint does not need any parameter.

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listProjectRegionKGWs"></a>
# **listProjectRegionKGWs**
> List listProjectRegionKGWs(projectId, regionId)



    Returns the IDs of KGW (Kowabunga Network Gateway) objects.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |
| **regionId** | **String**| The ID of the region. | [default to null] |

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="readKGW"></a>
# **readKGW**
> KGW readKGW(kgwId)



    Returns a KGW (Kowabunga Network Gateway).

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kgwId** | **String**| The ID of the KGW (Kowabunga Network Gateway). | [default to null] |

### Return type

[**KGW**](../Models/KGW.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateKGW"></a>
# **updateKGW**
> KGW updateKGW(kgwId, KGW)



    Updates a KGW (Kowabunga Network Gateway) configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kgwId** | **String**| The ID of the KGW (Kowabunga Network Gateway). | [default to null] |
| **KGW** | [**KGW**](../Models/KGW.md)| KGW payload. | |

### Return type

[**KGW**](../Models/KGW.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

