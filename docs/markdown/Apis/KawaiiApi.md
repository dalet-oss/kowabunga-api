# KawaiiApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createKawaiiIpSec**](KawaiiApi.md#createKawaiiIpSec) | **POST** /kawaii/{kawaiiId}/ipsec |  |
| [**createProjectRegionKawaii_1**](KawaiiApi.md#createProjectRegionKawaii_1) | **POST** /project/{projectId}/region/{regionId}/kawaii |  |
| [**deleteKawaii**](KawaiiApi.md#deleteKawaii) | **DELETE** /kawaii/{kawaiiId} |  |
| [**deleteKawaiiIpSec**](KawaiiApi.md#deleteKawaiiIpSec) | **DELETE** /kawaii/{kawaiiId}/ipsec/{KawaiiIpSecId} |  |
| [**listKawaiiIpSecs**](KawaiiApi.md#listKawaiiIpSecs) | **GET** /kawaii/{kawaiiId}/ipsec |  |
| [**listKawaiis**](KawaiiApi.md#listKawaiis) | **GET** /kawaii |  |
| [**listProjectRegionKawaiis_1**](KawaiiApi.md#listProjectRegionKawaiis_1) | **GET** /project/{projectId}/region/{regionId}/kawaiis |  |
| [**readKawaii**](KawaiiApi.md#readKawaii) | **GET** /kawaii/{kawaiiId} |  |
| [**readKawaiiIpSec**](KawaiiApi.md#readKawaiiIpSec) | **GET** /kawaii/{kawaiiId}/ipsec/{KawaiiIpSecId} |  |
| [**updateKawaii**](KawaiiApi.md#updateKawaii) | **PUT** /kawaii/{kawaiiId} |  |
| [**updateKawaiiIpSec**](KawaiiApi.md#updateKawaiiIpSec) | **PUT** /kawaii/{kawaiiId}/ipsec/{KawaiiIpSecId} |  |


<a name="createKawaiiIpSec"></a>
# **createKawaiiIpSec**
> KawaiiIpSec createKawaiiIpSec(kawaiiId, KawaiiIpSec)



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

<a name="createProjectRegionKawaii_1"></a>
# **createProjectRegionKawaii_1**
> Kawaii createProjectRegionKawaii_1(projectId, regionId, Kawaii)



    Creates a new Kawaii.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |
| **regionId** | **String**| The ID of the region. | [default to null] |
| **Kawaii** | [**Kawaii**](../Models/Kawaii.md)| Kawaii payload. | |

### Return type

[**Kawaii**](../Models/Kawaii.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteKawaii"></a>
# **deleteKawaii**
> deleteKawaii(kawaiiId)



    Deletes an existing Kawaii.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kawaiiId** | **String**| The ID of the Kawaii. | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="deleteKawaiiIpSec"></a>
# **deleteKawaiiIpSec**
> deleteKawaiiIpSec(kawaiiId, KawaiiIpSecId)



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

<a name="listKawaiiIpSecs"></a>
# **listKawaiiIpSecs**
> List listKawaiiIpSecs(kawaiiId)



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

<a name="listKawaiis"></a>
# **listKawaiis**
> List listKawaiis()



    Returns the IDs of Kawaii objects.

### Parameters
This endpoint does not need any parameter.

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listProjectRegionKawaiis_1"></a>
# **listProjectRegionKawaiis_1**
> List listProjectRegionKawaiis_1(projectId, regionId)



    Returns the IDs of Kawaii objects.

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

<a name="readKawaii"></a>
# **readKawaii**
> Kawaii readKawaii(kawaiiId)



    Returns a Kawaii.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kawaiiId** | **String**| The ID of the Kawaii. | [default to null] |

### Return type

[**Kawaii**](../Models/Kawaii.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="readKawaiiIpSec"></a>
# **readKawaiiIpSec**
> KawaiiIpSec readKawaiiIpSec(kawaiiId, KawaiiIpSecId)



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

<a name="updateKawaii"></a>
# **updateKawaii**
> Kawaii updateKawaii(kawaiiId, Kawaii)



    Updates a Kawaii configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kawaiiId** | **String**| The ID of the Kawaii. | [default to null] |
| **Kawaii** | [**Kawaii**](../Models/Kawaii.md)| Kawaii payload. | |

### Return type

[**Kawaii**](../Models/Kawaii.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updateKawaiiIpSec"></a>
# **updateKawaiiIpSec**
> KawaiiIpSec updateKawaiiIpSec(kawaiiId, KawaiiIpSecId, KawaiiIpSec)



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

