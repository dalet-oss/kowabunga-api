# Documentation for Kowabunga API documentation

<a name="documentation-for-api-endpoints"></a>
## Documentation for API Endpoints

All URIs are relative to */api/v1*

| Class | Method | HTTP request | Description |
|------------ | ------------- | ------------- | -------------|
| *AdapterApi* | [**createAdapter**](Apis/AdapterApi.md#createadapter) | **POST** /subnet/{ subnetId }/adapter | Creates a new network adapter. |
*AdapterApi* | [**deleteAdapter**](Apis/AdapterApi.md#deleteadapter) | **DELETE** /adapter/{ adapterId } | Deletes an existing network adapter. |
*AdapterApi* | [**listAdapters**](Apis/AdapterApi.md#listadapters) | **GET** /adapter | Returns the IDs of network adapter objects. |
*AdapterApi* | [**listSubnetAdapters**](Apis/AdapterApi.md#listsubnetadapters) | **GET** /subnet/{ subnetId }/adapters | Returns the IDs of network adapter objects. |
*AdapterApi* | [**readAdapter**](Apis/AdapterApi.md#readadapter) | **GET** /adapter/{ adapterId } | Returns a network adapter. |
*AdapterApi* | [**updateAdapter**](Apis/AdapterApi.md#updateadapter) | **PUT** /adapter/{ adapterId } | Updates a network adapter configuration. |
| *HostApi* | [**createHost**](Apis/HostApi.md#createhost) | **POST** /zone/{zoneId}/host | Creates a new host. |
*HostApi* | [**deleteHost**](Apis/HostApi.md#deletehost) | **DELETE** /host/{ hostId } | Deletes an existing computing host. |
*HostApi* | [**getZoneHosts**](Apis/HostApi.md#getzonehosts) | **GET** /zone/{zoneId}/hosts | Returns the IDs of the hosts existing in the zone. |
*HostApi* | [**listHostInstances**](Apis/HostApi.md#listhostinstances) | **GET** /host/{ hostId }/instances | Returns the IDs of virtual machine instance objects. |
*HostApi* | [**listHosts**](Apis/HostApi.md#listhosts) | **GET** /host | Returns the IDs of computing host objects. |
*HostApi* | [**readHost**](Apis/HostApi.md#readhost) | **GET** /host/{ hostId } | Returns a computing host. |
*HostApi* | [**readHostCaps**](Apis/HostApi.md#readhostcaps) | **GET** /host/{ hostId }/caps | Returns a computing host capability. |
*HostApi* | [**updateHost**](Apis/HostApi.md#updatehost) | **PUT** /host/{ hostId } | Updates a computing host configuration. |
| *InstanceApi* | [**createProjectZoneInstance**](Apis/InstanceApi.md#createprojectzoneinstance) | **POST** /project/{projectId}/zone/{zoneId}/instance | Creates a new virtual machine instance in specified zone. |
*InstanceApi* | [**deleteInstance**](Apis/InstanceApi.md#deleteinstance) | **DELETE** /instance/{ instanceId } | Deletes an existing virtual machine instance. |
*InstanceApi* | [**getProjectZoneInstances**](Apis/InstanceApi.md#getprojectzoneinstances) | **GET** /project/{projectId}/zone/{zoneId}/instances | Returns the IDs of the virtual machine instances existing in the project in the specified zone. |
*InstanceApi* | [**listHostInstances**](Apis/InstanceApi.md#listhostinstances) | **GET** /host/{ hostId }/instances | Returns the IDs of virtual machine instance objects. |
*InstanceApi* | [**listInstances**](Apis/InstanceApi.md#listinstances) | **GET** /instance | Returns the IDs of virtual machine instance objects. |
*InstanceApi* | [**readInstance**](Apis/InstanceApi.md#readinstance) | **GET** /instance/{ instanceId } | Returns a virtual machine instance. |
*InstanceApi* | [**readInstanceRemoteConnection**](Apis/InstanceApi.md#readinstanceremoteconnection) | **GET** /instance/{ instanceId }/connect | Returns a virtual machine instance remote access characteristics. |
*InstanceApi* | [**readInstanceState**](Apis/InstanceApi.md#readinstancestate) | **GET** /instance/{ instanceId }/state | Returns a virtual machine instance state. |
*InstanceApi* | [**rebootInstance**](Apis/InstanceApi.md#rebootinstance) | **PATCH** /instance/{ instanceId }/reboot | Performs a virtual machine instance software reboot. |
*InstanceApi* | [**resetInstance**](Apis/InstanceApi.md#resetinstance) | **PATCH** /instance/{ instanceId }/reset | Performs a virtual machine instance hardware reset. |
*InstanceApi* | [**resumeInstance**](Apis/InstanceApi.md#resumeinstance) | **PATCH** /instance/{ instanceId }/resume | Performs a virtual machine instance software PM resume. |
*InstanceApi* | [**shutdownInstance**](Apis/InstanceApi.md#shutdowninstance) | **PATCH** /instance/{ instanceId }/shutdown | Performs a virtual machine instance software shutdown. |
*InstanceApi* | [**startInstance**](Apis/InstanceApi.md#startinstance) | **PATCH** /instance/{ instanceId }/start | Performs a virtual machine instance hardware boot-up. |
*InstanceApi* | [**stopInstance**](Apis/InstanceApi.md#stopinstance) | **PATCH** /instance/{ instanceId }/stop | Performs a virtual machine instance hardware stop. |
*InstanceApi* | [**suspendInstance**](Apis/InstanceApi.md#suspendinstance) | **PATCH** /instance/{ instanceId }/suspend | Performs a virtual machine instance software PM suspend. |
*InstanceApi* | [**updateInstance**](Apis/InstanceApi.md#updateinstance) | **PUT** /instance/{ instanceId } | Updates a virtual machine instance configuration. |
| *KceApi* | [**createProjectZoneKce**](Apis/KceApi.md#createprojectzonekce) | **POST** /project/{projectId}/zone/{zoneId}/kce | Creates a new KCE virtual machine in specified zone. |
*KceApi* | [**deleteKCE**](Apis/KceApi.md#deletekce) | **DELETE** /kce/{ kceId } | Deletes an existing KCE (Kowabunga Compute Engine). |
*KceApi* | [**getProjectZoneKCEs**](Apis/KceApi.md#getprojectzonekces) | **GET** /project/{projectId}/zone/{zoneId}/kces | Returns the IDs of the KCE virtual machines existing in the project in the specified zone. |
*KceApi* | [**listKCEs**](Apis/KceApi.md#listkces) | **GET** /kce | Returns the IDs of KCE (Kowabunga Compute Engine) objects. |
*KceApi* | [**readKCE**](Apis/KceApi.md#readkce) | **GET** /kce/{ kceId } | Returns a KCE (Kowabunga Compute Engine). |
*KceApi* | [**readKCEState**](Apis/KceApi.md#readkcestate) | **GET** /kce/{ kceId }/state | Returns a virtual machine instance state. |
*KceApi* | [**rebootKCE**](Apis/KceApi.md#rebootkce) | **PATCH** /kce/{ kceId }/reboot | Performs a KCE (Kowabunga Compute Engine) software reboot. |
*KceApi* | [**resetKCE**](Apis/KceApi.md#resetkce) | **PATCH** /kce/{ kceId }/reset | Performs a KCE (Kowabunga Compute Engine) hardware reset. |
*KceApi* | [**resumeKCE**](Apis/KceApi.md#resumekce) | **PATCH** /kce/{ kceId }/resume | Performs a KCE (Kowabunga Compute Engine) software PM resume. |
*KceApi* | [**shutdownKCE**](Apis/KceApi.md#shutdownkce) | **PATCH** /kce/{ kceId }/shutdown | Performs a KCE (Kowabunga Compute Engine) software shutdown. |
*KceApi* | [**startKCE**](Apis/KceApi.md#startkce) | **PATCH** /kce/{ kceId }/start | Performs a KCE (Kowabunga Compute Engine) hardware boot-up. |
*KceApi* | [**stopKCE**](Apis/KceApi.md#stopkce) | **PATCH** /kce/{ kceId }/stop | Performs a KCE (Kowabunga Compute Engine) hardware stop. |
*KceApi* | [**suspendKCE**](Apis/KceApi.md#suspendkce) | **PATCH** /kce/{ kceId }/suspend | Performs a KCE (Kowabunga Compute Engine) software PM suspend. |
*KceApi* | [**updateKCE**](Apis/KceApi.md#updatekce) | **PUT** /kce/{ kceId } | Updates a KCE (Kowabunga Compute Engine) configuration. |
| *KfsApi* | [**createProjectZoneKfs**](Apis/KfsApi.md#createprojectzonekfs) | **POST** /project/{projectId}/zone/{zoneId}/kfs | Creates a new KFS storage volume in specified zone. |
*KfsApi* | [**deleteKFS**](Apis/KfsApi.md#deletekfs) | **DELETE** /kfs/{ kfsId } | Deletes an existing KFS (Kowabunga File System). |
*KfsApi* | [**getProjectZoneKfs**](Apis/KfsApi.md#getprojectzonekfs) | **GET** /project/{projectId}/zone/{zoneId}/kfs | Returns the IDs of the KFS storage volumes existing in the project in the specified zone. |
*KfsApi* | [**listKFSs**](Apis/KfsApi.md#listkfss) | **GET** /kfs | Returns the IDs of KFS (Kowabunga File System) objects. |
*KfsApi* | [**listStorageNFSKFSs**](Apis/KfsApi.md#liststoragenfskfss) | **GET** /nfs/{ nfsId }/kfs | Returns the IDs of KFS (Kowabunga File System) objects. |
*KfsApi* | [**readKFS**](Apis/KfsApi.md#readkfs) | **GET** /kfs/{ kfsId } | Returns a KFS (Kowabunga File System). |
*KfsApi* | [**updateKFS**](Apis/KfsApi.md#updatekfs) | **PUT** /kfs/{ kfsId } | Updates a KFS (Kowabunga File System) configuration. |
| *KgwApi* | [**createProjectZoneKgw**](Apis/KgwApi.md#createprojectzonekgw) | **POST** /project/{projectId}/zone/{zoneId}/kgw | Creates a new KGW in the specified zone. |
*KgwApi* | [**deleteKGW**](Apis/KgwApi.md#deletekgw) | **DELETE** /kgw/kgwId } | Deletes an existing KGW (Kowabunga Network Gateway). |
*KgwApi* | [**getProjectZoneKGWs**](Apis/KgwApi.md#getprojectzonekgws) | **GET** /project/{projectId}/zone/{zoneId}/kgws | Returns the IDs of the KGW existing in the project in the specified zone. |
*KgwApi* | [**listKGWs**](Apis/KgwApi.md#listkgws) | **GET** /kgw | Returns the IDs of KGW (Kowabunga Network Gateway) objects. |
*KgwApi* | [**readKGW**](Apis/KgwApi.md#readkgw) | **GET** /kgw/kgwId } | Returns a KGW (Kowabunga Network Gateway). |
*KgwApi* | [**updateKGW**](Apis/KgwApi.md#updatekgw) | **PUT** /kgw/kgwId } | Updates a KGW (Kowabunga Network Gateway) configuration. |
| *NetgwApi* | [**createNetGW**](Apis/NetgwApi.md#createnetgw) | **POST** /zone/{zoneId}/netgw | Creates a new network gateway. |
*NetgwApi* | [**deleteNetGW**](Apis/NetgwApi.md#deletenetgw) | **DELETE** /netgw/{ netgwId } | Deletes an existing Iris network gateway. |
*NetgwApi* | [**getZoneNetGWs**](Apis/NetgwApi.md#getzonenetgws) | **GET** /zone/{zoneId}/netgws | Returns the IDs of the hosts existing in the zone. |
*NetgwApi* | [**listNetGWs**](Apis/NetgwApi.md#listnetgws) | **GET** /netgw | Returns the IDs of Iris network gateway objects. |
*NetgwApi* | [**readNetGW**](Apis/NetgwApi.md#readnetgw) | **GET** /netgw/{ netgwId } | Returns a Iris network gateway. |
*NetgwApi* | [**updateNetGW**](Apis/NetgwApi.md#updatenetgw) | **PUT** /netgw/{ netgwId } | Updates a Iris network gateway configuration. |
| *NfsApi* | [**createNfsStorage**](Apis/NfsApi.md#createnfsstorage) | **POST** /zone/{zoneId}/nfs | Creates a new NFS storage. |
*NfsApi* | [**deleteStorageNFS**](Apis/NfsApi.md#deletestoragenfs) | **DELETE** /nfs/{ nfsId } | Deletes an existing NFS storage. |
*NfsApi* | [**getZoneNfsStorages**](Apis/NfsApi.md#getzonenfsstorages) | **GET** /zone/{zoneId}/nfs | Returns the IDs of the NFS storages existing in the zone. |
*NfsApi* | [**listStorageNFSKFSs**](Apis/NfsApi.md#liststoragenfskfss) | **GET** /nfs/{ nfsId }/kfs | Returns the IDs of KFS (Kowabunga File System) objects. |
*NfsApi* | [**listStorageNFSs**](Apis/NfsApi.md#liststoragenfss) | **GET** /nfs | Returns the IDs of NFS storage objects. |
*NfsApi* | [**readStorageNFS**](Apis/NfsApi.md#readstoragenfs) | **GET** /nfs/{ nfsId } | Returns a NFS storage. |
*NfsApi* | [**updateStorageNFS**](Apis/NfsApi.md#updatestoragenfs) | **PUT** /nfs/{ nfsId } | Updates a NFS storage configuration. |
*NfsApi* | [**updateZoneDefaultNfsStorage**](Apis/NfsApi.md#updatezonedefaultnfsstorage) | **PUT** /zone/{zoneId}/nfs/{nfsId}/default | Set a zone's default NFS storage. |
| *PoolApi* | [**createPool**](Apis/PoolApi.md#createpool) | **POST** /zone/{zoneId}/pool | Creates a new storage pool. |
*PoolApi* | [**createTemplate**](Apis/PoolApi.md#createtemplate) | **POST** /pool/{ poolId }/template | Creates a new image template. |
*PoolApi* | [**deleteStoragePool**](Apis/PoolApi.md#deletestoragepool) | **DELETE** /pool/{ poolId } | Deletes an existing storage pool. |
*PoolApi* | [**getZonePools**](Apis/PoolApi.md#getzonepools) | **GET** /zone/{zoneId}/pools | Returns the IDs of the pools existing in the zone. |
*PoolApi* | [**listStoragePoolTemplates**](Apis/PoolApi.md#liststoragepooltemplates) | **GET** /pool/{ poolId }/templates | Returns the IDs of image template objects. |
*PoolApi* | [**listStoragePoolVolumes**](Apis/PoolApi.md#liststoragepoolvolumes) | **GET** /pool/{ poolId }/volumes | Returns the IDs of storage volume objects. |
*PoolApi* | [**listStoragePools**](Apis/PoolApi.md#liststoragepools) | **GET** /pool | Returns the IDs of storage pool objects. |
*PoolApi* | [**readStoragePool**](Apis/PoolApi.md#readstoragepool) | **GET** /pool/{ poolId } | Returns a storage pool. |
*PoolApi* | [**setStoragePoolDefaultTemplate**](Apis/PoolApi.md#setstoragepooldefaulttemplate) | **PATCH** /pool/{ poolId }/template/{ templateId }/default | Performs a storage pool setting of default template. |
*PoolApi* | [**updateStoragePool**](Apis/PoolApi.md#updatestoragepool) | **PUT** /pool/{ poolId } | Updates a storage pool configuration. |
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
*RegionApi* | [**createZone**](Apis/RegionApi.md#createzone) | **POST** /region/{ regionId }/zone | Creates a new availability zone. |
*RegionApi* | [**deleteRegion**](Apis/RegionApi.md#deleteregion) | **DELETE** /region/{ regionId } | Deletes an existing region. |
*RegionApi* | [**listRegionZones**](Apis/RegionApi.md#listregionzones) | **GET** /region/{ regionId }/zones | Returns the IDs of availability zone objects. |
*RegionApi* | [**listRegions**](Apis/RegionApi.md#listregions) | **GET** /region | Returns the IDs of region objects. |
*RegionApi* | [**readRegion**](Apis/RegionApi.md#readregion) | **GET** /region/{ regionId } | Returns a region. |
*RegionApi* | [**updateRegion**](Apis/RegionApi.md#updateregion) | **PUT** /region/{ regionId } | Updates a region configuration. |
| *SubnetApi* | [**createAdapter**](Apis/SubnetApi.md#createadapter) | **POST** /subnet/{ subnetId }/adapter | Creates a new network adapter. |
*SubnetApi* | [**createSubnet**](Apis/SubnetApi.md#createsubnet) | **POST** /vnet/{ vnetId }/subnet | Creates a new network subnet. |
*SubnetApi* | [**deleteSubnet**](Apis/SubnetApi.md#deletesubnet) | **DELETE** /subnet/{ subnetId } | Deletes an existing network subnet. |
*SubnetApi* | [**listSubnetAdapters**](Apis/SubnetApi.md#listsubnetadapters) | **GET** /subnet/{ subnetId }/adapters | Returns the IDs of network adapter objects. |
*SubnetApi* | [**listSubnets**](Apis/SubnetApi.md#listsubnets) | **GET** /subnet | Returns the IDs of network subnet objects. |
*SubnetApi* | [**listVNetSubnets**](Apis/SubnetApi.md#listvnetsubnets) | **GET** /vnet/{ vnetId }/subnets | Returns the IDs of network subnet objects. |
*SubnetApi* | [**readSubnet**](Apis/SubnetApi.md#readsubnet) | **GET** /subnet/{ subnetId } | Returns a network subnet. |
*SubnetApi* | [**setVNetDefaultSubnet**](Apis/SubnetApi.md#setvnetdefaultsubnet) | **PATCH** /vnet/{ vnetId }/subnet/{ subnetId }/default | Performs a virtual network setting of default network subnet. |
*SubnetApi* | [**updateSubnet**](Apis/SubnetApi.md#updatesubnet) | **PUT** /subnet/{ subnetId } | Updates a network subnet configuration. |
| *TemplateApi* | [**createTemplate**](Apis/TemplateApi.md#createtemplate) | **POST** /pool/{ poolId }/template | Creates a new image template. |
*TemplateApi* | [**deleteTemplate**](Apis/TemplateApi.md#deletetemplate) | **DELETE** /template/{ templateId } | Deletes an existing image template. |
*TemplateApi* | [**listStoragePoolTemplates**](Apis/TemplateApi.md#liststoragepooltemplates) | **GET** /pool/{ poolId }/templates | Returns the IDs of image template objects. |
*TemplateApi* | [**listTemplates**](Apis/TemplateApi.md#listtemplates) | **GET** /template | Returns the IDs of image template objects. |
*TemplateApi* | [**readTemplate**](Apis/TemplateApi.md#readtemplate) | **GET** /template/{ templateId } | Returns a image template. |
*TemplateApi* | [**setStoragePoolDefaultTemplate**](Apis/TemplateApi.md#setstoragepooldefaulttemplate) | **PATCH** /pool/{ poolId }/template/{ templateId }/default | Performs a storage pool setting of default template. |
*TemplateApi* | [**updateTemplate**](Apis/TemplateApi.md#updatetemplate) | **PUT** /template/{ templateId } | Updates a image template configuration. |
| *VnetApi* | [**createSubnet**](Apis/VnetApi.md#createsubnet) | **POST** /vnet/{ vnetId }/subnet | Creates a new network subnet. |
*VnetApi* | [**createVNet**](Apis/VnetApi.md#createvnet) | **POST** /zone/{zoneId}/vnet | Creates a new virtual network. |
*VnetApi* | [**deleteVNet**](Apis/VnetApi.md#deletevnet) | **DELETE** /vnet/{ vnetId } | Deletes an existing virtual network. |
*VnetApi* | [**getZoneVNets**](Apis/VnetApi.md#getzonevnets) | **GET** /zone/{zoneId}/vnets | Returns the IDs of the virtual networks existing in the zone. |
*VnetApi* | [**listVNetSubnets**](Apis/VnetApi.md#listvnetsubnets) | **GET** /vnet/{ vnetId }/subnets | Returns the IDs of network subnet objects. |
*VnetApi* | [**listVNets**](Apis/VnetApi.md#listvnets) | **GET** /vnet | Returns the IDs of virtual network objects. |
*VnetApi* | [**readVNet**](Apis/VnetApi.md#readvnet) | **GET** /vnet/{ vnetId } | Returns a virtual network. |
*VnetApi* | [**setVNetDefaultSubnet**](Apis/VnetApi.md#setvnetdefaultsubnet) | **PATCH** /vnet/{ vnetId }/subnet/{ subnetId }/default | Performs a virtual network setting of default network subnet. |
*VnetApi* | [**updateVNet**](Apis/VnetApi.md#updatevnet) | **PUT** /vnet/{ vnetId } | Updates a virtual network configuration. |
| *VolumeApi* | [**createProjectZoneVolume**](Apis/VolumeApi.md#createprojectzonevolume) | **POST** /project/{projectId}/zone/{zoneId}/volume | Creates a new storage volume in specified zone. |
*VolumeApi* | [**deleteVolume**](Apis/VolumeApi.md#deletevolume) | **DELETE** /volume/{volumeId} | Deletes an existing storage volume. |
*VolumeApi* | [**getAllVolumes**](Apis/VolumeApi.md#getallvolumes) | **GET** /volume | Returns the IDs of storage volumes. |
*VolumeApi* | [**getProjectZoneVolumes**](Apis/VolumeApi.md#getprojectzonevolumes) | **GET** /project/{projectId}/zone/{zoneId}/volumes | Returns the IDs of the storage volumes existing in the project in the specified zone. |
*VolumeApi* | [**getVolume**](Apis/VolumeApi.md#getvolume) | **GET** /volume/{volumeId} | Returns a description of the storage volume. |
*VolumeApi* | [**listStoragePoolVolumes**](Apis/VolumeApi.md#liststoragepoolvolumes) | **GET** /pool/{ poolId }/volumes | Returns the IDs of storage volume objects. |
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
*ZoneApi* | [**createZone**](Apis/ZoneApi.md#createzone) | **POST** /region/{ regionId }/zone | Creates a new availability zone. |
*ZoneApi* | [**deleteZone**](Apis/ZoneApi.md#deletezone) | **DELETE** /zone/{zoneId} | Deletes an existing zone. |
*ZoneApi* | [**getAllZones**](Apis/ZoneApi.md#getallzones) | **GET** /zone | Returns the IDs of registered zones. |
*ZoneApi* | [**getProjectZoneInstances**](Apis/ZoneApi.md#getprojectzoneinstances) | **GET** /project/{projectId}/zone/{zoneId}/instances | Returns the IDs of the virtual machine instances existing in the project in the specified zone. |
*ZoneApi* | [**getProjectZoneKCEs**](Apis/ZoneApi.md#getprojectzonekces) | **GET** /project/{projectId}/zone/{zoneId}/kces | Returns the IDs of the KCE virtual machines existing in the project in the specified zone. |
*ZoneApi* | [**getProjectZoneKGWs**](Apis/ZoneApi.md#getprojectzonekgws) | **GET** /project/{projectId}/zone/{zoneId}/kgws | Returns the IDs of the KGW existing in the project in the specified zone. |
*ZoneApi* | [**getProjectZoneKfs**](Apis/ZoneApi.md#getprojectzonekfs) | **GET** /project/{projectId}/zone/{zoneId}/kfs | Returns the IDs of the KFS storage volumes existing in the project in the specified zone. |
*ZoneApi* | [**getProjectZoneVolumes**](Apis/ZoneApi.md#getprojectzonevolumes) | **GET** /project/{projectId}/zone/{zoneId}/volumes | Returns the IDs of the storage volumes existing in the project in the specified zone. |
*ZoneApi* | [**getZone**](Apis/ZoneApi.md#getzone) | **GET** /zone/{zoneId} | Returns a description of the zone |
*ZoneApi* | [**getZoneHosts**](Apis/ZoneApi.md#getzonehosts) | **GET** /zone/{zoneId}/hosts | Returns the IDs of the hosts existing in the zone. |
*ZoneApi* | [**getZoneNetGWs**](Apis/ZoneApi.md#getzonenetgws) | **GET** /zone/{zoneId}/netgws | Returns the IDs of the hosts existing in the zone. |
*ZoneApi* | [**getZoneNfsStorages**](Apis/ZoneApi.md#getzonenfsstorages) | **GET** /zone/{zoneId}/nfs | Returns the IDs of the NFS storages existing in the zone. |
*ZoneApi* | [**getZonePools**](Apis/ZoneApi.md#getzonepools) | **GET** /zone/{zoneId}/pools | Returns the IDs of the pools existing in the zone. |
*ZoneApi* | [**getZoneVNets**](Apis/ZoneApi.md#getzonevnets) | **GET** /zone/{zoneId}/vnets | Returns the IDs of the virtual networks existing in the zone. |
*ZoneApi* | [**listRegionZones**](Apis/ZoneApi.md#listregionzones) | **GET** /region/{ regionId }/zones | Returns the IDs of availability zone objects. |
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
 - [HostCPU](./Models/HostCPU.md)
 - [HostCaps](./Models/HostCaps.md)
 - [HostTLS](./Models/HostTLS.md)
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

