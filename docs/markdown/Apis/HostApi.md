# HostApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createHost**](HostApi.md#createHost) | **POST** /zone/{ zoneId }/host |  |
| [**deleteHost**](HostApi.md#deleteHost) | **DELETE** /host/{ hostId } |  |
| [**listHostInstances**](HostApi.md#listHostInstances) | **GET** /host/{ hostId }/instances |  |
| [**listHosts**](HostApi.md#listHosts) | **GET** /host |  |
| [**listZoneHosts**](HostApi.md#listZoneHosts) | **GET** /zone/{zoneId}/hosts |  |
| [**readHost**](HostApi.md#readHost) | **GET** /host/{ hostId } |  |
| [**readHostCaps**](HostApi.md#readHostCaps) | **GET** /host/{ hostId }/caps |  |
| [**updateHost**](HostApi.md#updateHost) | **PUT** /host/{ hostId } |  |


<a name="createHost"></a>
# **createHost**
> Host createHost(zoneId, Host)



    Creates a new computing host.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |
| **Host** | [**Host**](../Models/Host.md)| Host payload. | |

### Return type

[**Host**](../Models/Host.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteHost"></a>
# **deleteHost**
> deleteHost(hostId)



    Deletes an existing computing host.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **hostId** | **String**| The ID of the computing host. | [default to null] |

### Return type

null (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listHosts"></a>
# **listHosts**
> List listHosts()



    Returns the IDs of computing host objects.

### Parameters
This endpoint does not need any parameter.

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listZoneHosts"></a>
# **listZoneHosts**
> List listZoneHosts(zoneId)



    Returns the IDs of computing host objects.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="readHost"></a>
# **readHost**
> Host readHost(hostId)



    Returns a computing host.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **hostId** | **String**| The ID of the computing host. | [default to null] |

### Return type

[**Host**](../Models/Host.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="readHostCaps"></a>
# **readHostCaps**
> HostCaps readHostCaps(hostId)



    Returns a computing host capability.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **hostId** | **String**| The ID of the computing host. | [default to null] |

### Return type

[**HostCaps**](../Models/HostCaps.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateHost"></a>
# **updateHost**
> Host updateHost(hostId, Host)



    Updates a computing host configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **hostId** | **String**| The ID of the computing host. | [default to null] |
| **Host** | [**Host**](../Models/Host.md)| Host payload. | |

### Return type

[**Host**](../Models/Host.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

