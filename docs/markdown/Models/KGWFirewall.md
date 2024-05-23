# KGWFirewall
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **ingress** | [**List**](KGWFirewallIngressRule.md) | The KGW (Kowabunga Network Gateway) public firewall list of ingress rules. KGW default policy is to drop all incoming traffic, including ICMP. Specified ruleset will be explicitly accepted. | [optional] [default to null] |
| **egress\_policy** | **String** | The default public traffic egress policy. | [optional] [default to accept] |
| **egress** | [**List**](KGWFirewallEgressRule.md) | The KGW (Kowabunga Network Gateway) public firewall list of egress rules. KGW default policy is to accept all outgoing traffic, including ICMP. Specified ruleset will be explicitly dropped if egress_policy is set to accept, and explicitly accepted if egress policy is set to drop.. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

