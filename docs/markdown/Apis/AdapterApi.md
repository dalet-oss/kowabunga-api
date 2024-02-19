# AdapterApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createAdapter**](AdapterApi.md#createAdapter) | **POST** /subnet/{subnetId}/adapter |  |
| [**deleteAdapter**](AdapterApi.md#deleteAdapter) | **DELETE** /adapter/{ adapterId } |  |
| [**getSubnetAdapters**](AdapterApi.md#getSubnetAdapters) | **GET** /subnet/{subnetId}/adapters |  |
| [**listAdapters**](AdapterApi.md#listAdapters) | **GET** /adapter |  |
| [**readAdapter**](AdapterApi.md#readAdapter) | **GET** /adapter/{ adapterId } |  |
| [**updateAdapter**](AdapterApi.md#updateAdapter) | **PUT** /adapter/{ adapterId } |  |


<a name="createAdapter"></a>
# **createAdapter**
> Adapter createAdapter(subnetId, Adapter, assignIP)



    Creates a new network adapter.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **subnetId** | **String**| The ID of the network subnet. | [default to null] |
| **Adapter** | [**Adapter**](../Models/Adapter.md)| Adapter payload | |
| **assignIP** | **Boolean**| whether Kowabunga should pick and assign an IP address to this adapter. | [optional] [default to null] |

### Return type

[**Adapter**](../Models/Adapter.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteAdapter"></a>
# **deleteAdapter**
> deleteAdapter(adapterId)



    Deletes an existing network adapter.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **adapterId** | **String**| The ID of the network adapter. | [default to null] |

### Return type

null (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getSubnetAdapters"></a>
# **getSubnetAdapters**
> List getSubnetAdapters(subnetId)



    Returns the IDs of the network adapters existing in the subnet.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **subnetId** | **String**| The ID of the network subnet. | [default to null] |

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listAdapters"></a>
# **listAdapters**
> List listAdapters()



    Returns the IDs of network adapter objects.

### Parameters
This endpoint does not need any parameter.

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="readAdapter"></a>
# **readAdapter**
> Adapter readAdapter(adapterId)



    Returns a network adapter.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **adapterId** | **String**| The ID of the network adapter. | [default to null] |

### Return type

[**Adapter**](../Models/Adapter.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateAdapter"></a>
# **updateAdapter**
> Adapter updateAdapter(adapterId, Adapter)



    Updates a network adapter configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **adapterId** | **String**| The ID of the network adapter. | [default to null] |
| **Adapter** | [**Adapter**](../Models/Adapter.md)| Adapter payload. | |

### Return type

[**Adapter**](../Models/Adapter.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

