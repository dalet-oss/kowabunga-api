# KawaiiFirewall
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **ingress** | [**List**](KawaiiFirewallIngressRule.md) | The Kawaii public firewall list of ingress rules. Kawaii default policy is to drop all incoming traffic, including ICMP. Specified ruleset will be explicitly accepted. | [optional] [default to null] |
| **egress\_policy** | **String** | The default public traffic egress policy. | [optional] [default to accept] |
| **egress** | [**List**](KawaiiFirewallEgressRule.md) | The Kawaii public firewall list of egress rules. Kawaii default policy is to accept all outgoing traffic, including ICMP. Specified ruleset will be explicitly dropped if egress_policy is set to accept, and explicitly accepted if egress policy is set to drop.. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

