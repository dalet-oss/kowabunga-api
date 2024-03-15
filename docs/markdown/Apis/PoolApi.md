# PoolApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createStoragePool**](PoolApi.md#createStoragePool) | **POST** /zone/{zoneId}/pool |  |
| [**createTemplate**](PoolApi.md#createTemplate) | **POST** /pool/{poolId}/template |  |
| [**deleteStoragePool**](PoolApi.md#deleteStoragePool) | **DELETE** /pool/{poolId} |  |
| [**listStoragePoolTemplates**](PoolApi.md#listStoragePoolTemplates) | **GET** /pool/{poolId}/templates |  |
| [**listStoragePoolVolumes**](PoolApi.md#listStoragePoolVolumes) | **GET** /pool/{poolId}/volumes |  |
| [**listStoragePools**](PoolApi.md#listStoragePools) | **GET** /pool |  |
| [**listZoneStoragePools**](PoolApi.md#listZoneStoragePools) | **GET** /zone/{zoneId}/pools |  |
| [**readStoragePool**](PoolApi.md#readStoragePool) | **GET** /pool/{poolId} |  |
| [**setStoragePoolDefaultTemplate**](PoolApi.md#setStoragePoolDefaultTemplate) | **PATCH** /pool/{poolId}/template/{templateId}/default |  |
| [**setZoneDefaultStoragePool**](PoolApi.md#setZoneDefaultStoragePool) | **PATCH** /zone/{zoneId}/pool/{poolId}/default |  |
| [**updateStoragePool**](PoolApi.md#updateStoragePool) | **PUT** /pool/{poolId} |  |


<a name="createStoragePool"></a>
# **createStoragePool**
> StoragePool createStoragePool(zoneId, StoragePool)



    Creates a new storage pool.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |
| **StoragePool** | [**StoragePool**](../Models/StoragePool.md)| StoragePool payload. | |

### Return type

[**StoragePool**](../Models/StoragePool.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listZoneStoragePools"></a>
# **listZoneStoragePools**
> List listZoneStoragePools(zoneId)



    Returns the IDs of storage pool objects.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="setStoragePoolDefaultTemplate"></a>
# **setStoragePoolDefaultTemplate**
> setStoragePoolDefaultTemplate(poolId, templateId)



    Performs a storage pool setting of default template.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **poolId** | **String**| The ID of the storage pool. | [default to null] |
| **templateId** | **String**| The ID of the image template. | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="setZoneDefaultStoragePool"></a>
# **setZoneDefaultStoragePool**
> setZoneDefaultStoragePool(zoneId, poolId)



    Performs a availability zone setting of default storage pool.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |
| **poolId** | **String**| The ID of the storage pool. | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

