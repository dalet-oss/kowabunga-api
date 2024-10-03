# KiwiApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createKiwi**](KiwiApi.md#createKiwi) | **POST** /region/{regionId}/kiwi |  |
| [**deleteKiwi**](KiwiApi.md#deleteKiwi) | **DELETE** /kiwi/{kiwiId} |  |
| [**listKiwis**](KiwiApi.md#listKiwis) | **GET** /kiwi |  |
| [**listRegionKiwis**](KiwiApi.md#listRegionKiwis) | **GET** /region/{regionId}/kiwis |  |
| [**readKiwi**](KiwiApi.md#readKiwi) | **GET** /kiwi/{kiwiId} |  |
| [**updateKiwi**](KiwiApi.md#updateKiwi) | **PUT** /kiwi/{kiwiId} |  |


<a name="createKiwi"></a>
# **createKiwi**
> Kiwi createKiwi(regionId, Kiwi)



    Creates a new Kiwi (Kowabunga Inner Wan Interface) provides edge-network services..

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **regionId** | **String**| The ID of the region. | [default to null] |
| **Kiwi** | [**Kiwi**](../Models/Kiwi.md)| Kiwi payload. | |

### Return type

[**Kiwi**](../Models/Kiwi.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteKiwi"></a>
# **deleteKiwi**
> deleteKiwi(kiwiId)



    Deletes an existing Kiwi (Kowabunga Inner Wan Interface) provides edge-network services..

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kiwiId** | **String**| The ID of the Kiwi (Kowabunga Inner Wan Interface) provides edge-network services.. | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listKiwis"></a>
# **listKiwis**
> List listKiwis()



    Returns the IDs of Kiwi (Kowabunga Inner Wan Interface) provides edge-network services. objects.

### Parameters
This endpoint does not need any parameter.

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listRegionKiwis"></a>
# **listRegionKiwis**
> List listRegionKiwis(regionId)



    Returns the IDs of Kiwi (Kowabunga Inner Wan Interface) provides edge-network services. objects.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **regionId** | **String**| The ID of the region. | [default to null] |

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="readKiwi"></a>
# **readKiwi**
> Kiwi readKiwi(kiwiId)



    Returns a Kiwi (Kowabunga Inner Wan Interface) provides edge-network services..

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kiwiId** | **String**| The ID of the Kiwi (Kowabunga Inner Wan Interface) provides edge-network services.. | [default to null] |

### Return type

[**Kiwi**](../Models/Kiwi.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateKiwi"></a>
# **updateKiwi**
> Kiwi updateKiwi(kiwiId, Kiwi)



    Updates a Kiwi (Kowabunga Inner Wan Interface) provides edge-network services. configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kiwiId** | **String**| The ID of the Kiwi (Kowabunga Inner Wan Interface) provides edge-network services.. | [default to null] |
| **Kiwi** | [**Kiwi**](../Models/Kiwi.md)| Kiwi payload. | |

### Return type

[**Kiwi**](../Models/Kiwi.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

