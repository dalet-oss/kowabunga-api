# Project
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** | The project ID (auto-generated). | [optional] [default to null] |
| **name** | **String** | The project name. | [default to null] |
| **description** | **String** | The project description. | [optional] [default to null] |
| **owner** | **String** | The project&#39;s owner name. | [default to null] |
| **email** | **String** | The project associated email address, used to receive notifications. | [default to null] |
| **domain** | **String** | The project associated internal domain name (e.g. myproject.acme.com). | [optional] [default to null] |
| **root\_password** | **String** | The project default root password, set at cloud-init instance bootstrap phase. Will be randomly auto-generated at each instance creation if unspecified. | [optional] [default to null] |
| **bootstrap\_user** | **String** | The project default service user name, created at cloud-init instance bootstrap phase. Will use Kowabunga&#39;s default configuration one if unspecified. | [optional] [default to null] |
| **bootstrap\_pubkey** | **String** | The project default public SSH key, to be associated to bootstrap user. Will use Kowabunga&#39;s default configuration one if unspecified. | [optional] [default to null] |
| **tags** | **List** | A list of tags to be associated to the project. | [optional] [default to null] |
| **metadatas** | [**List**](Metadata.md) | A list of metadata to be associated to the project | [optional] [default to null] |
| **quotas** | [**ProjectResources**](.md) | The global project resource quotas (0 for unlimited) | [optional] [default to null] |
| **private\_subnets** | [**List**](ZoneSubnet.md) | The assigned project VPC private subnets IDs (read-only). | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

