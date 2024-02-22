# KgwApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createProjectZoneKGW**](KgwApi.md#createProjectZoneKGW) | **POST** /project/{projectId}/zone/{zoneId}/kgw |  |
| [**deleteKGW**](KgwApi.md#deleteKGW) | **DELETE** /kgw/{kgwId} |  |
| [**listKGWs**](KgwApi.md#listKGWs) | **GET** /kgw |  |
| [**listProjectZoneKGWs**](KgwApi.md#listProjectZoneKGWs) | **GET** /project/{projectId}/zone/{zoneId}/kgws |  |
| [**readKGW**](KgwApi.md#readKGW) | **GET** /kgw/{kgwId} |  |
| [**updateKGW**](KgwApi.md#updateKGW) | **PUT** /kgw/{kgwId} |  |


<a name="createProjectZoneKGW"></a>
# **createProjectZoneKGW**
> KGW createProjectZoneKGW(projectId, zoneId, KGW)



    Creates a new KGW (Kowabunga Network Gateway).

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |
| **KGW** | [**KGW**](../Models/KGW.md)| KGW payload. | |

### Return type

[**KGW**](../Models/KGW.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listProjectZoneKGWs"></a>
# **listProjectZoneKGWs**
> List listProjectZoneKGWs(projectId, zoneId)



    Returns the IDs of KGW (Kowabunga Network Gateway) objects.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

