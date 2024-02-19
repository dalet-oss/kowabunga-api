# HostApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createHost**](HostApi.md#createHost) | **POST** /zone/{zoneId}/host |  |
| [**deleteHost**](HostApi.md#deleteHost) | **DELETE** /host/{hostId} |  |
| [**getAllHosts**](HostApi.md#getAllHosts) | **GET** /host |  |
| [**getHost**](HostApi.md#getHost) | **GET** /host/{hostId} |  |
| [**getHostCaps**](HostApi.md#getHostCaps) | **GET** /host/{hostId}/caps |  |
| [**getHostInstances**](HostApi.md#getHostInstances) | **GET** /host/{hostId}/instances |  |
| [**getZoneHosts**](HostApi.md#getZoneHosts) | **GET** /zone/{zoneId}/hosts |  |
| [**updateHost**](HostApi.md#updateHost) | **PUT** /host/{hostId} |  |


<a name="createHost"></a>
# **createHost**
> Host createHost(zoneId, Host)



    Creates a new host.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **zoneId** | **String**| The ID of the zone. | [default to null] |
| **Host** | [**Host**](../Models/Host.md)| Host payload | |

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



    Deletes an existing host.

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

<a name="getAllHosts"></a>
# **getAllHosts**
> List getAllHosts()



    Returns the IDs of registered hosts.

### Parameters
This endpoint does not need any parameter.

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getHost"></a>
# **getHost**
> Host getHost(hostId)



    Returns a description of the host

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

<a name="getHostCaps"></a>
# **getHostCaps**
> HostCaps getHostCaps(hostId)



    Returns the host capabilities.

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

<a name="getZoneHosts"></a>
# **getZoneHosts**
> List getZoneHosts(zoneId)



    Returns the IDs of the hosts existing in the zone.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **zoneId** | **String**| The ID of the zone. | [default to null] |

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateHost"></a>
# **updateHost**
> Host updateHost(hostId, Host)



    Updates a host configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **hostId** | **String**| The ID of the computing host. | [default to null] |
| **Host** | [**Host**](../Models/Host.md)| Host payload | |

### Return type

[**Host**](../Models/Host.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

