# GroupApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createGroup**](GroupApi.md#createGroup) | **POST** /group |  |
| [**deleteGroup**](GroupApi.md#deleteGroup) | **DELETE** /group/{groupId} |  |
| [**listGroups**](GroupApi.md#listGroups) | **GET** /group |  |
| [**readGroup**](GroupApi.md#readGroup) | **GET** /group/{groupId} |  |
| [**updateGroup**](GroupApi.md#updateGroup) | **PUT** /group/{groupId} |  |


<a name="createGroup"></a>
# **createGroup**
> Group createGroup(Group)



    Creates a new Kowabunga users group.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **Group** | [**Group**](../Models/Group.md)| Group payload. | |

### Return type

[**Group**](../Models/Group.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteGroup"></a>
# **deleteGroup**
> deleteGroup(groupId)



    Deletes an existing Kowabunga users group.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **groupId** | **String**| The ID of the Kowabunga users group. | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listGroups"></a>
# **listGroups**
> List listGroups()



    Returns the IDs of Kowabunga users group objects.

### Parameters
This endpoint does not need any parameter.

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="readGroup"></a>
# **readGroup**
> Group readGroup(groupId)



    Returns a Kowabunga users group.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **groupId** | **String**| The ID of the Kowabunga users group. | [default to null] |

### Return type

[**Group**](../Models/Group.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateGroup"></a>
# **updateGroup**
> Group updateGroup(groupId, Group)



    Updates a Kowabunga users group configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **groupId** | **String**| The ID of the Kowabunga users group. | [default to null] |
| **Group** | [**Group**](../Models/Group.md)| Group payload. | |

### Return type

[**Group**](../Models/Group.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

