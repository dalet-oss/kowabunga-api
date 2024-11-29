# Kylo
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** | The Kylo ID (auto-generated). | [optional] [default to null] |
| **name** | **String** | The Kylo name. | [default to null] |
| **description** | **String** | The Kylo description. | [optional] [default to null] |
| **access** | **String** | The Kylo volume access type. | [optional] [default to RW] |
| **protocols** | **List** | The Kylo NFS protocol versions to be supported. | [optional] [default to [3, 4]] |
| **endpoint** | **String** | The Kylo endpoint FQDN (read-only). | [optional] [default to null] |
| **size** | **Long** | The Kylo volume bytes used (read-only). | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

