# NfsApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createStorageNFS_0**](NfsApi.md#createStorageNFS_0) | **POST** /region/{regionId}/nfs |  |
| [**deleteStorageNFS**](NfsApi.md#deleteStorageNFS) | **DELETE** /nfs/{nfsId} |  |
| [**listRegionStorageNFSs_0**](NfsApi.md#listRegionStorageNFSs_0) | **GET** /region/{regionId}/nfs |  |
| [**listStorageNFSKylos**](NfsApi.md#listStorageNFSKylos) | **GET** /nfs/{nfsId}/kylo |  |
| [**listStorageNFSs**](NfsApi.md#listStorageNFSs) | **GET** /nfs |  |
| [**readStorageNFS**](NfsApi.md#readStorageNFS) | **GET** /nfs/{nfsId} |  |
| [**setRegionDefaultStorageNFS_0**](NfsApi.md#setRegionDefaultStorageNFS_0) | **PATCH** /region/{regionId}/nfs/{nfsId}/default |  |
| [**updateStorageNFS**](NfsApi.md#updateStorageNFS) | **PUT** /nfs/{nfsId} |  |


<a name="createStorageNFS_0"></a>
# **createStorageNFS_0**
> StorageNFS createStorageNFS_0(regionId, StorageNFS, poolId)



    Creates a new NFS storage.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **regionId** | **String**| The ID of the region. | [default to null] |
| **StorageNFS** | [**StorageNFS**](../Models/StorageNFS.md)| StorageNFS payload. | |
| **poolId** | **String**| Storage pool ID (optional, region&#39;s default if unspecified). | [optional] [default to null] |

### Return type

[**StorageNFS**](../Models/StorageNFS.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteStorageNFS"></a>
# **deleteStorageNFS**
> deleteStorageNFS(nfsId)



    Deletes an existing NFS storage.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **nfsId** | **String**| The ID of the NFS storage. | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listRegionStorageNFSs_0"></a>
# **listRegionStorageNFSs_0**
> List listRegionStorageNFSs_0(regionId, poolId)



    Returns the IDs of NFS storage objects.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **regionId** | **String**| The ID of the region. | [default to null] |
| **poolId** | **String**| Storage pool ID (optional, region&#39;s default if unspecified). | [optional] [default to null] |

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

<a name="listStorageNFSs"></a>
# **listStorageNFSs**
> List listStorageNFSs()



    Returns the IDs of NFS storage objects.

### Parameters
This endpoint does not need any parameter.

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="readStorageNFS"></a>
# **readStorageNFS**
> StorageNFS readStorageNFS(nfsId)



    Returns a NFS storage.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **nfsId** | **String**| The ID of the NFS storage. | [default to null] |

### Return type

[**StorageNFS**](../Models/StorageNFS.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="setRegionDefaultStorageNFS_0"></a>
# **setRegionDefaultStorageNFS_0**
> setRegionDefaultStorageNFS_0(regionId, nfsId)



    Performs a region setting of default NFS storage.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **regionId** | **String**| The ID of the region. | [default to null] |
| **nfsId** | **String**| The ID of the NFS storage. | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateStorageNFS"></a>
# **updateStorageNFS**
> StorageNFS updateStorageNFS(nfsId, StorageNFS)



    Updates a NFS storage configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **nfsId** | **String**| The ID of the NFS storage. | [default to null] |
| **StorageNFS** | [**StorageNFS**](../Models/StorageNFS.md)| StorageNFS payload. | |

### Return type

[**StorageNFS**](../Models/StorageNFS.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

