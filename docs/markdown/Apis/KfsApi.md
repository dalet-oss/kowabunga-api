# KfsApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createProjectZoneKfs**](KfsApi.md#createProjectZoneKfs) | **POST** /project/{projectId}/zone/{zoneId}/kfs |  |
| [**deleteKFS**](KfsApi.md#deleteKFS) | **DELETE** /kfs/{kfsId} |  |
| [**getAllKFSs**](KfsApi.md#getAllKFSs) | **GET** /kfs |  |
| [**getKFS**](KfsApi.md#getKFS) | **GET** /kfs/{kfsId} |  |
| [**getNfsKfs**](KfsApi.md#getNfsKfs) | **GET** /nfs/{nfsId}/kfs |  |
| [**getProjectZoneKfs**](KfsApi.md#getProjectZoneKfs) | **GET** /project/{projectId}/zone/{zoneId}/kfs |  |
| [**updateKFS**](KfsApi.md#updateKFS) | **PUT** /kfs/{kfsId} |  |


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

<a name="deleteKFS"></a>
# **deleteKFS**
> deleteKFS(kfsId)



    Deletes an existing KFS storage volume.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kfsId** | **String**| The ID of the KFS storage volume. | [default to null] |

### Return type

null (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getAllKFSs"></a>
# **getAllKFSs**
> List getAllKFSs()



    Returns the IDs of registered KFS storage volumes.

### Parameters
This endpoint does not need any parameter.

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getKFS"></a>
# **getKFS**
> KFS getKFS(kfsId)



    Returns the description of the KFS storage volume.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kfsId** | **String**| The ID of the KFS storage volume. | [default to null] |

### Return type

[**KFS**](../Models/KFS.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getNfsKfs"></a>
# **getNfsKfs**
> List getNfsKfs(nfsId)



    Returns the IDs of the KFS volumes existing in the NFS storage.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **nfsId** | **String**| The ID of the NFS storage. | [default to null] |

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

<a name="updateKFS"></a>
# **updateKFS**
> KFS updateKFS(kfsId, KFS)



    Updates a KFS storage volume configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kfsId** | **String**| The ID of the KFS storage volume. | [default to null] |
| **KFS** | [**KFS**](../Models/KFS.md)| KFS payload | |

### Return type

[**KFS**](../Models/KFS.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

