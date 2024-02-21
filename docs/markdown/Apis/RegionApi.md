# RegionApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createRegion**](RegionApi.md#createRegion) | **POST** /region |  |
| [**createZone**](RegionApi.md#createZone) | **POST** /region/{regionId}/zone |  |
| [**deleteRegion**](RegionApi.md#deleteRegion) | **DELETE** /region/{regionId} |  |
| [**listRegionZones**](RegionApi.md#listRegionZones) | **GET** /region/{regionId}/zones |  |
| [**listRegions**](RegionApi.md#listRegions) | **GET** /region |  |
| [**readRegion**](RegionApi.md#readRegion) | **GET** /region/{regionId} |  |
| [**updateRegion**](RegionApi.md#updateRegion) | **PUT** /region/{regionId} |  |


<a name="createRegion"></a>
# **createRegion**
> Region createRegion(Region)



    Creates a new region.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **Region** | [**Region**](../Models/Region.md)| Region payload. | |

### Return type

[**Region**](../Models/Region.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="createZone"></a>
# **createZone**
> Zone createZone(regionId, Zone)



    Creates a new availability zone.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **regionId** | **String**| The ID of the region. | [default to null] |
| **Zone** | [**Zone**](../Models/Zone.md)| Zone payload. | |

### Return type

[**Zone**](../Models/Zone.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteRegion"></a>
# **deleteRegion**
> deleteRegion(regionId)



    Deletes an existing region.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **regionId** | **String**| The ID of the region. | [default to null] |

### Return type

null (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listRegionZones"></a>
# **listRegionZones**
> List listRegionZones(regionId)



    Returns the IDs of availability zone objects.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **regionId** | **String**| The ID of the region. | [default to null] |

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listRegions"></a>
# **listRegions**
> List listRegions()



    Returns the IDs of region objects.

### Parameters
This endpoint does not need any parameter.

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="readRegion"></a>
# **readRegion**
> Region readRegion(regionId)



    Returns a region.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **regionId** | **String**| The ID of the region. | [default to null] |

### Return type

[**Region**](../Models/Region.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateRegion"></a>
# **updateRegion**
> Region updateRegion(regionId, Region)



    Updates a region configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **regionId** | **String**| The ID of the region. | [default to null] |
| **Region** | [**Region**](../Models/Region.md)| Region payload. | |

### Return type

[**Region**](../Models/Region.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

