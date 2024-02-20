# TemplateApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createTemplate**](TemplateApi.md#createTemplate) | **POST** /pool/{ poolId }/template |  |
| [**deleteTemplate**](TemplateApi.md#deleteTemplate) | **DELETE** /template/{templateId} |  |
| [**getAllTemplates**](TemplateApi.md#getAllTemplates) | **GET** /template |  |
| [**getTemplate**](TemplateApi.md#getTemplate) | **GET** /template/{templateId} |  |
| [**listStoragePoolTemplates**](TemplateApi.md#listStoragePoolTemplates) | **GET** /pool/{ poolId }/templates |  |
| [**updateStoragePoolDefaultTemplate**](TemplateApi.md#updateStoragePoolDefaultTemplate) | **POST** /pool/{ poolId }/template/{ templateId }/default |  |
| [**updateTemplate**](TemplateApi.md#updateTemplate) | **PUT** /template/{templateId} |  |


<a name="createTemplate"></a>
# **createTemplate**
> Template createTemplate(poolId, Template)



    Creates a new image template.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **poolId** | **String**| The ID of the storage pool. | [default to null] |
| **Template** | [**Template**](../Models/Template.md)| Template payload. | |

### Return type

[**Template**](../Models/Template.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteTemplate"></a>
# **deleteTemplate**
> deleteTemplate(templateId)



    Deletes an existing volume template.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **templateId** | **String**| The ID of the volume template. | [default to null] |

### Return type

null (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getAllTemplates"></a>
# **getAllTemplates**
> List getAllTemplates()



    Returns the IDs of volume templates.

### Parameters
This endpoint does not need any parameter.

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="getTemplate"></a>
# **getTemplate**
> Template getTemplate(templateId)



    Returns a description of the volume template.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **templateId** | **String**| The ID of the volume template. | [default to null] |

### Return type

[**Template**](../Models/Template.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listStoragePoolTemplates"></a>
# **listStoragePoolTemplates**
> List listStoragePoolTemplates(poolId)



    Returns the IDs of image template objects.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **poolId** | **String**| The ID of the storage pool. | [default to null] |

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateStoragePoolDefaultTemplate"></a>
# **updateStoragePoolDefaultTemplate**
> updateStoragePoolDefaultTemplate(poolId, templateId)



    Performs a storage pool setting of default template.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **poolId** | **String**| The ID of the storage pool. | [default to null] |
| **templateId** | **String**| The ID of the volume template. | [default to null] |

### Return type

null (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateTemplate"></a>
# **updateTemplate**
> Template updateTemplate(templateId, Template)



    Updates a volume template configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **templateId** | **String**| The ID of the volume template. | [default to null] |
| **Template** | [**Template**](../Models/Template.md)| Template payload | |

### Return type

[**Template**](../Models/Template.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

