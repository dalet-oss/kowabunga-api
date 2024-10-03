# KawaiiApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createProjectRegionKawaii**](KawaiiApi.md#createProjectRegionKawaii) | **POST** /project/{projectId}/region/{regionId}/kawaii |  |
| [**deleteKawaii**](KawaiiApi.md#deleteKawaii) | **DELETE** /kawaii/{kawaiiId} |  |
| [**listKawaiis**](KawaiiApi.md#listKawaiis) | **GET** /kawaii |  |
| [**listProjectRegionKawaiis**](KawaiiApi.md#listProjectRegionKawaiis) | **GET** /project/{projectId}/region/{regionId}/kawaiis |  |
| [**readKawaii**](KawaiiApi.md#readKawaii) | **GET** /kawaii/{kawaiiId} |  |
| [**updateKawaii**](KawaiiApi.md#updateKawaii) | **PUT** /kawaii/{kawaiiId} |  |


<a name="createProjectRegionKawaii"></a>
# **createProjectRegionKawaii**
> Kawaii createProjectRegionKawaii(projectId, regionId, Kawaii)



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

<a name="listProjectRegionKawaiis"></a>
# **listProjectRegionKawaiis**
> List listProjectRegionKawaiis(projectId, regionId)



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

