# InstanceApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createProjectZoneInstance**](InstanceApi.md#createProjectZoneInstance) | **POST** /project/{projectId}/zone/{zoneId}/instance |  |
| [**deleteInstance**](InstanceApi.md#deleteInstance) | **DELETE** /instance/{instanceId} |  |
| [**getAllInstances**](InstanceApi.md#getAllInstances) | **GET** /instance |  |
| [**getHostInstances**](InstanceApi.md#getHostInstances) | **GET** /host/{hostId}/instances |  |
| [**getInstance**](InstanceApi.md#getInstance) | **GET** /instance/{instanceId} |  |
| [**getInstanceRemoteConnection**](InstanceApi.md#getInstanceRemoteConnection) | **GET** /instance/{instanceId}/connect |  |
| [**getInstanceState**](InstanceApi.md#getInstanceState) | **GET** /instance/{instanceId}/state |  |
| [**getProjectZoneInstances**](InstanceApi.md#getProjectZoneInstances) | **GET** /project/{projectId}/zone/{zoneId}/instances |  |
| [**rebootInstance**](InstanceApi.md#rebootInstance) | **POST** /instance/{instanceId}/reboot |  |
| [**resetInstance**](InstanceApi.md#resetInstance) | **POST** /instance/{instanceId}/reset |  |
| [**resumeInstance**](InstanceApi.md#resumeInstance) | **POST** /instance/{instanceId}/resume |  |
| [**shutdownInstance**](InstanceApi.md#shutdownInstance) | **POST** /instance/{instanceId}/shutdown |  |
| [**startInstance**](InstanceApi.md#startInstance) | **POST** /instance/{instanceId}/start |  |
| [**stopInstance**](InstanceApi.md#stopInstance) | **POST** /instance/{instanceId}/stop |  |
| [**suspendInstance**](InstanceApi.md#suspendInstance) | **POST** /instance/{instanceId}/suspend |  |
| [**updateInstance**](InstanceApi.md#updateInstance) | **PUT** /instance/{instanceId} |  |


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

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getAllInstances"></a>
# **getAllInstances**
> List getAllInstances()



    Returns the IDs of registered virtual machines.

### Parameters
This endpoint does not need any parameter.

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getHostInstances"></a>
# **getHostInstances**
> List getHostInstances(hostId)



    Returns the UUIDs of the virtual machines running on the host.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **hostId** | **String**| The ID of the computing host. | [default to null] |

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getInstance"></a>
# **getInstance**
> Instance getInstance(instanceId)



    Returns the description of the virtual machine.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **instanceId** | **String**| The ID of the virtual machine instance. | [default to null] |

### Return type

[**Instance**](../Models/Instance.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getInstanceRemoteConnection"></a>
# **getInstanceRemoteConnection**
> InstanceRemoteAccess getInstanceRemoteConnection(instanceId)



    Returns virtual machine remote access URL.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **instanceId** | **String**| The ID of the virtual machine instance. | [default to null] |

### Return type

[**InstanceRemoteAccess**](../Models/InstanceRemoteAccess.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getInstanceState"></a>
# **getInstanceState**
> InstanceState getInstanceState(instanceId)



    Returns the state of the virtual machine.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **instanceId** | **String**| The ID of the virtual machine instance. | [default to null] |

### Return type

[**InstanceState**](../Models/InstanceState.md)

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

<a name="rebootInstance"></a>
# **rebootInstance**
> rebootInstance(instanceId)



    Perform a virtual machine software reboot.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **instanceId** | **String**| The ID of the virtual machine instance. | [default to null] |

### Return type

null (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="resetInstance"></a>
# **resetInstance**
> resetInstance(instanceId)



    Perform a virtual machine hardware reset.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **instanceId** | **String**| The ID of the virtual machine instance. | [default to null] |

### Return type

null (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="resumeInstance"></a>
# **resumeInstance**
> resumeInstance(instanceId)



    Perform a virtual machine software PM resume.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **instanceId** | **String**| The ID of the virtual machine instance. | [default to null] |

### Return type

null (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="shutdownInstance"></a>
# **shutdownInstance**
> shutdownInstance(instanceId)



    Initiate a software shutdown of a virtual machine.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **instanceId** | **String**| The ID of the virtual machine instance. | [default to null] |

### Return type

null (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="startInstance"></a>
# **startInstance**
> startInstance(instanceId)



    Boot up a virtual machine.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **instanceId** | **String**| The ID of the virtual machine instance. | [default to null] |

### Return type

null (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="stopInstance"></a>
# **stopInstance**
> stopInstance(instanceId)



    Initiate a hardware stop of a virtual machine.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **instanceId** | **String**| The ID of the virtual machine instance. | [default to null] |

### Return type

null (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="suspendInstance"></a>
# **suspendInstance**
> suspendInstance(instanceId)



    Perform a virtual machine software PM suspend.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **instanceId** | **String**| The ID of the virtual machine instance. | [default to null] |

### Return type

null (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateInstance"></a>
# **updateInstance**
> Adapter updateInstance(instanceId, Instance)



    Updates a virtual machine configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **instanceId** | **String**| The ID of the virtual machine instance. | [default to null] |
| **Instance** | [**Instance**](../Models/Instance.md)| Adapter payload | |

### Return type

[**Adapter**](../Models/Adapter.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

