# UserApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createUser**](UserApi.md#createUser) | **POST** /user |  |
| [**deleteUser**](UserApi.md#deleteUser) | **DELETE** /user/{userId} |  |
| [**listUsers**](UserApi.md#listUsers) | **GET** /user |  |
| [**login**](UserApi.md#login) | **POST** /login |  |
| [**readUser**](UserApi.md#readUser) | **GET** /user/{userId} |  |
| [**resetUserPassword**](UserApi.md#resetUserPassword) | **PATCH** /user/{userId}/resetPassword |  |
| [**setUserApiToken**](UserApi.md#setUserApiToken) | **PATCH** /user/{userId}/token |  |
| [**setUserPassword**](UserApi.md#setUserPassword) | **PUT** /user/{userId}/password |  |
| [**updateUser**](UserApi.md#updateUser) | **PUT** /user/{userId} |  |


<a name="createUser"></a>
# **createUser**
> User createUser(User)



    Creates a new Kowabunga user.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **User** | [**User**](../Models/User.md)| User payload. | |

### Return type

[**User**](../Models/User.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteUser"></a>
# **deleteUser**
> deleteUser(userId)



    Deletes an existing Kowabunga user.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| The ID of the Kowabunga user. | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listUsers"></a>
# **listUsers**
> List listUsers()



    Returns the IDs of Kowabunga user objects.

### Parameters
This endpoint does not need any parameter.

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="login"></a>
# **login**
> UserCredentials login(UserCredentials)



    Creates a new Kowabunga user login credentials.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **UserCredentials** | [**UserCredentials**](../Models/UserCredentials.md)| UserCredentials payload. | |

### Return type

[**UserCredentials**](../Models/UserCredentials.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="readUser"></a>
# **readUser**
> User readUser(userId)



    Returns a Kowabunga user.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| The ID of the Kowabunga user. | [default to null] |

### Return type

[**User**](../Models/User.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="resetUserPassword"></a>
# **resetUserPassword**
> resetUserPassword(userId)



    Performs a Kowabunga user reset of password (server-side generated, will replace any existing one).

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| The ID of the Kowabunga user. | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="setUserApiToken"></a>
# **setUserApiToken**
> setUserApiToken(userId, expire, expiration\_date)



    Performs a Kowabunga user setting of API token (will replace any existing one).

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| The ID of the Kowabunga user. | [default to null] |
| **expire** | **Boolean**| Whether or not the token should expire. | [optional] [default to null] |
| **expiration\_date** | **date**| Token&#39;s expiration date (YYYY-MM-DD format). | [optional] [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="setUserPassword"></a>
# **setUserPassword**
> setUserPassword(userId, Password)



    Updates a Kowabunga user password (will replace any existing one) configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| The ID of the Kowabunga user. | [default to null] |
| **Password** | [**Password**](../Models/Password.md)| Password payload. | |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="updateUser"></a>
# **updateUser**
> User updateUser(userId, User)



    Updates a Kowabunga user configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **userId** | **String**| The ID of the Kowabunga user. | [default to null] |
| **User** | [**User**](../Models/User.md)| User payload. | |

### Return type

[**User**](../Models/User.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

