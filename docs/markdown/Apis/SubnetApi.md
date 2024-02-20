# SubnetApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createAdapter**](SubnetApi.md#createAdapter) | **POST** /subnet/{ subnetId }/adapter |  |
| [**createSubnet**](SubnetApi.md#createSubnet) | **POST** /vnet/{vnetId}/subnet |  |
| [**deleteSubnet**](SubnetApi.md#deleteSubnet) | **DELETE** /subnet/{ subnetId } |  |
| [**getVNetSubnets**](SubnetApi.md#getVNetSubnets) | **GET** /vnet/{vnetId}/subnets |  |
| [**listSubnetAdapters**](SubnetApi.md#listSubnetAdapters) | **GET** /subnet/{ subnetId }/adapters |  |
| [**listSubnets**](SubnetApi.md#listSubnets) | **GET** /subnet |  |
| [**readSubnet**](SubnetApi.md#readSubnet) | **GET** /subnet/{ subnetId } |  |
| [**updateSubnet**](SubnetApi.md#updateSubnet) | **PUT** /subnet/{ subnetId } |  |
| [**updateVNetDefaultSubnet**](SubnetApi.md#updateVNetDefaultSubnet) | **PUT** /vnet/{vnetId}/subnet/{subnetId}/default |  |


<a name="createAdapter"></a>
# **createAdapter**
> Adapter createAdapter(subnetId, Adapter, assignIP)



    Creates a new network adapter.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **subnetId** | **String**| The ID of the network subnet. | [default to null] |
| **Adapter** | [**Adapter**](../Models/Adapter.md)| Adapter payload. | |
| **assignIP** | **Boolean**| Whether Kowabunga should pick and assign an IP address to this adapter. | [optional] [default to null] |

### Return type

[**Adapter**](../Models/Adapter.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createSubnet"></a>
# **createSubnet**
> Subnet createSubnet(vnetId, Subnet)



    Creates a new subnet.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **vnetId** | **String**| The ID of the virtual network. | [default to null] |
| **Subnet** | [**Subnet**](../Models/Subnet.md)| Subnet payload | |

### Return type

[**Subnet**](../Models/Subnet.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteSubnet"></a>
# **deleteSubnet**
> deleteSubnet(subnetId)



    Deletes an existing network subnet.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **subnetId** | **String**| The ID of the network subnet. | [default to null] |

### Return type

null (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getVNetSubnets"></a>
# **getVNetSubnets**
> List getVNetSubnets(vnetId)



    Returns the IDs of the subnets existing in the virtual network.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **vnetId** | **String**| The ID of the virtual network. | [default to null] |

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listSubnetAdapters"></a>
# **listSubnetAdapters**
> List listSubnetAdapters(subnetId)



    Returns the IDs of network adapter objects.

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

<a name="listSubnets"></a>
# **listSubnets**
> List listSubnets()



    Returns the IDs of network subnet objects.

### Parameters
This endpoint does not need any parameter.

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="readSubnet"></a>
# **readSubnet**
> Subnet readSubnet(subnetId)



    Returns a network subnet.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **subnetId** | **String**| The ID of the network subnet. | [default to null] |

### Return type

[**Subnet**](../Models/Subnet.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateSubnet"></a>
# **updateSubnet**
> Subnet updateSubnet(subnetId, Subnet)



    Updates a network subnet configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **subnetId** | **String**| The ID of the network subnet. | [default to null] |
| **Subnet** | [**Subnet**](../Models/Subnet.md)| Subnet payload. | |

### Return type

[**Subnet**](../Models/Subnet.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updateVNetDefaultSubnet"></a>
# **updateVNetDefaultSubnet**
> updateVNetDefaultSubnet(vnetId, subnetId)



    Set a virtual network default subnet.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **vnetId** | **String**| The ID of the virtual network. | [default to null] |
| **subnetId** | **String**| The ID of the network subnet. | [default to null] |

### Return type

null (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

