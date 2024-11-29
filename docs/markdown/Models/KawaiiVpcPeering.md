# KawaiiVpcPeering
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **subnet** | **String** | Kowabunga Subnet ID to be peered with (subnet local IP addresses will be automatically assigned to Kawaii instances).. | [default to null] |
| **policy** | **String** | The default VPC traffic forwarding policy. | [optional] [default to drop] |
| **ingress** | [**List**](KawaiiVpcForwardRule.md) | The firewall list of forwarding ingress rules from VPC peered subnet. ICMP traffic is always accepted. The specified ruleset will be explicitly accepted if drop is the default policy (useless otherwise). | [optional] [default to null] |
| **egress** | [**List**](KawaiiVpcForwardRule.md) | The firewall list of forwarding egress rules to VPC peered subnet. ICMP traffic is always accepted. The specified ruleset will be explicitly accepted if drop is the default policy (useless otherwise). | [optional] [default to null] |
| **netip** | [**List**](KawaiiVpcNetIpZone.md) | The per-zone auto-assigned private IPs in peered subnet (read-only). | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

