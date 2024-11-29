# ZoneApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createKaktus**](ZoneApi.md#createKaktus) | **POST** /zone/{zoneId}/kaktus |  |
| [**createProjectZoneInstance_0**](ZoneApi.md#createProjectZoneInstance_0) | **POST** /project/{projectId}/zone/{zoneId}/instance |  |
| [**createProjectZoneKompute_0**](ZoneApi.md#createProjectZoneKompute_0) | **POST** /project/{projectId}/zone/{zoneId}/kompute |  |
| [**createProjectZoneKonvey_0**](ZoneApi.md#createProjectZoneKonvey_0) | **POST** /project/{projectId}/zone/{zoneId}/konvey |  |
| [**createZone_0**](ZoneApi.md#createZone_0) | **POST** /region/{regionId}/zone |  |
| [**deleteZone**](ZoneApi.md#deleteZone) | **DELETE** /zone/{zoneId} |  |
| [**listProjectZoneInstances_0**](ZoneApi.md#listProjectZoneInstances_0) | **GET** /project/{projectId}/zone/{zoneId}/instances |  |
| [**listProjectZoneKomputes_0**](ZoneApi.md#listProjectZoneKomputes_0) | **GET** /project/{projectId}/zone/{zoneId}/komputes |  |
| [**listProjectZoneKonveys_0**](ZoneApi.md#listProjectZoneKonveys_0) | **GET** /project/{projectId}/zone/{zoneId}/konveys |  |
| [**listRegionZones_0**](ZoneApi.md#listRegionZones_0) | **GET** /region/{regionId}/zones |  |
| [**listZoneKaktuses**](ZoneApi.md#listZoneKaktuses) | **GET** /zone/{zoneId}/kaktuses |  |
| [**listZones**](ZoneApi.md#listZones) | **GET** /zone |  |
| [**readZone**](ZoneApi.md#readZone) | **GET** /zone/{zoneId} |  |
| [**updateZone**](ZoneApi.md#updateZone) | **PUT** /zone/{zoneId} |  |


<a name="createKaktus"></a>
# **createKaktus**
> Kaktus createKaktus(zoneId, Kaktus)



    Creates a new Kaktus computing node.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |
| **Kaktus** | [**Kaktus**](../Models/Kaktus.md)| Kaktus payload. | |

### Return type

[**Kaktus**](../Models/Kaktus.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createProjectZoneInstance_0"></a>
# **createProjectZoneInstance_0**
> Instance createProjectZoneInstance_0(projectId, zoneId, Instance)



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

<a name="createProjectZoneKompute_0"></a>
# **createProjectZoneKompute_0**
> Kompute createProjectZoneKompute_0(projectId, zoneId, Kompute, poolId, templateId, public)



    Creates a new Kompute.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |
| **Kompute** | [**Kompute**](../Models/Kompute.md)| Kompute payload. | |
| **poolId** | **String**| Storage pool ID (optional, region&#39;s default if unspecified). | [optional] [default to null] |
| **templateId** | **String**| Template to clone the storage volume from (optional, region&#39;s default if unspecified). | [optional] [default to null] |
| **public** | **Boolean**| Should Kompute be exposed over public Internet ? (a public IPv4 address will then be auto-assigned, default to false). | [optional] [default to null] |

### Return type

[**Kompute**](../Models/Kompute.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createProjectZoneKonvey_0"></a>
# **createProjectZoneKonvey_0**
> Konvey createProjectZoneKonvey_0(projectId, zoneId, Konvey)



    Creates a new Konvey (Kowabunga Network Load-Balancer).

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |
| **Konvey** | [**Konvey**](../Models/Konvey.md)| Konvey payload. | |

### Return type

[**Konvey**](../Models/Konvey.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createZone_0"></a>
# **createZone_0**
> Zone createZone_0(regionId, Zone)



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

<a name="listProjectZoneInstances_0"></a>
# **listProjectZoneInstances_0**
> List listProjectZoneInstances_0(projectId, zoneId)



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

<a name="listProjectZoneKomputes_0"></a>
# **listProjectZoneKomputes_0**
> List listProjectZoneKomputes_0(projectId, zoneId)



    Returns the IDs of Kompute objects.

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

<a name="listProjectZoneKonveys_0"></a>
# **listProjectZoneKonveys_0**
> List listProjectZoneKonveys_0(projectId, zoneId)



    Returns the IDs of Konvey (Kowabunga Network Load-Balancer) objects.

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

<a name="listRegionZones_0"></a>
# **listRegionZones_0**
> List listRegionZones_0(regionId)



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

<a name="listZoneKaktuses"></a>
# **listZoneKaktuses**
> List listZoneKaktuses(zoneId)



    Returns the IDs of Kaktus computing node objects.

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

