# Subnet
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** | The subnet ID (auto-generated). | [optional] [default to null] |
| **name** | **String** | The subnet name. | [default to null] |
| **description** | **String** | The subnet description. | [optional] [default to null] |
| **cidr** | **String** | The subnet CIDR (e.g. 192.168.0.0/24). | [default to null] |
| **gateway** | **String** | The subnet router/gateway IP address (e.g. 192.168.0.254). | [default to null] |
| **dns** | **String** | The subnet DNS server IP address (gateway value if unspecified). | [optional] [default to null] |
| **extra\_routes** | **List** | The list of extra routes to be access through designated gateway (format is 10.0.0.0/8). | [optional] [default to null] |
| **reserved** | [**List**](IpRange.md) | The subnet list of reserved IPv4 ranges (i.e. no IP address can be assigned from there). | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

