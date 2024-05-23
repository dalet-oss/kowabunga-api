# KGW
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** | The KGW (Kowabunga Network Gateway) ID (auto-generated). | [optional] [default to null] |
| **name** | **String** | The KGW (Kowabunga Network Gateway) name. | [optional] [default to null] |
| **description** | **String** | The KGW (Kowabunga Network Gateway) description. | [optional] [default to null] |
| **netip** | [**List**](KGWNetIp.md) | The KGW (Kowabunga Network Gateway) list of assigned virtual IPs per-zone addresses (read-only). | [optional] [default to null] |
| **firewall** | [**List**](KGWFirewall.md) | The KGW (Kowabunga Network Gateway) firewall settings from/to public Internet). | [optional] [default to null] |
| **dnat** | [**List**](KGWDNatRule.md) | The KGW (Kowabunga Network Gateway) list of NAT forwarding entries. KGW will forward public Internet traffic from all public virtual IPs to requested private subnet IP addresses. | [optional] [default to null] |
| **vpc\_peerings** | [**List**](KGWVpcPeering.md) | The KGW (Kowabunga Network Gateway) list of Kowabunga private VPC subnet peering entries. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

