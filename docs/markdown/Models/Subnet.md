# Subnet
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** | The network subnet ID (auto-generated). | [optional] [default to null] |
| **name** | **String** | The network subnet name. | [default to null] |
| **description** | **String** | The network subnet description. | [optional] [default to null] |
| **cidr** | **String** | The network subnet CIDR (e.g. 192.168.0.0/24). | [default to null] |
| **gateway** | **String** | The network subnet router/gateway IP address (e.g. 192.168.0.254). | [default to null] |
| **dns** | **String** | The network subnet DNS server IP address (gateway value if unspecified). | [optional] [default to null] |
| **extra\_routes** | **List** | The list of extra routes to be access through designated gateway (format is 10.0.0.0/8). | [optional] [default to null] |
| **reserved** | [**List**](IpRange.md) | The network subnet reserved IPv4 ranges (i.e. no IP address can be assigned from there). | [optional] [default to null] |
| **gw\_pool** | [**List**](IpRange.md) | The network subnet IPv4 ranges reserved for per-zone local network gateways (range size must be at least equal to region number of zones). | [optional] [default to null] |
| **application** | **String** | Optional application service type. | [optional] [default to user] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

