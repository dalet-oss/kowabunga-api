# Host
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** | The computing host ID (auto-generated). | [optional] [default to null] |
| **name** | **String** | The computing host name. | [default to null] |
| **description** | **String** | The computing host description. | [optional] [default to null] |
| **protocol** | **String** | The protocol to use to issue libvirt connection. | [default to null] |
| **address** | **String** | The host libvirt&#39;s IPv4 address. | [default to null] |
| **port** | **Long** | The host libvirt&#39;s port. | [optional] [default to null] |
| **tls** | [**HostTLS**](.md) | The host libvirt&#39;s TLS configuration. | [optional] [default to null] |
| **cpu\_cost** | [**Cost**](.md) | Cost associated to the host&#39;s CPU resources. | [optional] [default to null] |
| **memory\_cost** | [**Cost**](.md) | Cost associated to the host&#39;s memoery resources. | [optional] [default to null] |
| **overcommit\_cpu\_ratio** | **Long** | The host CPU resource over-commit ratio. Overcommitting CPU resources for VMs means allocating more virtual CPUs (vCPUs) to the virtual machines (VMs) than the physical cores available on the host. This can help optimize the utilization of the host CPU and increase the density of VMs per host. | [optional] [default to 3] |
| **overcommit\_memory\_ratio** | **Long** | The host memory resource over-commit ratio. Memory overcommitment is a concept in computing that covers the assignment of more memory to virtual computing devices (or processes) than the physical machine they are hosted, or running on, actually has. | [optional] [default to 2] |
| **agents** | **List** | a list of existing remote agents managing the host. | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

