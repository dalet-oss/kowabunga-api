# ZoneApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createHost**](ZoneApi.md#createHost) | **POST** /zone/{zoneId}/host |  |
| [**createProjectZoneInstance**](ZoneApi.md#createProjectZoneInstance) | **POST** /project/{projectId}/zone/{zoneId}/instance |  |
| [**createProjectZoneKCE**](ZoneApi.md#createProjectZoneKCE) | **POST** /project/{projectId}/zone/{zoneId}/kce |  |
| [**createProjectZoneKFS**](ZoneApi.md#createProjectZoneKFS) | **POST** /project/{projectId}/zone/{zoneId}/kfs |  |
| [**createProjectZoneKGW**](ZoneApi.md#createProjectZoneKGW) | **POST** /project/{projectId}/zone/{zoneId}/kgw |  |
| [**createProjectZoneVolume**](ZoneApi.md#createProjectZoneVolume) | **POST** /project/{projectId}/zone/{zoneId}/volume |  |
| [**createZone**](ZoneApi.md#createZone) | **POST** /region/{regionId}/zone |  |
| [**deleteZone**](ZoneApi.md#deleteZone) | **DELETE** /zone/{zoneId} |  |
| [**listProjectZoneInstances**](ZoneApi.md#listProjectZoneInstances) | **GET** /project/{projectId}/zone/{zoneId}/instances |  |
| [**listProjectZoneKCEs**](ZoneApi.md#listProjectZoneKCEs) | **GET** /project/{projectId}/zone/{zoneId}/kces |  |
| [**listProjectZoneKFSs**](ZoneApi.md#listProjectZoneKFSs) | **GET** /project/{projectId}/zone/{zoneId}/kfs |  |
| [**listProjectZoneKGWs**](ZoneApi.md#listProjectZoneKGWs) | **GET** /project/{projectId}/zone/{zoneId}/kgws |  |
| [**listProjectZoneVolumes**](ZoneApi.md#listProjectZoneVolumes) | **GET** /project/{projectId}/zone/{zoneId}/volumes |  |
| [**listRegionZones**](ZoneApi.md#listRegionZones) | **GET** /region/{regionId}/zones |  |
| [**listZoneHosts**](ZoneApi.md#listZoneHosts) | **GET** /zone/{zoneId}/hosts |  |
| [**listZones**](ZoneApi.md#listZones) | **GET** /zone |  |
| [**readZone**](ZoneApi.md#readZone) | **GET** /zone/{zoneId} |  |
| [**updateZone**](ZoneApi.md#updateZone) | **PUT** /zone/{zoneId} |  |


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

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createProjectZoneInstance"></a>
# **createProjectZoneInstance**
> Instance createProjectZoneInstance(projectId, zoneId, Instance)



    Creates a new virtual machine instance.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |
| **Instance** | [**Instance**](../Models/Instance.md)| Instance payload. | |

### Return type

[**Instance**](../Models/Instance.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createProjectZoneKCE"></a>
# **createProjectZoneKCE**
> KCE createProjectZoneKCE(projectId, zoneId, KCE, poolId, templateId, public)



    Creates a new KCE (Kowabunga Compute Engine).

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |
| **KCE** | [**KCE**](../Models/KCE.md)| KCE payload. | |
| **poolId** | **String**| Storage pool ID (optional, region&#39;s default if unspecified). | [optional] [default to null] |
| **templateId** | **String**| Template to clone the storage volume from (optional, region&#39;s default if unspecified). | [optional] [default to null] |
| **public** | **Boolean**| Should KCE be exposed over public Internet ? (a public IPv4 address will then be auto-assigned, default to false). | [optional] [default to null] |

### Return type

[**KCE**](../Models/KCE.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createProjectZoneKFS"></a>
# **createProjectZoneKFS**
> KFS createProjectZoneKFS(projectId, zoneId, KFS, nfsId)



    Creates a new KFS (Kowabunga File System).

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |
| **KFS** | [**KFS**](../Models/KFS.md)| KFS payload. | |
| **nfsId** | **String**| NFS storage ID (optional, region&#39;s default if unspecified). | [optional] [default to null] |

### Return type

[**KFS**](../Models/KFS.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createProjectZoneKGW"></a>
# **createProjectZoneKGW**
> KGW createProjectZoneKGW(projectId, zoneId, KGW)



    Creates a new KGW (Kowabunga Network Gateway).

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |
| **KGW** | [**KGW**](../Models/KGW.md)| KGW payload. | |

### Return type

[**KGW**](../Models/KGW.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createProjectZoneVolume"></a>
# **createProjectZoneVolume**
> Volume createProjectZoneVolume(projectId, zoneId, Volume, poolId, templateId)



    Creates a new storage volume.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |
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

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listProjectZoneInstances"></a>
# **listProjectZoneInstances**
> List listProjectZoneInstances(projectId, zoneId)



    Returns the IDs of virtual machine instance objects.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listProjectZoneKCEs"></a>
# **listProjectZoneKCEs**
> List listProjectZoneKCEs(projectId, zoneId)



    Returns the IDs of KCE (Kowabunga Compute Engine) objects.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listProjectZoneKFSs"></a>
# **listProjectZoneKFSs**
> List listProjectZoneKFSs(projectId, zoneId, nfsId)



    Returns the IDs of KFS (Kowabunga File System) objects.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |
| **nfsId** | **String**| NFS storage ID (optional, region&#39;s default if unspecified). | [optional] [default to null] |

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listProjectZoneKGWs"></a>
# **listProjectZoneKGWs**
> List listProjectZoneKGWs(projectId, zoneId)



    Returns the IDs of KGW (Kowabunga Network Gateway) objects.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listProjectZoneVolumes"></a>
# **listProjectZoneVolumes**
> List listProjectZoneVolumes(projectId, zoneId)



    Returns the IDs of storage volume objects.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |

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

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

