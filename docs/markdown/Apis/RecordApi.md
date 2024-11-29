# RecordApi

All URIs are relative to */api/v1*

| Method | HTTP request | Description |
|------------- | ------------- | -------------|
| [**createProjectDnsRecord_0**](RecordApi.md#createProjectDnsRecord_0) | **POST** /project/{projectId}/record |  |
| [**deleteDnsRecord**](RecordApi.md#deleteDnsRecord) | **DELETE** /record/{recordId} |  |
| [**listProjectDnsRecords_0**](RecordApi.md#listProjectDnsRecords_0) | **GET** /project/{projectId}/records |  |
| [**readDnsRecord**](RecordApi.md#readDnsRecord) | **GET** /record/{recordId} |  |
| [**updateDnsRecord**](RecordApi.md#updateDnsRecord) | **PUT** /record/{recordId} |  |


<a name="createProjectDnsRecord_0"></a>
# **createProjectDnsRecord_0**
> DnsRecord createProjectDnsRecord_0(projectId, DnsRecord)



    Creates a new DNS record.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |
| **DnsRecord** | [**DnsRecord**](../Models/DnsRecord.md)| DnsRecord payload. | |

### Return type

[**DnsRecord**](../Models/DnsRecord.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

<a name="deleteDnsRecord"></a>
# **deleteDnsRecord**
> deleteDnsRecord(recordId)



    Deletes an existing DNS record.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **recordId** | **String**| The ID of the DNS record. | [default to null] |

### Return type

null (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="listProjectDnsRecords_0"></a>
# **listProjectDnsRecords_0**
> List listProjectDnsRecords_0(projectId)



    Returns the IDs of DNS record objects.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **projectId** | **String**| The ID of the project. | [default to null] |

### Return type

**List**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="readDnsRecord"></a>
# **readDnsRecord**
> DnsRecord readDnsRecord(recordId)



    Returns a DNS record.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **recordId** | **String**| The ID of the DNS record. | [default to null] |

### Return type

[**DnsRecord**](../Models/DnsRecord.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

<a name="updateDnsRecord"></a>
# **updateDnsRecord**
> DnsRecord updateDnsRecord(recordId, DnsRecord)



    Updates a DNS record configuration.

### Parameters

|Name | Type | Description  | Notes |
|------------- | ------------- | ------------- | -------------|
| **recordId** | **String**| The ID of the DNS record. | [default to null] |
| **DnsRecord** | [**DnsRecord**](../Models/DnsRecord.md)| DnsRecord payload. | |

### Return type

[**DnsRecord**](../Models/DnsRecord.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

