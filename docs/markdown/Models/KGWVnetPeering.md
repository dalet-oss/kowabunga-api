# KGWVnetPeering
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **vnet** | **String** | Kowabunga VNET ID to be peered with.. | [default to null] |
| **subnet** | **String** | Kowabunga Subnet ID to be peered with (IP addresses will be automatically assigned).. | [default to null] |
| **ports** | **String** | Ports to be reachable from peered subnet. Accept Ranges. Unless specified, traffic will be filtered.. | [optional] [default to null] |
| **ips** | **List** | The KGW (Kowabunga Network Gateway) auto-assigned private IPs in peered subnet (read-only). | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

