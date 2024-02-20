# TemplateApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createTemplate**](TemplateApi.md#createTemplate) | **POST** /pool/{ poolId }/template |  |
| [**deleteTemplate**](TemplateApi.md#deleteTemplate) | **DELETE** /template/{ templateId } |  |
| [**listStoragePoolTemplates**](TemplateApi.md#listStoragePoolTemplates) | **GET** /pool/{ poolId }/templates |  |
| [**listTemplates**](TemplateApi.md#listTemplates) | **GET** /template |  |
| [**readTemplate**](TemplateApi.md#readTemplate) | **GET** /template/{ templateId } |  |
| [**setStoragePoolDefaultTemplate**](TemplateApi.md#setStoragePoolDefaultTemplate) | **PATCH** /pool/{ poolId }/template/{ templateId }/default |  |
| [**updateTemplate**](TemplateApi.md#updateTemplate) | **PUT** /template/{ templateId } |  |


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



    Deletes an existing image template.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **templateId** | **String**| The ID of the image template. | [default to null] |

### Return type

null (empty response body)

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

<a name="listTemplates"></a>
# **listTemplates**
> List listTemplates()



    Returns the IDs of image template objects.

### Parameters
This endpoint does not need any parameter.

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="readTemplate"></a>
# **readTemplate**
> Template readTemplate(templateId)



    Returns a image template.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **templateId** | **String**| The ID of the image template. | [default to null] |

### Return type

[**Template**](../Models/Template.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="setStoragePoolDefaultTemplate"></a>
# **setStoragePoolDefaultTemplate**
> setStoragePoolDefaultTemplate(poolId, templateId)



    Performs a storage pool setting of default template.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **poolId** | **String**| The ID of the storage pool. | [default to null] |
| **templateId** | **String**| The ID of the image template. | [default to null] |

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



    Updates a image template configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **templateId** | **String**| The ID of the image template. | [default to null] |
| **Template** | [**Template**](../Models/Template.md)| Template payload. | |

### Return type

[**Template**](../Models/Template.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

