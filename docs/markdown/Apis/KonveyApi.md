# KonveyApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createProjectRegionKonvey_1**](KonveyApi.md#createProjectRegionKonvey_1) | **POST** /project/{projectId}/region/{regionId}/konvey |  |
| [**createProjectZoneKonvey_1**](KonveyApi.md#createProjectZoneKonvey_1) | **POST** /project/{projectId}/zone/{zoneId}/konvey |  |
| [**deleteKonvey**](KonveyApi.md#deleteKonvey) | **DELETE** /konvey/{konveyId} |  |
| [**listKonveys**](KonveyApi.md#listKonveys) | **GET** /konvey |  |
| [**listProjectRegionKonveys_1**](KonveyApi.md#listProjectRegionKonveys_1) | **GET** /project/{projectId}/region/{regionId}/konveys |  |
| [**listProjectZoneKonveys_1**](KonveyApi.md#listProjectZoneKonveys_1) | **GET** /project/{projectId}/zone/{zoneId}/konveys |  |
| [**readKonvey**](KonveyApi.md#readKonvey) | **GET** /konvey/{konveyId} |  |
| [**updateKonvey**](KonveyApi.md#updateKonvey) | **PUT** /konvey/{konveyId} |  |


<a name="createProjectRegionKonvey_1"></a>
# **createProjectRegionKonvey_1**
> Konvey createProjectRegionKonvey_1(projectId, regionId, Konvey)



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

<a name="createProjectZoneKonvey_1"></a>
# **createProjectZoneKonvey_1**
> Konvey createProjectZoneKonvey_1(projectId, zoneId, Konvey)



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

<a name="deleteKonvey"></a>
# **deleteKonvey**
> deleteKonvey(konveyId)



    Deletes an existing Konvey (Kowabunga Network Load-Balancer).

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **konveyId** | **String**| The ID of the Konvey (Kowabunga Network Load-Balancer). | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listKonveys"></a>
# **listKonveys**
> List listKonveys()



    Returns the IDs of Konvey (Kowabunga Network Load-Balancer) objects.

### Parameters
This endpoint does not need any parameter.

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listProjectRegionKonveys_1"></a>
# **listProjectRegionKonveys_1**
> List listProjectRegionKonveys_1(projectId, regionId)



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

<a name="listProjectZoneKonveys_1"></a>
# **listProjectZoneKonveys_1**
> List listProjectZoneKonveys_1(projectId, zoneId)



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

<a name="readKonvey"></a>
# **readKonvey**
> Konvey readKonvey(konveyId)



    Returns a Konvey (Kowabunga Network Load-Balancer).

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **konveyId** | **String**| The ID of the Konvey (Kowabunga Network Load-Balancer). | [default to null] |

### Return type

[**Konvey**](../Models/Konvey.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateKonvey"></a>
# **updateKonvey**
> Konvey updateKonvey(konveyId, Konvey)



    Updates a Konvey (Kowabunga Network Load-Balancer) configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **konveyId** | **String**| The ID of the Konvey (Kowabunga Network Load-Balancer). | [default to null] |
| **Konvey** | [**Konvey**](../Models/Konvey.md)| Konvey payload. | |

### Return type

[**Konvey**](../Models/Konvey.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

