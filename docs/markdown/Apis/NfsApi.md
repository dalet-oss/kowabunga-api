# NfsApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createStorageNFS**](NfsApi.md#createStorageNFS) | **POST** /zone/{ zoneId }/nfs |  |
| [**deleteStorageNFS**](NfsApi.md#deleteStorageNFS) | **DELETE** /nfs/{ nfsId } |  |
| [**listStorageNFSKFSs**](NfsApi.md#listStorageNFSKFSs) | **GET** /nfs/{ nfsId }/kfs |  |
| [**listStorageNFSs**](NfsApi.md#listStorageNFSs) | **GET** /nfs |  |
| [**listZoneStorageNFSs**](NfsApi.md#listZoneStorageNFSs) | **GET** /zone/{ zoneId }/nfs |  |
| [**readStorageNFS**](NfsApi.md#readStorageNFS) | **GET** /nfs/{ nfsId } |  |
| [**setZoneDefaultStorageNFS**](NfsApi.md#setZoneDefaultStorageNFS) | **PATCH** /zone/{ zoneId }/nfs/{ nfsId }/default |  |
| [**updateStorageNFS**](NfsApi.md#updateStorageNFS) | **PUT** /nfs/{ nfsId } |  |


<a name="createStorageNFS"></a>
# **createStorageNFS**
> StorageNFS createStorageNFS(zoneId, StorageNFS)



    Creates a new NFS storage.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |
| **StorageNFS** | [**StorageNFS**](../Models/StorageNFS.md)| StorageNFS payload. | |

### Return type

[**StorageNFS**](../Models/StorageNFS.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listStorageNFSKFSs"></a>
# **listStorageNFSKFSs**
> List listStorageNFSKFSs(nfsId)



    Returns the IDs of KFS (Kowabunga File System) objects.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **nfsId** | **String**| The ID of the NFS storage. | [default to null] |

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listZoneStorageNFSs"></a>
# **listZoneStorageNFSs**
> List listZoneStorageNFSs(zoneId)



    Returns the IDs of NFS storage objects.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="setZoneDefaultStorageNFS"></a>
# **setZoneDefaultStorageNFS**
> setZoneDefaultStorageNFS(zoneId, nfsId)



    Performs a availability zone setting of default NFS storage.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |
| **nfsId** | **String**| The ID of the NFS storage. | [default to null] |

### Return type

null (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

