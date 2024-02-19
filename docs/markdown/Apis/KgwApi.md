# KgwApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createProjectZoneKgw**](KgwApi.md#createProjectZoneKgw) | **POST** /project/{projectId}/zone/{zoneId}/kgw |  |
| [**deleteKGW**](KgwApi.md#deleteKGW) | **DELETE** /kgw/{kgwId} |  |
| [**getAllKgw**](KgwApi.md#getAllKgw) | **GET** /kgw |  |
| [**getKgw**](KgwApi.md#getKgw) | **GET** /kgw/{kgwId} |  |
| [**getProjectZoneKGWs**](KgwApi.md#getProjectZoneKGWs) | **GET** /project/{projectId}/zone/{zoneId}/kgws |  |
| [**updateKGW**](KgwApi.md#updateKGW) | **PUT** /kgw/{kgwId} |  |


<a name="createProjectZoneKgw"></a>
# **createProjectZoneKgw**
> KGW createProjectZoneKgw(projectId, zoneId, KGW)



    Creates a new KGW in the specified zone.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the network adapter. | [default to null] |
| **zoneId** | **String**| The ID of the zone. | [default to null] |
| **KGW** | [**KGW**](../Models/KGW.md)| KGW payload | |

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



    Deletes an existing KGW gateway.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kgwId** | **String**| The ID of the KGW network gateway. | [default to null] |

### Return type

null (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getAllKgw"></a>
# **getAllKgw**
> List getAllKgw()



    Returns the IDs of registered KGW

### Parameters
This endpoint does not need any parameter.

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getKgw"></a>
# **getKgw**
> KGW getKgw(kgwId)



    Returns the descirption of the registered KGW

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kgwId** | **String**| The ID of the KGW network gateway. | [default to null] |

### Return type

[**KGW**](../Models/KGW.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getProjectZoneKGWs"></a>
# **getProjectZoneKGWs**
> List getProjectZoneKGWs(projectId, zoneId)



    Returns the IDs of the KGW existing in the project in the specified zone.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the network adapter. | [default to null] |
| **zoneId** | **String**| The ID of the zone. | [default to null] |

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateKGW"></a>
# **updateKGW**
> KGW updateKGW(kgwId, KGW)



    Updates a KGW virtual machine configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kgwId** | **String**| The ID of the KGW network gateway. | [default to null] |
| **KGW** | [**KGW**](../Models/KGW.md)| KGW payload | |

### Return type

[**KGW**](../Models/KGW.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

