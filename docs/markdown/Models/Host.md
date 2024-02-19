# Host
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** | The computing hosts ID (auto-generated) | [optional] [default to null] |
| **name** | **String** | The host name | [default to null] |
| **description** | **String** | The host description | [optional] [default to null] |
| **protocol** | **String** | The protocol to use to issue libvirt connection. | [default to null] |
| **address** | **String** | The host libvirt&#39;s IPv4 address. | [default to null] |
| **port** | **Integer** | The host libvirt&#39;s port. | [optional] [default to null] |
| **tls** | [**Host_tls**](Host_tls.md) |  | [optional] [default to null] |
| **cost** | [**Cost**](.md) | Global cost associated to the host (deprecated, will be removed). | [optional] [default to null] |
| **cpu\_cost** | [**Cost**](.md) | Cost associated to the host&#39;s CPU resources. | [optional] [default to null] |
| **memory\_cost** | [**Cost**](.md) | Cost associated to the host&#39;s memory resources. | [optional] [default to null] |
| **overcommit\_cpu\_ratio** | **Integer** | The host CPU resource over-commit ratio. Overcommitting CPU resources for VMs means allocating more virtual CPUs (vCPUs) to the virtual machines (VMs) than the physical cores available on the host. This can help optimize the utilization of the host CPU and increase the density of VMs per host. | [optional] [default to 3] |
| **overcommit\_memory\_ratio** | **Integer** | The host memory resource over-commit ratio. Memory overcommitment is a concept in computing that covers the assignment of more memory to virtual computing devices (or processes) than the physical machine they are hosted, or running on, actually has. | [optional] [default to 2] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

