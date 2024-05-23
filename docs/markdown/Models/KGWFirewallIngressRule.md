# KGWFirewallIngressRule
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **source** | **String** | The source IP or CIDR to accept public traffic from. | [optional] [default to 0.0.0.0/0] |
| **protocol** | **String** | The transport layer protocol to accept public traffic from. | [optional] [default to tcp] |
| **ports** | **String** | The port (or list of ports) to accept public traffic from. Ranges are accepted. Format is a-b,c-d (e.g. 443; 22,80,443; 80,443,3000-3005). | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

