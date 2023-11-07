// This file is safe to edit. Once it exists it will not be overwritten

package server

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	swaggerapi "github.com/dalet-oss/kowabunga-api/server/api"
	"github.com/dalet-oss/kowabunga-api/server/api/adapter"
	"github.com/dalet-oss/kowabunga-api/server/api/host"
	"github.com/dalet-oss/kowabunga-api/server/api/instance"
	"github.com/dalet-oss/kowabunga-api/server/api/kce"
	"github.com/dalet-oss/kowabunga-api/server/api/kfs"
	"github.com/dalet-oss/kowabunga-api/server/api/netgw"
	"github.com/dalet-oss/kowabunga-api/server/api/nfs"
	"github.com/dalet-oss/kowabunga-api/server/api/pool"
	"github.com/dalet-oss/kowabunga-api/server/api/project"
	"github.com/dalet-oss/kowabunga-api/server/api/record"
	"github.com/dalet-oss/kowabunga-api/server/api/region"
	"github.com/dalet-oss/kowabunga-api/server/api/subnet"
	"github.com/dalet-oss/kowabunga-api/server/api/template"
	"github.com/dalet-oss/kowabunga-api/server/api/vnet"
	"github.com/dalet-oss/kowabunga-api/server/api/volume"
	"github.com/dalet-oss/kowabunga-api/server/api/zone"
)

//go:generate swagger generate server --target ../../kowabunga-api --name Kowabunga --spec ../swagger.generated.yml --api-package api --server-package server --principal interface{} --exclude-main

func configureFlags(api *swaggerapi.KowabungaAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *swaggerapi.KowabungaAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "x-token" header is set
	if api.KeyAuth == nil {
		api.KeyAuth = func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (key) x-token from header param [x-token] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.SubnetCreateAdapterHandler == nil {
		api.SubnetCreateAdapterHandler = subnet.CreateAdapterHandlerFunc(func(params subnet.CreateAdapterParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation subnet.CreateAdapter has not yet been implemented")
		})
	}
	if api.ZoneCreateHostHandler == nil {
		api.ZoneCreateHostHandler = zone.CreateHostHandlerFunc(func(params zone.CreateHostParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation zone.CreateHost has not yet been implemented")
		})
	}
	if api.ZoneCreateNetGWHandler == nil {
		api.ZoneCreateNetGWHandler = zone.CreateNetGWHandlerFunc(func(params zone.CreateNetGWParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation zone.CreateNetGW has not yet been implemented")
		})
	}
	if api.ZoneCreateNfsStorageHandler == nil {
		api.ZoneCreateNfsStorageHandler = zone.CreateNfsStorageHandlerFunc(func(params zone.CreateNfsStorageParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation zone.CreateNfsStorage has not yet been implemented")
		})
	}
	if api.ZoneCreatePoolHandler == nil {
		api.ZoneCreatePoolHandler = zone.CreatePoolHandlerFunc(func(params zone.CreatePoolParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation zone.CreatePool has not yet been implemented")
		})
	}
	if api.ProjectCreateProjectHandler == nil {
		api.ProjectCreateProjectHandler = project.CreateProjectHandlerFunc(func(params project.CreateProjectParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation project.CreateProject has not yet been implemented")
		})
	}
	if api.ProjectCreateProjectDNSRecordHandler == nil {
		api.ProjectCreateProjectDNSRecordHandler = project.CreateProjectDNSRecordHandlerFunc(func(params project.CreateProjectDNSRecordParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation project.CreateProjectDNSRecord has not yet been implemented")
		})
	}
	if api.ProjectCreateProjectZoneInstanceHandler == nil {
		api.ProjectCreateProjectZoneInstanceHandler = project.CreateProjectZoneInstanceHandlerFunc(func(params project.CreateProjectZoneInstanceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation project.CreateProjectZoneInstance has not yet been implemented")
		})
	}
	if api.ProjectCreateProjectZoneKceHandler == nil {
		api.ProjectCreateProjectZoneKceHandler = project.CreateProjectZoneKceHandlerFunc(func(params project.CreateProjectZoneKceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation project.CreateProjectZoneKce has not yet been implemented")
		})
	}
	if api.ProjectCreateProjectZoneKfsHandler == nil {
		api.ProjectCreateProjectZoneKfsHandler = project.CreateProjectZoneKfsHandlerFunc(func(params project.CreateProjectZoneKfsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation project.CreateProjectZoneKfs has not yet been implemented")
		})
	}
	if api.ProjectCreateProjectZoneVolumeHandler == nil {
		api.ProjectCreateProjectZoneVolumeHandler = project.CreateProjectZoneVolumeHandlerFunc(func(params project.CreateProjectZoneVolumeParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation project.CreateProjectZoneVolume has not yet been implemented")
		})
	}
	if api.RegionCreateRegionHandler == nil {
		api.RegionCreateRegionHandler = region.CreateRegionHandlerFunc(func(params region.CreateRegionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation region.CreateRegion has not yet been implemented")
		})
	}
	if api.VnetCreateSubnetHandler == nil {
		api.VnetCreateSubnetHandler = vnet.CreateSubnetHandlerFunc(func(params vnet.CreateSubnetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation vnet.CreateSubnet has not yet been implemented")
		})
	}
	if api.PoolCreateTemplateHandler == nil {
		api.PoolCreateTemplateHandler = pool.CreateTemplateHandlerFunc(func(params pool.CreateTemplateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation pool.CreateTemplate has not yet been implemented")
		})
	}
	if api.ZoneCreateVNetHandler == nil {
		api.ZoneCreateVNetHandler = zone.CreateVNetHandlerFunc(func(params zone.CreateVNetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation zone.CreateVNet has not yet been implemented")
		})
	}
	if api.RegionCreateZoneHandler == nil {
		api.RegionCreateZoneHandler = region.CreateZoneHandlerFunc(func(params region.CreateZoneParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation region.CreateZone has not yet been implemented")
		})
	}
	if api.AdapterDeleteAdapterHandler == nil {
		api.AdapterDeleteAdapterHandler = adapter.DeleteAdapterHandlerFunc(func(params adapter.DeleteAdapterParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation adapter.DeleteAdapter has not yet been implemented")
		})
	}
	if api.RecordDeleteDNSRecordHandler == nil {
		api.RecordDeleteDNSRecordHandler = record.DeleteDNSRecordHandlerFunc(func(params record.DeleteDNSRecordParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation record.DeleteDNSRecord has not yet been implemented")
		})
	}
	if api.HostDeleteHostHandler == nil {
		api.HostDeleteHostHandler = host.DeleteHostHandlerFunc(func(params host.DeleteHostParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation host.DeleteHost has not yet been implemented")
		})
	}
	if api.InstanceDeleteInstanceHandler == nil {
		api.InstanceDeleteInstanceHandler = instance.DeleteInstanceHandlerFunc(func(params instance.DeleteInstanceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation instance.DeleteInstance has not yet been implemented")
		})
	}
	if api.KceDeleteKCEHandler == nil {
		api.KceDeleteKCEHandler = kce.DeleteKCEHandlerFunc(func(params kce.DeleteKCEParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation kce.DeleteKCE has not yet been implemented")
		})
	}
	if api.KfsDeleteKFSHandler == nil {
		api.KfsDeleteKFSHandler = kfs.DeleteKFSHandlerFunc(func(params kfs.DeleteKFSParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation kfs.DeleteKFS has not yet been implemented")
		})
	}
	if api.NetgwDeleteNetGWHandler == nil {
		api.NetgwDeleteNetGWHandler = netgw.DeleteNetGWHandlerFunc(func(params netgw.DeleteNetGWParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation netgw.DeleteNetGW has not yet been implemented")
		})
	}
	if api.NfsDeleteNfsStorageHandler == nil {
		api.NfsDeleteNfsStorageHandler = nfs.DeleteNfsStorageHandlerFunc(func(params nfs.DeleteNfsStorageParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation nfs.DeleteNfsStorage has not yet been implemented")
		})
	}
	if api.PoolDeletePoolHandler == nil {
		api.PoolDeletePoolHandler = pool.DeletePoolHandlerFunc(func(params pool.DeletePoolParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation pool.DeletePool has not yet been implemented")
		})
	}
	if api.ProjectDeleteProjectHandler == nil {
		api.ProjectDeleteProjectHandler = project.DeleteProjectHandlerFunc(func(params project.DeleteProjectParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation project.DeleteProject has not yet been implemented")
		})
	}
	if api.RegionDeleteRegionHandler == nil {
		api.RegionDeleteRegionHandler = region.DeleteRegionHandlerFunc(func(params region.DeleteRegionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation region.DeleteRegion has not yet been implemented")
		})
	}
	if api.SubnetDeleteSubnetHandler == nil {
		api.SubnetDeleteSubnetHandler = subnet.DeleteSubnetHandlerFunc(func(params subnet.DeleteSubnetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation subnet.DeleteSubnet has not yet been implemented")
		})
	}
	if api.TemplateDeleteTemplateHandler == nil {
		api.TemplateDeleteTemplateHandler = template.DeleteTemplateHandlerFunc(func(params template.DeleteTemplateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation template.DeleteTemplate has not yet been implemented")
		})
	}
	if api.VnetDeleteVNetHandler == nil {
		api.VnetDeleteVNetHandler = vnet.DeleteVNetHandlerFunc(func(params vnet.DeleteVNetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation vnet.DeleteVNet has not yet been implemented")
		})
	}
	if api.VolumeDeleteVolumeHandler == nil {
		api.VolumeDeleteVolumeHandler = volume.DeleteVolumeHandlerFunc(func(params volume.DeleteVolumeParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation volume.DeleteVolume has not yet been implemented")
		})
	}
	if api.ZoneDeleteZoneHandler == nil {
		api.ZoneDeleteZoneHandler = zone.DeleteZoneHandlerFunc(func(params zone.DeleteZoneParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation zone.DeleteZone has not yet been implemented")
		})
	}
	if api.AdapterGetAdapterHandler == nil {
		api.AdapterGetAdapterHandler = adapter.GetAdapterHandlerFunc(func(params adapter.GetAdapterParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation adapter.GetAdapter has not yet been implemented")
		})
	}
	if api.AdapterGetAllAdaptersHandler == nil {
		api.AdapterGetAllAdaptersHandler = adapter.GetAllAdaptersHandlerFunc(func(params adapter.GetAllAdaptersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation adapter.GetAllAdapters has not yet been implemented")
		})
	}
	if api.HostGetAllHostsHandler == nil {
		api.HostGetAllHostsHandler = host.GetAllHostsHandlerFunc(func(params host.GetAllHostsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation host.GetAllHosts has not yet been implemented")
		})
	}
	if api.InstanceGetAllInstancesHandler == nil {
		api.InstanceGetAllInstancesHandler = instance.GetAllInstancesHandlerFunc(func(params instance.GetAllInstancesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation instance.GetAllInstances has not yet been implemented")
		})
	}
	if api.KceGetAllKCEsHandler == nil {
		api.KceGetAllKCEsHandler = kce.GetAllKCEsHandlerFunc(func(params kce.GetAllKCEsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation kce.GetAllKCEs has not yet been implemented")
		})
	}
	if api.KfsGetAllKFSsHandler == nil {
		api.KfsGetAllKFSsHandler = kfs.GetAllKFSsHandlerFunc(func(params kfs.GetAllKFSsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation kfs.GetAllKFSs has not yet been implemented")
		})
	}
	if api.NetgwGetAllNetGWsHandler == nil {
		api.NetgwGetAllNetGWsHandler = netgw.GetAllNetGWsHandlerFunc(func(params netgw.GetAllNetGWsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation netgw.GetAllNetGWs has not yet been implemented")
		})
	}
	if api.NfsGetAllNfsStoragesHandler == nil {
		api.NfsGetAllNfsStoragesHandler = nfs.GetAllNfsStoragesHandlerFunc(func(params nfs.GetAllNfsStoragesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation nfs.GetAllNfsStorages has not yet been implemented")
		})
	}
	if api.PoolGetAllPoolsHandler == nil {
		api.PoolGetAllPoolsHandler = pool.GetAllPoolsHandlerFunc(func(params pool.GetAllPoolsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation pool.GetAllPools has not yet been implemented")
		})
	}
	if api.ProjectGetAllProjectsHandler == nil {
		api.ProjectGetAllProjectsHandler = project.GetAllProjectsHandlerFunc(func(params project.GetAllProjectsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation project.GetAllProjects has not yet been implemented")
		})
	}
	if api.RegionGetAllRegionsHandler == nil {
		api.RegionGetAllRegionsHandler = region.GetAllRegionsHandlerFunc(func(params region.GetAllRegionsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation region.GetAllRegions has not yet been implemented")
		})
	}
	if api.SubnetGetAllSubnetsHandler == nil {
		api.SubnetGetAllSubnetsHandler = subnet.GetAllSubnetsHandlerFunc(func(params subnet.GetAllSubnetsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation subnet.GetAllSubnets has not yet been implemented")
		})
	}
	if api.TemplateGetAllTemplatesHandler == nil {
		api.TemplateGetAllTemplatesHandler = template.GetAllTemplatesHandlerFunc(func(params template.GetAllTemplatesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation template.GetAllTemplates has not yet been implemented")
		})
	}
	if api.VnetGetAllVNetsHandler == nil {
		api.VnetGetAllVNetsHandler = vnet.GetAllVNetsHandlerFunc(func(params vnet.GetAllVNetsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation vnet.GetAllVNets has not yet been implemented")
		})
	}
	if api.VolumeGetAllVolumesHandler == nil {
		api.VolumeGetAllVolumesHandler = volume.GetAllVolumesHandlerFunc(func(params volume.GetAllVolumesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation volume.GetAllVolumes has not yet been implemented")
		})
	}
	if api.ZoneGetAllZonesHandler == nil {
		api.ZoneGetAllZonesHandler = zone.GetAllZonesHandlerFunc(func(params zone.GetAllZonesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation zone.GetAllZones has not yet been implemented")
		})
	}
	if api.RecordGetDNSRecordHandler == nil {
		api.RecordGetDNSRecordHandler = record.GetDNSRecordHandlerFunc(func(params record.GetDNSRecordParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation record.GetDNSRecord has not yet been implemented")
		})
	}
	if api.HostGetHostHandler == nil {
		api.HostGetHostHandler = host.GetHostHandlerFunc(func(params host.GetHostParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation host.GetHost has not yet been implemented")
		})
	}
	if api.HostGetHostCapsHandler == nil {
		api.HostGetHostCapsHandler = host.GetHostCapsHandlerFunc(func(params host.GetHostCapsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation host.GetHostCaps has not yet been implemented")
		})
	}
	if api.HostGetHostInstancesHandler == nil {
		api.HostGetHostInstancesHandler = host.GetHostInstancesHandlerFunc(func(params host.GetHostInstancesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation host.GetHostInstances has not yet been implemented")
		})
	}
	if api.InstanceGetInstanceHandler == nil {
		api.InstanceGetInstanceHandler = instance.GetInstanceHandlerFunc(func(params instance.GetInstanceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation instance.GetInstance has not yet been implemented")
		})
	}
	if api.InstanceGetInstanceRemoteConnectionHandler == nil {
		api.InstanceGetInstanceRemoteConnectionHandler = instance.GetInstanceRemoteConnectionHandlerFunc(func(params instance.GetInstanceRemoteConnectionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation instance.GetInstanceRemoteConnection has not yet been implemented")
		})
	}
	if api.InstanceGetInstanceStateHandler == nil {
		api.InstanceGetInstanceStateHandler = instance.GetInstanceStateHandlerFunc(func(params instance.GetInstanceStateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation instance.GetInstanceState has not yet been implemented")
		})
	}
	if api.KceGetKCEHandler == nil {
		api.KceGetKCEHandler = kce.GetKCEHandlerFunc(func(params kce.GetKCEParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation kce.GetKCE has not yet been implemented")
		})
	}
	if api.KceGetKCEStateHandler == nil {
		api.KceGetKCEStateHandler = kce.GetKCEStateHandlerFunc(func(params kce.GetKCEStateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation kce.GetKCEState has not yet been implemented")
		})
	}
	if api.KfsGetKFSHandler == nil {
		api.KfsGetKFSHandler = kfs.GetKFSHandlerFunc(func(params kfs.GetKFSParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation kfs.GetKFS has not yet been implemented")
		})
	}
	if api.NetgwGetNetGWHandler == nil {
		api.NetgwGetNetGWHandler = netgw.GetNetGWHandlerFunc(func(params netgw.GetNetGWParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation netgw.GetNetGW has not yet been implemented")
		})
	}
	if api.NfsGetNfsKfsHandler == nil {
		api.NfsGetNfsKfsHandler = nfs.GetNfsKfsHandlerFunc(func(params nfs.GetNfsKfsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation nfs.GetNfsKfs has not yet been implemented")
		})
	}
	if api.NfsGetNfsStorageHandler == nil {
		api.NfsGetNfsStorageHandler = nfs.GetNfsStorageHandlerFunc(func(params nfs.GetNfsStorageParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation nfs.GetNfsStorage has not yet been implemented")
		})
	}
	if api.PoolGetPoolHandler == nil {
		api.PoolGetPoolHandler = pool.GetPoolHandlerFunc(func(params pool.GetPoolParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation pool.GetPool has not yet been implemented")
		})
	}
	if api.PoolGetPoolTemplatesHandler == nil {
		api.PoolGetPoolTemplatesHandler = pool.GetPoolTemplatesHandlerFunc(func(params pool.GetPoolTemplatesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation pool.GetPoolTemplates has not yet been implemented")
		})
	}
	if api.PoolGetPoolVolumesHandler == nil {
		api.PoolGetPoolVolumesHandler = pool.GetPoolVolumesHandlerFunc(func(params pool.GetPoolVolumesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation pool.GetPoolVolumes has not yet been implemented")
		})
	}
	if api.ProjectGetProjectHandler == nil {
		api.ProjectGetProjectHandler = project.GetProjectHandlerFunc(func(params project.GetProjectParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation project.GetProject has not yet been implemented")
		})
	}
	if api.ProjectGetProjectCostHandler == nil {
		api.ProjectGetProjectCostHandler = project.GetProjectCostHandlerFunc(func(params project.GetProjectCostParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation project.GetProjectCost has not yet been implemented")
		})
	}
	if api.ProjectGetProjectDNSRecordsHandler == nil {
		api.ProjectGetProjectDNSRecordsHandler = project.GetProjectDNSRecordsHandlerFunc(func(params project.GetProjectDNSRecordsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation project.GetProjectDNSRecords has not yet been implemented")
		})
	}
	if api.ProjectGetProjectUsageHandler == nil {
		api.ProjectGetProjectUsageHandler = project.GetProjectUsageHandlerFunc(func(params project.GetProjectUsageParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation project.GetProjectUsage has not yet been implemented")
		})
	}
	if api.ProjectGetProjectZoneInstancesHandler == nil {
		api.ProjectGetProjectZoneInstancesHandler = project.GetProjectZoneInstancesHandlerFunc(func(params project.GetProjectZoneInstancesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation project.GetProjectZoneInstances has not yet been implemented")
		})
	}
	if api.ProjectGetProjectZoneKCEsHandler == nil {
		api.ProjectGetProjectZoneKCEsHandler = project.GetProjectZoneKCEsHandlerFunc(func(params project.GetProjectZoneKCEsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation project.GetProjectZoneKCEs has not yet been implemented")
		})
	}
	if api.ProjectGetProjectZoneKfsHandler == nil {
		api.ProjectGetProjectZoneKfsHandler = project.GetProjectZoneKfsHandlerFunc(func(params project.GetProjectZoneKfsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation project.GetProjectZoneKfs has not yet been implemented")
		})
	}
	if api.ProjectGetProjectZoneVolumesHandler == nil {
		api.ProjectGetProjectZoneVolumesHandler = project.GetProjectZoneVolumesHandlerFunc(func(params project.GetProjectZoneVolumesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation project.GetProjectZoneVolumes has not yet been implemented")
		})
	}
	if api.RegionGetRegionHandler == nil {
		api.RegionGetRegionHandler = region.GetRegionHandlerFunc(func(params region.GetRegionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation region.GetRegion has not yet been implemented")
		})
	}
	if api.RegionGetRegionZonesHandler == nil {
		api.RegionGetRegionZonesHandler = region.GetRegionZonesHandlerFunc(func(params region.GetRegionZonesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation region.GetRegionZones has not yet been implemented")
		})
	}
	if api.SubnetGetSubnetHandler == nil {
		api.SubnetGetSubnetHandler = subnet.GetSubnetHandlerFunc(func(params subnet.GetSubnetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation subnet.GetSubnet has not yet been implemented")
		})
	}
	if api.SubnetGetSubnetAdaptersHandler == nil {
		api.SubnetGetSubnetAdaptersHandler = subnet.GetSubnetAdaptersHandlerFunc(func(params subnet.GetSubnetAdaptersParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation subnet.GetSubnetAdapters has not yet been implemented")
		})
	}
	if api.TemplateGetTemplateHandler == nil {
		api.TemplateGetTemplateHandler = template.GetTemplateHandlerFunc(func(params template.GetTemplateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation template.GetTemplate has not yet been implemented")
		})
	}
	if api.VnetGetVNetHandler == nil {
		api.VnetGetVNetHandler = vnet.GetVNetHandlerFunc(func(params vnet.GetVNetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation vnet.GetVNet has not yet been implemented")
		})
	}
	if api.VnetGetVNetSubnetsHandler == nil {
		api.VnetGetVNetSubnetsHandler = vnet.GetVNetSubnetsHandlerFunc(func(params vnet.GetVNetSubnetsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation vnet.GetVNetSubnets has not yet been implemented")
		})
	}
	if api.VolumeGetVolumeHandler == nil {
		api.VolumeGetVolumeHandler = volume.GetVolumeHandlerFunc(func(params volume.GetVolumeParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation volume.GetVolume has not yet been implemented")
		})
	}
	if api.ZoneGetZoneHandler == nil {
		api.ZoneGetZoneHandler = zone.GetZoneHandlerFunc(func(params zone.GetZoneParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation zone.GetZone has not yet been implemented")
		})
	}
	if api.ZoneGetZoneHostsHandler == nil {
		api.ZoneGetZoneHostsHandler = zone.GetZoneHostsHandlerFunc(func(params zone.GetZoneHostsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation zone.GetZoneHosts has not yet been implemented")
		})
	}
	if api.ZoneGetZoneNetGWsHandler == nil {
		api.ZoneGetZoneNetGWsHandler = zone.GetZoneNetGWsHandlerFunc(func(params zone.GetZoneNetGWsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation zone.GetZoneNetGWs has not yet been implemented")
		})
	}
	if api.ZoneGetZoneNfsStoragesHandler == nil {
		api.ZoneGetZoneNfsStoragesHandler = zone.GetZoneNfsStoragesHandlerFunc(func(params zone.GetZoneNfsStoragesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation zone.GetZoneNfsStorages has not yet been implemented")
		})
	}
	if api.ZoneGetZonePoolsHandler == nil {
		api.ZoneGetZonePoolsHandler = zone.GetZonePoolsHandlerFunc(func(params zone.GetZonePoolsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation zone.GetZonePools has not yet been implemented")
		})
	}
	if api.ZoneGetZoneVNetsHandler == nil {
		api.ZoneGetZoneVNetsHandler = zone.GetZoneVNetsHandlerFunc(func(params zone.GetZoneVNetsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation zone.GetZoneVNets has not yet been implemented")
		})
	}
	if api.InstanceRebootInstanceHandler == nil {
		api.InstanceRebootInstanceHandler = instance.RebootInstanceHandlerFunc(func(params instance.RebootInstanceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation instance.RebootInstance has not yet been implemented")
		})
	}
	if api.KceRebootKCEHandler == nil {
		api.KceRebootKCEHandler = kce.RebootKCEHandlerFunc(func(params kce.RebootKCEParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation kce.RebootKCE has not yet been implemented")
		})
	}
	if api.InstanceResetInstanceHandler == nil {
		api.InstanceResetInstanceHandler = instance.ResetInstanceHandlerFunc(func(params instance.ResetInstanceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation instance.ResetInstance has not yet been implemented")
		})
	}
	if api.KceResetKCEHandler == nil {
		api.KceResetKCEHandler = kce.ResetKCEHandlerFunc(func(params kce.ResetKCEParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation kce.ResetKCE has not yet been implemented")
		})
	}
	if api.InstanceResumeInstanceHandler == nil {
		api.InstanceResumeInstanceHandler = instance.ResumeInstanceHandlerFunc(func(params instance.ResumeInstanceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation instance.ResumeInstance has not yet been implemented")
		})
	}
	if api.KceResumeKCEHandler == nil {
		api.KceResumeKCEHandler = kce.ResumeKCEHandlerFunc(func(params kce.ResumeKCEParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation kce.ResumeKCE has not yet been implemented")
		})
	}
	if api.InstanceShutdownInstanceHandler == nil {
		api.InstanceShutdownInstanceHandler = instance.ShutdownInstanceHandlerFunc(func(params instance.ShutdownInstanceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation instance.ShutdownInstance has not yet been implemented")
		})
	}
	if api.KceShutdownKCEHandler == nil {
		api.KceShutdownKCEHandler = kce.ShutdownKCEHandlerFunc(func(params kce.ShutdownKCEParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation kce.ShutdownKCE has not yet been implemented")
		})
	}
	if api.InstanceStartInstanceHandler == nil {
		api.InstanceStartInstanceHandler = instance.StartInstanceHandlerFunc(func(params instance.StartInstanceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation instance.StartInstance has not yet been implemented")
		})
	}
	if api.KceStartKCEHandler == nil {
		api.KceStartKCEHandler = kce.StartKCEHandlerFunc(func(params kce.StartKCEParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation kce.StartKCE has not yet been implemented")
		})
	}
	if api.InstanceStopInstanceHandler == nil {
		api.InstanceStopInstanceHandler = instance.StopInstanceHandlerFunc(func(params instance.StopInstanceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation instance.StopInstance has not yet been implemented")
		})
	}
	if api.KceStopKCEHandler == nil {
		api.KceStopKCEHandler = kce.StopKCEHandlerFunc(func(params kce.StopKCEParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation kce.StopKCE has not yet been implemented")
		})
	}
	if api.InstanceSuspendInstanceHandler == nil {
		api.InstanceSuspendInstanceHandler = instance.SuspendInstanceHandlerFunc(func(params instance.SuspendInstanceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation instance.SuspendInstance has not yet been implemented")
		})
	}
	if api.KceSuspendKCEHandler == nil {
		api.KceSuspendKCEHandler = kce.SuspendKCEHandlerFunc(func(params kce.SuspendKCEParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation kce.SuspendKCE has not yet been implemented")
		})
	}
	if api.AdapterUpdateAdapterHandler == nil {
		api.AdapterUpdateAdapterHandler = adapter.UpdateAdapterHandlerFunc(func(params adapter.UpdateAdapterParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation adapter.UpdateAdapter has not yet been implemented")
		})
	}
	if api.RecordUpdateDNSRecordHandler == nil {
		api.RecordUpdateDNSRecordHandler = record.UpdateDNSRecordHandlerFunc(func(params record.UpdateDNSRecordParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation record.UpdateDNSRecord has not yet been implemented")
		})
	}
	if api.HostUpdateHostHandler == nil {
		api.HostUpdateHostHandler = host.UpdateHostHandlerFunc(func(params host.UpdateHostParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation host.UpdateHost has not yet been implemented")
		})
	}
	if api.InstanceUpdateInstanceHandler == nil {
		api.InstanceUpdateInstanceHandler = instance.UpdateInstanceHandlerFunc(func(params instance.UpdateInstanceParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation instance.UpdateInstance has not yet been implemented")
		})
	}
	if api.KceUpdateKCEHandler == nil {
		api.KceUpdateKCEHandler = kce.UpdateKCEHandlerFunc(func(params kce.UpdateKCEParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation kce.UpdateKCE has not yet been implemented")
		})
	}
	if api.KfsUpdateKFSHandler == nil {
		api.KfsUpdateKFSHandler = kfs.UpdateKFSHandlerFunc(func(params kfs.UpdateKFSParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation kfs.UpdateKFS has not yet been implemented")
		})
	}
	if api.NetgwUpdateNetGWHandler == nil {
		api.NetgwUpdateNetGWHandler = netgw.UpdateNetGWHandlerFunc(func(params netgw.UpdateNetGWParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation netgw.UpdateNetGW has not yet been implemented")
		})
	}
	if api.NfsUpdateNfsStorageHandler == nil {
		api.NfsUpdateNfsStorageHandler = nfs.UpdateNfsStorageHandlerFunc(func(params nfs.UpdateNfsStorageParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation nfs.UpdateNfsStorage has not yet been implemented")
		})
	}
	if api.PoolUpdatePoolHandler == nil {
		api.PoolUpdatePoolHandler = pool.UpdatePoolHandlerFunc(func(params pool.UpdatePoolParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation pool.UpdatePool has not yet been implemented")
		})
	}
	if api.PoolUpdatePoolDefaultTemplateHandler == nil {
		api.PoolUpdatePoolDefaultTemplateHandler = pool.UpdatePoolDefaultTemplateHandlerFunc(func(params pool.UpdatePoolDefaultTemplateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation pool.UpdatePoolDefaultTemplate has not yet been implemented")
		})
	}
	if api.ProjectUpdateProjectHandler == nil {
		api.ProjectUpdateProjectHandler = project.UpdateProjectHandlerFunc(func(params project.UpdateProjectParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation project.UpdateProject has not yet been implemented")
		})
	}
	if api.RegionUpdateRegionHandler == nil {
		api.RegionUpdateRegionHandler = region.UpdateRegionHandlerFunc(func(params region.UpdateRegionParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation region.UpdateRegion has not yet been implemented")
		})
	}
	if api.SubnetUpdateSubnetHandler == nil {
		api.SubnetUpdateSubnetHandler = subnet.UpdateSubnetHandlerFunc(func(params subnet.UpdateSubnetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation subnet.UpdateSubnet has not yet been implemented")
		})
	}
	if api.TemplateUpdateTemplateHandler == nil {
		api.TemplateUpdateTemplateHandler = template.UpdateTemplateHandlerFunc(func(params template.UpdateTemplateParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation template.UpdateTemplate has not yet been implemented")
		})
	}
	if api.VnetUpdateVNetHandler == nil {
		api.VnetUpdateVNetHandler = vnet.UpdateVNetHandlerFunc(func(params vnet.UpdateVNetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation vnet.UpdateVNet has not yet been implemented")
		})
	}
	if api.VnetUpdateVNetDefaultSubnetHandler == nil {
		api.VnetUpdateVNetDefaultSubnetHandler = vnet.UpdateVNetDefaultSubnetHandlerFunc(func(params vnet.UpdateVNetDefaultSubnetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation vnet.UpdateVNetDefaultSubnet has not yet been implemented")
		})
	}
	if api.VolumeUpdateVolumeHandler == nil {
		api.VolumeUpdateVolumeHandler = volume.UpdateVolumeHandlerFunc(func(params volume.UpdateVolumeParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation volume.UpdateVolume has not yet been implemented")
		})
	}
	if api.ZoneUpdateZoneHandler == nil {
		api.ZoneUpdateZoneHandler = zone.UpdateZoneHandlerFunc(func(params zone.UpdateZoneParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation zone.UpdateZone has not yet been implemented")
		})
	}
	if api.ZoneUpdateZoneDefaultNfsStorageHandler == nil {
		api.ZoneUpdateZoneDefaultNfsStorageHandler = zone.UpdateZoneDefaultNfsStorageHandlerFunc(func(params zone.UpdateZoneDefaultNfsStorageParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation zone.UpdateZoneDefaultNfsStorage has not yet been implemented")
		})
	}
	if api.ZoneUpdateZoneDefaultPoolHandler == nil {
		api.ZoneUpdateZoneDefaultPoolHandler = zone.UpdateZoneDefaultPoolHandlerFunc(func(params zone.UpdateZoneDefaultPoolParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation zone.UpdateZoneDefaultPool has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
