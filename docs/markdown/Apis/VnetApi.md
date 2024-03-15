# VnetApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createSubnet**](VnetApi.md#createSubnet) | **POST** /vnet/{vnetId}/subnet |  |
| [**createVNet**](VnetApi.md#createVNet) | **POST** /zone/{zoneId}/vnet |  |
| [**deleteVNet**](VnetApi.md#deleteVNet) | **DELETE** /vnet/{vnetId} |  |
| [**listVNetSubnets**](VnetApi.md#listVNetSubnets) | **GET** /vnet/{vnetId}/subnets |  |
| [**listVNets**](VnetApi.md#listVNets) | **GET** /vnet |  |
| [**listZoneVNets**](VnetApi.md#listZoneVNets) | **GET** /zone/{zoneId}/vnets |  |
| [**readVNet**](VnetApi.md#readVNet) | **GET** /vnet/{vnetId} |  |
| [**setVNetDefaultSubnet**](VnetApi.md#setVNetDefaultSubnet) | **PATCH** /vnet/{vnetId}/subnet/{subnetId}/default |  |
| [**updateVNet**](VnetApi.md#updateVNet) | **PUT** /vnet/{vnetId} |  |


<a name="createSubnet"></a>
# **createSubnet**
> Subnet createSubnet(vnetId, Subnet)



    Creates a new network subnet.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **vnetId** | **String**| The ID of the virtual network. | [default to null] |
| **Subnet** | [**Subnet**](../Models/Subnet.md)| Subnet payload. | |

### Return type

[**Subnet**](../Models/Subnet.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createVNet"></a>
# **createVNet**
> VNet createVNet(zoneId, VNet)



    Creates a new virtual network.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |
| **VNet** | [**VNet**](../Models/VNet.md)| VNet payload. | |

### Return type

[**VNet**](../Models/VNet.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteVNet"></a>
# **deleteVNet**
> deleteVNet(vnetId)



    Deletes an existing virtual network.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **vnetId** | **String**| The ID of the virtual network. | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listVNetSubnets"></a>
# **listVNetSubnets**
> List listVNetSubnets(vnetId)



    Returns the IDs of network subnet objects.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **vnetId** | **String**| The ID of the virtual network. | [default to null] |

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listVNets"></a>
# **listVNets**
> List listVNets()



    Returns the IDs of virtual network objects.

### Parameters
This endpoint does not need any parameter.

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listZoneVNets"></a>
# **listZoneVNets**
> List listZoneVNets(zoneId)



    Returns the IDs of virtual network objects.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **zoneId** | **String**| The ID of the availability zone. | [default to null] |

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="readVNet"></a>
# **readVNet**
> VNet readVNet(vnetId)



    Returns a virtual network.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **vnetId** | **String**| The ID of the virtual network. | [default to null] |

### Return type

[**VNet**](../Models/VNet.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="setVNetDefaultSubnet"></a>
# **setVNetDefaultSubnet**
> setVNetDefaultSubnet(vnetId, subnetId)



    Performs a virtual network setting of default network subnet.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **vnetId** | **String**| The ID of the virtual network. | [default to null] |
| **subnetId** | **String**| The ID of the network subnet. | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateVNet"></a>
# **updateVNet**
> VNet updateVNet(vnetId, VNet)



    Updates a virtual network configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **vnetId** | **String**| The ID of the virtual network. | [default to null] |
| **VNet** | [**VNet**](../Models/VNet.md)| VNet payload. | |

### Return type

[**VNet**](../Models/VNet.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

