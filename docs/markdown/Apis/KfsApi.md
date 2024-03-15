# KfsApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createProjectZoneKFS**](KfsApi.md#createProjectZoneKFS) | **POST** /project/{projectId}/zone/{zoneId}/kfs |  |
| [**deleteKFS**](KfsApi.md#deleteKFS) | **DELETE** /kfs/{kfsId} |  |
| [**listKFSs**](KfsApi.md#listKFSs) | **GET** /kfs |  |
| [**listProjectZoneKFSs**](KfsApi.md#listProjectZoneKFSs) | **GET** /project/{projectId}/zone/{zoneId}/kfs |  |
| [**listStorageNFSKFSs**](KfsApi.md#listStorageNFSKFSs) | **GET** /nfs/{nfsId}/kfs |  |
| [**readKFS**](KfsApi.md#readKFS) | **GET** /kfs/{kfsId} |  |
| [**updateKFS**](KfsApi.md#updateKFS) | **PUT** /kfs/{kfsId} |  |


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
| **nfsId** | **String**| NFS storage ID (optional, zone&#39;s default if unspecified). | [optional] [default to null] |

### Return type

[**KFS**](../Models/KFS.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteKFS"></a>
# **deleteKFS**
> deleteKFS(kfsId)



    Deletes an existing KFS (Kowabunga File System).

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kfsId** | **String**| The ID of the KFS (Kowabunga File System). | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listKFSs"></a>
# **listKFSs**
> List listKFSs()



    Returns the IDs of KFS (Kowabunga File System) objects.

### Parameters
This endpoint does not need any parameter.

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
| **nfsId** | **String**| NFS storage ID (optional, zone&#39;s default if unspecified). | [optional] [default to null] |

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listStorageNFSKFSs"></a>
# **listStorageNFSKFSs**
> List listStorageNFSKFSs(nfsId)



    Returns the IDs of KFS (Kowabunga File System) objects.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **nfsId** | **String**| The ID of the NFS storage. | [default to null] |

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="readKFS"></a>
# **readKFS**
> KFS readKFS(kfsId)



    Returns a KFS (Kowabunga File System).

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kfsId** | **String**| The ID of the KFS (Kowabunga File System). | [default to null] |

### Return type

[**KFS**](../Models/KFS.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateKFS"></a>
# **updateKFS**
> KFS updateKFS(kfsId, KFS)



    Updates a KFS (Kowabunga File System) configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kfsId** | **String**| The ID of the KFS (Kowabunga File System). | [default to null] |
| **KFS** | [**KFS**](../Models/KFS.md)| KFS payload. | |

### Return type

[**KFS**](../Models/KFS.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

