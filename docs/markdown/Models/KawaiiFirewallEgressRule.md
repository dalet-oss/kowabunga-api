# KawaiiFirewallEgressRule
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **destination** | **String** | The destination IP or CIDR to accept/drop public traffic to. | [optional] [default to 0.0.0.0/0] |
| **protocol** | **String** | The transport layer protocol to accept/drop public traffic to. | [optional] [default to tcp] |
| **ports** | **String** | The port (or list of ports) to accept/drop public traffic from. Ranges are accepted. Format is a-b,c-d (e.g. 443; 22,80,443; 80,443,3000-3005). | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

