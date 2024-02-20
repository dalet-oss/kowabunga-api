# NfsApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createNfsStorage**](NfsApi.md#createNfsStorage) | **POST** /zone/{zoneId}/nfs |  |
| [**deleteStorageNFS**](NfsApi.md#deleteStorageNFS) | **DELETE** /nfs/{ nfsId } |  |
| [**getZoneNfsStorages**](NfsApi.md#getZoneNfsStorages) | **GET** /zone/{zoneId}/nfs |  |
| [**listStorageNFSKFSs**](NfsApi.md#listStorageNFSKFSs) | **GET** /nfs/{ nfsId }/kfs |  |
| [**listStorageNFSs**](NfsApi.md#listStorageNFSs) | **GET** /nfs |  |
| [**readStorageNFS**](NfsApi.md#readStorageNFS) | **GET** /nfs/{ nfsId } |  |
| [**updateStorageNFS**](NfsApi.md#updateStorageNFS) | **PUT** /nfs/{ nfsId } |  |
| [**updateZoneDefaultNfsStorage**](NfsApi.md#updateZoneDefaultNfsStorage) | **PUT** /zone/{zoneId}/nfs/{nfsId}/default |  |


<a name="createNfsStorage"></a>
# **createNfsStorage**
> StorageNFS createNfsStorage(zoneId, StorageNFS)



    Creates a new NFS storage.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **zoneId** | **String**| The ID of the zone. | [default to null] |
| **StorageNFS** | [**StorageNFS**](../Models/StorageNFS.md)| NFS payload | |

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

<a name="getZoneNfsStorages"></a>
# **getZoneNfsStorages**
> List getZoneNfsStorages(zoneId)



    Returns the IDs of the NFS storages existing in the zone.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **zoneId** | **String**| The ID of the zone. | [default to null] |

### Return type

**List**

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

<a name="updateZoneDefaultNfsStorage"></a>
# **updateZoneDefaultNfsStorage**
> updateZoneDefaultNfsStorage(zoneId, nfsId)



    Set a zone&#39;s default NFS storage.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **zoneId** | **String**| The ID of the zone. | [default to null] |
| **nfsId** | **String**| The ID of the NFS storage. | [default to null] |

### Return type

null (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

