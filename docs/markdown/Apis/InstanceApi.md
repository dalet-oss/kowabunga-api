# InstanceApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createProjectZoneInstance**](InstanceApi.md#createProjectZoneInstance) | **POST** /project/{projectId}/zone/{zoneId}/instance |  |
| [**deleteInstance**](InstanceApi.md#deleteInstance) | **DELETE** /instance/{instanceId} |  |
| [**listHostInstances**](InstanceApi.md#listHostInstances) | **GET** /host/{hostId}/instances |  |
| [**listInstances**](InstanceApi.md#listInstances) | **GET** /instance |  |
| [**listProjectZoneInstances**](InstanceApi.md#listProjectZoneInstances) | **GET** /project/{projectId}/zone/{zoneId}/instances |  |
| [**readInstance**](InstanceApi.md#readInstance) | **GET** /instance/{instanceId} |  |
| [**readInstanceRemoteConnection**](InstanceApi.md#readInstanceRemoteConnection) | **GET** /instance/{instanceId}/connect |  |
| [**readInstanceState**](InstanceApi.md#readInstanceState) | **GET** /instance/{instanceId}/state |  |
| [**rebootInstance**](InstanceApi.md#rebootInstance) | **PATCH** /instance/{instanceId}/reboot |  |
| [**resetInstance**](InstanceApi.md#resetInstance) | **PATCH** /instance/{instanceId}/reset |  |
| [**resumeInstance**](InstanceApi.md#resumeInstance) | **PATCH** /instance/{instanceId}/resume |  |
| [**shutdownInstance**](InstanceApi.md#shutdownInstance) | **PATCH** /instance/{instanceId}/shutdown |  |
| [**startInstance**](InstanceApi.md#startInstance) | **PATCH** /instance/{instanceId}/start |  |
| [**stopInstance**](InstanceApi.md#stopInstance) | **PATCH** /instance/{instanceId}/stop |  |
| [**suspendInstance**](InstanceApi.md#suspendInstance) | **PATCH** /instance/{instanceId}/suspend |  |
| [**updateInstance**](InstanceApi.md#updateInstance) | **PUT** /instance/{instanceId} |  |


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

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteInstance"></a>
# **deleteInstance**
> deleteInstance(instanceId)



    Deletes an existing virtual machine instance.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **instanceId** | **String**| The ID of the virtual machine instance. | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listHostInstances"></a>
# **listHostInstances**
> List listHostInstances(hostId)



    Returns the IDs of virtual machine instance objects.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **hostId** | **String**| The ID of the computing host. | [default to null] |

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listInstances"></a>
# **listInstances**
> List listInstances()



    Returns the IDs of virtual machine instance objects.

### Parameters
This endpoint does not need any parameter.

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

<a name="readInstance"></a>
# **readInstance**
> Instance readInstance(instanceId)



    Returns a virtual machine instance.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **instanceId** | **String**| The ID of the virtual machine instance. | [default to null] |

### Return type

[**Instance**](../Models/Instance.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="readInstanceRemoteConnection"></a>
# **readInstanceRemoteConnection**
> InstanceRemoteAccess readInstanceRemoteConnection(instanceId)



    Returns a virtual machine instance remote access characteristics.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **instanceId** | **String**| The ID of the virtual machine instance. | [default to null] |

### Return type

[**InstanceRemoteAccess**](../Models/InstanceRemoteAccess.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="readInstanceState"></a>
# **readInstanceState**
> InstanceState readInstanceState(instanceId)



    Returns a virtual machine instance state.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **instanceId** | **String**| The ID of the virtual machine instance. | [default to null] |

### Return type

[**InstanceState**](../Models/InstanceState.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="rebootInstance"></a>
# **rebootInstance**
> rebootInstance(instanceId)



    Performs a virtual machine instance software reboot.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **instanceId** | **String**| The ID of the virtual machine instance. | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="resetInstance"></a>
# **resetInstance**
> resetInstance(instanceId)



    Performs a virtual machine instance hardware reset.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **instanceId** | **String**| The ID of the virtual machine instance. | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="resumeInstance"></a>
# **resumeInstance**
> resumeInstance(instanceId)



    Performs a virtual machine instance software PM resume.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **instanceId** | **String**| The ID of the virtual machine instance. | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="shutdownInstance"></a>
# **shutdownInstance**
> shutdownInstance(instanceId)



    Performs a virtual machine instance software shutdown.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **instanceId** | **String**| The ID of the virtual machine instance. | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="startInstance"></a>
# **startInstance**
> startInstance(instanceId)



    Performs a virtual machine instance hardware boot-up.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **instanceId** | **String**| The ID of the virtual machine instance. | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="stopInstance"></a>
# **stopInstance**
> stopInstance(instanceId)



    Performs a virtual machine instance hardware stop.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **instanceId** | **String**| The ID of the virtual machine instance. | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="suspendInstance"></a>
# **suspendInstance**
> suspendInstance(instanceId)



    Performs a virtual machine instance software PM suspend.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **instanceId** | **String**| The ID of the virtual machine instance. | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateInstance"></a>
# **updateInstance**
> Instance updateInstance(instanceId, Instance)



    Updates a virtual machine instance configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **instanceId** | **String**| The ID of the virtual machine instance. | [default to null] |
| **Instance** | [**Instance**](../Models/Instance.md)| Instance payload. | |

### Return type

[**Instance**](../Models/Instance.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

