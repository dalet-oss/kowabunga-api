# KceApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createProjectZoneKce**](KceApi.md#createProjectZoneKce) | **POST** /project/{projectId}/zone/{zoneId}/kce |  |
| [**deleteKCE**](KceApi.md#deleteKCE) | **DELETE** /kce/{kceId} |  |
| [**getAllKCEs**](KceApi.md#getAllKCEs) | **GET** /kce |  |
| [**getKCE**](KceApi.md#getKCE) | **GET** /kce/{kceId} |  |
| [**getKCEState**](KceApi.md#getKCEState) | **GET** /kce/{kceId}/state |  |
| [**getProjectZoneKCEs**](KceApi.md#getProjectZoneKCEs) | **GET** /project/{projectId}/zone/{zoneId}/kces |  |
| [**rebootKCE**](KceApi.md#rebootKCE) | **POST** /kce/{kceId}/reboot |  |
| [**resetKCE**](KceApi.md#resetKCE) | **POST** /kce/{kceId}/reset |  |
| [**resumeKCE**](KceApi.md#resumeKCE) | **POST** /kce/{kceId}/resume |  |
| [**shutdownKCE**](KceApi.md#shutdownKCE) | **POST** /kce/{kceId}/shutdown |  |
| [**startKCE**](KceApi.md#startKCE) | **POST** /kce/{kceId}/start |  |
| [**stopKCE**](KceApi.md#stopKCE) | **POST** /kce/{kceId}/stop |  |
| [**suspendKCE**](KceApi.md#suspendKCE) | **POST** /kce/{kceId}/suspend |  |
| [**updateKCE**](KceApi.md#updateKCE) | **PUT** /kce/{kceId} |  |


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
| **notify** | **Boolean**| Whether or not to send a notification email at resource creation | [optional] [default to null] |

### Return type

[**KCE**](../Models/KCE.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteKCE"></a>
# **deleteKCE**
> deleteKCE(kceId)



    Deletes an existing KCE virtual machine.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kceId** | **String**| The ID of the KCE virtual machine. | [default to null] |

### Return type

null (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getAllKCEs"></a>
# **getAllKCEs**
> List getAllKCEs()



    Returns the IDs of registered KCE virtual machines.

### Parameters
This endpoint does not need any parameter.

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getKCE"></a>
# **getKCE**
> KCE getKCE(kceId)



    Returns the description of the KCE virtual machine.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kceId** | **String**| The ID of the KCE virtual machine. | [default to null] |

### Return type

[**KCE**](../Models/KCE.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getKCEState"></a>
# **getKCEState**
> InstanceState getKCEState(kceId)



    Returns the state of the KCE virtual machine.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kceId** | **String**| The ID of the KCE virtual machine. | [default to null] |

### Return type

[**InstanceState**](../Models/InstanceState.md)

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

<a name="rebootKCE"></a>
# **rebootKCE**
> rebootKCE(kceId)



    Perform a KCE virtual machine software reboot.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kceId** | **String**| The ID of the KCE virtual machine. | [default to null] |

### Return type

null (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="resetKCE"></a>
# **resetKCE**
> resetKCE(kceId)



    Perform a KCE virtual machine hardware reset.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kceId** | **String**| The ID of the KCE virtual machine. | [default to null] |

### Return type

null (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="resumeKCE"></a>
# **resumeKCE**
> resumeKCE(kceId)



    Perform a KCE virtual machine software PM resume.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kceId** | **String**| The ID of the KCE virtual machine. | [default to null] |

### Return type

null (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="shutdownKCE"></a>
# **shutdownKCE**
> shutdownKCE(kceId)



    Initiate a software shutdown of a KCE virtual machine.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kceId** | **String**| The ID of the KCE virtual machine. | [default to null] |

### Return type

null (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="startKCE"></a>
# **startKCE**
> startKCE(kceId)



    Boot up a KCE virtual machine.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kceId** | **String**| The ID of the KCE virtual machine. | [default to null] |

### Return type

null (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="stopKCE"></a>
# **stopKCE**
> stopKCE(kceId)



    Initiate a hardware stop of a KCE virtual machine.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kceId** | **String**| The ID of the KCE virtual machine. | [default to null] |

### Return type

null (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="suspendKCE"></a>
# **suspendKCE**
> suspendKCE(kceId)



    Perform a KCE virtual machine software PM suspend.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kceId** | **String**| The ID of the KCE virtual machine. | [default to null] |

### Return type

null (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateKCE"></a>
# **updateKCE**
> KCE updateKCE(kceId, KCE)



    Updates a KCE virtual machine configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kceId** | **String**| The ID of the KCE virtual machine. | [default to null] |
| **KCE** | [**KCE**](../Models/KCE.md)| KCE payload | |

### Return type

[**KCE**](../Models/KCE.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

