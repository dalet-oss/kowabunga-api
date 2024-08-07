# RegionApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createNetGW**](RegionApi.md#createNetGW) | **POST** /region/{regionId}/netgw |  |
| [**createProjectRegionKFS**](RegionApi.md#createProjectRegionKFS) | **POST** /project/{projectId}/region/{regionId}/kfs |  |
| [**createProjectRegionKGW**](RegionApi.md#createProjectRegionKGW) | **POST** /project/{projectId}/region/{regionId}/kgw |  |
| [**createProjectRegionKonvey**](RegionApi.md#createProjectRegionKonvey) | **POST** /project/{projectId}/region/{regionId}/konvey |  |
| [**createProjectRegionVolume**](RegionApi.md#createProjectRegionVolume) | **POST** /project/{projectId}/region/{regionId}/volume |  |
| [**createRegion**](RegionApi.md#createRegion) | **POST** /region |  |
| [**createStorageNFS**](RegionApi.md#createStorageNFS) | **POST** /region/{regionId}/nfs |  |
| [**createStoragePool**](RegionApi.md#createStoragePool) | **POST** /region/{regionId}/pool |  |
| [**createVNet**](RegionApi.md#createVNet) | **POST** /region/{regionId}/vnet |  |
| [**createZone**](RegionApi.md#createZone) | **POST** /region/{regionId}/zone |  |
| [**deleteRegion**](RegionApi.md#deleteRegion) | **DELETE** /region/{regionId} |  |
| [**listProjectRegionKFSs**](RegionApi.md#listProjectRegionKFSs) | **GET** /project/{projectId}/region/{regionId}/kfs |  |
| [**listProjectRegionKGWs**](RegionApi.md#listProjectRegionKGWs) | **GET** /project/{projectId}/region/{regionId}/kgws |  |
| [**listProjectRegionKonveys**](RegionApi.md#listProjectRegionKonveys) | **GET** /project/{projectId}/region/{regionId}/konveys |  |
| [**listProjectRegionVolumes**](RegionApi.md#listProjectRegionVolumes) | **GET** /project/{projectId}/region/{regionId}/volumes |  |
| [**listRegionNetGWs**](RegionApi.md#listRegionNetGWs) | **GET** /region/{regionId}/netgws |  |
| [**listRegionStorageNFSs**](RegionApi.md#listRegionStorageNFSs) | **GET** /region/{regionId}/nfs |  |
| [**listRegionStoragePools**](RegionApi.md#listRegionStoragePools) | **GET** /region/{regionId}/pools |  |
| [**listRegionVNets**](RegionApi.md#listRegionVNets) | **GET** /region/{regionId}/vnets |  |
| [**listRegionZones**](RegionApi.md#listRegionZones) | **GET** /region/{regionId}/zones |  |
| [**listRegions**](RegionApi.md#listRegions) | **GET** /region |  |
| [**readRegion**](RegionApi.md#readRegion) | **GET** /region/{regionId} |  |
| [**setRegionDefaultStorageNFS**](RegionApi.md#setRegionDefaultStorageNFS) | **PATCH** /region/{regionId}/nfs/{nfsId}/default |  |
| [**setRegionDefaultStoragePool**](RegionApi.md#setRegionDefaultStoragePool) | **PATCH** /region/{regionId}/pool/{poolId}/default |  |
| [**updateRegion**](RegionApi.md#updateRegion) | **PUT** /region/{regionId} |  |


<a name="createNetGW"></a>
# **createNetGW**
> NetGW createNetGW(regionId, NetGW)



    Creates a new Iris network gateway.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **regionId** | **String**| The ID of the region. | [default to null] |
| **NetGW** | [**NetGW**](../Models/NetGW.md)| NetGW payload. | |

### Return type

[**NetGW**](../Models/NetGW.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createProjectRegionKFS"></a>
# **createProjectRegionKFS**
> KFS createProjectRegionKFS(projectId, regionId, KFS, nfsId)



    Creates a new KFS (Kowabunga File System).

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |
| **regionId** | **String**| The ID of the region. | [default to null] |
| **KFS** | [**KFS**](../Models/KFS.md)| KFS payload. | |
| **nfsId** | **String**| NFS storage ID (optional, region&#39;s default if unspecified). | [optional] [default to null] |

### Return type

[**KFS**](../Models/KFS.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createProjectRegionKGW"></a>
# **createProjectRegionKGW**
> KGW createProjectRegionKGW(projectId, regionId, KGW)



    Creates a new KGW (Kowabunga Network Gateway).

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |
| **regionId** | **String**| The ID of the region. | [default to null] |
| **KGW** | [**KGW**](../Models/KGW.md)| KGW payload. | |

### Return type

[**KGW**](../Models/KGW.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createProjectRegionKonvey"></a>
# **createProjectRegionKonvey**
> Konvey createProjectRegionKonvey(projectId, regionId, Konvey)



    Creates a new Konvey (Kowabunga Network Load-Balancer).

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |
| **regionId** | **String**| The ID of the region. | [default to null] |
| **Konvey** | [**Konvey**](../Models/Konvey.md)| Konvey payload. | |

### Return type

[**Konvey**](../Models/Konvey.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createProjectRegionVolume"></a>
# **createProjectRegionVolume**
> Volume createProjectRegionVolume(projectId, regionId, Volume, poolId, templateId)



    Creates a new storage volume.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |
| **regionId** | **String**| The ID of the region. | [default to null] |
| **Volume** | [**Volume**](../Models/Volume.md)| Volume payload. | |
| **poolId** | **String**| Storage pool ID (optional, region&#39;s default if unspecified). | [optional] [default to null] |
| **templateId** | **String**| Template to clone the storage volume from (optional, region&#39;s default if unspecified). | [optional] [default to null] |

### Return type

[**Volume**](../Models/Volume.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createRegion"></a>
# **createRegion**
> Region createRegion(Region)



    Creates a new region.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **Region** | [**Region**](../Models/Region.md)| Region payload. | |

### Return type

[**Region**](../Models/Region.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createStorageNFS"></a>
# **createStorageNFS**
> StorageNFS createStorageNFS(regionId, StorageNFS, poolId)



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

<a name="createStoragePool"></a>
# **createStoragePool**
> StoragePool createStoragePool(regionId, StoragePool)



    Creates a new storage pool.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **regionId** | **String**| The ID of the region. | [default to null] |
| **StoragePool** | [**StoragePool**](../Models/StoragePool.md)| StoragePool payload. | |

### Return type

[**StoragePool**](../Models/StoragePool.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createVNet"></a>
# **createVNet**
> VNet createVNet(regionId, VNet)



    Creates a new virtual network.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **regionId** | **String**| The ID of the region. | [default to null] |
| **VNet** | [**VNet**](../Models/VNet.md)| VNet payload. | |

### Return type

[**VNet**](../Models/VNet.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createZone"></a>
# **createZone**
> Zone createZone(regionId, Zone)



    Creates a new availability zone.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **regionId** | **String**| The ID of the region. | [default to null] |
| **Zone** | [**Zone**](../Models/Zone.md)| Zone payload. | |

### Return type

[**Zone**](../Models/Zone.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteRegion"></a>
# **deleteRegion**
> deleteRegion(regionId)



    Deletes an existing region.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **regionId** | **String**| The ID of the region. | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listProjectRegionKFSs"></a>
# **listProjectRegionKFSs**
> List listProjectRegionKFSs(projectId, regionId, nfsId)



    Returns the IDs of KFS (Kowabunga File System) objects.

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

<a name="listProjectRegionKGWs"></a>
# **listProjectRegionKGWs**
> List listProjectRegionKGWs(projectId, regionId)



    Returns the IDs of KGW (Kowabunga Network Gateway) objects.

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

<a name="listProjectRegionKonveys"></a>
# **listProjectRegionKonveys**
> List listProjectRegionKonveys(projectId, regionId)



    Returns the IDs of Konvey (Kowabunga Network Load-Balancer) objects.

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

<a name="listProjectRegionVolumes"></a>
# **listProjectRegionVolumes**
> List listProjectRegionVolumes(projectId, regionId)



    Returns the IDs of storage volume objects.

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

<a name="listRegionNetGWs"></a>
# **listRegionNetGWs**
> List listRegionNetGWs(regionId)



    Returns the IDs of Iris network gateway objects.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **regionId** | **String**| The ID of the region. | [default to null] |

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listRegionStorageNFSs"></a>
# **listRegionStorageNFSs**
> List listRegionStorageNFSs(regionId, poolId)



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

<a name="listRegionStoragePools"></a>
# **listRegionStoragePools**
> List listRegionStoragePools(regionId)



    Returns the IDs of storage pool objects.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **regionId** | **String**| The ID of the region. | [default to null] |

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listRegionVNets"></a>
# **listRegionVNets**
> List listRegionVNets(regionId)



    Returns the IDs of virtual network objects.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **regionId** | **String**| The ID of the region. | [default to null] |

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listRegionZones"></a>
# **listRegionZones**
> List listRegionZones(regionId)



    Returns the IDs of availability zone objects.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **regionId** | **String**| The ID of the region. | [default to null] |

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listRegions"></a>
# **listRegions**
> List listRegions()



    Returns the IDs of region objects.

### Parameters
This endpoint does not need any parameter.

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="readRegion"></a>
# **readRegion**
> Region readRegion(regionId)



    Returns a region.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **regionId** | **String**| The ID of the region. | [default to null] |

### Return type

[**Region**](../Models/Region.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="setRegionDefaultStorageNFS"></a>
# **setRegionDefaultStorageNFS**
> setRegionDefaultStorageNFS(regionId, nfsId)



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

<a name="setRegionDefaultStoragePool"></a>
# **setRegionDefaultStoragePool**
> setRegionDefaultStoragePool(regionId, poolId)



    Performs a region setting of default storage pool.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **regionId** | **String**| The ID of the region. | [default to null] |
| **poolId** | **String**| The ID of the storage pool. | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateRegion"></a>
# **updateRegion**
> Region updateRegion(regionId, Region)



    Updates a region configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **regionId** | **String**| The ID of the region. | [default to null] |
| **Region** | [**Region**](../Models/Region.md)| Region payload. | |

### Return type

[**Region**](../Models/Region.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

