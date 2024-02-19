# StorageNFS
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** | The NFS storage ID (auto-generated). | [optional] [default to null] |
| **name** | **String** | The NFS storage name. | [default to null] |
| **description** | **String** | The NFS storage description. | [optional] [default to null] |
| **endpoint** | **String** | The associated NFS endpoint FQDN. | [default to null] |
| **fs** | **String** | The underlying associated CephFS volume name. | [optional] [default to nfs] |
| **backends** | **List** | List of NFS Ganesha API server IP addresses. | [optional] [default to null] |
| **port** | **Integer** | NFS Ganesha API server port (default 54934). | [optional] [default to 54934] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

