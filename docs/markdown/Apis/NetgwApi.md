# NetgwApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createNetGW**](NetgwApi.md#createNetGW) | **POST** /zone/{zoneId}/netgw |  |
| [**deleteNetGW**](NetgwApi.md#deleteNetGW) | **DELETE** /netgw/{netgwId} |  |
| [**getAllNetGWs**](NetgwApi.md#getAllNetGWs) | **GET** /netgw |  |
| [**getNetGW**](NetgwApi.md#getNetGW) | **GET** /netgw/{netgwId} |  |
| [**getZoneNetGWs**](NetgwApi.md#getZoneNetGWs) | **GET** /zone/{zoneId}/netgws |  |
| [**updateNetGW**](NetgwApi.md#updateNetGW) | **PUT** /netgw/{netgwId} |  |


<a name="createNetGW"></a>
# **createNetGW**
> NetGW createNetGW(zoneId, NetGW)



    Creates a new network gateway.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **zoneId** | **String**| The ID of the zone. | [default to null] |
| **NetGW** | [**NetGW**](../Models/NetGW.md)| NetGW payload | |

### Return type

[**NetGW**](../Models/NetGW.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteNetGW"></a>
# **deleteNetGW**
> deleteNetGW(netgwId)



    Deletes an existing network gateway.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **netgwId** | **String**| The ID of the network gateway. | [default to null] |

### Return type

null (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getAllNetGWs"></a>
# **getAllNetGWs**
> List getAllNetGWs()



    Returns the IDs of network gateways.

### Parameters
This endpoint does not need any parameter.

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getNetGW"></a>
# **getNetGW**
> NetGW getNetGW(netgwId)



    Returns a description of the network gateway

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **netgwId** | **String**| The ID of the network gateway. | [default to null] |

### Return type

[**NetGW**](../Models/NetGW.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getZoneNetGWs"></a>
# **getZoneNetGWs**
> List getZoneNetGWs(zoneId)



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

<a name="updateNetGW"></a>
# **updateNetGW**
> NetGW updateNetGW(netgwId, NetGW)



    Updates a network gateway configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **netgwId** | **String**| The ID of the network gateway. | [default to null] |
| **NetGW** | [**NetGW**](../Models/NetGW.md)| NetGW payload | |

### Return type

[**NetGW**](../Models/NetGW.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

