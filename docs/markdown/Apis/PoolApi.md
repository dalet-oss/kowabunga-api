# PoolApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createPool**](PoolApi.md#createPool) | **POST** /zone/{zoneId}/pool |  |
| [**createTemplate**](PoolApi.md#createTemplate) | **POST** /pool/{poolId}/template |  |
| [**deletePool**](PoolApi.md#deletePool) | **DELETE** /pool/{poolId} |  |
| [**getAllPools**](PoolApi.md#getAllPools) | **GET** /pool |  |
| [**getPool**](PoolApi.md#getPool) | **GET** /pool/{poolId} |  |
| [**getPoolTemplates**](PoolApi.md#getPoolTemplates) | **GET** /pool/{poolId}/templates |  |
| [**getPoolVolumes**](PoolApi.md#getPoolVolumes) | **GET** /pool/{poolId}/volumes |  |
| [**getZonePools**](PoolApi.md#getZonePools) | **GET** /zone/{zoneId}/pools |  |
| [**updatePool**](PoolApi.md#updatePool) | **PUT** /pool/{poolId} |  |
| [**updatePoolDefaultTemplate**](PoolApi.md#updatePoolDefaultTemplate) | **PUT** /pool/{poolId}/template/{templateId}/default |  |
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



    Creates a new volume template.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **poolId** | **String**| The ID of the storage pool. | [default to null] |
| **Template** | [**Template**](../Models/Template.md)| Template payload | |

### Return type

[**Template**](../Models/Template.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deletePool"></a>
# **deletePool**
> deletePool(poolId)



    Deletes an existing pool.

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

<a name="getAllPools"></a>
# **getAllPools**
> List getAllPools()



    Returns the IDs of registered pools.

### Parameters
This endpoint does not need any parameter.

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getPool"></a>
# **getPool**
> StoragePool getPool(poolId)



    Returns a description of the pool

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

<a name="getPoolTemplates"></a>
# **getPoolTemplates**
> List getPoolTemplates(poolId)



    Returns the IDs of the volume templates existing in the storage pool.

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

<a name="getPoolVolumes"></a>
# **getPoolVolumes**
> List getPoolVolumes(poolId)



    Returns the IDs of the storage volumes existing in the pool.

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

<a name="updatePool"></a>
# **updatePool**
> StoragePool updatePool(poolId, StoragePool)



    Updates a pool configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **poolId** | **String**| The ID of the storage pool. | [default to null] |
| **StoragePool** | [**StoragePool**](../Models/StoragePool.md)| Storage pool payload | |

### Return type

[**StoragePool**](../Models/StoragePool.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updatePoolDefaultTemplate"></a>
# **updatePoolDefaultTemplate**
> updatePoolDefaultTemplate(poolId, templateId)



    Set a storage pool default volume template.

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

