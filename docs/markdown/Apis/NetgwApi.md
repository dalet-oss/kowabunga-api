# NetgwApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createNetGW**](NetgwApi.md#createNetGW) | **POST** /zone/{zoneId}/netgw |  |
| [**deleteNetGW**](NetgwApi.md#deleteNetGW) | **DELETE** /netgw/{netgwId} |  |
| [**listNetGWs**](NetgwApi.md#listNetGWs) | **GET** /netgw |  |
| [**listZoneNetGWs**](NetgwApi.md#listZoneNetGWs) | **GET** /zone/{zoneId}/netgws |  |
| [**readNetGW**](NetgwApi.md#readNetGW) | **GET** /netgw/{netgwId} |  |
| [**updateNetGW**](NetgwApi.md#updateNetGW) | **PUT** /netgw/{netgwId} |  |


<a name="createNetGW"></a>
# **createNetGW**
> NetGW createNetGW(zoneId, NetGW)



    Creates a new Iris network gateway.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |
| **NetGW** | [**NetGW**](../Models/NetGW.md)| NetGW payload. | |

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



    Deletes an existing Iris network gateway.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **netgwId** | **String**| The ID of the Iris network gateway. | [default to null] |

### Return type

null (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listNetGWs"></a>
# **listNetGWs**
> List listNetGWs()



    Returns the IDs of Iris network gateway objects.

### Parameters
This endpoint does not need any parameter.

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listZoneNetGWs"></a>
# **listZoneNetGWs**
> List listZoneNetGWs(zoneId)



    Returns the IDs of Iris network gateway objects.

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

<a name="readNetGW"></a>
# **readNetGW**
> NetGW readNetGW(netgwId)



    Returns a Iris network gateway.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **netgwId** | **String**| The ID of the Iris network gateway. | [default to null] |

### Return type

[**NetGW**](../Models/NetGW.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateNetGW"></a>
# **updateNetGW**
> NetGW updateNetGW(netgwId, NetGW)



    Updates a Iris network gateway configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **netgwId** | **String**| The ID of the Iris network gateway. | [default to null] |
| **NetGW** | [**NetGW**](../Models/NetGW.md)| NetGW payload. | |

### Return type

[**NetGW**](../Models/NetGW.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

