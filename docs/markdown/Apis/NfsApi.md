# NfsApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createNfsStorage**](NfsApi.md#createNfsStorage) | **POST** /zone/{zoneId}/nfs |  |
| [**deleteNfsStorage**](NfsApi.md#deleteNfsStorage) | **DELETE** /nfs/{nfsId} |  |
| [**getAllNfsStorages**](NfsApi.md#getAllNfsStorages) | **GET** /nfs |  |
| [**getNfsKfs**](NfsApi.md#getNfsKfs) | **GET** /nfs/{nfsId}/kfs |  |
| [**getNfsStorage**](NfsApi.md#getNfsStorage) | **GET** /nfs/{nfsId} |  |
| [**getZoneNfsStorages**](NfsApi.md#getZoneNfsStorages) | **GET** /zone/{zoneId}/nfs |  |
| [**updateNfsStorage**](NfsApi.md#updateNfsStorage) | **PUT** /nfs/{nfsId} |  |
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

<a name="deleteNfsStorage"></a>
# **deleteNfsStorage**
> deleteNfsStorage(nfsId)



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

<a name="getAllNfsStorages"></a>
# **getAllNfsStorages**
> List getAllNfsStorages()



    Returns the IDs of registered NFS storages.

### Parameters
This endpoint does not need any parameter.

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getNfsKfs"></a>
# **getNfsKfs**
> List getNfsKfs(nfsId)



    Returns the IDs of the KFS volumes existing in the NFS storage.

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

<a name="getNfsStorage"></a>
# **getNfsStorage**
> StorageNFS getNfsStorage(nfsId)



    Returns a description of the NFS storage.

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

<a name="updateNfsStorage"></a>
# **updateNfsStorage**
> StorageNFS updateNfsStorage(nfsId, StorageNFS)



    Updates an NFS storage configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **nfsId** | **String**| The ID of the NFS storage. | [default to null] |
| **StorageNFS** | [**StorageNFS**](../Models/StorageNFS.md)| NFS storage payload | |

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

