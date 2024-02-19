# Adapter
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** | The network adapter ID (auto-generated) | [optional] [default to null] |
| **name** | **String** | The network adapter name | [default to null] |
| **description** | **String** | The network adapter description | [optional] [default to null] |
| **mac** | **String** | The network adapter hardware address (e.g. 00:11:22:33:44:55). Auto-generated if unspecified | [optional] [default to null] |
| **addresses** | **List** | The network adapter list of associated IPv4 addresses | [optional] [default to null] |
| **reserved** | **Boolean** | The network adapter is a reserved adapter (e.g. router), where the same hardware address can be reused over several subnets | [optional] [default to false] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

