# ProjectApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createProject**](ProjectApi.md#createProject) | **POST** /project |  |
| [**createProjectDnsRecord**](ProjectApi.md#createProjectDnsRecord) | **POST** /project/{projectId}/record |  |
| [**createProjectRegionKawaii**](ProjectApi.md#createProjectRegionKawaii) | **POST** /project/{projectId}/region/{regionId}/kawaii |  |
| [**createProjectRegionKonvey**](ProjectApi.md#createProjectRegionKonvey) | **POST** /project/{projectId}/region/{regionId}/konvey |  |
| [**createProjectRegionKylo**](ProjectApi.md#createProjectRegionKylo) | **POST** /project/{projectId}/region/{regionId}/kylo |  |
| [**createProjectRegionVolume**](ProjectApi.md#createProjectRegionVolume) | **POST** /project/{projectId}/region/{regionId}/volume |  |
| [**createProjectZoneInstance**](ProjectApi.md#createProjectZoneInstance) | **POST** /project/{projectId}/zone/{zoneId}/instance |  |
| [**createProjectZoneKompute**](ProjectApi.md#createProjectZoneKompute) | **POST** /project/{projectId}/zone/{zoneId}/kompute |  |
| [**createProjectZoneKonvey**](ProjectApi.md#createProjectZoneKonvey) | **POST** /project/{projectId}/zone/{zoneId}/konvey |  |
| [**deleteProject**](ProjectApi.md#deleteProject) | **DELETE** /project/{projectId} |  |
| [**listProjectDnsRecords**](ProjectApi.md#listProjectDnsRecords) | **GET** /project/{projectId}/records |  |
| [**listProjectRegionKawaiis**](ProjectApi.md#listProjectRegionKawaiis) | **GET** /project/{projectId}/region/{regionId}/kawaiis |  |
| [**listProjectRegionKonveys**](ProjectApi.md#listProjectRegionKonveys) | **GET** /project/{projectId}/region/{regionId}/konveys |  |
| [**listProjectRegionKylos**](ProjectApi.md#listProjectRegionKylos) | **GET** /project/{projectId}/region/{regionId}/kylo |  |
| [**listProjectRegionVolumes**](ProjectApi.md#listProjectRegionVolumes) | **GET** /project/{projectId}/region/{regionId}/volumes |  |
| [**listProjectZoneInstances**](ProjectApi.md#listProjectZoneInstances) | **GET** /project/{projectId}/zone/{zoneId}/instances |  |
| [**listProjectZoneKomputes**](ProjectApi.md#listProjectZoneKomputes) | **GET** /project/{projectId}/zone/{zoneId}/komputes |  |
| [**listProjectZoneKonveys**](ProjectApi.md#listProjectZoneKonveys) | **GET** /project/{projectId}/zone/{zoneId}/konveys |  |
| [**listProjects**](ProjectApi.md#listProjects) | **GET** /project |  |
| [**readProject**](ProjectApi.md#readProject) | **GET** /project/{projectId} |  |
| [**readProjectCost**](ProjectApi.md#readProjectCost) | **GET** /project/{projectId}/cost |  |
| [**readProjectUsage**](ProjectApi.md#readProjectUsage) | **GET** /project/{projectId}/usage |  |
| [**updateProject**](ProjectApi.md#updateProject) | **PUT** /project/{projectId} |  |


<a name="createProject"></a>
# **createProject**
> Project createProject(Project, subnetSize)



    Creates a new project.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **Project** | [**Project**](../Models/Project.md)| Project payload. | |
| **subnetSize** | **Integer**| The minimum VPC subnet size to be affected to the project. WARNING, this cannot be changed later. | [optional] [default to null] |

### Return type

[**Project**](../Models/Project.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createProjectDnsRecord"></a>
# **createProjectDnsRecord**
> DnsRecord createProjectDnsRecord(projectId, DnsRecord)



    Creates a new DNS record.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |
| **DnsRecord** | [**DnsRecord**](../Models/DnsRecord.md)| DnsRecord payload. | |

### Return type

[**DnsRecord**](../Models/DnsRecord.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createProjectRegionKawaii"></a>
# **createProjectRegionKawaii**
> Kawaii createProjectRegionKawaii(projectId, regionId, Kawaii)



    Creates a new Kawaii.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |
| **regionId** | **String**| The ID of the region. | [default to null] |
| **Kawaii** | [**Kawaii**](../Models/Kawaii.md)| Kawaii payload. | |

### Return type

[**Kawaii**](../Models/Kawaii.md)

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

<a name="createProjectRegionKylo"></a>
# **createProjectRegionKylo**
> Kylo createProjectRegionKylo(projectId, regionId, Kylo, nfsId)



    Creates a new Kylo.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |
| **regionId** | **String**| The ID of the region. | [default to null] |
| **Kylo** | [**Kylo**](../Models/Kylo.md)| Kylo payload. | |
| **nfsId** | **String**| NFS storage ID (optional, region&#39;s default if unspecified). | [optional] [default to null] |

### Return type

[**Kylo**](../Models/Kylo.md)

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

<a name="createProjectZoneKompute"></a>
# **createProjectZoneKompute**
> Kompute createProjectZoneKompute(projectId, zoneId, Kompute, poolId, templateId, public)



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

<a name="createProjectZoneKonvey"></a>
# **createProjectZoneKonvey**
> Konvey createProjectZoneKonvey(projectId, zoneId, Konvey)



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

<a name="deleteProject"></a>
# **deleteProject**
> deleteProject(projectId)



    Deletes an existing project.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listProjectDnsRecords"></a>
# **listProjectDnsRecords**
> List listProjectDnsRecords(projectId)



    Returns the IDs of DNS record objects.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listProjectRegionKawaiis"></a>
# **listProjectRegionKawaiis**
> List listProjectRegionKawaiis(projectId, regionId)



    Returns the IDs of Kawaii objects.

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

<a name="listProjectRegionKylos"></a>
# **listProjectRegionKylos**
> List listProjectRegionKylos(projectId, regionId, nfsId)



    Returns the IDs of Kylo objects.

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

<a name="listProjectZoneKomputes"></a>
# **listProjectZoneKomputes**
> List listProjectZoneKomputes(projectId, zoneId)



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

<a name="listProjectZoneKonveys"></a>
# **listProjectZoneKonveys**
> List listProjectZoneKonveys(projectId, zoneId)



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

<a name="listProjects"></a>
# **listProjects**
> List listProjects(subnetSize)



    Returns the IDs of project objects.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **subnetSize** | **Integer**| The minimum VPC subnet size to be affected to the project. WARNING, this cannot be changed later. | [optional] [default to null] |

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="readProject"></a>
# **readProject**
> Project readProject(projectId)



    Returns a project.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |

### Return type

[**Project**](../Models/Project.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="readProjectCost"></a>
# **readProjectCost**
> Cost readProjectCost(projectId)



    Returns a resource cost.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |

### Return type

[**Cost**](../Models/Cost.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="readProjectUsage"></a>
# **readProjectUsage**
> ProjectResources readProjectUsage(projectId)



    Returns a global project resource quotas/usage (0 for unlimited).

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |

### Return type

[**ProjectResources**](../Models/ProjectResources.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateProject"></a>
# **updateProject**
> Project updateProject(projectId, Project)



    Updates a project configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |
| **Project** | [**Project**](../Models/Project.md)| Project payload. | |

### Return type

[**Project**](../Models/Project.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

