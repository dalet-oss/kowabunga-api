# Kaktus
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** | The Kaktus computing node ID (auto-generated). | [optional] [default to null] |
| **name** | **String** | The Kaktus computing node name. | [default to null] |
| **description** | **String** | The Kaktus computing node description. | [optional] [default to null] |
| **cpu\_cost** | [**Cost**](.md) | Cost associated to the Kaktus node&#39;s CPU resources. | [optional] [default to null] |
| **memory\_cost** | [**Cost**](.md) | Cost associated to the Kaktus node&#39;s memory resources. | [optional] [default to null] |
| **overcommit\_cpu\_ratio** | **Long** | The Kaktus node CPU resource over-commit ratio. Overcommitting CPU resources for VMs means allocating more virtual CPUs (vCPUs) to the virtual machines (VMs) than the physical cores available on the node. This can help optimize the utilization of the node CPU and increase the density of VMs per node. | [optional] [default to 3] |
| **overcommit\_memory\_ratio** | **Long** | The Kaktus node memory resource over-commit ratio. Memory overcommitment is a concept in computing that covers the assignment of more memory to virtual computing devices (or processes) than the physical machine they are hosted, or running on, actually has. | [optional] [default to 2] |
| **agents** | **List** | a list of existing remote agents managing the Kaktus node. | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

