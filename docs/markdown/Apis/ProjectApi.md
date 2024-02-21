# ProjectApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createProject**](ProjectApi.md#createProject) | **POST** /project |  |
| [**createProjectDnsRecord**](ProjectApi.md#createProjectDnsRecord) | **POST** /project/{projectId}/record |  |
| [**createProjectZoneInstance**](ProjectApi.md#createProjectZoneInstance) | **POST** /project/{projectId}/zone/{zoneId}/instance |  |
| [**createProjectZoneKCE**](ProjectApi.md#createProjectZoneKCE) | **POST** /project/{projectId}/zone/{zoneId}/kce |  |
| [**createProjectZoneKFS**](ProjectApi.md#createProjectZoneKFS) | **POST** /project/{projectId}/zone/{zoneId}/kfs |  |
| [**createProjectZoneKGW**](ProjectApi.md#createProjectZoneKGW) | **POST** /project/{projectId}/zone/{zoneId}/kgw |  |
| [**createProjectZoneVolume**](ProjectApi.md#createProjectZoneVolume) | **POST** /project/{projectId}/zone/{zoneId}/volume |  |
| [**deleteProject**](ProjectApi.md#deleteProject) | **DELETE** /project/{projectId} |  |
| [**listProjectDnsRecords**](ProjectApi.md#listProjectDnsRecords) | **GET** /project/{projectId}/records |  |
| [**listProjectZoneInstances**](ProjectApi.md#listProjectZoneInstances) | **GET** /project/{projectId}/zone/{zoneId}/instances |  |
| [**listProjectZoneKCEs**](ProjectApi.md#listProjectZoneKCEs) | **GET** /project/{projectId}/zone/{zoneId}/kces |  |
| [**listProjectZoneKFSs**](ProjectApi.md#listProjectZoneKFSs) | **GET** /project/{projectId}/zone/{zoneId}/kfs |  |
| [**listProjectZoneKGWs**](ProjectApi.md#listProjectZoneKGWs) | **GET** /project/{projectId}/zone/{zoneId}/kgws |  |
| [**listProjectZoneVolumes**](ProjectApi.md#listProjectZoneVolumes) | **GET** /project/{projectId}/zone/{zoneId}/volumes |  |
| [**listProjects**](ProjectApi.md#listProjects) | **GET** /project |  |
| [**readProject**](ProjectApi.md#readProject) | **GET** /project/{projectId} |  |
| [**readProjectCost**](ProjectApi.md#readProjectCost) | **GET** /project/{projectId}/cost |  |
| [**readProjectUsage**](ProjectApi.md#readProjectUsage) | **GET** /project/{projectId}/usage |  |
| [**updateProject**](ProjectApi.md#updateProject) | **PUT** /project/{projectId} |  |


<a name="createProject"></a>
# **createProject**
> Project createProject(Project, subnetSize, notify)



    Creates a new project.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **Project** | [**Project**](../Models/Project.md)| Project payload. | |
| **subnetSize** | **Integer**| The minimum VPC subnet size to be affected to the project. WARNING, this cannot be changed later. | [optional] [default to null] |
| **notify** | **Boolean**| Whether or not to send a notification email at resource creation. | [optional] [default to null] |

### Return type

[**Project**](../Models/Project.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createProjectZoneInstance"></a>
# **createProjectZoneInstance**
> Instance createProjectZoneInstance(projectId, zoneId, Instance, notify)



    Creates a new virtual machine instance.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |
| **Instance** | [**Instance**](../Models/Instance.md)| Instance payload. | |
| **notify** | **Boolean**| Whether or not to send a notification email at resource creation. | [optional] [default to null] |

### Return type

[**Instance**](../Models/Instance.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createProjectZoneKCE"></a>
# **createProjectZoneKCE**
> KCE createProjectZoneKCE(projectId, zoneId, KCE, poolId, templateId, public, notify)



    Creates a new KCE (Kowabunga Compute Engine).

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |
| **KCE** | [**KCE**](../Models/KCE.md)| KCE payload. | |
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

<a name="createProjectZoneKFS"></a>
# **createProjectZoneKFS**
> KFS createProjectZoneKFS(projectId, zoneId, KFS, nfsId, notify)



    Creates a new KFS (Kowabunga File System).

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |
| **KFS** | [**KFS**](../Models/KFS.md)| KFS payload. | |
| **nfsId** | **String**| NFS storage ID (optional, zone&#39;s default if unspecified). | [optional] [default to null] |
| **notify** | **Boolean**| Whether or not to send a notification email at resource creation. | [optional] [default to null] |

### Return type

[**KFS**](../Models/KFS.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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
| **poolId** | **String**| Storage pool ID (optional, zone&#39;s default if unspecified). | [optional] [default to null] |
| **templateId** | **String**| Template to clone the storage volume from (optional, zone&#39;s default if unspecified). | [optional] [default to null] |

### Return type

[**Volume**](../Models/Volume.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listProjectZoneKFSs"></a>
# **listProjectZoneKFSs**
> List listProjectZoneKFSs(projectId, zoneId, nfsId, notify)



    Returns the IDs of KFS (Kowabunga File System) objects.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |
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

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listProjects"></a>
# **listProjects**
> List listProjects(subnetSize, notify)



    Returns the IDs of project objects.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **subnetSize** | **Integer**| The minimum VPC subnet size to be affected to the project. WARNING, this cannot be changed later. | [optional] [default to null] |
| **notify** | **Boolean**| Whether or not to send a notification email at resource creation. | [optional] [default to null] |

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

