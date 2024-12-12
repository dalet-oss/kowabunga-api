# KawaiiIpsecApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createKawaiiIpSec_0**](KawaiiIpsecApi.md#createKawaiiIpSec_0) | **POST** /kawaii/{kawaiiId}/ipsec |  |
| [**deleteKawaiiIpSec_0**](KawaiiIpsecApi.md#deleteKawaiiIpSec_0) | **DELETE** /kawaii/{kawaiiId}/ipsec/{KawaiiIpSecId} |  |
| [**listKawaiiIpSecs_0**](KawaiiIpsecApi.md#listKawaiiIpSecs_0) | **GET** /kawaii/{kawaiiId}/ipsec |  |
| [**readKawaiiIpSec_0**](KawaiiIpsecApi.md#readKawaiiIpSec_0) | **GET** /kawaii/{kawaiiId}/ipsec/{KawaiiIpSecId} |  |
| [**updateKawaiiIpSec_0**](KawaiiIpsecApi.md#updateKawaiiIpSec_0) | **PUT** /kawaii/{kawaiiId}/ipsec/{KawaiiIpSecId} |  |


<a name="createKawaiiIpSec_0"></a>
# **createKawaiiIpSec_0**
> KawaiiIpSec createKawaiiIpSec_0(kawaiiId, KawaiiIpSec)



    Creates a new Kawaii IPsec connection.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kawaiiId** | **String**| The ID of the Kawaii. | [default to null] |
| **KawaiiIpSec** | [**KawaiiIpSec**](../Models/KawaiiIpSec.md)| KawaiiIpSec payload. | |

### Return type

[**KawaiiIpSec**](../Models/KawaiiIpSec.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteKawaiiIpSec_0"></a>
# **deleteKawaiiIpSec_0**
> deleteKawaiiIpSec_0(kawaiiId, KawaiiIpSecId)



    Deletes an existing Kawaii IPsec connection.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kawaiiId** | **String**| The ID of the Kawaii. | [default to null] |
| **KawaiiIpSecId** | **String**| The ID of the Kawaii IPsec connection. | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listKawaiiIpSecs_0"></a>
# **listKawaiiIpSecs_0**
> List listKawaiiIpSecs_0(kawaiiId)



    Returns the IDs of Kawaii IPsec connection objects.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kawaiiId** | **String**| The ID of the Kawaii. | [default to null] |

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="readKawaiiIpSec_0"></a>
# **readKawaiiIpSec_0**
> KawaiiIpSec readKawaiiIpSec_0(kawaiiId, KawaiiIpSecId)



    Returns a Kawaii IPsec connection.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kawaiiId** | **String**| The ID of the Kawaii. | [default to null] |
| **KawaiiIpSecId** | **String**| The ID of the Kawaii IPsec connection. | [default to null] |

### Return type

[**KawaiiIpSec**](../Models/KawaiiIpSec.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateKawaiiIpSec_0"></a>
# **updateKawaiiIpSec_0**
> KawaiiIpSec updateKawaiiIpSec_0(kawaiiId, KawaiiIpSecId, KawaiiIpSec)



    Updates a Kawaii IPsec connection configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kawaiiId** | **String**| The ID of the Kawaii. | [default to null] |
| **KawaiiIpSecId** | **String**| The ID of the Kawaii IPsec connection. | [default to null] |
| **KawaiiIpSec** | [**KawaiiIpSec**](../Models/KawaiiIpSec.md)| KawaiiIpSec payload. | |

### Return type

[**KawaiiIpSec**](../Models/KawaiiIpSec.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

