# StoragePool
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** | The storage pool ID (auto-generated). | [optional] [default to null] |
| **name** | **String** | The storage pool name. | [default to null] |
| **description** | **String** | The storage pool description. | [optional] [default to null] |
| **type** | **String** | The storage pool type. | [optional] [default to rbd] |
| **pool** | **String** | The libvirt pool name. | [default to null] |
| **ceph\_address** | **String** | The local Ceph Monitor(s) address or FQDN, empty for local pool type. | [optional] [default to localhost] |
| **ceph\_port** | **Long** | The local Ceph Monitor(s) port (default 3300), empty for local pool type. | [optional] [default to 3300] |
| **ceph\_secret\_uuid** | **String** | The libvirt secret UUID for CephX authentication, empty for local pool type. | [optional] [default to null] |
| **cost** | [**Cost**](.md) | Cost associated to the storage pool. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

