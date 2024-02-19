# Documentation for Kowabunga API documentation

<a name="documentation-for-api-endpoints"></a>
## Documentation for API Endpoints

All URIs are relative to */api/v1*

| Class | Method | HTTP request | Description |
|------------ | ------------- | ------------- | -------------|
| *AdapterApi* | [**createAdapter**](Apis/AdapterApi.md#createadapter) | **POST** /subnet/{subnetId}/adapter | Creates a new network adapter. |
*AdapterApi* | [**deleteAdapter**](Apis/AdapterApi.md#deleteadapter) | **DELETE** /adapter/{ adapterId } | Deletes an existing network adapter. |
*AdapterApi* | [**getSubnetAdapters**](Apis/AdapterApi.md#getsubnetadapters) | **GET** /subnet/{subnetId}/adapters | Returns the IDs of the network adapters existing in the subnet. |
*AdapterApi* | [**listAdapters**](Apis/AdapterApi.md#listadapters) | **GET** /adapter | Returns the IDs of network adapter objects. |
*AdapterApi* | [**readAdapter**](Apis/AdapterApi.md#readadapter) | **GET** /adapter/{ adapterId } | Returns a network adapter. |
*AdapterApi* | [**updateAdapter**](Apis/AdapterApi.md#updateadapter) | **PUT** /adapter/{ adapterId } | Updates a network adapter configuration. |
| *HostApi* | [**createHost**](Apis/HostApi.md#createhost) | **POST** /zone/{zoneId}/host | Creates a new host. |
*HostApi* | [**deleteHost**](Apis/HostApi.md#deletehost) | **DELETE** /host/{ hostId } | Deletes an existing computing host. |
*HostApi* | [**getZoneHosts**](Apis/HostApi.md#getzonehosts) | **GET** /zone/{zoneId}/hosts | Returns the IDs of the hosts existing in the zone. |
*HostApi* | [**listHostInstances**](Apis/HostApi.md#listhostinstances) | **GET** /host/{hostId}/instances | Returns the IDs of virtual machine instance objects. |
*HostApi* | [**listHosts**](Apis/HostApi.md#listhosts) | **GET** /host | Returns the IDs of computing host objects. |
*HostApi* | [**readHost**](Apis/HostApi.md#readhost) | **GET** /host/{ hostId } | Returns a computing host. |
*HostApi* | [**readHostCaps**](Apis/HostApi.md#readhostcaps) | **GET** /host/{hostId}/caps | Returns a computing host capabilities. |
*HostApi* | [**updateHost**](Apis/HostApi.md#updatehost) | **PUT** /host/{ hostId } | Updates a computing host configuration. |
| *InstanceApi* | [**createProjectZoneInstance**](Apis/InstanceApi.md#createprojectzoneinstance) | **POST** /project/{projectId}/zone/{zoneId}/instance | Creates a new virtual machine instance in specified zone. |
*InstanceApi* | [**deleteInstance**](Apis/InstanceApi.md#deleteinstance) | **DELETE** /instance/{instanceId} | Deletes an existing virtual machine instance. |
*InstanceApi* | [**getAllInstances**](Apis/InstanceApi.md#getallinstances) | **GET** /instance | Returns the IDs of registered virtual machines. |
*InstanceApi* | [**getInstance**](Apis/InstanceApi.md#getinstance) | **GET** /instance/{instanceId} | Returns the description of the virtual machine. |
*InstanceApi* | [**getInstanceRemoteConnection**](Apis/InstanceApi.md#getinstanceremoteconnection) | **GET** /instance/{instanceId}/connect | Returns virtual machine remote access URL. |
*InstanceApi* | [**getInstanceState**](Apis/InstanceApi.md#getinstancestate) | **GET** /instance/{instanceId}/state | Returns the state of the virtual machine. |
*InstanceApi* | [**getProjectZoneInstances**](Apis/InstanceApi.md#getprojectzoneinstances) | **GET** /project/{projectId}/zone/{zoneId}/instances | Returns the IDs of the virtual machine instances existing in the project in the specified zone. |
*InstanceApi* | [**listHostInstances**](Apis/InstanceApi.md#listhostinstances) | **GET** /host/{hostId}/instances | Returns the IDs of virtual machine instance objects. |
*InstanceApi* | [**rebootInstance**](Apis/InstanceApi.md#rebootinstance) | **POST** /instance/{instanceId}/reboot | Perform a virtual machine software reboot. |
*InstanceApi* | [**resetInstance**](Apis/InstanceApi.md#resetinstance) | **POST** /instance/{instanceId}/reset | Perform a virtual machine hardware reset. |
*InstanceApi* | [**resumeInstance**](Apis/InstanceApi.md#resumeinstance) | **POST** /instance/{instanceId}/resume | Perform a virtual machine software PM resume. |
*InstanceApi* | [**shutdownInstance**](Apis/InstanceApi.md#shutdowninstance) | **POST** /instance/{instanceId}/shutdown | Initiate a software shutdown of a virtual machine. |
*InstanceApi* | [**startInstance**](Apis/InstanceApi.md#startinstance) | **POST** /instance/{instanceId}/start | Boot up a virtual machine. |
*InstanceApi* | [**stopInstance**](Apis/InstanceApi.md#stopinstance) | **POST** /instance/{instanceId}/stop | Initiate a hardware stop of a virtual machine. |
*InstanceApi* | [**suspendInstance**](Apis/InstanceApi.md#suspendinstance) | **POST** /instance/{instanceId}/suspend | Perform a virtual machine software PM suspend. |
*InstanceApi* | [**updateInstance**](Apis/InstanceApi.md#updateinstance) | **PUT** /instance/{instanceId} | Updates a virtual machine configuration. |
| *KceApi* | [**createProjectZoneKce**](Apis/KceApi.md#createprojectzonekce) | **POST** /project/{projectId}/zone/{zoneId}/kce | Creates a new KCE virtual machine in specified zone. |
*KceApi* | [**deleteKCE**](Apis/KceApi.md#deletekce) | **DELETE** /kce/{kceId} | Deletes an existing KCE virtual machine. |
*KceApi* | [**getAllKCEs**](Apis/KceApi.md#getallkces) | **GET** /kce | Returns the IDs of registered KCE virtual machines. |
*KceApi* | [**getKCE**](Apis/KceApi.md#getkce) | **GET** /kce/{kceId} | Returns the description of the KCE virtual machine. |
*KceApi* | [**getKCEState**](Apis/KceApi.md#getkcestate) | **GET** /kce/{kceId}/state | Returns the state of the KCE virtual machine. |
*KceApi* | [**getProjectZoneKCEs**](Apis/KceApi.md#getprojectzonekces) | **GET** /project/{projectId}/zone/{zoneId}/kces | Returns the IDs of the KCE virtual machines existing in the project in the specified zone. |
*KceApi* | [**rebootKCE**](Apis/KceApi.md#rebootkce) | **POST** /kce/{kceId}/reboot | Perform a KCE virtual machine software reboot. |
*KceApi* | [**resetKCE**](Apis/KceApi.md#resetkce) | **POST** /kce/{kceId}/reset | Perform a KCE virtual machine hardware reset. |
*KceApi* | [**resumeKCE**](Apis/KceApi.md#resumekce) | **POST** /kce/{kceId}/resume | Perform a KCE virtual machine software PM resume. |
*KceApi* | [**shutdownKCE**](Apis/KceApi.md#shutdownkce) | **POST** /kce/{kceId}/shutdown | Initiate a software shutdown of a KCE virtual machine. |
*KceApi* | [**startKCE**](Apis/KceApi.md#startkce) | **POST** /kce/{kceId}/start | Boot up a KCE virtual machine. |
*KceApi* | [**stopKCE**](Apis/KceApi.md#stopkce) | **POST** /kce/{kceId}/stop | Initiate a hardware stop of a KCE virtual machine. |
*KceApi* | [**suspendKCE**](Apis/KceApi.md#suspendkce) | **POST** /kce/{kceId}/suspend | Perform a KCE virtual machine software PM suspend. |
*KceApi* | [**updateKCE**](Apis/KceApi.md#updatekce) | **PUT** /kce/{kceId} | Updates a KCE virtual machine configuration. |
| *KfsApi* | [**createProjectZoneKfs**](Apis/KfsApi.md#createprojectzonekfs) | **POST** /project/{projectId}/zone/{zoneId}/kfs | Creates a new KFS storage volume in specified zone. |
*KfsApi* | [**deleteKFS**](Apis/KfsApi.md#deletekfs) | **DELETE** /kfs/{kfsId} | Deletes an existing KFS storage volume. |
*KfsApi* | [**getAllKFSs**](Apis/KfsApi.md#getallkfss) | **GET** /kfs | Returns the IDs of registered KFS storage volumes. |
*KfsApi* | [**getKFS**](Apis/KfsApi.md#getkfs) | **GET** /kfs/{kfsId} | Returns the description of the KFS storage volume. |
*KfsApi* | [**getNfsKfs**](Apis/KfsApi.md#getnfskfs) | **GET** /nfs/{nfsId}/kfs | Returns the IDs of the KFS volumes existing in the NFS storage. |
*KfsApi* | [**getProjectZoneKfs**](Apis/KfsApi.md#getprojectzonekfs) | **GET** /project/{projectId}/zone/{zoneId}/kfs | Returns the IDs of the KFS storage volumes existing in the project in the specified zone. |
*KfsApi* | [**updateKFS**](Apis/KfsApi.md#updatekfs) | **PUT** /kfs/{kfsId} | Updates a KFS storage volume configuration. |
| *KgwApi* | [**createProjectZoneKgw**](Apis/KgwApi.md#createprojectzonekgw) | **POST** /project/{projectId}/zone/{zoneId}/kgw | Creates a new KGW in the specified zone. |
*KgwApi* | [**deleteKGW**](Apis/KgwApi.md#deletekgw) | **DELETE** /kgw/{kgwId} | Deletes an existing KGW gateway. |
*KgwApi* | [**getAllKgw**](Apis/KgwApi.md#getallkgw) | **GET** /kgw | Returns the IDs of registered KGW |
*KgwApi* | [**getKgw**](Apis/KgwApi.md#getkgw) | **GET** /kgw/{kgwId} | Returns the descirption of the registered KGW |
*KgwApi* | [**getProjectZoneKGWs**](Apis/KgwApi.md#getprojectzonekgws) | **GET** /project/{projectId}/zone/{zoneId}/kgws | Returns the IDs of the KGW existing in the project in the specified zone. |
*KgwApi* | [**updateKGW**](Apis/KgwApi.md#updatekgw) | **PUT** /kgw/{kgwId} | Updates a KGW virtual machine configuration. |
| *NetgwApi* | [**createNetGW**](Apis/NetgwApi.md#createnetgw) | **POST** /zone/{zoneId}/netgw | Creates a new network gateway. |
*NetgwApi* | [**deleteNetGW**](Apis/NetgwApi.md#deletenetgw) | **DELETE** /netgw/{netgwId} | Deletes an existing network gateway. |
*NetgwApi* | [**getAllNetGWs**](Apis/NetgwApi.md#getallnetgws) | **GET** /netgw | Returns the IDs of network gateways. |
*NetgwApi* | [**getNetGW**](Apis/NetgwApi.md#getnetgw) | **GET** /netgw/{netgwId} | Returns a description of the network gateway |
*NetgwApi* | [**getZoneNetGWs**](Apis/NetgwApi.md#getzonenetgws) | **GET** /zone/{zoneId}/netgws | Returns the IDs of the hosts existing in the zone. |
*NetgwApi* | [**updateNetGW**](Apis/NetgwApi.md#updatenetgw) | **PUT** /netgw/{netgwId} | Updates a network gateway configuration. |
| *NfsApi* | [**createNfsStorage**](Apis/NfsApi.md#createnfsstorage) | **POST** /zone/{zoneId}/nfs | Creates a new NFS storage. |
*NfsApi* | [**deleteNfsStorage**](Apis/NfsApi.md#deletenfsstorage) | **DELETE** /nfs/{nfsId} | Deletes an existing NFS storage. |
*NfsApi* | [**getAllNfsStorages**](Apis/NfsApi.md#getallnfsstorages) | **GET** /nfs | Returns the IDs of registered NFS storages. |
*NfsApi* | [**getNfsKfs**](Apis/NfsApi.md#getnfskfs) | **GET** /nfs/{nfsId}/kfs | Returns the IDs of the KFS volumes existing in the NFS storage. |
*NfsApi* | [**getNfsStorage**](Apis/NfsApi.md#getnfsstorage) | **GET** /nfs/{nfsId} | Returns a description of the NFS storage. |
*NfsApi* | [**getZoneNfsStorages**](Apis/NfsApi.md#getzonenfsstorages) | **GET** /zone/{zoneId}/nfs | Returns the IDs of the NFS storages existing in the zone. |
*NfsApi* | [**updateNfsStorage**](Apis/NfsApi.md#updatenfsstorage) | **PUT** /nfs/{nfsId} | Updates an NFS storage configuration. |
*NfsApi* | [**updateZoneDefaultNfsStorage**](Apis/NfsApi.md#updatezonedefaultnfsstorage) | **PUT** /zone/{zoneId}/nfs/{nfsId}/default | Set a zone's default NFS storage. |
| *PoolApi* | [**createPool**](Apis/PoolApi.md#createpool) | **POST** /zone/{zoneId}/pool | Creates a new storage pool. |
*PoolApi* | [**createTemplate**](Apis/PoolApi.md#createtemplate) | **POST** /pool/{poolId}/template | Creates a new volume template. |
*PoolApi* | [**deletePool**](Apis/PoolApi.md#deletepool) | **DELETE** /pool/{poolId} | Deletes an existing pool. |
*PoolApi* | [**getAllPools**](Apis/PoolApi.md#getallpools) | **GET** /pool | Returns the IDs of registered pools. |
*PoolApi* | [**getPool**](Apis/PoolApi.md#getpool) | **GET** /pool/{poolId} | Returns a description of the pool |
*PoolApi* | [**getPoolTemplates**](Apis/PoolApi.md#getpooltemplates) | **GET** /pool/{poolId}/templates | Returns the IDs of the volume templates existing in the storage pool. |
*PoolApi* | [**getPoolVolumes**](Apis/PoolApi.md#getpoolvolumes) | **GET** /pool/{poolId}/volumes | Returns the IDs of the storage volumes existing in the pool. |
*PoolApi* | [**getZonePools**](Apis/PoolApi.md#getzonepools) | **GET** /zone/{zoneId}/pools | Returns the IDs of the pools existing in the zone. |
*PoolApi* | [**updatePool**](Apis/PoolApi.md#updatepool) | **PUT** /pool/{poolId} | Updates a pool configuration. |
*PoolApi* | [**updatePoolDefaultTemplate**](Apis/PoolApi.md#updatepooldefaulttemplate) | **PUT** /pool/{poolId}/template/{templateId}/default | Set a storage pool default volume template. |
*PoolApi* | [**updateZoneDefaultPool**](Apis/PoolApi.md#updatezonedefaultpool) | **PUT** /zone/{zoneId}/pool/{poolId}/default | Set a zone's default storage pool. |
| *ProjectApi* | [**createProject**](Apis/ProjectApi.md#createproject) | **POST** /project | Creates a new project. |
*ProjectApi* | [**createProjectDnsRecord**](Apis/ProjectApi.md#createprojectdnsrecord) | **POST** /project/{projectId}/record | Creates a new DNS record in specified project. |
*ProjectApi* | [**createProjectZoneInstance**](Apis/ProjectApi.md#createprojectzoneinstance) | **POST** /project/{projectId}/zone/{zoneId}/instance | Creates a new virtual machine instance in specified zone. |
*ProjectApi* | [**createProjectZoneKce**](Apis/ProjectApi.md#createprojectzonekce) | **POST** /project/{projectId}/zone/{zoneId}/kce | Creates a new KCE virtual machine in specified zone. |
*ProjectApi* | [**createProjectZoneKfs**](Apis/ProjectApi.md#createprojectzonekfs) | **POST** /project/{projectId}/zone/{zoneId}/kfs | Creates a new KFS storage volume in specified zone. |
*ProjectApi* | [**createProjectZoneKgw**](Apis/ProjectApi.md#createprojectzonekgw) | **POST** /project/{projectId}/zone/{zoneId}/kgw | Creates a new KGW in the specified zone. |
*ProjectApi* | [**createProjectZoneVolume**](Apis/ProjectApi.md#createprojectzonevolume) | **POST** /project/{projectId}/zone/{zoneId}/volume | Creates a new storage volume in specified zone. |
*ProjectApi* | [**deleteProject**](Apis/ProjectApi.md#deleteproject) | **DELETE** /project/{projectId} | Deletes an existing project. |
*ProjectApi* | [**getAllProjects**](Apis/ProjectApi.md#getallprojects) | **GET** /project | Returns the IDs of registered projects. |
*ProjectApi* | [**getProject**](Apis/ProjectApi.md#getproject) | **GET** /project/{projectId} | Returns a description of the project |
*ProjectApi* | [**getProjectCost**](Apis/ProjectApi.md#getprojectcost) | **GET** /project/{projectId}/cost | Returns the current cost for the project. |
*ProjectApi* | [**getProjectDnsRecords**](Apis/ProjectApi.md#getprojectdnsrecords) | **GET** /project/{projectId}/records | Returns the IDs of the DNS records existing in the project. |
*ProjectApi* | [**getProjectUsage**](Apis/ProjectApi.md#getprojectusage) | **GET** /project/{projectId}/usage | Returns the current resources usage for the project. |
*ProjectApi* | [**getProjectZoneInstances**](Apis/ProjectApi.md#getprojectzoneinstances) | **GET** /project/{projectId}/zone/{zoneId}/instances | Returns the IDs of the virtual machine instances existing in the project in the specified zone. |
*ProjectApi* | [**getProjectZoneKCEs**](Apis/ProjectApi.md#getprojectzonekces) | **GET** /project/{projectId}/zone/{zoneId}/kces | Returns the IDs of the KCE virtual machines existing in the project in the specified zone. |
*ProjectApi* | [**getProjectZoneKGWs**](Apis/ProjectApi.md#getprojectzonekgws) | **GET** /project/{projectId}/zone/{zoneId}/kgws | Returns the IDs of the KGW existing in the project in the specified zone. |
*ProjectApi* | [**getProjectZoneKfs**](Apis/ProjectApi.md#getprojectzonekfs) | **GET** /project/{projectId}/zone/{zoneId}/kfs | Returns the IDs of the KFS storage volumes existing in the project in the specified zone. |
*ProjectApi* | [**getProjectZoneVolumes**](Apis/ProjectApi.md#getprojectzonevolumes) | **GET** /project/{projectId}/zone/{zoneId}/volumes | Returns the IDs of the storage volumes existing in the project in the specified zone. |
*ProjectApi* | [**updateProject**](Apis/ProjectApi.md#updateproject) | **PUT** /project/{projectId} | Updates a project configuration. |
| *RecordApi* | [**createProjectDnsRecord**](Apis/RecordApi.md#createprojectdnsrecord) | **POST** /project/{projectId}/record | Creates a new DNS record in specified project. |
*RecordApi* | [**deleteDnsRecord**](Apis/RecordApi.md#deletednsrecord) | **DELETE** /record/{ recordId } | Deletes an existing DNS record. |
*RecordApi* | [**getProjectDnsRecords**](Apis/RecordApi.md#getprojectdnsrecords) | **GET** /project/{projectId}/records | Returns the IDs of the DNS records existing in the project. |
*RecordApi* | [**readDnsRecord**](Apis/RecordApi.md#readdnsrecord) | **GET** /record/{ recordId } | Returns a DNS record. |
*RecordApi* | [**updateDnsRecord**](Apis/RecordApi.md#updatednsrecord) | **PUT** /record/{ recordId } | Updates a DNS record configuration. |
| *RegionApi* | [**createRegion**](Apis/RegionApi.md#createregion) | **POST** /region | Creates a new region. |
*RegionApi* | [**createZone**](Apis/RegionApi.md#createzone) | **POST** /region/{regionId}/zone | Creates a new zone. |
*RegionApi* | [**deleteRegion**](Apis/RegionApi.md#deleteregion) | **DELETE** /region/{regionId} | Deletes an existing region. |
*RegionApi* | [**getAllRegions**](Apis/RegionApi.md#getallregions) | **GET** /region | Returns the IDs of registered regions. |
*RegionApi* | [**getRegion**](Apis/RegionApi.md#getregion) | **GET** /region/{regionId} | Returns a description of the region |
*RegionApi* | [**getRegionZones**](Apis/RegionApi.md#getregionzones) | **GET** /region/{regionId}/zones | Returns the IDs of the availability zones existing in the region. |
*RegionApi* | [**updateRegion**](Apis/RegionApi.md#updateregion) | **PUT** /region/{regionId} | Updates a region configuration. |
| *SubnetApi* | [**createAdapter**](Apis/SubnetApi.md#createadapter) | **POST** /subnet/{subnetId}/adapter | Creates a new network adapter. |
*SubnetApi* | [**createSubnet**](Apis/SubnetApi.md#createsubnet) | **POST** /vnet/{vnetId}/subnet | Creates a new subnet. |
*SubnetApi* | [**deleteSubnet**](Apis/SubnetApi.md#deletesubnet) | **DELETE** /subnet/{subnetId} | Deletes an existing subnet. |
*SubnetApi* | [**getAllSubnets**](Apis/SubnetApi.md#getallsubnets) | **GET** /subnet | Returns the IDs of subnets. |
*SubnetApi* | [**getSubnet**](Apis/SubnetApi.md#getsubnet) | **GET** /subnet/{subnetId} | Returns a description of the subnet. |
*SubnetApi* | [**getSubnetAdapters**](Apis/SubnetApi.md#getsubnetadapters) | **GET** /subnet/{subnetId}/adapters | Returns the IDs of the network adapters existing in the subnet. |
*SubnetApi* | [**getVNetSubnets**](Apis/SubnetApi.md#getvnetsubnets) | **GET** /vnet/{vnetId}/subnets | Returns the IDs of the subnets existing in the virtual network. |
*SubnetApi* | [**updateSubnet**](Apis/SubnetApi.md#updatesubnet) | **PUT** /subnet/{subnetId} | Updates a subnet configuration. |
*SubnetApi* | [**updateVNetDefaultSubnet**](Apis/SubnetApi.md#updatevnetdefaultsubnet) | **PUT** /vnet/{vnetId}/subnet/{subnetId}/default | Set a virtual network default subnet. |
| *TemplateApi* | [**createTemplate**](Apis/TemplateApi.md#createtemplate) | **POST** /pool/{poolId}/template | Creates a new volume template. |
*TemplateApi* | [**deleteTemplate**](Apis/TemplateApi.md#deletetemplate) | **DELETE** /template/{templateId} | Deletes an existing volume template. |
*TemplateApi* | [**getAllTemplates**](Apis/TemplateApi.md#getalltemplates) | **GET** /template | Returns the IDs of volume templates. |
*TemplateApi* | [**getTemplate**](Apis/TemplateApi.md#gettemplate) | **GET** /template/{templateId} | Returns a description of the volume template. |
*TemplateApi* | [**updatePoolDefaultTemplate**](Apis/TemplateApi.md#updatepooldefaulttemplate) | **PUT** /pool/{poolId}/template/{templateId}/default | Set a storage pool default volume template. |
*TemplateApi* | [**updateTemplate**](Apis/TemplateApi.md#updatetemplate) | **PUT** /template/{templateId} | Updates a volume template configuration. |
| *TemplatesApi* | [**getPoolTemplates**](Apis/TemplatesApi.md#getpooltemplates) | **GET** /pool/{poolId}/templates | Returns the IDs of the volume templates existing in the storage pool. |
| *VnetApi* | [**createSubnet**](Apis/VnetApi.md#createsubnet) | **POST** /vnet/{vnetId}/subnet | Creates a new subnet. |
*VnetApi* | [**createVNet**](Apis/VnetApi.md#createvnet) | **POST** /zone/{zoneId}/vnet | Creates a new virtual network. |
*VnetApi* | [**deleteVNet**](Apis/VnetApi.md#deletevnet) | **DELETE** /vnet/{vnetId} | Deletes an existing virtual network. |
*VnetApi* | [**getAllVNets**](Apis/VnetApi.md#getallvnets) | **GET** /vnet | Returns the IDs of virtual networks. |
*VnetApi* | [**getVNet**](Apis/VnetApi.md#getvnet) | **GET** /vnet/{vnetId} | Returns a description of the virtual network |
*VnetApi* | [**getVNetSubnets**](Apis/VnetApi.md#getvnetsubnets) | **GET** /vnet/{vnetId}/subnets | Returns the IDs of the subnets existing in the virtual network. |
*VnetApi* | [**getZoneVNets**](Apis/VnetApi.md#getzonevnets) | **GET** /zone/{zoneId}/vnets | Returns the IDs of the virtual networks existing in the zone. |
*VnetApi* | [**updateVNet**](Apis/VnetApi.md#updatevnet) | **PUT** /vnet/{vnetId} | Updates a virtual network configuration. |
*VnetApi* | [**updateVNetDefaultSubnet**](Apis/VnetApi.md#updatevnetdefaultsubnet) | **PUT** /vnet/{vnetId}/subnet/{subnetId}/default | Set a virtual network default subnet. |
| *VolumeApi* | [**createProjectZoneVolume**](Apis/VolumeApi.md#createprojectzonevolume) | **POST** /project/{projectId}/zone/{zoneId}/volume | Creates a new storage volume in specified zone. |
*VolumeApi* | [**deleteVolume**](Apis/VolumeApi.md#deletevolume) | **DELETE** /volume/{volumeId} | Deletes an existing storage volume. |
*VolumeApi* | [**getAllVolumes**](Apis/VolumeApi.md#getallvolumes) | **GET** /volume | Returns the IDs of storage volumes. |
*VolumeApi* | [**getPoolVolumes**](Apis/VolumeApi.md#getpoolvolumes) | **GET** /pool/{poolId}/volumes | Returns the IDs of the storage volumes existing in the pool. |
*VolumeApi* | [**getProjectZoneVolumes**](Apis/VolumeApi.md#getprojectzonevolumes) | **GET** /project/{projectId}/zone/{zoneId}/volumes | Returns the IDs of the storage volumes existing in the project in the specified zone. |
*VolumeApi* | [**getVolume**](Apis/VolumeApi.md#getvolume) | **GET** /volume/{volumeId} | Returns a description of the storage volume. |
*VolumeApi* | [**updateVolume**](Apis/VolumeApi.md#updatevolume) | **PUT** /volume/{volumeId} | Updates/resizes a storage volume configuration. |
| *ZoneApi* | [**createHost**](Apis/ZoneApi.md#createhost) | **POST** /zone/{zoneId}/host | Creates a new host. |
*ZoneApi* | [**createNetGW**](Apis/ZoneApi.md#createnetgw) | **POST** /zone/{zoneId}/netgw | Creates a new network gateway. |
*ZoneApi* | [**createNfsStorage**](Apis/ZoneApi.md#createnfsstorage) | **POST** /zone/{zoneId}/nfs | Creates a new NFS storage. |
*ZoneApi* | [**createPool**](Apis/ZoneApi.md#createpool) | **POST** /zone/{zoneId}/pool | Creates a new storage pool. |
*ZoneApi* | [**createProjectZoneInstance**](Apis/ZoneApi.md#createprojectzoneinstance) | **POST** /project/{projectId}/zone/{zoneId}/instance | Creates a new virtual machine instance in specified zone. |
*ZoneApi* | [**createProjectZoneKce**](Apis/ZoneApi.md#createprojectzonekce) | **POST** /project/{projectId}/zone/{zoneId}/kce | Creates a new KCE virtual machine in specified zone. |
*ZoneApi* | [**createProjectZoneKfs**](Apis/ZoneApi.md#createprojectzonekfs) | **POST** /project/{projectId}/zone/{zoneId}/kfs | Creates a new KFS storage volume in specified zone. |
*ZoneApi* | [**createProjectZoneKgw**](Apis/ZoneApi.md#createprojectzonekgw) | **POST** /project/{projectId}/zone/{zoneId}/kgw | Creates a new KGW in the specified zone. |
*ZoneApi* | [**createProjectZoneVolume**](Apis/ZoneApi.md#createprojectzonevolume) | **POST** /project/{projectId}/zone/{zoneId}/volume | Creates a new storage volume in specified zone. |
*ZoneApi* | [**createVNet**](Apis/ZoneApi.md#createvnet) | **POST** /zone/{zoneId}/vnet | Creates a new virtual network. |
*ZoneApi* | [**createZone**](Apis/ZoneApi.md#createzone) | **POST** /region/{regionId}/zone | Creates a new zone. |
*ZoneApi* | [**deleteZone**](Apis/ZoneApi.md#deletezone) | **DELETE** /zone/{zoneId} | Deletes an existing zone. |
*ZoneApi* | [**getAllZones**](Apis/ZoneApi.md#getallzones) | **GET** /zone | Returns the IDs of registered zones. |
*ZoneApi* | [**getProjectZoneInstances**](Apis/ZoneApi.md#getprojectzoneinstances) | **GET** /project/{projectId}/zone/{zoneId}/instances | Returns the IDs of the virtual machine instances existing in the project in the specified zone. |
*ZoneApi* | [**getProjectZoneKCEs**](Apis/ZoneApi.md#getprojectzonekces) | **GET** /project/{projectId}/zone/{zoneId}/kces | Returns the IDs of the KCE virtual machines existing in the project in the specified zone. |
*ZoneApi* | [**getProjectZoneKGWs**](Apis/ZoneApi.md#getprojectzonekgws) | **GET** /project/{projectId}/zone/{zoneId}/kgws | Returns the IDs of the KGW existing in the project in the specified zone. |
*ZoneApi* | [**getProjectZoneKfs**](Apis/ZoneApi.md#getprojectzonekfs) | **GET** /project/{projectId}/zone/{zoneId}/kfs | Returns the IDs of the KFS storage volumes existing in the project in the specified zone. |
*ZoneApi* | [**getProjectZoneVolumes**](Apis/ZoneApi.md#getprojectzonevolumes) | **GET** /project/{projectId}/zone/{zoneId}/volumes | Returns the IDs of the storage volumes existing in the project in the specified zone. |
*ZoneApi* | [**getRegionZones**](Apis/ZoneApi.md#getregionzones) | **GET** /region/{regionId}/zones | Returns the IDs of the availability zones existing in the region. |
*ZoneApi* | [**getZone**](Apis/ZoneApi.md#getzone) | **GET** /zone/{zoneId} | Returns a description of the zone |
*ZoneApi* | [**getZoneHosts**](Apis/ZoneApi.md#getzonehosts) | **GET** /zone/{zoneId}/hosts | Returns the IDs of the hosts existing in the zone. |
*ZoneApi* | [**getZoneNetGWs**](Apis/ZoneApi.md#getzonenetgws) | **GET** /zone/{zoneId}/netgws | Returns the IDs of the hosts existing in the zone. |
*ZoneApi* | [**getZoneNfsStorages**](Apis/ZoneApi.md#getzonenfsstorages) | **GET** /zone/{zoneId}/nfs | Returns the IDs of the NFS storages existing in the zone. |
*ZoneApi* | [**getZonePools**](Apis/ZoneApi.md#getzonepools) | **GET** /zone/{zoneId}/pools | Returns the IDs of the pools existing in the zone. |
*ZoneApi* | [**getZoneVNets**](Apis/ZoneApi.md#getzonevnets) | **GET** /zone/{zoneId}/vnets | Returns the IDs of the virtual networks existing in the zone. |
*ZoneApi* | [**updateZone**](Apis/ZoneApi.md#updatezone) | **PUT** /zone/{zoneId} | Updates a zone configuration. |
*ZoneApi* | [**updateZoneDefaultNfsStorage**](Apis/ZoneApi.md#updatezonedefaultnfsstorage) | **PUT** /zone/{zoneId}/nfs/{nfsId}/default | Set a zone's default NFS storage. |
*ZoneApi* | [**updateZoneDefaultPool**](Apis/ZoneApi.md#updatezonedefaultpool) | **PUT** /zone/{zoneId}/pool/{poolId}/default | Set a zone's default storage pool. |


<a name="documentation-for-models"></a>
## Documentation for Models

 - [Adapter](./Models/Adapter.md)
 - [ApiErrorBadRequest](./Models/ApiErrorBadRequest.md)
 - [ApiErrorConflict](./Models/ApiErrorConflict.md)
 - [ApiErrorForbidden](./Models/ApiErrorForbidden.md)
 - [ApiErrorInsufficientStorage](./Models/ApiErrorInsufficientStorage.md)
 - [ApiErrorNotFound](./Models/ApiErrorNotFound.md)
 - [ApiErrorUnauthorized](./Models/ApiErrorUnauthorized.md)
 - [ApiErrorUnprocessableEntity](./Models/ApiErrorUnprocessableEntity.md)
 - [Cost](./Models/Cost.md)
 - [DnsRecord](./Models/DnsRecord.md)
 - [Host](./Models/Host.md)
 - [HostCaps](./Models/HostCaps.md)
 - [HostCaps_cpu](./Models/HostCaps_cpu.md)
 - [Host_tls](./Models/Host_tls.md)
 - [Instance](./Models/Instance.md)
 - [InstanceRemoteAccess](./Models/InstanceRemoteAccess.md)
 - [InstanceState](./Models/InstanceState.md)
 - [IpRange](./Models/IpRange.md)
 - [KCE](./Models/KCE.md)
 - [KFS](./Models/KFS.md)
 - [KGW](./Models/KGW.md)
 - [KGWNat](./Models/KGWNat.md)
 - [Metadata](./Models/Metadata.md)
 - [NetGW](./Models/NetGW.md)
 - [Project](./Models/Project.md)
 - [ProjectResources](./Models/ProjectResources.md)
 - [Region](./Models/Region.md)
 - [StorageNFS](./Models/StorageNFS.md)
 - [StoragePool](./Models/StoragePool.md)
 - [Subnet](./Models/Subnet.md)
 - [Template](./Models/Template.md)
 - [VNet](./Models/VNet.md)
 - [Volume](./Models/Volume.md)
 - [Zone](./Models/Zone.md)
 - [ZoneSubnet](./Models/ZoneSubnet.md)


<a name="documentation-for-authorization"></a>
## Documentation for Authorization

<a name="ApiKeyAuth"></a>
### ApiKeyAuth

- **Type**: API key
- **API key parameter name**: X-API-Key
- **Location**: HTTP header

<a name="BearerAuth"></a>
### BearerAuth

- **Type**: HTTP Bearer Token authentication

<a name="TokenAuth"></a>
### TokenAuth

- **Type**: API key
- **API key parameter name**: x-token
- **Location**: HTTP header

