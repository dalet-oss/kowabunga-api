# Documentation for Kowabunga API documentation

<a name="documentation-for-api-endpoints"></a>
## Documentation for API Endpoints

All URIs are relative to */api/v1*

| Class | Method | HTTP request | Description |
|------------ | ------------- | ------------- | -------------|
| *AdapterApi* | [**createAdapter_0**](Apis/AdapterApi.md#createadapter_0) | **POST** /subnet/{subnetId}/adapter | Creates a new network adapter. |
*AdapterApi* | [**deleteAdapter**](Apis/AdapterApi.md#deleteadapter) | **DELETE** /adapter/{adapterId} | Deletes an existing network adapter. |
*AdapterApi* | [**listAdapters**](Apis/AdapterApi.md#listadapters) | **GET** /adapter | Returns the IDs of network adapter objects. |
*AdapterApi* | [**listSubnetAdapters_0**](Apis/AdapterApi.md#listsubnetadapters_0) | **GET** /subnet/{subnetId}/adapters | Returns the IDs of network adapter objects. |
*AdapterApi* | [**readAdapter**](Apis/AdapterApi.md#readadapter) | **GET** /adapter/{adapterId} | Returns a network adapter. |
*AdapterApi* | [**updateAdapter**](Apis/AdapterApi.md#updateadapter) | **PUT** /adapter/{adapterId} | Updates a network adapter configuration. |
| *AgentApi* | [**createAgent**](Apis/AgentApi.md#createagent) | **POST** /agent | Creates a new Kowabunga remote agent. |
*AgentApi* | [**deleteAgent**](Apis/AgentApi.md#deleteagent) | **DELETE** /agent/{agentId} | Deletes an existing Kowabunga remote agent. |
*AgentApi* | [**listAgents**](Apis/AgentApi.md#listagents) | **GET** /agent | Returns the IDs of Kowabunga remote agent objects. |
*AgentApi* | [**readAgent**](Apis/AgentApi.md#readagent) | **GET** /agent/{agentId} | Returns a Kowabunga remote agent. |
*AgentApi* | [**setAgentApiToken**](Apis/AgentApi.md#setagentapitoken) | **PATCH** /agent/{agentId}/token | Performs a Kowabunga remote agent setting of API token (will replace any existing one). |
*AgentApi* | [**updateAgent**](Apis/AgentApi.md#updateagent) | **PUT** /agent/{agentId} | Updates a Kowabunga remote agent configuration. |
| *InstanceApi* | [**createProjectZoneInstance_1**](Apis/InstanceApi.md#createprojectzoneinstance_1) | **POST** /project/{projectId}/zone/{zoneId}/instance | Creates a new virtual machine instance. |
*InstanceApi* | [**deleteInstance**](Apis/InstanceApi.md#deleteinstance) | **DELETE** /instance/{instanceId} | Deletes an existing virtual machine instance. |
*InstanceApi* | [**listInstances**](Apis/InstanceApi.md#listinstances) | **GET** /instance | Returns the IDs of virtual machine instance objects. |
*InstanceApi* | [**listKaktusInstances_0**](Apis/InstanceApi.md#listkaktusinstances_0) | **GET** /kaktus/{kaktusId}/instances | Returns the IDs of virtual machine instance objects. |
*InstanceApi* | [**listProjectZoneInstances_1**](Apis/InstanceApi.md#listprojectzoneinstances_1) | **GET** /project/{projectId}/zone/{zoneId}/instances | Returns the IDs of virtual machine instance objects. |
*InstanceApi* | [**readInstance**](Apis/InstanceApi.md#readinstance) | **GET** /instance/{instanceId} | Returns a virtual machine instance. |
*InstanceApi* | [**readInstanceRemoteConnection**](Apis/InstanceApi.md#readinstanceremoteconnection) | **GET** /instance/{instanceId}/connect | Returns a virtual machine instance remote access characteristics. |
*InstanceApi* | [**readInstanceState**](Apis/InstanceApi.md#readinstancestate) | **GET** /instance/{instanceId}/state | Returns a virtual machine instance state. |
*InstanceApi* | [**rebootInstance**](Apis/InstanceApi.md#rebootinstance) | **PATCH** /instance/{instanceId}/reboot | Performs a virtual machine instance software reboot. |
*InstanceApi* | [**resetInstance**](Apis/InstanceApi.md#resetinstance) | **PATCH** /instance/{instanceId}/reset | Performs a virtual machine instance hardware reset. |
*InstanceApi* | [**resumeInstance**](Apis/InstanceApi.md#resumeinstance) | **PATCH** /instance/{instanceId}/resume | Performs a virtual machine instance software PM resume. |
*InstanceApi* | [**shutdownInstance**](Apis/InstanceApi.md#shutdowninstance) | **PATCH** /instance/{instanceId}/shutdown | Performs a virtual machine instance software shutdown. |
*InstanceApi* | [**startInstance**](Apis/InstanceApi.md#startinstance) | **PATCH** /instance/{instanceId}/start | Performs a virtual machine instance hardware boot-up. |
*InstanceApi* | [**stopInstance**](Apis/InstanceApi.md#stopinstance) | **PATCH** /instance/{instanceId}/stop | Performs a virtual machine instance hardware stop. |
*InstanceApi* | [**suspendInstance**](Apis/InstanceApi.md#suspendinstance) | **PATCH** /instance/{instanceId}/suspend | Performs a virtual machine instance software PM suspend. |
*InstanceApi* | [**updateInstance**](Apis/InstanceApi.md#updateinstance) | **PUT** /instance/{instanceId} | Updates a virtual machine instance configuration. |
| *KaktusApi* | [**createKaktus_0**](Apis/KaktusApi.md#createkaktus_0) | **POST** /zone/{zoneId}/kaktus | Creates a new Kaktus computing node. |
*KaktusApi* | [**deleteKaktus**](Apis/KaktusApi.md#deletekaktus) | **DELETE** /kaktus/{kaktusId} | Deletes an existing Kaktus computing node. |
*KaktusApi* | [**listKaktusInstances**](Apis/KaktusApi.md#listkaktusinstances) | **GET** /kaktus/{kaktusId}/instances | Returns the IDs of virtual machine instance objects. |
*KaktusApi* | [**listKaktuss**](Apis/KaktusApi.md#listkaktuss) | **GET** /kaktus | Returns the IDs of Kaktus computing node objects. |
*KaktusApi* | [**listZoneKaktuses_0**](Apis/KaktusApi.md#listzonekaktuses_0) | **GET** /zone/{zoneId}/kaktuses | Returns the IDs of Kaktus computing node objects. |
*KaktusApi* | [**readKaktus**](Apis/KaktusApi.md#readkaktus) | **GET** /kaktus/{kaktusId} | Returns a Kaktus computing node. |
*KaktusApi* | [**readKaktusCaps**](Apis/KaktusApi.md#readkaktuscaps) | **GET** /kaktus/{kaktusId}/caps | Returns a Kaktus computing node capability. |
*KaktusApi* | [**updateKaktus**](Apis/KaktusApi.md#updatekaktus) | **PUT** /kaktus/{kaktusId} | Updates a Kaktus computing node configuration. |
| *KawaiiApi* | [**createProjectRegionKawaii_1**](Apis/KawaiiApi.md#createprojectregionkawaii_1) | **POST** /project/{projectId}/region/{regionId}/kawaii | Creates a new Kawaii. |
*KawaiiApi* | [**deleteKawaii**](Apis/KawaiiApi.md#deletekawaii) | **DELETE** /kawaii/{kawaiiId} | Deletes an existing Kawaii. |
*KawaiiApi* | [**listKawaiis**](Apis/KawaiiApi.md#listkawaiis) | **GET** /kawaii | Returns the IDs of Kawaii objects. |
*KawaiiApi* | [**listProjectRegionKawaiis_1**](Apis/KawaiiApi.md#listprojectregionkawaiis_1) | **GET** /project/{projectId}/region/{regionId}/kawaiis | Returns the IDs of Kawaii objects. |
*KawaiiApi* | [**readKawaii**](Apis/KawaiiApi.md#readkawaii) | **GET** /kawaii/{kawaiiId} | Returns a Kawaii. |
*KawaiiApi* | [**updateKawaii**](Apis/KawaiiApi.md#updatekawaii) | **PUT** /kawaii/{kawaiiId} | Updates a Kawaii configuration. |
| *KiwiApi* | [**createKiwi_0**](Apis/KiwiApi.md#createkiwi_0) | **POST** /region/{regionId}/kiwi | Creates a new Kiwi (Kowabunga Inner Wan Interface) provides edge-network services.. |
*KiwiApi* | [**deleteKiwi**](Apis/KiwiApi.md#deletekiwi) | **DELETE** /kiwi/{kiwiId} | Deletes an existing Kiwi (Kowabunga Inner Wan Interface) provides edge-network services.. |
*KiwiApi* | [**listKiwis**](Apis/KiwiApi.md#listkiwis) | **GET** /kiwi | Returns the IDs of Kiwi (Kowabunga Inner Wan Interface) provides edge-network services. objects. |
*KiwiApi* | [**listRegionKiwis_0**](Apis/KiwiApi.md#listregionkiwis_0) | **GET** /region/{regionId}/kiwis | Returns the IDs of Kiwi (Kowabunga Inner Wan Interface) provides edge-network services. objects. |
*KiwiApi* | [**readKiwi**](Apis/KiwiApi.md#readkiwi) | **GET** /kiwi/{kiwiId} | Returns a Kiwi (Kowabunga Inner Wan Interface) provides edge-network services.. |
*KiwiApi* | [**updateKiwi**](Apis/KiwiApi.md#updatekiwi) | **PUT** /kiwi/{kiwiId} | Updates a Kiwi (Kowabunga Inner Wan Interface) provides edge-network services. configuration. |
| *KomputeApi* | [**createProjectZoneKompute_1**](Apis/KomputeApi.md#createprojectzonekompute_1) | **POST** /project/{projectId}/zone/{zoneId}/kompute | Creates a new Kompute. |
*KomputeApi* | [**deleteKompute**](Apis/KomputeApi.md#deletekompute) | **DELETE** /kompute/{komputeId} | Deletes an existing Kompute. |
*KomputeApi* | [**listKomputes**](Apis/KomputeApi.md#listkomputes) | **GET** /kompute | Returns the IDs of Kompute objects. |
*KomputeApi* | [**listProjectZoneKomputes_1**](Apis/KomputeApi.md#listprojectzonekomputes_1) | **GET** /project/{projectId}/zone/{zoneId}/komputes | Returns the IDs of Kompute objects. |
*KomputeApi* | [**readKompute**](Apis/KomputeApi.md#readkompute) | **GET** /kompute/{komputeId} | Returns a Kompute. |
*KomputeApi* | [**readKomputeState**](Apis/KomputeApi.md#readkomputestate) | **GET** /kompute/{komputeId}/state | Returns a virtual machine instance state. |
*KomputeApi* | [**rebootKompute**](Apis/KomputeApi.md#rebootkompute) | **PATCH** /kompute/{komputeId}/reboot | Performs a Kompute software reboot. |
*KomputeApi* | [**resetKompute**](Apis/KomputeApi.md#resetkompute) | **PATCH** /kompute/{komputeId}/reset | Performs a Kompute hardware reset. |
*KomputeApi* | [**resumeKompute**](Apis/KomputeApi.md#resumekompute) | **PATCH** /kompute/{komputeId}/resume | Performs a Kompute software PM resume. |
*KomputeApi* | [**shutdownKompute**](Apis/KomputeApi.md#shutdownkompute) | **PATCH** /kompute/{komputeId}/shutdown | Performs a Kompute software shutdown. |
*KomputeApi* | [**startKompute**](Apis/KomputeApi.md#startkompute) | **PATCH** /kompute/{komputeId}/start | Performs a Kompute hardware boot-up. |
*KomputeApi* | [**stopKompute**](Apis/KomputeApi.md#stopkompute) | **PATCH** /kompute/{komputeId}/stop | Performs a Kompute hardware stop. |
*KomputeApi* | [**suspendKompute**](Apis/KomputeApi.md#suspendkompute) | **PATCH** /kompute/{komputeId}/suspend | Performs a Kompute software PM suspend. |
*KomputeApi* | [**updateKompute**](Apis/KomputeApi.md#updatekompute) | **PUT** /kompute/{komputeId} | Updates a Kompute configuration. |
| *KonveyApi* | [**createProjectRegionKonvey_1**](Apis/KonveyApi.md#createprojectregionkonvey_1) | **POST** /project/{projectId}/region/{regionId}/konvey | Creates a new Konvey (Kowabunga Network Load-Balancer). |
*KonveyApi* | [**createProjectZoneKonvey_1**](Apis/KonveyApi.md#createprojectzonekonvey_1) | **POST** /project/{projectId}/zone/{zoneId}/konvey | Creates a new Konvey (Kowabunga Network Load-Balancer). |
*KonveyApi* | [**deleteKonvey**](Apis/KonveyApi.md#deletekonvey) | **DELETE** /konvey/{konveyId} | Deletes an existing Konvey (Kowabunga Network Load-Balancer). |
*KonveyApi* | [**listKonveys**](Apis/KonveyApi.md#listkonveys) | **GET** /konvey | Returns the IDs of Konvey (Kowabunga Network Load-Balancer) objects. |
*KonveyApi* | [**listProjectRegionKonveys_1**](Apis/KonveyApi.md#listprojectregionkonveys_1) | **GET** /project/{projectId}/region/{regionId}/konveys | Returns the IDs of Konvey (Kowabunga Network Load-Balancer) objects. |
*KonveyApi* | [**listProjectZoneKonveys_1**](Apis/KonveyApi.md#listprojectzonekonveys_1) | **GET** /project/{projectId}/zone/{zoneId}/konveys | Returns the IDs of Konvey (Kowabunga Network Load-Balancer) objects. |
*KonveyApi* | [**readKonvey**](Apis/KonveyApi.md#readkonvey) | **GET** /konvey/{konveyId} | Returns a Konvey (Kowabunga Network Load-Balancer). |
*KonveyApi* | [**updateKonvey**](Apis/KonveyApi.md#updatekonvey) | **PUT** /konvey/{konveyId} | Updates a Konvey (Kowabunga Network Load-Balancer) configuration. |
| *KyloApi* | [**createProjectRegionKylo_1**](Apis/KyloApi.md#createprojectregionkylo_1) | **POST** /project/{projectId}/region/{regionId}/kylo | Creates a new Kylo. |
*KyloApi* | [**deleteKylo**](Apis/KyloApi.md#deletekylo) | **DELETE** /kylo/{kyloId} | Deletes an existing Kylo. |
*KyloApi* | [**listKylos**](Apis/KyloApi.md#listkylos) | **GET** /kylo | Returns the IDs of Kylo objects. |
*KyloApi* | [**listProjectRegionKylos_1**](Apis/KyloApi.md#listprojectregionkylos_1) | **GET** /project/{projectId}/region/{regionId}/kylo | Returns the IDs of Kylo objects. |
*KyloApi* | [**listStorageNFSKylos_0**](Apis/KyloApi.md#liststoragenfskylos_0) | **GET** /nfs/{nfsId}/kylo | Returns the IDs of Kylo objects. |
*KyloApi* | [**readKylo**](Apis/KyloApi.md#readkylo) | **GET** /kylo/{kyloId} | Returns a Kylo. |
*KyloApi* | [**updateKylo**](Apis/KyloApi.md#updatekylo) | **PUT** /kylo/{kyloId} | Updates a Kylo configuration. |
| *NfsApi* | [**createStorageNFS_0**](Apis/NfsApi.md#createstoragenfs_0) | **POST** /region/{regionId}/nfs | Creates a new NFS storage. |
*NfsApi* | [**deleteStorageNFS**](Apis/NfsApi.md#deletestoragenfs) | **DELETE** /nfs/{nfsId} | Deletes an existing NFS storage. |
*NfsApi* | [**listRegionStorageNFSs_0**](Apis/NfsApi.md#listregionstoragenfss_0) | **GET** /region/{regionId}/nfs | Returns the IDs of NFS storage objects. |
*NfsApi* | [**listStorageNFSKylos**](Apis/NfsApi.md#liststoragenfskylos) | **GET** /nfs/{nfsId}/kylo | Returns the IDs of Kylo objects. |
*NfsApi* | [**listStorageNFSs**](Apis/NfsApi.md#liststoragenfss) | **GET** /nfs | Returns the IDs of NFS storage objects. |
*NfsApi* | [**readStorageNFS**](Apis/NfsApi.md#readstoragenfs) | **GET** /nfs/{nfsId} | Returns a NFS storage. |
*NfsApi* | [**setRegionDefaultStorageNFS_0**](Apis/NfsApi.md#setregiondefaultstoragenfs_0) | **PATCH** /region/{regionId}/nfs/{nfsId}/default | Performs a region setting of default NFS storage. |
*NfsApi* | [**updateStorageNFS**](Apis/NfsApi.md#updatestoragenfs) | **PUT** /nfs/{nfsId} | Updates a NFS storage configuration. |
| *PoolApi* | [**createStoragePool_0**](Apis/PoolApi.md#createstoragepool_0) | **POST** /region/{regionId}/pool | Creates a new storage pool. |
*PoolApi* | [**createTemplate**](Apis/PoolApi.md#createtemplate) | **POST** /pool/{poolId}/template | Creates a new image template. |
*PoolApi* | [**deleteStoragePool**](Apis/PoolApi.md#deletestoragepool) | **DELETE** /pool/{poolId} | Deletes an existing storage pool. |
*PoolApi* | [**listRegionStoragePools_0**](Apis/PoolApi.md#listregionstoragepools_0) | **GET** /region/{regionId}/pools | Returns the IDs of storage pool objects. |
*PoolApi* | [**listStoragePoolTemplates**](Apis/PoolApi.md#liststoragepooltemplates) | **GET** /pool/{poolId}/templates | Returns the IDs of image template objects. |
*PoolApi* | [**listStoragePoolVolumes**](Apis/PoolApi.md#liststoragepoolvolumes) | **GET** /pool/{poolId}/volumes | Returns the IDs of storage volume objects. |
*PoolApi* | [**listStoragePools**](Apis/PoolApi.md#liststoragepools) | **GET** /pool | Returns the IDs of storage pool objects. |
*PoolApi* | [**readStoragePool**](Apis/PoolApi.md#readstoragepool) | **GET** /pool/{poolId} | Returns a storage pool. |
*PoolApi* | [**setRegionDefaultStoragePool_0**](Apis/PoolApi.md#setregiondefaultstoragepool_0) | **PATCH** /region/{regionId}/pool/{poolId}/default | Performs a region setting of default storage pool. |
*PoolApi* | [**setStoragePoolDefaultTemplate**](Apis/PoolApi.md#setstoragepooldefaulttemplate) | **PATCH** /pool/{poolId}/template/{templateId}/default | Performs a storage pool setting of default template. |
*PoolApi* | [**updateStoragePool**](Apis/PoolApi.md#updatestoragepool) | **PUT** /pool/{poolId} | Updates a storage pool configuration. |
| *ProjectApi* | [**createProject**](Apis/ProjectApi.md#createproject) | **POST** /project | Creates a new project. |
*ProjectApi* | [**createProjectDnsRecord**](Apis/ProjectApi.md#createprojectdnsrecord) | **POST** /project/{projectId}/record | Creates a new DNS record. |
*ProjectApi* | [**createProjectRegionKawaii**](Apis/ProjectApi.md#createprojectregionkawaii) | **POST** /project/{projectId}/region/{regionId}/kawaii | Creates a new Kawaii. |
*ProjectApi* | [**createProjectRegionKonvey**](Apis/ProjectApi.md#createprojectregionkonvey) | **POST** /project/{projectId}/region/{regionId}/konvey | Creates a new Konvey (Kowabunga Network Load-Balancer). |
*ProjectApi* | [**createProjectRegionKylo**](Apis/ProjectApi.md#createprojectregionkylo) | **POST** /project/{projectId}/region/{regionId}/kylo | Creates a new Kylo. |
*ProjectApi* | [**createProjectRegionVolume**](Apis/ProjectApi.md#createprojectregionvolume) | **POST** /project/{projectId}/region/{regionId}/volume | Creates a new storage volume. |
*ProjectApi* | [**createProjectZoneInstance**](Apis/ProjectApi.md#createprojectzoneinstance) | **POST** /project/{projectId}/zone/{zoneId}/instance | Creates a new virtual machine instance. |
*ProjectApi* | [**createProjectZoneKompute**](Apis/ProjectApi.md#createprojectzonekompute) | **POST** /project/{projectId}/zone/{zoneId}/kompute | Creates a new Kompute. |
*ProjectApi* | [**createProjectZoneKonvey**](Apis/ProjectApi.md#createprojectzonekonvey) | **POST** /project/{projectId}/zone/{zoneId}/konvey | Creates a new Konvey (Kowabunga Network Load-Balancer). |
*ProjectApi* | [**deleteProject**](Apis/ProjectApi.md#deleteproject) | **DELETE** /project/{projectId} | Deletes an existing project. |
*ProjectApi* | [**listProjectDnsRecords**](Apis/ProjectApi.md#listprojectdnsrecords) | **GET** /project/{projectId}/records | Returns the IDs of DNS record objects. |
*ProjectApi* | [**listProjectRegionKawaiis**](Apis/ProjectApi.md#listprojectregionkawaiis) | **GET** /project/{projectId}/region/{regionId}/kawaiis | Returns the IDs of Kawaii objects. |
*ProjectApi* | [**listProjectRegionKonveys**](Apis/ProjectApi.md#listprojectregionkonveys) | **GET** /project/{projectId}/region/{regionId}/konveys | Returns the IDs of Konvey (Kowabunga Network Load-Balancer) objects. |
*ProjectApi* | [**listProjectRegionKylos**](Apis/ProjectApi.md#listprojectregionkylos) | **GET** /project/{projectId}/region/{regionId}/kylo | Returns the IDs of Kylo objects. |
*ProjectApi* | [**listProjectRegionVolumes**](Apis/ProjectApi.md#listprojectregionvolumes) | **GET** /project/{projectId}/region/{regionId}/volumes | Returns the IDs of storage volume objects. |
*ProjectApi* | [**listProjectZoneInstances**](Apis/ProjectApi.md#listprojectzoneinstances) | **GET** /project/{projectId}/zone/{zoneId}/instances | Returns the IDs of virtual machine instance objects. |
*ProjectApi* | [**listProjectZoneKomputes**](Apis/ProjectApi.md#listprojectzonekomputes) | **GET** /project/{projectId}/zone/{zoneId}/komputes | Returns the IDs of Kompute objects. |
*ProjectApi* | [**listProjectZoneKonveys**](Apis/ProjectApi.md#listprojectzonekonveys) | **GET** /project/{projectId}/zone/{zoneId}/konveys | Returns the IDs of Konvey (Kowabunga Network Load-Balancer) objects. |
*ProjectApi* | [**listProjects**](Apis/ProjectApi.md#listprojects) | **GET** /project | Returns the IDs of project objects. |
*ProjectApi* | [**readProject**](Apis/ProjectApi.md#readproject) | **GET** /project/{projectId} | Returns a project. |
*ProjectApi* | [**readProjectCost**](Apis/ProjectApi.md#readprojectcost) | **GET** /project/{projectId}/cost | Returns a resource cost. |
*ProjectApi* | [**readProjectUsage**](Apis/ProjectApi.md#readprojectusage) | **GET** /project/{projectId}/usage | Returns a global project resource quotas/usage (0 for unlimited). |
*ProjectApi* | [**updateProject**](Apis/ProjectApi.md#updateproject) | **PUT** /project/{projectId} | Updates a project configuration. |
| *RecordApi* | [**createProjectDnsRecord_0**](Apis/RecordApi.md#createprojectdnsrecord_0) | **POST** /project/{projectId}/record | Creates a new DNS record. |
*RecordApi* | [**deleteDnsRecord**](Apis/RecordApi.md#deletednsrecord) | **DELETE** /record/{recordId} | Deletes an existing DNS record. |
*RecordApi* | [**listProjectDnsRecords_0**](Apis/RecordApi.md#listprojectdnsrecords_0) | **GET** /project/{projectId}/records | Returns the IDs of DNS record objects. |
*RecordApi* | [**readDnsRecord**](Apis/RecordApi.md#readdnsrecord) | **GET** /record/{recordId} | Returns a DNS record. |
*RecordApi* | [**updateDnsRecord**](Apis/RecordApi.md#updatednsrecord) | **PUT** /record/{recordId} | Updates a DNS record configuration. |
| *RegionApi* | [**createKiwi**](Apis/RegionApi.md#createkiwi) | **POST** /region/{regionId}/kiwi | Creates a new Kiwi (Kowabunga Inner Wan Interface) provides edge-network services.. |
*RegionApi* | [**createProjectRegionKawaii_0**](Apis/RegionApi.md#createprojectregionkawaii_0) | **POST** /project/{projectId}/region/{regionId}/kawaii | Creates a new Kawaii. |
*RegionApi* | [**createProjectRegionKonvey_0**](Apis/RegionApi.md#createprojectregionkonvey_0) | **POST** /project/{projectId}/region/{regionId}/konvey | Creates a new Konvey (Kowabunga Network Load-Balancer). |
*RegionApi* | [**createProjectRegionKylo_0**](Apis/RegionApi.md#createprojectregionkylo_0) | **POST** /project/{projectId}/region/{regionId}/kylo | Creates a new Kylo. |
*RegionApi* | [**createProjectRegionVolume_0**](Apis/RegionApi.md#createprojectregionvolume_0) | **POST** /project/{projectId}/region/{regionId}/volume | Creates a new storage volume. |
*RegionApi* | [**createRegion**](Apis/RegionApi.md#createregion) | **POST** /region | Creates a new region. |
*RegionApi* | [**createStorageNFS**](Apis/RegionApi.md#createstoragenfs) | **POST** /region/{regionId}/nfs | Creates a new NFS storage. |
*RegionApi* | [**createStoragePool**](Apis/RegionApi.md#createstoragepool) | **POST** /region/{regionId}/pool | Creates a new storage pool. |
*RegionApi* | [**createVNet**](Apis/RegionApi.md#createvnet) | **POST** /region/{regionId}/vnet | Creates a new virtual network. |
*RegionApi* | [**createZone**](Apis/RegionApi.md#createzone) | **POST** /region/{regionId}/zone | Creates a new availability zone. |
*RegionApi* | [**deleteRegion**](Apis/RegionApi.md#deleteregion) | **DELETE** /region/{regionId} | Deletes an existing region. |
*RegionApi* | [**listProjectRegionKawaiis_0**](Apis/RegionApi.md#listprojectregionkawaiis_0) | **GET** /project/{projectId}/region/{regionId}/kawaiis | Returns the IDs of Kawaii objects. |
*RegionApi* | [**listProjectRegionKonveys_0**](Apis/RegionApi.md#listprojectregionkonveys_0) | **GET** /project/{projectId}/region/{regionId}/konveys | Returns the IDs of Konvey (Kowabunga Network Load-Balancer) objects. |
*RegionApi* | [**listProjectRegionKylos_0**](Apis/RegionApi.md#listprojectregionkylos_0) | **GET** /project/{projectId}/region/{regionId}/kylo | Returns the IDs of Kylo objects. |
*RegionApi* | [**listProjectRegionVolumes_0**](Apis/RegionApi.md#listprojectregionvolumes_0) | **GET** /project/{projectId}/region/{regionId}/volumes | Returns the IDs of storage volume objects. |
*RegionApi* | [**listRegionKiwis**](Apis/RegionApi.md#listregionkiwis) | **GET** /region/{regionId}/kiwis | Returns the IDs of Kiwi (Kowabunga Inner Wan Interface) provides edge-network services. objects. |
*RegionApi* | [**listRegionStorageNFSs**](Apis/RegionApi.md#listregionstoragenfss) | **GET** /region/{regionId}/nfs | Returns the IDs of NFS storage objects. |
*RegionApi* | [**listRegionStoragePools**](Apis/RegionApi.md#listregionstoragepools) | **GET** /region/{regionId}/pools | Returns the IDs of storage pool objects. |
*RegionApi* | [**listRegionVNets**](Apis/RegionApi.md#listregionvnets) | **GET** /region/{regionId}/vnets | Returns the IDs of virtual network objects. |
*RegionApi* | [**listRegionZones**](Apis/RegionApi.md#listregionzones) | **GET** /region/{regionId}/zones | Returns the IDs of availability zone objects. |
*RegionApi* | [**listRegions**](Apis/RegionApi.md#listregions) | **GET** /region | Returns the IDs of region objects. |
*RegionApi* | [**readRegion**](Apis/RegionApi.md#readregion) | **GET** /region/{regionId} | Returns a region. |
*RegionApi* | [**setRegionDefaultStorageNFS**](Apis/RegionApi.md#setregiondefaultstoragenfs) | **PATCH** /region/{regionId}/nfs/{nfsId}/default | Performs a region setting of default NFS storage. |
*RegionApi* | [**setRegionDefaultStoragePool**](Apis/RegionApi.md#setregiondefaultstoragepool) | **PATCH** /region/{regionId}/pool/{poolId}/default | Performs a region setting of default storage pool. |
*RegionApi* | [**updateRegion**](Apis/RegionApi.md#updateregion) | **PUT** /region/{regionId} | Updates a region configuration. |
| *SubnetApi* | [**createAdapter**](Apis/SubnetApi.md#createadapter) | **POST** /subnet/{subnetId}/adapter | Creates a new network adapter. |
*SubnetApi* | [**createSubnet_0**](Apis/SubnetApi.md#createsubnet_0) | **POST** /vnet/{vnetId}/subnet | Creates a new network subnet. |
*SubnetApi* | [**deleteSubnet**](Apis/SubnetApi.md#deletesubnet) | **DELETE** /subnet/{subnetId} | Deletes an existing network subnet. |
*SubnetApi* | [**listSubnetAdapters**](Apis/SubnetApi.md#listsubnetadapters) | **GET** /subnet/{subnetId}/adapters | Returns the IDs of network adapter objects. |
*SubnetApi* | [**listSubnets**](Apis/SubnetApi.md#listsubnets) | **GET** /subnet | Returns the IDs of network subnet objects. |
*SubnetApi* | [**listVNetSubnets_0**](Apis/SubnetApi.md#listvnetsubnets_0) | **GET** /vnet/{vnetId}/subnets | Returns the IDs of network subnet objects. |
*SubnetApi* | [**readSubnet**](Apis/SubnetApi.md#readsubnet) | **GET** /subnet/{subnetId} | Returns a network subnet. |
*SubnetApi* | [**setVNetDefaultSubnet_0**](Apis/SubnetApi.md#setvnetdefaultsubnet_0) | **PATCH** /vnet/{vnetId}/subnet/{subnetId}/default | Performs a virtual network setting of default network subnet. |
*SubnetApi* | [**updateSubnet**](Apis/SubnetApi.md#updatesubnet) | **PUT** /subnet/{subnetId} | Updates a network subnet configuration. |
| *TeamApi* | [**createTeam**](Apis/TeamApi.md#createteam) | **POST** /team | Creates a new Kowabunga users team. |
*TeamApi* | [**deleteTeam**](Apis/TeamApi.md#deleteteam) | **DELETE** /team/{teamId} | Deletes an existing Kowabunga users team. |
*TeamApi* | [**listTeams**](Apis/TeamApi.md#listteams) | **GET** /team | Returns the IDs of Kowabunga users team objects. |
*TeamApi* | [**readTeam**](Apis/TeamApi.md#readteam) | **GET** /team/{teamId} | Returns a Kowabunga users team. |
*TeamApi* | [**updateTeam**](Apis/TeamApi.md#updateteam) | **PUT** /team/{teamId} | Updates a Kowabunga users team configuration. |
| *TemplateApi* | [**createTemplate_0**](Apis/TemplateApi.md#createtemplate_0) | **POST** /pool/{poolId}/template | Creates a new image template. |
*TemplateApi* | [**deleteTemplate**](Apis/TemplateApi.md#deletetemplate) | **DELETE** /template/{templateId} | Deletes an existing image template. |
*TemplateApi* | [**listStoragePoolTemplates_0**](Apis/TemplateApi.md#liststoragepooltemplates_0) | **GET** /pool/{poolId}/templates | Returns the IDs of image template objects. |
*TemplateApi* | [**listTemplates**](Apis/TemplateApi.md#listtemplates) | **GET** /template | Returns the IDs of image template objects. |
*TemplateApi* | [**readTemplate**](Apis/TemplateApi.md#readtemplate) | **GET** /template/{templateId} | Returns a image template. |
*TemplateApi* | [**setStoragePoolDefaultTemplate_0**](Apis/TemplateApi.md#setstoragepooldefaulttemplate_0) | **PATCH** /pool/{poolId}/template/{templateId}/default | Performs a storage pool setting of default template. |
*TemplateApi* | [**updateTemplate**](Apis/TemplateApi.md#updatetemplate) | **PUT** /template/{templateId} | Updates a image template configuration. |
| *TokenApi* | [**deleteApiToken**](Apis/TokenApi.md#deleteapitoken) | **DELETE** /token/{tokenId} | Deletes an existing server-to-server authentication security token. |
*TokenApi* | [**listApiTokens**](Apis/TokenApi.md#listapitokens) | **GET** /token | Returns the IDs of server-to-server authentication security token objects. |
*TokenApi* | [**readApiToken**](Apis/TokenApi.md#readapitoken) | **GET** /token/{tokenId} | Returns a server-to-server authentication security token. |
*TokenApi* | [**setAgentApiToken_0**](Apis/TokenApi.md#setagentapitoken_0) | **PATCH** /agent/{agentId}/token | Performs a Kowabunga remote agent setting of API token (will replace any existing one). |
*TokenApi* | [**setUserApiToken_0**](Apis/TokenApi.md#setuserapitoken_0) | **PATCH** /user/{userId}/token | Performs a Kowabunga user setting of API token (will replace any existing one). |
*TokenApi* | [**updateApiToken**](Apis/TokenApi.md#updateapitoken) | **PUT** /token/{tokenId} | Updates a server-to-server authentication security token configuration. |
| *UserApi* | [**createUser**](Apis/UserApi.md#createuser) | **POST** /user | Creates a new Kowabunga user. |
*UserApi* | [**deleteUser**](Apis/UserApi.md#deleteuser) | **DELETE** /user/{userId} | Deletes an existing Kowabunga user. |
*UserApi* | [**listUsers**](Apis/UserApi.md#listusers) | **GET** /user | Returns the IDs of Kowabunga user objects. |
*UserApi* | [**login**](Apis/UserApi.md#login) | **POST** /login | Creates a new Kowabunga user login credentials. |
*UserApi* | [**logout**](Apis/UserApi.md#logout) | **POST** /logout | Creates a new Kowabunga user logout session. |
*UserApi* | [**readUser**](Apis/UserApi.md#readuser) | **GET** /user/{userId} | Returns a Kowabunga user. |
*UserApi* | [**resetPassword**](Apis/UserApi.md#resetpassword) | **PUT** /resetPassword | Updates a Kowabunga userreset of password for the provided email (server-side generated, will replace any existing one) configuration. |
*UserApi* | [**resetUserPassword**](Apis/UserApi.md#resetuserpassword) | **PATCH** /user/{userId}/resetPassword | Performs a Kowabunga user reset of password (server-side generated, will replace any existing one). |
*UserApi* | [**setUserApiToken**](Apis/UserApi.md#setuserapitoken) | **PATCH** /user/{userId}/token | Performs a Kowabunga user setting of API token (will replace any existing one). |
*UserApi* | [**setUserPassword**](Apis/UserApi.md#setuserpassword) | **PUT** /user/{userId}/password | Updates a Kowabunga user password (will replace any existing one) configuration. |
*UserApi* | [**updateUser**](Apis/UserApi.md#updateuser) | **PUT** /user/{userId} | Updates a Kowabunga user configuration. |
| *VnetApi* | [**createSubnet**](Apis/VnetApi.md#createsubnet) | **POST** /vnet/{vnetId}/subnet | Creates a new network subnet. |
*VnetApi* | [**createVNet_0**](Apis/VnetApi.md#createvnet_0) | **POST** /region/{regionId}/vnet | Creates a new virtual network. |
*VnetApi* | [**deleteVNet**](Apis/VnetApi.md#deletevnet) | **DELETE** /vnet/{vnetId} | Deletes an existing virtual network. |
*VnetApi* | [**listRegionVNets_0**](Apis/VnetApi.md#listregionvnets_0) | **GET** /region/{regionId}/vnets | Returns the IDs of virtual network objects. |
*VnetApi* | [**listVNetSubnets**](Apis/VnetApi.md#listvnetsubnets) | **GET** /vnet/{vnetId}/subnets | Returns the IDs of network subnet objects. |
*VnetApi* | [**listVNets**](Apis/VnetApi.md#listvnets) | **GET** /vnet | Returns the IDs of virtual network objects. |
*VnetApi* | [**readVNet**](Apis/VnetApi.md#readvnet) | **GET** /vnet/{vnetId} | Returns a virtual network. |
*VnetApi* | [**setVNetDefaultSubnet**](Apis/VnetApi.md#setvnetdefaultsubnet) | **PATCH** /vnet/{vnetId}/subnet/{subnetId}/default | Performs a virtual network setting of default network subnet. |
*VnetApi* | [**updateVNet**](Apis/VnetApi.md#updatevnet) | **PUT** /vnet/{vnetId} | Updates a virtual network configuration. |
| *VolumeApi* | [**createProjectRegionVolume_1**](Apis/VolumeApi.md#createprojectregionvolume_1) | **POST** /project/{projectId}/region/{regionId}/volume | Creates a new storage volume. |
*VolumeApi* | [**deleteVolume**](Apis/VolumeApi.md#deletevolume) | **DELETE** /volume/{volumeId} | Deletes an existing storage volume. |
*VolumeApi* | [**listProjectRegionVolumes_1**](Apis/VolumeApi.md#listprojectregionvolumes_1) | **GET** /project/{projectId}/region/{regionId}/volumes | Returns the IDs of storage volume objects. |
*VolumeApi* | [**listStoragePoolVolumes_0**](Apis/VolumeApi.md#liststoragepoolvolumes_0) | **GET** /pool/{poolId}/volumes | Returns the IDs of storage volume objects. |
*VolumeApi* | [**listVolumes**](Apis/VolumeApi.md#listvolumes) | **GET** /volume | Returns the IDs of storage volume objects. |
*VolumeApi* | [**readVolume**](Apis/VolumeApi.md#readvolume) | **GET** /volume/{volumeId} | Returns a storage volume. |
*VolumeApi* | [**updateVolume**](Apis/VolumeApi.md#updatevolume) | **PUT** /volume/{volumeId} | Updates a storage volume configuration. |
| *ZoneApi* | [**createKaktus**](Apis/ZoneApi.md#createkaktus) | **POST** /zone/{zoneId}/kaktus | Creates a new Kaktus computing node. |
*ZoneApi* | [**createProjectZoneInstance_0**](Apis/ZoneApi.md#createprojectzoneinstance_0) | **POST** /project/{projectId}/zone/{zoneId}/instance | Creates a new virtual machine instance. |
*ZoneApi* | [**createProjectZoneKompute_0**](Apis/ZoneApi.md#createprojectzonekompute_0) | **POST** /project/{projectId}/zone/{zoneId}/kompute | Creates a new Kompute. |
*ZoneApi* | [**createProjectZoneKonvey_0**](Apis/ZoneApi.md#createprojectzonekonvey_0) | **POST** /project/{projectId}/zone/{zoneId}/konvey | Creates a new Konvey (Kowabunga Network Load-Balancer). |
*ZoneApi* | [**createZone_0**](Apis/ZoneApi.md#createzone_0) | **POST** /region/{regionId}/zone | Creates a new availability zone. |
*ZoneApi* | [**deleteZone**](Apis/ZoneApi.md#deletezone) | **DELETE** /zone/{zoneId} | Deletes an existing availability zone. |
*ZoneApi* | [**listProjectZoneInstances_0**](Apis/ZoneApi.md#listprojectzoneinstances_0) | **GET** /project/{projectId}/zone/{zoneId}/instances | Returns the IDs of virtual machine instance objects. |
*ZoneApi* | [**listProjectZoneKomputes_0**](Apis/ZoneApi.md#listprojectzonekomputes_0) | **GET** /project/{projectId}/zone/{zoneId}/komputes | Returns the IDs of Kompute objects. |
*ZoneApi* | [**listProjectZoneKonveys_0**](Apis/ZoneApi.md#listprojectzonekonveys_0) | **GET** /project/{projectId}/zone/{zoneId}/konveys | Returns the IDs of Konvey (Kowabunga Network Load-Balancer) objects. |
*ZoneApi* | [**listRegionZones_0**](Apis/ZoneApi.md#listregionzones_0) | **GET** /region/{regionId}/zones | Returns the IDs of availability zone objects. |
*ZoneApi* | [**listZoneKaktuses**](Apis/ZoneApi.md#listzonekaktuses) | **GET** /zone/{zoneId}/kaktuses | Returns the IDs of Kaktus computing node objects. |
*ZoneApi* | [**listZones**](Apis/ZoneApi.md#listzones) | **GET** /zone | Returns the IDs of availability zone objects. |
*ZoneApi* | [**readZone**](Apis/ZoneApi.md#readzone) | **GET** /zone/{zoneId} | Returns a availability zone. |
*ZoneApi* | [**updateZone**](Apis/ZoneApi.md#updatezone) | **PUT** /zone/{zoneId} | Updates a availability zone configuration. |


<a name="documentation-for-models"></a>
## Documentation for Models

 - [Adapter](./Models/Adapter.md)
 - [Agent](./Models/Agent.md)
 - [ApiErrorBadRequest](./Models/ApiErrorBadRequest.md)
 - [ApiErrorConflict](./Models/ApiErrorConflict.md)
 - [ApiErrorForbidden](./Models/ApiErrorForbidden.md)
 - [ApiErrorInsufficientStorage](./Models/ApiErrorInsufficientStorage.md)
 - [ApiErrorNotFound](./Models/ApiErrorNotFound.md)
 - [ApiErrorUnauthorized](./Models/ApiErrorUnauthorized.md)
 - [ApiErrorUnprocessableEntity](./Models/ApiErrorUnprocessableEntity.md)
 - [ApiToken](./Models/ApiToken.md)
 - [Cost](./Models/Cost.md)
 - [DnsRecord](./Models/DnsRecord.md)
 - [Instance](./Models/Instance.md)
 - [InstanceRemoteAccess](./Models/InstanceRemoteAccess.md)
 - [InstanceState](./Models/InstanceState.md)
 - [IpRange](./Models/IpRange.md)
 - [Kaktus](./Models/Kaktus.md)
 - [KaktusCPU](./Models/KaktusCPU.md)
 - [KaktusCaps](./Models/KaktusCaps.md)
 - [Kawaii](./Models/Kawaii.md)
 - [KawaiiDNatRule](./Models/KawaiiDNatRule.md)
 - [KawaiiFirewall](./Models/KawaiiFirewall.md)
 - [KawaiiFirewallEgressRule](./Models/KawaiiFirewallEgressRule.md)
 - [KawaiiFirewallIngressRule](./Models/KawaiiFirewallIngressRule.md)
 - [KawaiiNetIp](./Models/KawaiiNetIp.md)
 - [KawaiiNetIpZone](./Models/KawaiiNetIpZone.md)
 - [KawaiiVpcForwardRule](./Models/KawaiiVpcForwardRule.md)
 - [KawaiiVpcNetIpZone](./Models/KawaiiVpcNetIpZone.md)
 - [KawaiiVpcPeering](./Models/KawaiiVpcPeering.md)
 - [Kiwi](./Models/Kiwi.md)
 - [Kompute](./Models/Kompute.md)
 - [Konvey](./Models/Konvey.md)
 - [KonveyBackends](./Models/KonveyBackends.md)
 - [KonveyEndpoint](./Models/KonveyEndpoint.md)
 - [Kylo](./Models/Kylo.md)
 - [Metadata](./Models/Metadata.md)
 - [Password](./Models/Password.md)
 - [Project](./Models/Project.md)
 - [ProjectResources](./Models/ProjectResources.md)
 - [Region](./Models/Region.md)
 - [RegionSubnet](./Models/RegionSubnet.md)
 - [StorageNFS](./Models/StorageNFS.md)
 - [StoragePool](./Models/StoragePool.md)
 - [Subnet](./Models/Subnet.md)
 - [Team](./Models/Team.md)
 - [Template](./Models/Template.md)
 - [User](./Models/User.md)
 - [UserCredentials](./Models/UserCredentials.md)
 - [UserEmail](./Models/UserEmail.md)
 - [VNet](./Models/VNet.md)
 - [Volume](./Models/Volume.md)
 - [Zone](./Models/Zone.md)


<a name="documentation-for-authorization"></a>
## Documentation for Authorization

<a name="ApiKeyAuth"></a>
### ApiKeyAuth

- **Type**: API key
- **API key parameter name**: X-API-Key
- **Location**: HTTP header

<a name="BearerAuth"></a>
### BearerAuth

- **Type**: HTTP Bearer Token authentication (JWT)

