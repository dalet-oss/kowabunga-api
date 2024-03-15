# KceApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createProjectZoneKCE**](KceApi.md#createProjectZoneKCE) | **POST** /project/{projectId}/zone/{zoneId}/kce |  |
| [**deleteKCE**](KceApi.md#deleteKCE) | **DELETE** /kce/{kceId} |  |
| [**listKCEs**](KceApi.md#listKCEs) | **GET** /kce |  |
| [**listProjectZoneKCEs**](KceApi.md#listProjectZoneKCEs) | **GET** /project/{projectId}/zone/{zoneId}/kces |  |
| [**readKCE**](KceApi.md#readKCE) | **GET** /kce/{kceId} |  |
| [**readKCEState**](KceApi.md#readKCEState) | **GET** /kce/{kceId}/state |  |
| [**rebootKCE**](KceApi.md#rebootKCE) | **PATCH** /kce/{kceId}/reboot |  |
| [**resetKCE**](KceApi.md#resetKCE) | **PATCH** /kce/{kceId}/reset |  |
| [**resumeKCE**](KceApi.md#resumeKCE) | **PATCH** /kce/{kceId}/resume |  |
| [**shutdownKCE**](KceApi.md#shutdownKCE) | **PATCH** /kce/{kceId}/shutdown |  |
| [**startKCE**](KceApi.md#startKCE) | **PATCH** /kce/{kceId}/start |  |
| [**stopKCE**](KceApi.md#stopKCE) | **PATCH** /kce/{kceId}/stop |  |
| [**suspendKCE**](KceApi.md#suspendKCE) | **PATCH** /kce/{kceId}/suspend |  |
| [**updateKCE**](KceApi.md#updateKCE) | **PUT** /kce/{kceId} |  |


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
| **public** | **Boolean**| Should KCE be exposed over public Internet ? (a public IPv4 address will then be auto-assigned, default to false). | [optional] [default to null] |
| **notify** | **Boolean**| Whether or not to send a notification email at resource creation. | [optional] [default to null] |

### Return type

[**KCE**](../Models/KCE.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteKCE"></a>
# **deleteKCE**
> deleteKCE(kceId)



    Deletes an existing KCE (Kowabunga Compute Engine).

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kceId** | **String**| The ID of the KCE (Kowabunga Compute Engine). | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listKCEs"></a>
# **listKCEs**
> List listKCEs()



    Returns the IDs of KCE (Kowabunga Compute Engine) objects.

### Parameters
This endpoint does not need any parameter.

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="readKCE"></a>
# **readKCE**
> KCE readKCE(kceId)



    Returns a KCE (Kowabunga Compute Engine).

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kceId** | **String**| The ID of the KCE (Kowabunga Compute Engine). | [default to null] |

### Return type

[**KCE**](../Models/KCE.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="readKCEState"></a>
# **readKCEState**
> InstanceState readKCEState(kceId)



    Returns a virtual machine instance state.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kceId** | **String**| The ID of the KCE (Kowabunga Compute Engine). | [default to null] |

### Return type

[**InstanceState**](../Models/InstanceState.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="rebootKCE"></a>
# **rebootKCE**
> rebootKCE(kceId)



    Performs a KCE (Kowabunga Compute Engine) software reboot.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kceId** | **String**| The ID of the KCE (Kowabunga Compute Engine). | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="resetKCE"></a>
# **resetKCE**
> resetKCE(kceId)



    Performs a KCE (Kowabunga Compute Engine) hardware reset.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kceId** | **String**| The ID of the KCE (Kowabunga Compute Engine). | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="resumeKCE"></a>
# **resumeKCE**
> resumeKCE(kceId)



    Performs a KCE (Kowabunga Compute Engine) software PM resume.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kceId** | **String**| The ID of the KCE (Kowabunga Compute Engine). | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="shutdownKCE"></a>
# **shutdownKCE**
> shutdownKCE(kceId)



    Performs a KCE (Kowabunga Compute Engine) software shutdown.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kceId** | **String**| The ID of the KCE (Kowabunga Compute Engine). | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="startKCE"></a>
# **startKCE**
> startKCE(kceId)



    Performs a KCE (Kowabunga Compute Engine) hardware boot-up.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kceId** | **String**| The ID of the KCE (Kowabunga Compute Engine). | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="stopKCE"></a>
# **stopKCE**
> stopKCE(kceId)



    Performs a KCE (Kowabunga Compute Engine) hardware stop.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kceId** | **String**| The ID of the KCE (Kowabunga Compute Engine). | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="suspendKCE"></a>
# **suspendKCE**
> suspendKCE(kceId)



    Performs a KCE (Kowabunga Compute Engine) software PM suspend.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kceId** | **String**| The ID of the KCE (Kowabunga Compute Engine). | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateKCE"></a>
# **updateKCE**
> KCE updateKCE(kceId, KCE)



    Updates a KCE (Kowabunga Compute Engine) configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **kceId** | **String**| The ID of the KCE (Kowabunga Compute Engine). | [default to null] |
| **KCE** | [**KCE**](../Models/KCE.md)| KCE payload. | |

### Return type

[**KCE**](../Models/KCE.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

