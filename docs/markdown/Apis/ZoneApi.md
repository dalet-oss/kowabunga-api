# ZoneApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createHost**](ZoneApi.md#createHost) | **POST** /zone/{ zoneId }/host |  |
| [**createNetGW**](ZoneApi.md#createNetGW) | **POST** /zone/{ zoneId }/netgw |  |
| [**createProjectZoneInstance**](ZoneApi.md#createProjectZoneInstance) | **POST** /project/{projectId}/zone/{zoneId}/instance |  |
| [**createProjectZoneKce**](ZoneApi.md#createProjectZoneKce) | **POST** /project/{projectId}/zone/{zoneId}/kce |  |
| [**createProjectZoneKfs**](ZoneApi.md#createProjectZoneKfs) | **POST** /project/{projectId}/zone/{zoneId}/kfs |  |
| [**createProjectZoneKgw**](ZoneApi.md#createProjectZoneKgw) | **POST** /project/{projectId}/zone/{zoneId}/kgw |  |
| [**createProjectZoneVolume**](ZoneApi.md#createProjectZoneVolume) | **POST** /project/{projectId}/zone/{zoneId}/volume |  |
| [**createStorageNFS**](ZoneApi.md#createStorageNFS) | **POST** /zone/{ zoneId }/nfs |  |
| [**createStoragePool**](ZoneApi.md#createStoragePool) | **POST** /zone/{ zoneId }/pool |  |
| [**createVNet**](ZoneApi.md#createVNet) | **POST** /zone/{ zoneId }/vnet |  |
| [**createZone**](ZoneApi.md#createZone) | **POST** /region/{ regionId }/zone |  |
| [**deleteZone**](ZoneApi.md#deleteZone) | **DELETE** /zone/{ zoneId } |  |
| [**getProjectZoneInstances**](ZoneApi.md#getProjectZoneInstances) | **GET** /project/{projectId}/zone/{zoneId}/instances |  |
| [**getProjectZoneKCEs**](ZoneApi.md#getProjectZoneKCEs) | **GET** /project/{projectId}/zone/{zoneId}/kces |  |
| [**getProjectZoneKGWs**](ZoneApi.md#getProjectZoneKGWs) | **GET** /project/{projectId}/zone/{zoneId}/kgws |  |
| [**getProjectZoneKfs**](ZoneApi.md#getProjectZoneKfs) | **GET** /project/{projectId}/zone/{zoneId}/kfs |  |
| [**getProjectZoneVolumes**](ZoneApi.md#getProjectZoneVolumes) | **GET** /project/{projectId}/zone/{zoneId}/volumes |  |
| [**listRegionZones**](ZoneApi.md#listRegionZones) | **GET** /region/{ regionId }/zones |  |
| [**listZoneHosts**](ZoneApi.md#listZoneHosts) | **GET** /zone/{zoneId}/hosts |  |
| [**listZoneNetGWs**](ZoneApi.md#listZoneNetGWs) | **GET** /zone/{ zoneId }/netgws |  |
| [**listZoneStorageNFSs**](ZoneApi.md#listZoneStorageNFSs) | **GET** /zone/{ zoneId }/nfs |  |
| [**listZoneStoragePools**](ZoneApi.md#listZoneStoragePools) | **GET** /zone/{ zoneId }/pools |  |
| [**listZoneVNets**](ZoneApi.md#listZoneVNets) | **GET** /zone/{ zoneId }/vnets |  |
| [**listZones**](ZoneApi.md#listZones) | **GET** /zone |  |
| [**readZone**](ZoneApi.md#readZone) | **GET** /zone/{ zoneId } |  |
| [**setZoneDefaultStorageNFS**](ZoneApi.md#setZoneDefaultStorageNFS) | **PATCH** /zone/{ zoneId }/nfs/{ nfsId }/default |  |
| [**setZoneDefaultStoragePool**](ZoneApi.md#setZoneDefaultStoragePool) | **PATCH** /zone/{ zoneId }/pool/{ poolId }/default |  |
| [**updateZone**](ZoneApi.md#updateZone) | **PUT** /zone/{ zoneId } |  |


<a name="createHost"></a>
# **createHost**
> Host createHost(zoneId, Host)



    Creates a new computing host.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |
| **Host** | [**Host**](../Models/Host.md)| Host payload. | |

### Return type

[**Host**](../Models/Host.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createNetGW"></a>
# **createNetGW**
> NetGW createNetGW(zoneId, NetGW)



    Creates a new Iris network gateway.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |
| **NetGW** | [**NetGW**](../Models/NetGW.md)| NetGW payload. | |

### Return type

[**NetGW**](../Models/NetGW.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createProjectZoneInstance"></a>
# **createProjectZoneInstance**
> Instance createProjectZoneInstance(projectId, zoneId, Instance, notify)



    Creates a new virtual machine instance in specified zone.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the network adapter. | [default to null] |
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |
| **Instance** | [**Instance**](../Models/Instance.md)| Instance payload | |
| **notify** | **Boolean**| Whether or not to send a notification email at resource creation. | [optional] [default to null] |

### Return type

[**Instance**](../Models/Instance.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createProjectZoneKce"></a>
# **createProjectZoneKce**
> KCE createProjectZoneKce(projectId, zoneId, KCE, poolId, templateId, public, notify)



    Creates a new KCE virtual machine in specified zone.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the network adapter. | [default to null] |
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |
| **KCE** | [**KCE**](../Models/KCE.md)| KCE payload | |
| **poolId** | **String**| Storage pool ID (optional, zone&#39;s default if unspecified). | [optional] [default to null] |
| **templateId** | **String**| Template to clone the storage volume from (optional, zone&#39;s default if unspecified). | [optional] [default to null] |
| **public** | **String**| Should KCE be exposed over public Internet ? (a public IPv4 address will then be auto-assigned, default to false). | [optional] [default to null] |
| **notify** | **Boolean**| Whether or not to send a notification email at resource creation. | [optional] [default to null] |

### Return type

[**KCE**](../Models/KCE.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createProjectZoneKfs"></a>
# **createProjectZoneKfs**
> KFS createProjectZoneKfs(projectId, zoneId, KFS, nfsId, notify)



    Creates a new KFS storage volume in specified zone.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the network adapter. | [default to null] |
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |
| **KFS** | [**KFS**](../Models/KFS.md)| KFS payload | |
| **nfsId** | **String**| NFS storage ID (optional, zone&#39;s default if unspecified). | [optional] [default to null] |
| **notify** | **Boolean**| Whether or not to send a notification email at resource creation. | [optional] [default to null] |

### Return type

[**KFS**](../Models/KFS.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createProjectZoneKgw"></a>
# **createProjectZoneKgw**
> KGW createProjectZoneKgw(projectId, zoneId, KGW)



    Creates a new KGW in the specified zone.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the network adapter. | [default to null] |
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |
| **KGW** | [**KGW**](../Models/KGW.md)| KGW payload | |

### Return type

[**KGW**](../Models/KGW.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createProjectZoneVolume"></a>
# **createProjectZoneVolume**
> Volume createProjectZoneVolume(projectId, zoneId, Volume, poolId, templateId)



    Creates a new storage volume in specified zone.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the network adapter. | [default to null] |
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |
| **Volume** | [**Volume**](../Models/Volume.md)| Volume payload | |
| **poolId** | **String**| Storage pool ID (optional, zone&#39;s default if unspecified). | [optional] [default to null] |
| **templateId** | **String**| Template to clone the storage volume from (optional, zone&#39;s default if unspecified). | [optional] [default to null] |

### Return type

[**Volume**](../Models/Volume.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

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

<a name="createStoragePool"></a>
# **createStoragePool**
> StoragePool createStoragePool(zoneId, StoragePool, hostId)



    Creates a new storage pool.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |
| **StoragePool** | [**StoragePool**](../Models/StoragePool.md)| StoragePool payload. | |
| **hostId** | **String**| The ID of the computing host (useless for RBD pools, mandatory for local ones). | [optional] [default to null] |

### Return type

[**StoragePool**](../Models/StoragePool.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createVNet"></a>
# **createVNet**
> VNet createVNet(zoneId, VNet)



    Creates a new virtual network.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |
| **VNet** | [**VNet**](../Models/VNet.md)| VNet payload. | |

### Return type

[**VNet**](../Models/VNet.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteZone"></a>
# **deleteZone**
> deleteZone(zoneId)



    Deletes an existing availability zone.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |

### Return type

null (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getProjectZoneInstances"></a>
# **getProjectZoneInstances**
> List getProjectZoneInstances(projectId, zoneId)



    Returns the IDs of the virtual machine instances existing in the project in the specified zone.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the network adapter. | [default to null] |
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getProjectZoneKCEs"></a>
# **getProjectZoneKCEs**
> List getProjectZoneKCEs(projectId, zoneId)



    Returns the IDs of the KCE virtual machines existing in the project in the specified zone.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the network adapter. | [default to null] |
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getProjectZoneKGWs"></a>
# **getProjectZoneKGWs**
> List getProjectZoneKGWs(projectId, zoneId)



    Returns the IDs of the KGW existing in the project in the specified zone.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the network adapter. | [default to null] |
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getProjectZoneKfs"></a>
# **getProjectZoneKfs**
> List getProjectZoneKfs(projectId, zoneId, nfsId, notify)



    Returns the IDs of the KFS storage volumes existing in the project in the specified zone.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the network adapter. | [default to null] |
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |
| **nfsId** | **String**| NFS storage ID (optional, zone&#39;s default if unspecified). | [optional] [default to null] |
| **notify** | **Boolean**| Whether or not to send a notification email at resource creation. | [optional] [default to null] |

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
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listZoneHosts"></a>
# **listZoneHosts**
> List listZoneHosts(zoneId)



    Returns the IDs of computing host objects.

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

<a name="listZoneNetGWs"></a>
# **listZoneNetGWs**
> List listZoneNetGWs(zoneId)



    Returns the IDs of Iris network gateway objects.

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

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listZoneVNets"></a>
# **listZoneVNets**
> List listZoneVNets(zoneId)



    Returns the IDs of virtual network objects.

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

<a name="listZones"></a>
# **listZones**
> List listZones()



    Returns the IDs of availability zone objects.

### Parameters
This endpoint does not need any parameter.

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="readZone"></a>
# **readZone**
> Zone readZone(zoneId)



    Returns a availability zone.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |

### Return type

[**Zone**](../Models/Zone.md)

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

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateZone"></a>
# **updateZone**
> Zone updateZone(zoneId, Zone)



    Updates a availability zone configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |
| **Zone** | [**Zone**](../Models/Zone.md)| Zone payload. | |

### Return type

[**Zone**](../Models/Zone.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

