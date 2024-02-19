# SubnetApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createAdapter**](SubnetApi.md#createAdapter) | **POST** /subnet/{subnetId}/adapter |  |
| [**createSubnet**](SubnetApi.md#createSubnet) | **POST** /vnet/{vnetId}/subnet |  |
| [**deleteSubnet**](SubnetApi.md#deleteSubnet) | **DELETE** /subnet/{subnetId} |  |
| [**getAllSubnets**](SubnetApi.md#getAllSubnets) | **GET** /subnet |  |
| [**getSubnet**](SubnetApi.md#getSubnet) | **GET** /subnet/{subnetId} |  |
| [**getSubnetAdapters**](SubnetApi.md#getSubnetAdapters) | **GET** /subnet/{subnetId}/adapters |  |
| [**getVNetSubnets**](SubnetApi.md#getVNetSubnets) | **GET** /vnet/{vnetId}/subnets |  |
| [**updateSubnet**](SubnetApi.md#updateSubnet) | **PUT** /subnet/{subnetId} |  |
| [**updateVNetDefaultSubnet**](SubnetApi.md#updateVNetDefaultSubnet) | **PUT** /vnet/{vnetId}/subnet/{subnetId}/default |  |


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



    Deletes an existing subnet.

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

<a name="getAllSubnets"></a>
# **getAllSubnets**
> List getAllSubnets()



    Returns the IDs of subnets.

### Parameters
This endpoint does not need any parameter.

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getSubnet"></a>
# **getSubnet**
> Subnet getSubnet(subnetId)



    Returns a description of the subnet.

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

<a name="updateSubnet"></a>
# **updateSubnet**
> Subnet updateSubnet(subnetId, Subnet)



    Updates a subnet configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **subnetId** | **String**| The ID of the network subnet. | [default to null] |
| **Subnet** | [**Subnet**](../Models/Subnet.md)| Subnet payload | |

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

