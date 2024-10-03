# KomputeApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createProjectZoneKompute**](KomputeApi.md#createProjectZoneKompute) | **POST** /project/{projectId}/zone/{zoneId}/kompute |  |
| [**deleteKompute**](KomputeApi.md#deleteKompute) | **DELETE** /kompute/{komputeId} |  |
| [**listKomputes**](KomputeApi.md#listKomputes) | **GET** /kompute |  |
| [**listProjectZoneKomputes**](KomputeApi.md#listProjectZoneKomputes) | **GET** /project/{projectId}/zone/{zoneId}/komputes |  |
| [**readKompute**](KomputeApi.md#readKompute) | **GET** /kompute/{komputeId} |  |
| [**readKomputeState**](KomputeApi.md#readKomputeState) | **GET** /kompute/{komputeId}/state |  |
| [**rebootKompute**](KomputeApi.md#rebootKompute) | **PATCH** /kompute/{komputeId}/reboot |  |
| [**resetKompute**](KomputeApi.md#resetKompute) | **PATCH** /kompute/{komputeId}/reset |  |
| [**resumeKompute**](KomputeApi.md#resumeKompute) | **PATCH** /kompute/{komputeId}/resume |  |
| [**shutdownKompute**](KomputeApi.md#shutdownKompute) | **PATCH** /kompute/{komputeId}/shutdown |  |
| [**startKompute**](KomputeApi.md#startKompute) | **PATCH** /kompute/{komputeId}/start |  |
| [**stopKompute**](KomputeApi.md#stopKompute) | **PATCH** /kompute/{komputeId}/stop |  |
| [**suspendKompute**](KomputeApi.md#suspendKompute) | **PATCH** /kompute/{komputeId}/suspend |  |
| [**updateKompute**](KomputeApi.md#updateKompute) | **PUT** /kompute/{komputeId} |  |


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

<a name="deleteKompute"></a>
# **deleteKompute**
> deleteKompute(komputeId)



    Deletes an existing Kompute.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **komputeId** | **String**| The ID of the Kompute. | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listKomputes"></a>
# **listKomputes**
> List listKomputes()



    Returns the IDs of Kompute objects.

### Parameters
This endpoint does not need any parameter.

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

<a name="readKompute"></a>
# **readKompute**
> Kompute readKompute(komputeId)



    Returns a Kompute.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **komputeId** | **String**| The ID of the Kompute. | [default to null] |

### Return type

[**Kompute**](../Models/Kompute.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="readKomputeState"></a>
# **readKomputeState**
> InstanceState readKomputeState(komputeId)



    Returns a virtual machine instance state.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **komputeId** | **String**| The ID of the Kompute. | [default to null] |

### Return type

[**InstanceState**](../Models/InstanceState.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="rebootKompute"></a>
# **rebootKompute**
> rebootKompute(komputeId)



    Performs a Kompute software reboot.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **komputeId** | **String**| The ID of the Kompute. | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="resetKompute"></a>
# **resetKompute**
> resetKompute(komputeId)



    Performs a Kompute hardware reset.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **komputeId** | **String**| The ID of the Kompute. | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="resumeKompute"></a>
# **resumeKompute**
> resumeKompute(komputeId)



    Performs a Kompute software PM resume.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **komputeId** | **String**| The ID of the Kompute. | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="shutdownKompute"></a>
# **shutdownKompute**
> shutdownKompute(komputeId)



    Performs a Kompute software shutdown.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **komputeId** | **String**| The ID of the Kompute. | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="startKompute"></a>
# **startKompute**
> startKompute(komputeId)



    Performs a Kompute hardware boot-up.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **komputeId** | **String**| The ID of the Kompute. | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="stopKompute"></a>
# **stopKompute**
> stopKompute(komputeId)



    Performs a Kompute hardware stop.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **komputeId** | **String**| The ID of the Kompute. | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="suspendKompute"></a>
# **suspendKompute**
> suspendKompute(komputeId)



    Performs a Kompute software PM suspend.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **komputeId** | **String**| The ID of the Kompute. | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateKompute"></a>
# **updateKompute**
> Kompute updateKompute(komputeId, Kompute)



    Updates a Kompute configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **komputeId** | **String**| The ID of the Kompute. | [default to null] |
| **Kompute** | [**Kompute**](../Models/Kompute.md)| Kompute payload. | |

### Return type

[**Kompute**](../Models/Kompute.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

