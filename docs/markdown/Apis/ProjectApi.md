# ProjectApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createProject**](ProjectApi.md#createProject) | **POST** /project |  |
| [**createProjectDnsRecord**](ProjectApi.md#createProjectDnsRecord) | **POST** /project/{projectId}/record |  |
| [**createProjectZoneInstance**](ProjectApi.md#createProjectZoneInstance) | **POST** /project/{projectId}/zone/{zoneId}/instance |  |
| [**createProjectZoneKce**](ProjectApi.md#createProjectZoneKce) | **POST** /project/{projectId}/zone/{zoneId}/kce |  |
| [**createProjectZoneKfs**](ProjectApi.md#createProjectZoneKfs) | **POST** /project/{projectId}/zone/{zoneId}/kfs |  |
| [**createProjectZoneKgw**](ProjectApi.md#createProjectZoneKgw) | **POST** /project/{projectId}/zone/{zoneId}/kgw |  |
| [**createProjectZoneVolume**](ProjectApi.md#createProjectZoneVolume) | **POST** /project/{projectId}/zone/{zoneId}/volume |  |
| [**deleteProject**](ProjectApi.md#deleteProject) | **DELETE** /project/{projectId} |  |
| [**getAllProjects**](ProjectApi.md#getAllProjects) | **GET** /project |  |
| [**getProject**](ProjectApi.md#getProject) | **GET** /project/{projectId} |  |
| [**getProjectCost**](ProjectApi.md#getProjectCost) | **GET** /project/{projectId}/cost |  |
| [**getProjectDnsRecords**](ProjectApi.md#getProjectDnsRecords) | **GET** /project/{projectId}/records |  |
| [**getProjectUsage**](ProjectApi.md#getProjectUsage) | **GET** /project/{projectId}/usage |  |
| [**getProjectZoneInstances**](ProjectApi.md#getProjectZoneInstances) | **GET** /project/{projectId}/zone/{zoneId}/instances |  |
| [**getProjectZoneKCEs**](ProjectApi.md#getProjectZoneKCEs) | **GET** /project/{projectId}/zone/{zoneId}/kces |  |
| [**getProjectZoneKGWs**](ProjectApi.md#getProjectZoneKGWs) | **GET** /project/{projectId}/zone/{zoneId}/kgws |  |
| [**getProjectZoneKfs**](ProjectApi.md#getProjectZoneKfs) | **GET** /project/{projectId}/zone/{zoneId}/kfs |  |
| [**getProjectZoneVolumes**](ProjectApi.md#getProjectZoneVolumes) | **GET** /project/{projectId}/zone/{zoneId}/volumes |  |
| [**updateProject**](ProjectApi.md#updateProject) | **PUT** /project/{projectId} |  |


<a name="createProject"></a>
# **createProject**
> Project createProject(Project, subnetSize, notify)



    Creates a new project.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **Project** | [**Project**](../Models/Project.md)| Project payload | |
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



    Creates a new DNS record in specified project.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the network adapter. | [default to null] |
| **DnsRecord** | [**DnsRecord**](../Models/DnsRecord.md)| DNS record payload | |

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



    Creates a new virtual machine instance in specified zone.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the network adapter. | [default to null] |
| **zoneId** | **String**| The ID of the zone. | [default to null] |
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
| **zoneId** | **String**| The ID of the zone. | [default to null] |
| **KCE** | [**KCE**](../Models/KCE.md)| KCE payload | |
| **poolId** | **String**| Storage pool ID (optional, zone&#39;s default if unspecified). | [optional] [default to null] |
| **templateId** | **String**| Template to clone the storage volume from (optional, zone&#39;s default if unspecified) | [optional] [default to null] |
| **public** | **Boolean**| Should KCE be exposed over public Internet ? (a public IPv4 address will then be auto-assigned, default to false). | [optional] [default to null] |
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
| **zoneId** | **String**| The ID of the zone. | [default to null] |
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
| **zoneId** | **String**| The ID of the zone. | [default to null] |
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

<a name="deleteProject"></a>
# **deleteProject**
> deleteProject(projectId)



    Deletes an existing project.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the network adapter. | [default to null] |

### Return type

null (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getAllProjects"></a>
# **getAllProjects**
> List getAllProjects()



    Returns the IDs of registered projects.

### Parameters
This endpoint does not need any parameter.

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getProject"></a>
# **getProject**
> Project getProject(projectId)



    Returns a description of the project

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the network adapter. | [default to null] |

### Return type

[**Project**](../Models/Project.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getProjectCost"></a>
# **getProjectCost**
> Cost getProjectCost(projectId)



    Returns the current cost for the project.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the network adapter. | [default to null] |

### Return type

[**Cost**](../Models/Cost.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getProjectDnsRecords"></a>
# **getProjectDnsRecords**
> List getProjectDnsRecords(projectId)



    Returns the IDs of the DNS records existing in the project.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the network adapter. | [default to null] |

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getProjectUsage"></a>
# **getProjectUsage**
> ProjectResources getProjectUsage(projectId)



    Returns the current resources usage for the project.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the network adapter. | [default to null] |

### Return type

[**ProjectResources**](../Models/ProjectResources.md)

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
| **zoneId** | **String**| The ID of the zone. | [default to null] |

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
| **zoneId** | **String**| The ID of the zone. | [default to null] |

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
| **zoneId** | **String**| The ID of the zone. | [default to null] |

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
| **zoneId** | **String**| The ID of the zone. | [default to null] |
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
| **zoneId** | **String**| The ID of the zone. | [default to null] |

### Return type

**List**

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
| **projectId** | **String**| The ID of the network adapter. | [default to null] |
| **Project** | [**Project**](../Models/Project.md)| Project payload | |

### Return type

[**Project**](../Models/Project.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

