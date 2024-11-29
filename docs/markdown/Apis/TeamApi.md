# TeamApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createTeam**](TeamApi.md#createTeam) | **POST** /team |  |
| [**deleteTeam**](TeamApi.md#deleteTeam) | **DELETE** /team/{teamId} |  |
| [**listTeams**](TeamApi.md#listTeams) | **GET** /team |  |
| [**readTeam**](TeamApi.md#readTeam) | **GET** /team/{teamId} |  |
| [**updateTeam**](TeamApi.md#updateTeam) | **PUT** /team/{teamId} |  |


<a name="createTeam"></a>
# **createTeam**
> Team createTeam(Team)



    Creates a new Kowabunga users team.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **Team** | [**Team**](../Models/Team.md)| Team payload. | |

### Return type

[**Team**](../Models/Team.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteTeam"></a>
# **deleteTeam**
> deleteTeam(teamId)



    Deletes an existing Kowabunga users team.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **teamId** | **String**| The ID of the Kowabunga users team. | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listTeams"></a>
# **listTeams**
> List listTeams()



    Returns the IDs of Kowabunga users team objects.

### Parameters
This endpoint does not need any parameter.

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="readTeam"></a>
# **readTeam**
> Team readTeam(teamId)



    Returns a Kowabunga users team.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **teamId** | **String**| The ID of the Kowabunga users team. | [default to null] |

### Return type

[**Team**](../Models/Team.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateTeam"></a>
# **updateTeam**
> Team updateTeam(teamId, Team)



    Updates a Kowabunga users team configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **teamId** | **String**| The ID of the Kowabunga users team. | [default to null] |
| **Team** | [**Team**](../Models/Team.md)| Team payload. | |

### Return type

[**Team**](../Models/Team.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

