# PoolApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createPool**](PoolApi.md#createPool) | **POST** /zone/{zoneId}/pool |  |
| [**createTemplate**](PoolApi.md#createTemplate) | **POST** /pool/{ poolId }/template |  |
| [**deleteStoragePool**](PoolApi.md#deleteStoragePool) | **DELETE** /pool/{ poolId } |  |
| [**getZonePools**](PoolApi.md#getZonePools) | **GET** /zone/{zoneId}/pools |  |
| [**listStoragePoolTemplates**](PoolApi.md#listStoragePoolTemplates) | **GET** /pool/{ poolId }/templates |  |
| [**listStoragePoolVolumes**](PoolApi.md#listStoragePoolVolumes) | **GET** /pool/{ poolId }/volumes |  |
| [**listStoragePools**](PoolApi.md#listStoragePools) | **GET** /pool |  |
| [**readStoragePool**](PoolApi.md#readStoragePool) | **GET** /pool/{ poolId } |  |
| [**updateStoragePool**](PoolApi.md#updateStoragePool) | **PUT** /pool/{ poolId } |  |
| [**updateStoragePoolDefaultTemplate**](PoolApi.md#updateStoragePoolDefaultTemplate) | **PATCH** /pool/{ poolId }/template/{ templateId }/default |  |
| [**updateZoneDefaultPool**](PoolApi.md#updateZoneDefaultPool) | **PUT** /zone/{zoneId}/pool/{poolId}/default |  |


<a name="createPool"></a>
# **createPool**
> StoragePool createPool(zoneId, StoragePool, hostId)



    Creates a new storage pool.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **zoneId** | **String**| The ID of the zone. | [default to null] |
| **StoragePool** | [**StoragePool**](../Models/StoragePool.md)| Pool payload | |
| **hostId** | **String**| The ID of the computing host (useless for RBD pools, mandatory for local ones). | [optional] [default to null] |

### Return type

[**StoragePool**](../Models/StoragePool.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createTemplate"></a>
# **createTemplate**
> Template createTemplate(poolId, Template)



    Creates a new image template.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **poolId** | **String**| The ID of the storage pool. | [default to null] |
| **Template** | [**Template**](../Models/Template.md)| Template payload. | |

### Return type

[**Template**](../Models/Template.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteStoragePool"></a>
# **deleteStoragePool**
> deleteStoragePool(poolId)



    Deletes an existing storage pool.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **poolId** | **String**| The ID of the storage pool. | [default to null] |

### Return type

null (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getZonePools"></a>
# **getZonePools**
> List getZonePools(zoneId)



    Returns the IDs of the pools existing in the zone.

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

<a name="listStoragePoolTemplates"></a>
# **listStoragePoolTemplates**
> List listStoragePoolTemplates(poolId)



    Returns the IDs of image template objects.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **poolId** | **String**| The ID of the storage pool. | [default to null] |

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listStoragePoolVolumes"></a>
# **listStoragePoolVolumes**
> List listStoragePoolVolumes(poolId)



    Returns the IDs of storage volume objects.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **poolId** | **String**| The ID of the storage pool. | [default to null] |

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listStoragePools"></a>
# **listStoragePools**
> List listStoragePools()



    Returns the IDs of storage pool objects.

### Parameters
This endpoint does not need any parameter.

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="readStoragePool"></a>
# **readStoragePool**
> StoragePool readStoragePool(poolId)



    Returns a storage pool.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **poolId** | **String**| The ID of the storage pool. | [default to null] |

### Return type

[**StoragePool**](../Models/StoragePool.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateStoragePool"></a>
# **updateStoragePool**
> StoragePool updateStoragePool(poolId, StoragePool)



    Updates a storage pool configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **poolId** | **String**| The ID of the storage pool. | [default to null] |
| **StoragePool** | [**StoragePool**](../Models/StoragePool.md)| StoragePool payload. | |

### Return type

[**StoragePool**](../Models/StoragePool.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updateStoragePoolDefaultTemplate"></a>
# **updateStoragePoolDefaultTemplate**
> updateStoragePoolDefaultTemplate(poolId, templateId)



    Performs a storage pool setting of default template.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **poolId** | **String**| The ID of the storage pool. | [default to null] |
| **templateId** | **String**| The ID of the volume template. | [default to null] |

### Return type

null (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateZoneDefaultPool"></a>
# **updateZoneDefaultPool**
> updateZoneDefaultPool(zoneId, poolId)



    Set a zone&#39;s default storage pool.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **zoneId** | **String**| The ID of the zone. | [default to null] |
| **poolId** | **String**| The ID of the storage pool. | [default to null] |

### Return type

null (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

