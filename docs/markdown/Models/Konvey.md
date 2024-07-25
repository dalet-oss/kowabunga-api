# Konvey
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** | The Konvey (Kowabunga Network Load-Balancer) ID (auto-generated). | [optional] [default to null] |
| **name** | **String** | The Konvey (Kowabunga Network Load-Balancer) name. | [optional] [default to null] |
| **description** | **String** | The Konvey (Kowabunga Network Load-Balancer) description. | [optional] [default to null] |
| **vip** | **String** | The Konvey (Kowabunga Network Load-Balancer) assigned private virtual IP address (read-only). | [optional] [default to null] |
| **failover** | **Boolean** | Whether Konvey (Kowabunga Network Load-Balancer) must be deployed in a highly-available replicated state to support service failover. | [optional] [default to true] |
| **endpoints** | [**List**](KonveyEndpoint.md) | The Konvey (Kowabunga Network Load-Balancer) list of load-balanced endpoints. | [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

