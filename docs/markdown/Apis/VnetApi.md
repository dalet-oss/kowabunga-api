# VnetApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createSubnet**](VnetApi.md#createSubnet) | **POST** /vnet/{vnetId}/subnet |  |
| [**createVNet**](VnetApi.md#createVNet) | **POST** /zone/{zoneId}/vnet |  |
| [**deleteVNet**](VnetApi.md#deleteVNet) | **DELETE** /vnet/{vnetId} |  |
| [**getAllVNets**](VnetApi.md#getAllVNets) | **GET** /vnet |  |
| [**getVNet**](VnetApi.md#getVNet) | **GET** /vnet/{vnetId} |  |
| [**getVNetSubnets**](VnetApi.md#getVNetSubnets) | **GET** /vnet/{vnetId}/subnets |  |
| [**getZoneVNets**](VnetApi.md#getZoneVNets) | **GET** /zone/{zoneId}/vnets |  |
| [**updateVNet**](VnetApi.md#updateVNet) | **PUT** /vnet/{vnetId} |  |
| [**updateVNetDefaultSubnet**](VnetApi.md#updateVNetDefaultSubnet) | **PUT** /vnet/{vnetId}/subnet/{subnetId}/default |  |


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

<a name="createVNet"></a>
# **createVNet**
> VNet createVNet(zoneId, VNet)



    Creates a new virtual network.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **zoneId** | **String**| The ID of the zone. | [default to null] |
| **VNet** | [**VNet**](../Models/VNet.md)| VNet payload | |

### Return type

[**VNet**](../Models/VNet.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getAllVNets"></a>
# **getAllVNets**
> List getAllVNets()



    Returns the IDs of virtual networks.

### Parameters
This endpoint does not need any parameter.

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getVNet"></a>
# **getVNet**
> VNet getVNet(vnetId)



    Returns a description of the virtual network

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **vnetId** | **String**| The ID of the virtual network. | [default to null] |

### Return type

[**VNet**](../Models/VNet.md)

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

<a name="getZoneVNets"></a>
# **getZoneVNets**
> List getZoneVNets(zoneId)



    Returns the IDs of the virtual networks existing in the zone.

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

<a name="updateVNet"></a>
# **updateVNet**
> VNet updateVNet(vnetId, VNet)



    Updates a virtual network configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **vnetId** | **String**| The ID of the virtual network. | [default to null] |
| **VNet** | [**VNet**](../Models/VNet.md)| VNet payload | |

### Return type

[**VNet**](../Models/VNet.md)

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

