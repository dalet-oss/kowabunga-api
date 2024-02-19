# VolumeApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createProjectZoneVolume**](VolumeApi.md#createProjectZoneVolume) | **POST** /project/{projectId}/zone/{zoneId}/volume |  |
| [**deleteVolume**](VolumeApi.md#deleteVolume) | **DELETE** /volume/{volumeId} |  |
| [**getAllVolumes**](VolumeApi.md#getAllVolumes) | **GET** /volume |  |
| [**getPoolVolumes**](VolumeApi.md#getPoolVolumes) | **GET** /pool/{poolId}/volumes |  |
| [**getProjectZoneVolumes**](VolumeApi.md#getProjectZoneVolumes) | **GET** /project/{projectId}/zone/{zoneId}/volumes |  |
| [**getVolume**](VolumeApi.md#getVolume) | **GET** /volume/{volumeId} |  |
| [**updateVolume**](VolumeApi.md#updateVolume) | **PUT** /volume/{volumeId} |  |


<a name="createProjectZoneVolume"></a>
# **createProjectZoneVolume**
> Volume createProjectZoneVolume(projectId, zoneId, Volume, poolId, templateId)



    Creates a new storage volume in specified zone.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the network adapter. | [default to null] |
| **zoneId** | **String**| The ID of the zone. | [default to null] |
| **Volume** | [**Volume**](../Models/Volume.md)| Volume payload | |
| **poolId** | **String**| Storage pool ID (optional, zone&#39;s default if unspecified). | [optional] [default to null] |
| **templateId** | **String**| Template to clone the storage volume from (optional, zone&#39;s default if unspecified) | [optional] [default to null] |

### Return type

[**Volume**](../Models/Volume.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteVolume"></a>
# **deleteVolume**
> deleteVolume(volumeId)



    Deletes an existing storage volume.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **volumeId** | **String**| The ID of the storage volume. | [default to null] |

### Return type

null (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getAllVolumes"></a>
# **getAllVolumes**
> List getAllVolumes()



    Returns the IDs of storage volumes.

### Parameters
This endpoint does not need any parameter.

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

<a name="getProjectZoneVolumes"></a>
# **getProjectZoneVolumes**
> List getProjectZoneVolumes(projectId, zoneId)



    Returns the IDs of the storage volumes existing in the project in the specified zone.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the network adapter. | [default to null] |
| **zoneId** | **String**| The ID of the zone. | [default to null] |

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getVolume"></a>
# **getVolume**
> Volume getVolume(volumeId)



    Returns a description of the storage volume.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **volumeId** | **String**| The ID of the storage volume. | [default to null] |

### Return type

[**Volume**](../Models/Volume.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateVolume"></a>
# **updateVolume**
> Volume updateVolume(volumeId, Volume)



    Updates/resizes a storage volume configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **volumeId** | **String**| The ID of the storage volume. | [default to null] |
| **Volume** | [**Volume**](../Models/Volume.md)| Volume payload | |

### Return type

[**Volume**](../Models/Volume.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

