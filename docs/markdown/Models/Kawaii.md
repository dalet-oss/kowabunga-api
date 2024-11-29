# Kawaii
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** | The Kawaii ID (auto-generated). | [optional] [default to null] |
| **name** | **String** | The Kawaii name. | [optional] [default to null] |
| **description** | **String** | The Kawaii description. | [optional] [default to null] |
| **netip** | [**KawaiiNetIp**](.md) | The Kawaii list of assigned virtual IPs per-zone addresses (read-only). | [optional] [default to null] |
| **firewall** | [**KawaiiFirewall**](.md) | The Kawaii firewall settings from/to public Internet). | [optional] [default to null] |
| **dnat** | [**List**](KawaiiDNatRule.md) | The Kawaii list of NAT forwarding entries. Kawaii will forward public Internet traffic from all public virtual IPs to requested private subnet IP addresses. | [optional] [default to null] |
| **vpc\_peerings** | [**List**](KawaiiVpcPeering.md) | The Kawaii list of Kowabunga private VPC subnet peering entries. | [optional] [default to null] |
| **ipsec\_connections** | [**List**](KawaiiIpSec.md) | The Kawaii list of Kowabunga ipsecs connections. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

