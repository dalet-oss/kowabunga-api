# TokenApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**deleteApiToken**](TokenApi.md#deleteApiToken) | **DELETE** /token/{tokenId} |  |
| [**listApiTokens**](TokenApi.md#listApiTokens) | **GET** /token |  |
| [**readApiToken**](TokenApi.md#readApiToken) | **GET** /token/{tokenId} |  |
| [**setApiToken**](TokenApi.md#setApiToken) | **PATCH** /agent/{agentId}/token |  |
| [**updateApiToken**](TokenApi.md#updateApiToken) | **PUT** /token/{tokenId} |  |


<a name="deleteApiToken"></a>
# **deleteApiToken**
> deleteApiToken(tokenId)



    Deletes an existing server-to-server authentication security token.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **tokenId** | **String**| The ID of the server-to-server authentication security token. | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listApiTokens"></a>
# **listApiTokens**
> List listApiTokens()



    Returns the IDs of server-to-server authentication security token objects.

### Parameters
This endpoint does not need any parameter.

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="readApiToken"></a>
# **readApiToken**
> ApiToken readApiToken(tokenId)



    Returns a server-to-server authentication security token.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **tokenId** | **String**| The ID of the server-to-server authentication security token. | [default to null] |

### Return type

[**ApiToken**](../Models/ApiToken.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="setApiToken"></a>
# **setApiToken**
> ApiToken setApiToken(agentId, expire, expiration\_date)



    Performs a Kowabunga remote agent setting of API token (will replace any existing one).

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **agentId** | **String**| The ID of the Kowabunga remote agent. | [default to null] |
| **expire** | **Boolean**| Whether or not the token should expire. | [optional] [default to null] |
| **expiration\_date** | **date**| Token&#39;s expiration date (YYYY-MM-DD format). | [optional] [default to null] |

### Return type

[**ApiToken**](../Models/ApiToken.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateApiToken"></a>
# **updateApiToken**
> ApiToken updateApiToken(tokenId, ApiToken)



    Updates a server-to-server authentication security token configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **tokenId** | **String**| The ID of the server-to-server authentication security token. | [default to null] |
| **ApiToken** | [**ApiToken**](../Models/ApiToken.md)| ApiToken payload. | |

### Return type

[**ApiToken**](../Models/ApiToken.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

