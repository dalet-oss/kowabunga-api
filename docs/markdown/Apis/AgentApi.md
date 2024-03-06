# AgentApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createAgent**](AgentApi.md#createAgent) | **POST** /agent |  |
| [**deleteAgent**](AgentApi.md#deleteAgent) | **DELETE** /agent/{agentId} |  |
| [**listAgents**](AgentApi.md#listAgents) | **GET** /agent |  |
| [**readAgent**](AgentApi.md#readAgent) | **GET** /agent/{agentId} |  |
| [**setApiToken**](AgentApi.md#setApiToken) | **PATCH** /agent/{agentId}/token |  |
| [**updateAgent**](AgentApi.md#updateAgent) | **PUT** /agent/{agentId} |  |


<a name="createAgent"></a>
# **createAgent**
> Agent createAgent(Agent)



    Creates a new Kowabunga remote agent.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **Agent** | [**Agent**](../Models/Agent.md)| Agent payload. | |

### Return type

[**Agent**](../Models/Agent.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteAgent"></a>
# **deleteAgent**
> deleteAgent(agentId)



    Deletes an existing Kowabunga remote agent.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **agentId** | **String**| The ID of the Kowabunga remote agent. | [default to null] |

### Return type

null (empty response body)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listAgents"></a>
# **listAgents**
> List listAgents()



    Returns the IDs of Kowabunga remote agent objects.

### Parameters
This endpoint does not need any parameter.

### Return type

**List**

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="readAgent"></a>
# **readAgent**
> Agent readAgent(agentId)



    Returns a Kowabunga remote agent.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **agentId** | **String**| The ID of the Kowabunga remote agent. | [default to null] |

### Return type

[**Agent**](../Models/Agent.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

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

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateAgent"></a>
# **updateAgent**
> Agent updateAgent(agentId, Agent)



    Updates a Kowabunga remote agent configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **agentId** | **String**| The ID of the Kowabunga remote agent. | [default to null] |
| **Agent** | [**Agent**](../Models/Agent.md)| Agent payload. | |

### Return type

[**Agent**](../Models/Agent.md)

### Authorization

[TokenAuth](../README.md#TokenAuth), [ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

