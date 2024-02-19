# KFS
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** | The KFS ID (auto-generated). | [optional] [default to null] |
| **name** | **String** | The KFS storage volume name. | [default to null] |
| **description** | **String** | The KFS storage volume description. | [optional] [default to null] |
| **access** | **String** | The KFS storage volume access type. | [optional] [default to RW] |
| **protocols** | **List** | The KFS storage volume&#39;s NFS protocol versions to be supported. | [optional] [default to [3, 4]] |
| **endpoint** | **String** | The KFS endpoint FQDN (read-only). | [optional] [default to null] |
| **size** | **Integer** | The KFS storage volume bytes used (read-only). | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

