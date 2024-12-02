# KawaiiIpsecApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createKawaiiIpSec**](KawaiiIpsecApi.md#createKawaiiIpSec) | **POST** /kawaii/{kawaiiId}/ipsec |  |
| [**deleteKawaiiIpSec**](KawaiiIpsecApi.md#deleteKawaiiIpSec) | **DELETE** /kawaii/{kawaiiId}/ipsec/{KawaiiIpSecId} |  |
| [**listKawaiiIpSecs**](KawaiiIpsecApi.md#listKawaiiIpSecs) | **GET** /kawaii/{kawaiiId}/ipsec |  |
| [**readKawaiiIpSec**](KawaiiIpsecApi.md#readKawaiiIpSec) | **GET** /kawaii/{kawaiiId}/ipsec/{KawaiiIpSecId} |  |
| [**updateKawaiiIpSec**](KawaiiIpsecApi.md#updateKawaiiIpSec) | **PUT** /kawaii/{kawaiiId}/ipsec/{KawaiiIpSecId} |  |


<a name="createKawaiiIpSec"></a>
# **createKawaiiIpSec**
> KawaiiIpSec createKawaiiIpSec(KawaiiIpSecId, KawaiiIpSec)



    Creates a new Kawaii IPsec connection.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **KawaiiIpSecId** | **String**| The ID of the Kawaii IPsec connection. | [default to null] |
| **KawaiiIpSec** | [**KawaiiIpSec**](../Models/KawaiiIpSec.md)| KawaiiIpSec payload. | |

### Return type

[**KawaiiIpSec**](../Models/KawaiiIpSec.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteKawaiiIpSec"></a>
# **deleteKawaiiIpSec**
> deleteKawaiiIpSec(KawaiiIpSecId)



    Deletes an existing Kawaii IPsec connection.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **KawaiiIpSecId** | **String**| The ID of the Kawaii IPsec connection. | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listKawaiiIpSecs"></a>
# **listKawaiiIpSecs**
> List listKawaiiIpSecs(KawaiiIpSecId)



    Returns the IDs of Kawaii IPsec connection objects.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **KawaiiIpSecId** | **String**| The ID of the Kawaii IPsec connection. | [default to null] |

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="readKawaiiIpSec"></a>
# **readKawaiiIpSec**
> KawaiiIpSec readKawaiiIpSec(KawaiiIpSecId)



    Returns a Kawaii IPsec connection.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **KawaiiIpSecId** | **String**| The ID of the Kawaii IPsec connection. | [default to null] |

### Return type

[**KawaiiIpSec**](../Models/KawaiiIpSec.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateKawaiiIpSec"></a>
# **updateKawaiiIpSec**
> KawaiiIpSec updateKawaiiIpSec(KawaiiIpSecId, KawaiiIpSec)



    Updates a Kawaii IPsec connection configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **KawaiiIpSecId** | **String**| The ID of the Kawaii IPsec connection. | [default to null] |
| **KawaiiIpSec** | [**KawaiiIpSec**](../Models/KawaiiIpSec.md)| KawaiiIpSec payload. | |

### Return type

[**KawaiiIpSec**](../Models/KawaiiIpSec.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

