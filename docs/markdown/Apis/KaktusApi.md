# KaktusApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createKaktus**](KaktusApi.md#createKaktus) | **POST** /zone/{zoneId}/kaktus |  |
| [**deleteKaktus**](KaktusApi.md#deleteKaktus) | **DELETE** /kaktus/{kaktusId} |  |
| [**listKaktusInstances**](KaktusApi.md#listKaktusInstances) | **GET** /kaktus/{kaktusId}/instances |  |
| [**listKaktuss**](KaktusApi.md#listKaktuss) | **GET** /kaktus |  |
| [**listZoneKaktuss**](KaktusApi.md#listZoneKaktuss) | **GET** /zone/{zoneId}/kaktuss |  |
| [**readKaktus**](KaktusApi.md#readKaktus) | **GET** /kaktus/{kaktusId} |  |
| [**readKaktusCaps**](KaktusApi.md#readKaktusCaps) | **GET** /kaktus/{kaktusId}/caps |  |
| [**updateKaktus**](KaktusApi.md#updateKaktus) | **PUT** /kaktus/{kaktusId} |  |


<a name="createKaktus"></a>
# **createKaktus**
> Kaktus createKaktus(zoneId, Kaktus)



    Creates a new Kaktus computing node.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |
| **Kaktus** | [**Kaktus**](../Models/Kaktus.md)| Kaktus payload. | |

### Return type

[**Kaktus**](../Models/Kaktus.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteKaktus"></a>
# **deleteKaktus**
> deleteKaktus(kaktusId)



    Deletes an existing Kaktus computing node.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kaktusId** | **String**| The ID of the Kaktus computing node. | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listKaktusInstances"></a>
# **listKaktusInstances**
> List listKaktusInstances(kaktusId)



    Returns the IDs of virtual machine instance objects.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kaktusId** | **String**| The ID of the Kaktus computing node. | [default to null] |

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listKaktuss"></a>
# **listKaktuss**
> List listKaktuss()



    Returns the IDs of Kaktus computing node objects.

### Parameters
This endpoint does not need any parameter.

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listZoneKaktuss"></a>
# **listZoneKaktuss**
> List listZoneKaktuss(zoneId)



    Returns the IDs of Kaktus computing node objects.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="readKaktus"></a>
# **readKaktus**
> Kaktus readKaktus(kaktusId)



    Returns a Kaktus computing node.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kaktusId** | **String**| The ID of the Kaktus computing node. | [default to null] |

### Return type

[**Kaktus**](../Models/Kaktus.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="readKaktusCaps"></a>
# **readKaktusCaps**
> KaktusCaps readKaktusCaps(kaktusId)



    Returns a Kaktus computing node capability.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kaktusId** | **String**| The ID of the Kaktus computing node. | [default to null] |

### Return type

[**KaktusCaps**](../Models/KaktusCaps.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateKaktus"></a>
# **updateKaktus**
> Kaktus updateKaktus(kaktusId, Kaktus)



    Updates a Kaktus computing node configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kaktusId** | **String**| The ID of the Kaktus computing node. | [default to null] |
| **Kaktus** | [**Kaktus**](../Models/Kaktus.md)| Kaktus payload. | |

### Return type

[**Kaktus**](../Models/Kaktus.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

