# KyloApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createProjectRegionKylo**](KyloApi.md#createProjectRegionKylo) | **POST** /project/{projectId}/region/{regionId}/kylo |  |
| [**deleteKylo**](KyloApi.md#deleteKylo) | **DELETE** /kylo/{kyloId} |  |
| [**listKylos**](KyloApi.md#listKylos) | **GET** /kylo |  |
| [**listProjectRegionKylos**](KyloApi.md#listProjectRegionKylos) | **GET** /project/{projectId}/region/{regionId}/kylo |  |
| [**listStorageNFSKylos**](KyloApi.md#listStorageNFSKylos) | **GET** /nfs/{nfsId}/kylo |  |
| [**readKylo**](KyloApi.md#readKylo) | **GET** /kylo/{kyloId} |  |
| [**updateKylo**](KyloApi.md#updateKylo) | **PUT** /kylo/{kyloId} |  |


<a name="createProjectRegionKylo"></a>
# **createProjectRegionKylo**
> Kylo createProjectRegionKylo(projectId, regionId, Kylo, nfsId)



    Creates a new Kylo.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |
| **regionId** | **String**| The ID of the region. | [default to null] |
| **Kylo** | [**Kylo**](../Models/Kylo.md)| Kylo payload. | |
| **nfsId** | **String**| NFS storage ID (optional, region&#39;s default if unspecified). | [optional] [default to null] |

### Return type

[**Kylo**](../Models/Kylo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteKylo"></a>
# **deleteKylo**
> deleteKylo(kyloId)



    Deletes an existing Kylo.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kyloId** | **String**| The ID of the Kylo. | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listKylos"></a>
# **listKylos**
> List listKylos()



    Returns the IDs of Kylo objects.

### Parameters
This endpoint does not need any parameter.

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listProjectRegionKylos"></a>
# **listProjectRegionKylos**
> List listProjectRegionKylos(projectId, regionId, nfsId)



    Returns the IDs of Kylo objects.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |
| **regionId** | **String**| The ID of the region. | [default to null] |
| **nfsId** | **String**| NFS storage ID (optional, region&#39;s default if unspecified). | [optional] [default to null] |

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listStorageNFSKylos"></a>
# **listStorageNFSKylos**
> List listStorageNFSKylos(nfsId)



    Returns the IDs of Kylo objects.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **nfsId** | **String**| The ID of the NFS storage. | [default to null] |

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="readKylo"></a>
# **readKylo**
> Kylo readKylo(kyloId)



    Returns a Kylo.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kyloId** | **String**| The ID of the Kylo. | [default to null] |

### Return type

[**Kylo**](../Models/Kylo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateKylo"></a>
# **updateKylo**
> Kylo updateKylo(kyloId, Kylo)



    Updates a Kylo configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kyloId** | **String**| The ID of the Kylo. | [default to null] |
| **Kylo** | [**Kylo**](../Models/Kylo.md)| Kylo payload. | |

### Return type

[**Kylo**](../Models/Kylo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

