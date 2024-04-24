# StoragePool
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** | The storage pool ID (auto-generated). | [optional] [default to null] |
| **name** | **String** | The storage pool name. | [default to null] |
| **description** | **String** | The storage pool description. | [optional] [default to null] |
| **pool** | **String** | Ceph pool name. | [default to null] |
| **ceph\_address** | **String** | Ceph Monitor(s) address or FQDN. | [optional] [default to localhost] |
| **ceph\_port** | **Long** | Ceph Monitor(s) port (default 3300). | [optional] [default to 3300] |
| **ceph\_secret\_uuid** | **String** | The libvirt secret UUID for CephX authentication. | [optional] [default to null] |
| **cost** | [**Cost**](.md) | Cost associated to the storage pool. | [optional] [default to null] |
| **agents** | **List** | a list of existing remote agents managing the storage pool. | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

